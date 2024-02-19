package main

import (
	"fmt"
	"sync"
	"time"
)

func SayA(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("A")
	wg.Done()
}

func main() {
	startSyncTime := time.Now()
	// Synchronous or Sequential
	func() {
		time.Sleep(1 * time.Second)
		fmt.Println("1")
	}()
	func() {
		time.Sleep(1 * time.Second)
		fmt.Println("2")
	}()
	func() {
		time.Sleep(1 * time.Second)
		fmt.Println("3")
	}()
	fmt.Println("Total Run Time : ", time.Since(startSyncTime))
	fmt.Println("================")
	time.Sleep(2 * time.Second)
	startAsyncTime := time.Now()
	// Asynchronous
	var wg sync.WaitGroup // สร้างตัวแปร WaitGroup
	wg.Add(3)             // เพิ่มจำนวน goroutine ที่ต้องรอใน WaitGroup
	// ===== แบบนี้ไม่ได้เพราะอยู่นอก Scope =====
	// go fmt.Println("Hello, world : A")
	// wg.Done()
	// ====================================
	go SayA(&wg) // ส่งทางการอ้างอิง *sync.WaitGroup
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("B")
		wg.Done() // บอก WaitGroup ว่า goroutine นี้เสร็จสิ้นแล้ว
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("C")
		wg.Done()
	}()
	wg.Wait() // รอให้ทุก goroutine ใน WaitGroup เสร็จสิ้น
	//fmt.Println("All go routines have finished")
	fmt.Println("Total Run Time : ", time.Since(startAsyncTime))
}
