// goroutines
// Concurrency and parallelism are different concepts.
// Concurrency is about managing multiple tasks that can make progress
// independently, possibly by interleaving their execution on one CPU core.
// Parallelism is about executing multiple tasks at the same time, typically
// using multiple CPU cores or other execution units.

// EXAMPLES:

package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"data1", "data2", "data3", "data4", "data5"}
var results = []string{}

func fetch() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
		// this "go keyword" will run the dbCall function in a
		// separate goroutine, and it will not block the main thread
	}
	wg.Wait() // this will block the main thread until all goroutines are done
	fmt.Println("Total time taken to fetch all data => ", time.Since(t0))
	fmt.Printf("The results are => %v\n", results)
}

/*

pthread_mutex_t mu = PTHREAD_MUTEX_INITIALIZER;
int counter;

void increment() {
    pthread_mutex_lock(&mu);
    counter++;
    pthread_mutex_unlock(&mu);
}

This is the mutex equivalent in C, and it is used to protect the
counter variable from being accessed by multiple threads
at the same time.

*/

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database call is => ", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done() // this will decrement the counter of the WaitGroup by 1
}

func hiCall() {
	runHiInMultipleGoroutines()
}

func runHiInMultipleGoroutines() {
	var i int = 0
	for i < 10 {
		wg.Add(1)
		go hi()
		i++
	}
	wg.Wait()
}

func hi() {
	var name string = "tarokh"
	fmt.Println(name)
	wg.Done()

}
