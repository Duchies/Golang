A Structured Approach to Optimization

CPU Profiling: `go tool pprof` to find where the CPU cycles are being spent. This often points to inefficient algorithms, excessive serialization, or tight loops.

Memory Profiling: `go tool pprof` (for heap allocations) to find memory leaks and excessive allocation pressure. This is a huge source of slowdowns in Go due to Garbage Collection (GC).

Benchmarking: Use the built-in testing.B benchmarking framework. "I would write benchmarks for suspected hot paths to get precise, reproducible performance data."

Tracing: `go tool trace` to find concurrency issues like goroutine bottlenecks, poor parallelism, or excessive blocking.
