package main

import "fmt"
import "net"

type Sequence struct {
    name string
    host string
    ports []int
    action Action
    portit int
    conn net.Conn
    rcvd chan bool
}

// Easy constructor for Sequence structs.
func NewSequence(name string, host string, ports []int, action Action) *Sequence{
    return &Sequence{name, host, ports, action, -1, nil, make(chan bool, 1)}
}

// Returns the next available port according to the sequence.
func (seq *Sequence) nextport() int {
    seq.portit = (seq.portit + 1) % len(seq.ports)
    return seq.ports[seq.portit]
}

// Returns the formatted host and port.
func (seq *Sequence) addr() string {
    return fmt.Sprintf("%s:%d", seq.host, seq.nextport())
}

// Returns the next available port according to the sequence.
func (seq *Sequence) reset() {
    seq.portit = -1
}

// Accepts the initial knock. Does not have any timeout.
func (seq *Sequence) acceptKnock() {
    addr := seq.addr()
    fmt.Println(addr)
    tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
    if err != nil {
        // If we can't resolve the host, we'll panic!
        panic(err)
    }

    netListen, err := net.Listen(tcpAddr.Network(), tcpAddr.String())
    defer netListen.Close()
    if err != nil {
        // If this port is taken, let's just go to the next one.
        fmt.Println(err)
        seq.acceptKnock()
    }
    
    conn, err := netListen.Accept()
    defer conn.Close()
    
    if err != nil {
        // Client error, reseting the sequence.
        fmt.Println(err)
        seq.reset()
    }else if seq.portit != len(seq.ports) - 1 {
        fmt.Println("Successful connection.")
    }else{
        // Launch action!
        seq.action.launch(conn)
        // Once this returns let's reset the sequence to the original setting.
        seq.reset()
    }
    // Moving onto the next port after successful knock and/or action.
    conn.Close()
    netListen.Close()
    seq.acceptKnock()
}

// Starts the initial listening for the sequence.
func (seq *Sequence) ListenCarefully(){
    seq.acceptKnock()
}