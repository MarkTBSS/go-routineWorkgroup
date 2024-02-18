package main

import (
	"fmt"
	"sync"
)

func SayB(wg *sync.WaitGroup) {
	fmt.Println("Hello, world : B")
	wg.Done()
}

func main() {
	// Synchronous or Sequential
	fmt.Println("Hello, world : 1")
	fmt.Println("Hello, world : 2")
	fmt.Println("Hello, world : 3")
	fmt.Println("Hello, world : 4")
	fmt.Println("Hello, world : 5")
	fmt.Println("================")
	// Asynchronous
	var wg sync.WaitGroup // สร้างตัวแปร WaitGroup
	wg.Add(4)             // เพิ่มจำนวน goroutine ที่ต้องรอใน WaitGroup
	// ===== แบบนี้ไม่ได้เพราะอยู่นอก Scope =====
	// go fmt.Println("Hello, world : A")
	// wg.Done()
	// ====================================
	go func() {
		fmt.Println("Hello, world : A")
		wg.Done() // บอก WaitGroup ว่า goroutine นี้เสร็จสิ้นแล้ว
	}()

	go SayB(&wg) // ส่งทางการอ้างอิง *sync.WaitGroup

	go func() {
		fmt.Println("Hello, world : C")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello, world : D")
		wg.Done()
	}()

	wg.Wait() // รอให้ทุก goroutine ใน WaitGroup เสร็จสิ้น
	fmt.Println("All go routines have finished")
}
