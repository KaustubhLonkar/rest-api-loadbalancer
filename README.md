# rest-api-loadbalancer
The objective is to create a load balancer that would return different locations randomly whenever the  API call is invoked.

# How to run the application:-
Clone the application from my github repo.
Run the below command:- go run server.go
Open a browser and enter the url as:- http://localhost:8080/api/location
Go back to the terminal where you ran the `go run` command.
The application would keep on running endlessly unless terminated manually.Just in case if you wish to terminate the program you will have to manually enter crtl + c.Upon doing so you would see a message in the browser that application has been terminated.

# The below API is used
http://localhost:8080/api/location

# Implementation
The objective shall be accomplished by using worker pool and buffered channels.

# The following are the core functionalities of our worker pool

Creation of a pool of Goroutines which listen on an input buffered channel waiting for jobs to be assigned.
Addition of jobs to the input buffered channel.
Writing results to an output buffered channel after job completion.
Read and print results from the output buffered channel.

The first step will be creation of the structs representing the job and the result.

Each Job struct has a id and a randomno for which the location has to be computed.
The Result struct has a job field which is the job for which it holds the result  in the location field.
The next step is to create the buffered channels for receiving the jobs and writing the output.

var jobs = make(chan Job, 10)  
var results = make(chan Result, 10) 

Worker Goroutines listen for new tasks on the jobs buffered channel. Once a task is complete, the result is written to the results buffered channel.

The location function does the actual job of generating the location based on random request and returning it.

We then write a function which creates a worker Goroutine which reads from the jobs channel, creates a Result struct using the current job and the location to be returned and then writes the result to the results buffered channel. This function takes a WaitGroup wg as parameter on which it will call the Done() method when all jobs have been completed.

The createWorkerPool function will create a pool of worker Goroutines.

The allocate function above takes the number of jobs to be created as input parameter.

We the create result function that reads the results channel and prints the output.