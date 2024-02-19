package main

import (
	"fmt"
	"sync"
	"time"
)

func SayA(wg *sync.WaitGroup) {
	startA := time.Now()
	fmt.Println("Start : A")
	time.Sleep(1 * time.Second)
	fmt.Println("End : A = ", time.Since(startA))
	wg.Done()
}

func main() {
	startSyncTime := time.Now()
	// Synchronous or Sequential
	func() {
		start1 := time.Now()
		fmt.Println("Start : 1")
		time.Sleep(1 * time.Second)
		fmt.Println("End : 1 = ", time.Since(start1))
	}()
	func() {
		start2 := time.Now()
		fmt.Println("Start : 2")
		time.Sleep(1 * time.Second)
		fmt.Println("End : 2 = ", time.Since(start2))
	}()
	func() {
		start3 := time.Now()
		fmt.Println("Start : 3")
		time.Sleep(1 * time.Second)
		fmt.Println("End : 3 = ", time.Since(start3))
	}()
	fmt.Println("Total Run Time : ", time.Since(startSyncTime))
	fmt.Println("================")
	time.Sleep(2 * time.Second)
	startAsyncTime := time.Now()
	// Asynchronous
	var wg sync.WaitGroup // สร้างตัวแปร WaitGroup
	wg.Add(3)             // เพิ่มจำนวน goroutine ที่ต้องรอใน WaitGroup
	go SayA(&wg)          // ส่งทางการอ้างอิง *sync.WaitGroup
	go func() {
		startB := time.Now()
		fmt.Println("Start : B")
		time.Sleep(1 * time.Second)
		fmt.Println("End : B = ", time.Since(startB))
		wg.Done() // บอก WaitGroup ว่า goroutine นี้เสร็จสิ้นแล้ว
	}()
	go func() {
		startC := time.Now()
		fmt.Println("Start : C")
		time.Sleep(1 * time.Second)
		fmt.Println("End : C = ", time.Since(startC))
		wg.Done()
	}()
	wg.Wait() // รอให้ทุก goroutine ใน WaitGroup เสร็จสิ้น
	//fmt.Println("All go routines have finished")
	fmt.Println("Total Run Time : ", time.Since(startAsyncTime))
}
