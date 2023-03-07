package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("worker %d Recovered. Error:%++v\n", id, r)
		}
		wg.Done()
	}()
	fmt.Println("worker", id, "init")
	for j := range jobs {
		if j%2 == 0 {
			fmt.Println("worker", id, "started  job", j)
			time.Sleep(time.Second)
			fmt.Println("worker", id, "finished job", j)
			results <- j
		} else if id%2 == 0 {
			s, _ := fmt.Printf("worker %d panic\n", id)
			panic(s)
		}
	}
}

func CallWorkerPool() {
	const numJobs = 5
	// 使用bufferd channel
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	wg := new(sync.WaitGroup)

	// 启动了3个worker，同时处理5个job
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, wg)
	}
	// 向job channel发送数据
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// 关闭jobs channel
	close(jobs)
	wg.Wait()
	close(results)
	for result := range results {
		fmt.Println("result:", result)
	}
}
