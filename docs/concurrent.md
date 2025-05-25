# Concurrency

This section include examples of concurrency patterns.

## Goroutine Producer

This pattern is where a function **returns a channel**, and a **goroutine is started inside the function** to produce values on that channel asynchronously.

Common Use Cases:

* Performing background computations
* Streaming data (e.g., from a file, network, or other long-running source)
* Decoupling data generation from data consumption

Why it’s used:

* **Non-blocking API design**: The function immediately returns, allowing the caller to start consuming from the channel without waiting for the entire operation to complete.
* **Concurrency**: The actual work is done in a goroutine, allowing the Go scheduler to manage it alongside other concurrent tasks.

Working examples:

* [Main package](../examples/concurrent/ex1/main.go)
* [Producer](../internal/csvutil/csvutil.go)

## Fan Out and Worker Pool

A Fan Out pattern has these characteristics:

* Multiple Goroutines
* A shared input channel
* A shared output channel
* Coordinated completion (via WaitGroup)
* Aggregated result  

A worker pool is used to support multiple processing Goroutines.

Here is a [working example](../examples/concurrent/ex2/main.go)
