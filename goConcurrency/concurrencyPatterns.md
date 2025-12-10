# Concurrency Patterns

Established design aproaches for structuring concurrent programs using goroutines and channels. Help build efficient, scalable apps while avoiding race conditions and deadlocks

## Fan-in

Concurrency pattern merging multiple input channels into single output channel. Allows collecting results from multiple goroutines. Typically implemented with select statement or separate goroutines for each input. Useful for aggregating parallel processing result. 

## Fan-out

Concurrency pattern distributing work from single source to multiple workers. Typically uses one input channel feeding multiple goroutines. Each worker processes items independently. Useful for parallelizing CPU-intensive tasks and increasing throughput throught parallel processing.

## Pipeline

Concurrency pattern chaining processing stages where output of one stage becomes input of next. Each stage runs concurrently using goroutines and channels. Common in data processing, transformation workflows, and streaming applications.