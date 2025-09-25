# transaction-management

This repository explore ideas of managing the database transaction in Go. 

My primary requirements is that I want the transaction management to be explicit, therefore I ignore the ideas of passing Transaction via Go's Context. 

An interesting article from threedots.tech; [Database Transactions in Go with Layered Architecture](https://threedots.tech/post/database-transactions-in-go/) show various approaches and examples. It demonstrate the good and the bad for each approach. I like the approach `The Transaction Provider` the most. 

There is another approach proposed by **Metalfmmetalfm** in the Medium article; [Yet Another Way to Handle Transactions in Go Using Clean Architecture](https://medium.com/@metalfmmetalfm/yet-another-way-to-handle-transactions-in-go-using-clean-architecture-fe45d0ebbdd5). 

This article proposed adding another abstract layer for containing repos' methods and repositories needed in a service with transaction. 

Although this approach add the transaction layer with clear separation, the methods introduce a lot of abstraction interface, and may not suitable for every project. It also utilize the Generic with Interface, which add more Cognitive Load, High complexity to understand the flow.


Another common approach is passing Transaction via context, which I choose to ignore, from this article [SQL Transactions in Go: The Good Way](https://blog.thibaut-rousseau.com/blog/sql-transactions-in-go-the-good-way/). It create a transaction provider, which have Transaction function that takes a context, wraps the repo functions, inject the given context with a Transaction, and the repofunction will retrieve the transaction inside the repo function.  

