package godistrabutor

import(
	"fmt"
	"math/rand"
	// "sync"
	"time"
)

const (
	numberGorutines = 4
	taskLoad = 10
)

//var wg sync.WaitGroup

func init() {
	fmt.Println("---> start to init")
	rand.Seed(time.Now().Unix())
	fmt.Println("---> end to init")
}

func Distrabutor03() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGorutines)

	for gr := 1; gr <= numberGorutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	close(tasks)

	wg.Wait()

}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <- tasks
		if !ok {
			fmt.Printf("worker: %d : shutting down\n", worker)
			return
		}
		fmt.Printf("Worker: %d : Start %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
		fmt.Println("---> GlobleInt ", globleInt)
	}
}