package main

func main(){
    // #####
    // Define the sequence and its action here.
    knockseq := NewSequence("Echo action", "127.0.0.1", []Knock{{1024, -1}, {1025, 5}, {1026, 5}}, Echo{}, true)
    // #####
    immortality := make(chan bool, 1)
    // Main infinite loop to make sure we're always running.
    // Also allows for multiple clients.
    go knockseq.ListenCarefully()
    <- immortality 
}