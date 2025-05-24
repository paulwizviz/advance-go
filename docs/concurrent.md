# Concurrency

This section include examples of concurrency patterns.

## CSV Files and Channels

This [example](../examples/concurrent/ex1/main.go) demonstrates extraction of csv file content via goroutine and channeling it for extraction by main routine.

## Fan Out and Worker Pool

A Fan Out pattern has these characteristics:

* Multiple Goroutines
* A shared input channel
* A shared output channel
* Coordinated completion (via WaitGroup)
* Aggregated result  

A worker pool is used to support multiple processing Goroutines.

Here is a [working example](../examples/concurrent/ex2/main.go)
