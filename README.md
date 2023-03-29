## ChatGPT Golang Exercises

A series of exercises ChatGPT generated to help me learn idiomatic patterns of programming in Go

## OOP

- Create a package that defines a Shape interface with two methods Area() and Perimeter(). Then, create structs for Rectangle, Circle, and Triangle that implement the Shape interface. Use these structs to calculate the area and perimeter of various shapes.

- Create a package that defines a Vehicle interface with two methods Start() and Stop(). Then, create structs for Car, Truck, and Motorcycle that implement the Vehicle interface. Use these structs to start and stop different types of vehicles.

- Create a package that defines a Person struct with public and private methods. The public methods should be GetName(), SetName(), GetAge(), and SetAge(). The private methods should be used to validate the input for SetName() and SetAge(). Write tests to ensure that the public methods behave correctly.

- Create a package that defines a BankAccount struct with public and private methods. The public methods should be Deposit(), Withdraw(), and Balance(). The private methods should be used to validate the input for Deposit() and Withdraw(). Write tests to ensure that the public methods behave correctly.

- Create a package that defines a Game interface with two methods Play() and GameOver(). Then, create structs for TicTacToe, Chess, and Checkers that implement the Game interface. Use these structs to play different types of games and determine when the game is over.


## Concurrency

- Write a program that simulates a bank account with multiple concurrent transactions. Each transaction will either deposit or withdraw a random amount of money from the account. Use a mutex to protect the account balance from concurrent access and a waitgroup to ensure that all transactions have completed before printing the final balance.

- Create a program that generates a large amount of data and then processes it concurrently using a fixed number of goroutines. Use a channel to communicate between the producer and consumer goroutines and use a waitgroup to ensure that all goroutines have completed before exiting the program.

- Write a program that simulates a race condition by concurrently incrementing a variable. Use a mutex to protect the variable from concurrent access and print the final value of the variable after all goroutines have completed.

- Create a program that downloads multiple files concurrently using goroutines. Use a waitgroup to ensure that all downloads have completed before exiting the program. Use a mutex to protect the downloaded files from concurrent writes.

- Write a program that simulates a restaurant with multiple tables and customers. Use a waitgroup to ensure that all customers have finished eating before closing the restaurant. Use a mutex to protect the tables from concurrent access and ensure that each customer is assigned to a free table before being served.

## Data structures, slicing, and maps


