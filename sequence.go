package main

import "fmt"
import "net"
import "time"

type Knock struct {
    port int
    timeout int
}

type Sequence struct {
    name string
    host string
    ports []Knock
    action Action
    portit int
    conn net.Conn
    debug bool
}

// Easy constructor for Sequence structs.
func NewSequence(name string, host string, ports []Knock, action Action, debug bool) *Sequence{
    return &Sequence{name, host, ports, action, -1, nil, debug}
}

func NewLocalSequence(ports []Knock, action Action, debug bool) *Sequence{
    return &Sequence{"Local sequence", "127.0.0.1", ports, action, -1, nil, debug}
}

// Returns the next available port according to the sequence.
func (seq *Sequence) next() Knock {
    seq.portit = (seq.portit + 1) % len(seq.ports)
    return seq.ports[seq.portit]
}

// Returns the formatted host and port.
func (seq *Sequence) addr() (string, int){
    kn := seq.next()
    return fmt.Sprintf("%s:%d", seq.host, kn.port), kn.timeout
}

// Returns the next available port according to the sequence.
func (seq *Sequence) reset() {
    seq.portit = -1
}

// Accepts the initial knock. Does not have any timeout.
func (seq *Sequence) acceptKnock() {
    addr, timeout := seq.addr()
    tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
    if err != nil {
        // If we can't resolve the host, we'll panic!
        panic(err)
    }
    netListen, err := net.ListenTCP("tcp", tcpAddr)
    defer netListen.Close()
    if timeout > 0 {
        netListen.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
    }
    
    if err != nil {
        // If this port is taken, let's just go to the next one.
        if seq.debug{
            fmt.Println(err)
        }
        netListen.Close()
        seq.acceptKnock()
    }
    
    if seq.debug{
        fmt.Println("Listening on", addr, ".")
    }

    conn, err := netListen.Accept()
    defer conn.Close()
    
    if err != nil {
        // Client error, reseting the sequence.
        fmt.Println(err)
        seq.reset()
        if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
            fmt.Println("Timed out")
        }

    }else{
        if seq.portit != len(seq.ports) - 1 && seq.debug {
            fmt.Println("Successful connection.")
        }else{
            if seq.debug{
                fmt.Println("Launching action %s.", seq.name)
            }
            // Launch action!
            seq.action.launch(conn)
            // Once this returns let's reset the sequence to the original setting.
            seq.reset()
        }
        conn.Close()
    }
    // Moving onto the next port after successful knock and/or action.
    netListen.Close()
    seq.acceptKnock()
}

// Starts the initial listening for the sequence.
func (seq *Sequence) ListenCarefully(){
    seq.acceptKnock()
}