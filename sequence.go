package main

import "fmt"
import "net"

type Sequence struct {
    name string
    host string
    ports []int
    action Action
    portit int
}

func (seq Sequence) nextport() int {
    seq.portit = (seq.portit + 1) % len(seq.ports)
    return seq.ports[seq.portit]
}

func (seq Sequence) addr() string {
    return fmt.Sprintf("127.0.0.1:%d", seq.nextport())
}

func (seq Sequence) AcceptKnock() net.Conn {
    fmt.Println(seq.addr())
    tcpAddr, err := net.ResolveTCPAddr("tcp", seq.addr())
    if err != nil {
        // If we can't resolve the host, we'll panic!
        panic(err)
    }
    netListen, err := net.Listen(tcpAddr.Network(), tcpAddr.String())
    if err != nil {
        // If this port is taken, let's just go to the next one.
        fmt.Errorf("", err)
    }
    conn, err := netListen.Accept()
    if err != nil {
        fmt.Errorf("Client error", err)
        return nil
    }else{
        return conn
    }
}