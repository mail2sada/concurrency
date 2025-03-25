## 🔹 Core Concurrency Patterns  

### Pipeline – Streamlined Data Processing  
📌 **Break tasks into multiple stages**, where each stage processes and passes data downstream.  
🔹 **Use Case:** Data streaming, ETL, analytics pipelines.  

### Worker Pool – Controlled Parallelism  
📌 **Restrict the number of concurrent goroutines** to avoid system overload.  
🔹 **Use Case:** Web crawlers, database queries.  

### Fan-Out, Fan-In – Parallel Execution with Aggregation  
📌 **Fan-Out:** Distribute work across multiple goroutines.  
📌 **Fan-In:** Merge results into a single output.  
🔹 **Use Case:** Concurrent API calls, parallel computations.  

### Pub-Sub – Message Broadcasting  
📌 **Publisher sends messages** to multiple independent subscribers asynchronously.  
🔹 **Use Case:** Event-driven systems, distributed messaging.  

### Or-Done – Graceful Cancellation  
📌 **Exit goroutines cleanly** when a cancellation signal is received.  
🔹 **Use Case:** Network requests, background jobs.  