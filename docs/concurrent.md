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

## Fan Out and Fan in

The Fan-Out amd Fan-In Concurrency pattern in Go is a concurrent design pattern where a single job or input is fanned out (distributed) to multiple workers (goroutines), allowing tasks to be processed in parallel. This pattern is especially useful for increasing throughput and making full use of multi-core processors.

Key characteristics:

* Multiple Goroutines
* A shared input channel
* A shared output channel
* Coordinated completion (via WaitGroup)
* Aggregated result  

Here is a [working example](../examples/concurrent/ex2/main.go)
