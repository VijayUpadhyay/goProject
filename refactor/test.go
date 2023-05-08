package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	users []UserTest
	mutex sync.Mutex
}

func InitStat() *Queue {
	return &Queue{users: []UserTest{}}
}

type UserTest struct {
	Name string
}

func (q *Queue) Enqueue(user UserTest) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.users = append(q.users, user)
	fmt.Println("Enqueued:", user)
}

func (q *Queue) Dequeue() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.users) == 1 {
		q.users = []UserTest{}
	}
	q.users = q.users[1:]
}

func worker(id int, tasks <-chan UserTest, results chan<- UserTest) {
	for j := range tasks {
		fmt.Println("worker", id, "started task for id:", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished task for id:", j)
		results <- UserTest{Name: j.Name + "_new"}
	}
}

var once sync.Once

func main() {
	var queue *Queue
	once.Do(func() {
		fmt.Println("Init the queue")
		queue = InitStat()
	})

	go fillDataInQueue(queue)
	time.Sleep(1 * time.Second)
	fmt.Println(queue)
	startProcessing(queue)
	fmt.Println(queue)
}

func startProcessing(queue *Queue) {
	const workerpoolSize = 10
	jobs := make(chan UserTest, workerpoolSize)
	results := make(chan UserTest, workerpoolSize)
	for {
		for w := 1; w <= workerpoolSize; w++ {
			go worker(w, jobs, results)
		}
		for j := 0; j < workerpoolSize; j++ {
			jobs <- queue.users[j]
		}
		//defer close(jobs)
		// receive results
		for a := 0; a < workerpoolSize; a++ {
			fmt.Printf("Result for index: %d is: %+v\n", a, <-results)
		}
		for i := 0; i < workerpoolSize; i++ {
			queue.Dequeue()
		}
	}
}

func fillDataInQueue(queue *Queue) {
	for {
		for i := 0; i < 10; i++ {
			u := UserTest{Name: fmt.Sprintf("SLI_%d", i)}
			queue.Enqueue(u)
		}
		time.Sleep(1 * time.Second)
	}
}
