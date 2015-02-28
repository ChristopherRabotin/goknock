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
func NewSequence(name string, host string, ports []int, action Action) Sequence{
    return Sequence{name, host, ports, action, -1, nil, make(chan bool, 1)}
}

// Returns the next available port according to the sequence.
func (seq Sequence) nextport() int {
    seq.portit = (seq.portit + 1) % len(seq.ports)
    return seq.ports[seq.portit]
}

// Returns the formatted host and port.
func (seq Sequence) addr() string {
    return fmt.Sprintf("%s:%d", seq.host, seq.nextport())
}

// Accepts the initial knock. Does not have any timeout.
func (seq Sequence) acceptKnock() {
    <- seq.rcvd
    tcpAddr, err := net.ResolveTCPAddr("tcp", seq.addr())
    if err != nil {
        // If we can't resolve the host, we'll panic!
        panic(err)
    }
    fmt.Println(seq.addr())
    netListen, err := net.Listen(tcpAddr.Network(), tcpAddr.String())
    if err != nil {
        // If this port is taken, let's just go to the next one.
        fmt.Errorf("", err)
    }
    conn, err := netListen.Accept()
    if err != nil {
        fmt.Errorf("Client error", err)
    }else if seq.portit != len(seq.ports) - 1 {
        fmt.Println("Successful connection.")
        // We've got a good connection, let's close it and move to the next one,
        // or launch the action.
        conn.Close()
        fmt.Println("Yup yup")
        go seq.acceptKnock()
        seq.rcvd <- true
    }else{
        // Launch action!
        seq.action.launch(conn)
        // Once this returns let's reset the sequence to the original setting.
        seq.portit = 0
        go seq.acceptKnock()
    }
}

// Starts the initial listening for the sequence.
func (seq Sequence) ListenCarefully(){
    seq.rcvd <- true
    seq.acceptKnock()
}