package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// ============================================================================
// Example 1: backoffAlgo.go
// ============================================================================

func example1_backoffalgo() {
	fmt.Print(1 << 3)
}

// ============================================================================
// Example 2: check.go
// ============================================================================

func example2_check() {
	val1 := "2023-01-24T17:04:46.145306+05:30"
	kv := make(map[string]interface{})
	kv["CompletedDate"] = val1

	val, perr := time.Parse(time.RFC3339, kv["CompletedDate"].(string))
	fmt.Println(val)
	fmt.Println(perr)
}

// ============================================================================
// Example 3: counter.go
// ============================================================================

var n int64

func example3_counter() {
	n = 100
	atomic.AddInt64(&n, 1)
	fmt.Println(n)
}

// ============================================================================
// Example 4: doneChannel_test.go
// ============================================================================

func testChannel(t *testing.T) {

}

// ============================================================================
// Example 5: doneChannel.go
// ============================================================================

func example5_donechannel() {
	done := make(chan bool)
	//done <- false
	go isDone(done)
	<-done
	fmt.Println("completed")
}

func isDone(done chan bool) {
	//time.Sleep(3 * time.Second)
	done <- true
}

// ============================================================================
// Example 6: jsonParser.go
// ============================================================================

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func example6_jsonparser() {
	// read our opened jsonFile as a byte array.
	var jsonFile = "users.json"
	byteValue, _ := os.ReadFile(jsonFile)

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err := json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}

// ============================================================================
// Example 7: jsonParser2.go
// ============================================================================

type User1 struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Age    int     `json:"Age"`
	Social Social1 `json:"social"`
}
type Social1 struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func example7_jsonparser2() {
	var jsonFile = "user2.json"
	byteValue, _ := os.ReadFile(jsonFile)

	// we initialize our Users array
	var users map[string]User1

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err := json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	fmt.Println(users["user2a"])
}

// ============================================================================
// Example 8: monthdata.go
// ============================================================================

const quarter int = 3

func example8_monthdata() {

	/*fmt.Println(int(t.Month()))
	fmt.Println(int(time.Now().UTC().Month()))*/
	fmt.Println(GetCurrentQuarterAndYear())

}

func GetCurrentQuarterAndYear() (int, int) {
	t := time.Date(2022, time.Month(10), 11, 12, 13, 44, 213, time.Local)
	return int(math.Ceil(float64((int(t.UTC().Month())-1)/quarter))) + 1, t.Year()
}

// ============================================================================
// Example 9: queue.go
// ============================================================================

// LockfreeQueue is a goroutine-safe Queue implementation.
// The overall performance of LockfreeQueue is much better than List+Mutex(standard package).
type LockfreeQueue[T any] struct {
	head  unsafe.Pointer
	tail  unsafe.Pointer
	dummy lfqNode[T]
}

// NewLockfreeQueue is the only way to get a new, ready-to-use LockfreeQueue.
//
// Example:
//
//	lfq := queue.NewLockfreeQueue[int]()
//	lfq.Push(100)
//	v, ok := lfq.Pop()
func NewLockfreeQueue[T any]() *LockfreeQueue[T] {
	var lfq LockfreeQueue[T]
	lfq.head = unsafe.Pointer(&lfq.dummy)
	lfq.tail = lfq.head
	return &lfq
}

// Pop returns (and removes) an element from the front of the queue and true if the queue is not empty,
// otherwise it returns a default value and false if the queue is empty.
// It performs about 100% better than list.List.Front() and list.List.Remove() with sync.Mutex.
func (lfq *LockfreeQueue[T]) Pop() (T, bool) {
	for {
		h := atomic.LoadPointer(&lfq.head)
		rh := (*lfqNode[T])(h)
		n := (*lfqNode[T])(atomic.LoadPointer(&rh.next))
		if n != nil {
			if atomic.CompareAndSwapPointer(&lfq.head, h, rh.next) {
				return n.val, true
			} else {
				continue
			}
		} else {
			var v T
			return v, false
		}
	}
}

// Push inserts an element to the back of the queue.
// It performs exactly the same as list.List.PushBack() with sync.Mutex.
func (lfq *LockfreeQueue[T]) Push(val T) {
	node := unsafe.Pointer(&lfqNode[T]{val: val})
	for {
		rt := (*lfqNode[T])(atomic.LoadPointer(&lfq.tail))
		//t := atomic.LoadPointer(&lfq.tail)
		//rt := (*lfqNode[T])(t)
		if atomic.CompareAndSwapPointer(&rt.next, nil, node) {
			atomic.StorePointer(&lfq.tail, node)
			// If dead loop occurs, use CompareAndSwapPointer instead of StorePointer
			// atomic.CompareAndSwapPointer(&lfq.tail, t, node)
			return
		} else {
			continue
		}
	}
}

type lfqNode[T any] struct {
	val  T
	next unsafe.Pointer
}

// ============================================================================
// Example 10: regextest.go
// ============================================================================

func example10_regextest() {
	//str1 := "Hello X42 I'm a Y-32.35 string Z30"
	str1 := "GrossMargin"
	fmt.Println("getSectionName(str1): ", getSectionName(str1))
	fmt.Println("name: ", getName("model.BalanceSheet"))
	getNameTrimR()
}

func getSectionName(str1 string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)

	fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

	submatchall := re.FindAllString(str1, -1)
	/*	for _, element := range submatchall {
		fmt.Println(element)
	}*/
	return strings.Join(submatchall, " ")
}

func getName(s string) string {
	return strings.TrimPrefix(s, "model.")
}

func getNameTrimR() {
	fmt.Println(strings.TrimSuffix("invalid section: model.Test", "model."))
}

// ============================================================================
// Example 11: strTrim.go
// ============================================================================

func example11_strtrim() {
	str := "$VA_REVENUE_TOTAL_5YR_CAGR$"
	fmt.Print(strings.Trim(str, "$"))
}

// ============================================================================
// Example 12: syncMap.go
// ============================================================================

type Foo struct {
	A int
	B string
}
type Bar struct {
	C int
	D string
}

func example12_syncmap() {
	var temp sync.Map
	foo := Foo{A: 5, B: "bar"}
	bar := Bar{C: 5, D: "bar"}
	err := json.Unmarshal([]byte(fmt.Sprintf("%v", foo)), temp)
	if err != nil {
		return
	}
	fmt.Println(temp)
	err = json.Unmarshal([]byte(fmt.Sprintf("%v", bar)), temp)
	if err != nil {
		return
	}
	fmt.Println(temp)
}

// ============================================================================
// Example 13: test.go
// ============================================================================

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

func example13_test() {
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

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 11 examples from refactor")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_backoffalgo ---")
	example1_backoffalgo()

	fmt.Println("\n--- example2_check ---")
	example2_check()

	fmt.Println("\n--- example3_counter ---")
	example3_counter()

	fmt.Println("\n--- example5_donechannel ---")
	example5_donechannel()

	fmt.Println("\n--- example6_jsonparser ---")
	example6_jsonparser()

	fmt.Println("\n--- example7_jsonparser2 ---")
	example7_jsonparser2()

	fmt.Println("\n--- example8_monthdata ---")
	example8_monthdata()

	fmt.Println("\n--- example10_regextest ---")
	example10_regextest()

	fmt.Println("\n--- example11_strtrim ---")
	example11_strtrim()

	fmt.Println("\n--- example12_syncmap ---")
	example12_syncmap()

	fmt.Println("\n--- example13_test ---")
	example13_test()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
	fmt.Println("Note: doneChannel_test.go and queue.go had no main() functions")
}
