package loadbalancer

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

//Job  Struct
type Job struct {
	id       int
	randomno int
}

//Result Struct
type Result struct {
	job      Job
	location string
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// GetLocation func to generate location based on request
func GetLocation(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

}

//The location function does the actual job of generating the location based on random request and returning it.

func location(number int) string {
	var location string
	switch {
	case number <= 500 && number >= 1:
		location = "Hyderabad"
	case number >= 501 && number <= 999:
		location = "Bangalore"
	default:
		fmt.Println("Invalid")
	}
	time.Sleep(2 * time.Second)
	return location
}

//Worker Goroutine reads from the jobs channel, creates a Result struct using the current job and the location to be returned and then writes the result to the results buffered channel.
//This function takes a WaitGroup wg as parameter on which it will call the Done() method when all jobs have been completed.
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, location(job.randomno)}
		results <- output
	}
	wg.Done()
}

//The createWorkerPool function will create a pool of worker Goroutines
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

//The allocate function takes the number of jobs to be created as input parameter.
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

//function that reads the results channel and prints the output.
func result(done chan bool) {
	for result := range results {
		fmt.Println("Job id , input random no corresponding to the request received and the location obtained corresponding to request ", result.job.id, result.job.randomno, result.location)
	}
	done <- true
}
