package main

func main() {
	// #####
	// Define sequences and their actions here.
    sequences := []*Sequence{NewSequence("Echo action", "127.0.0.1", []Knock{{1024, -1}, {1025, -1}, {1026, 5}}, Echo{}, true),
                             NewSequence("Echo 2", "127.0.0.1", []Knock{{2024, -1}, {2025, -1}, {2026, 5}}, Echo{}, true)}
	// #####
	immortality := make(chan bool, 1)
	// Main infinite loop to make sure we're always running.
	// Also allows for multiple clients.
    for i := 0; i < len(sequences); i++ {
        go sequences[i].ListenCarefully()
    }
	<-immortality
}
