package main

//imoport some packages
import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func mystery(x int) int {
	if x == 0 {
		return 0
	}
	return x + mystery(x-1)
}

// create func using each library
func timeNow() {
	fmt.Println(time.Now())
}
func randInt() {
	fmt.Println(rand.Int())
}
func computeLoss(x float64, y float64) {
	fmt.Println(math.Abs(x - y))
}

// use a database
func useDatabase() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		runtime.Gosched()
	}()
	wg.Wait()
}

func main() {
	// call the functions
	timeNow()
	randInt()
	computeLoss(1.0, 2.0)
	useDatabase()
	// call the add function
	fmt.Println(add(1, 2))
	// call the mystery function
	fmt.Println(mystery(5))
	// print the arguments
	fmt.Println(os.Args)
	// print the environment variables
	fmt.Println(os.Environ())
	// print the number of CPUs
	fmt.Println(runtime.NumCPU())
	// print the number of goroutines
	fmt.Println(runtime.NumGoroutine())
	// print the GOMAXPROCS
	fmt.Println(runtime.GOMAXPROCS(0))
	// print the number of cgo calls
	fmt.Println(runtime.NumCgoCall())
	// print the number of memory allocations
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Println(memStats.Alloc)
	// print the number of memory allocations
	fmt.Println(memStats.Frees)
}
