package main

func main(){
    // #####
    // Define the sequence and its action here.
    knockseq := Sequence{"Echo action", "127.0.0.1", []int{1024, 1025, 1026}, Echo{}, -1}
    // #####
    // Main infinite loop to make sure we're always running.
    for {
        // Accepts the initial knock without any timeout.
        HandleConnection(knockseq.AcceptKnock())
    }
}