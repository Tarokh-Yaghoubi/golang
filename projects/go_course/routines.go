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
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var dbData = []string{"data1", "data2", "data3", "data4", "data5"}

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
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database call is => ", dbData[i])
	wg.Done() // this will decrement the counter of the WaitGroup by 1
}
