package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2) // Izinkan 2 CPU core

	messages := make(chan int, 3) // Buffered channel dengan kapasitas 3

	// Receiver goroutine
	go func() {
		time.Sleep(1 * time.Second) // Delay agar pengirim lebih dulu jalan
		for {
			val := <-messages
			fmt.Println("🟢 RECEIVED:", val)
			time.Sleep(500 * time.Millisecond) // Simulasi lambat
		}
	}()

	// Sender (utama)
	for i := 0; i < 5; i++ {
		fmt.Println("📤 SENDING :", i)
		messages <- i // Akan block jika buffer penuh
		fmt.Println("✅ SENT    :", i)
	}

	// Biarkan goroutine receiver selesai
	time.Sleep(5 * time.Second)
}
