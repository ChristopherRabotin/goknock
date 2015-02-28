package main

func main(){
    // #####
    // Define the sequence and its action here.
    knockseq := NewSequence("Echo action", "127.0.0.1", []int{1024, 1025, 1026}, Echo{})
    // #####
    // Main infinite loop to make sure we're always running.
    // Also allows for multiple clients.
    for {
        knockseq.AcceptKnocks()
    }
}