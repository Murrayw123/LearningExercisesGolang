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

- Write a program that simulates a restaurant with multiple tables and customers. Use a waitgroup to ensure that all customers have finished eating before closing the restaurant. Use a mutex to protect the tables from concurrent access and ensure that each customer is assigned to a free table before being served.

### Idiomatic Concurrency

- Write a function called fanIn that takes two input channels, c1 and c2, and returns a single channel that combines the values from both input channels. The function should read values from both channels and send them to the output channel as they arrive. Finally, you can close the output channel after both input channels have been closed.

- Implement a simple producer-consumer scenario using channels. Create a producer goroutine that generates a sequence of numbers and sends them to a channel. Create multiple consumer goroutines that receive numbers from the channel and perform some computation on them.

- Write a function called fanOut that takes an input channel in and a slice of output channels outs. The function should read values from the input channel and distribute them to all output channels in a round-robin fashion. When the input channel is closed, all output channels should be closed as well.

- Write a function called timeout that takes an input channel in and a timeout duration. The function should read from the input channel, but if no value is received within the specified timeout, it should print a timeout message and return. You can achieve this by combining a select statement with a time.After channel.

## Lists, slicing, and maps

- convert a string array of numbers to a slice, and then calculates the average of the numbers in the slice.

- Implement a function that takes a slice of integers and returns a new slice that contains only the unique elements of the original slice.

- Write a program that reads a string from the command line, stores the frequency of each word in the string in a map, and then prints out the word frequencies.

- Implement a function that takes a map[string]int and returns a new map with the keys and values swapped (i.e., the original values become keys, and the original keys become values).

- Write a program that reads a list of words from the command line, stores them in a slice, and then sorts the slice in alphabetical order.

### Testing

- String Reversal Function: Write a function that takes in a string and returns it reversed. Write unit tests to confirm that your function is working correctly. 

- Write a function that takes in a slice of integers and a function that modifies each integer. This function should return a new slice with each integer modified by the function.
