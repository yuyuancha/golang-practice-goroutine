package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

func main() {
	fmt.Println("程式執行開始...")
	fmt.Printf("現在時間為: %s\n", time.Now().Format("2006-01-02 03:04:05"))

	start := time.Now().Unix()
	count := 10

	wg := new(sync.WaitGroup)
	wg.Add(count)

	for i := 1; i <= count; i++ {
		go wait(i, wg)
	}

	wg.Wait()

	end := time.Now().Unix()

	fmt.Println("程式執行結束...")
	fmt.Printf("現在時間為: %s\n", time.Now().Format("2006-01-02 03:04:05"))
	fmt.Printf("程式總共花費 %d 秒", end-start)
}

func wait(round int, wg *sync.WaitGroup) {
	defer func() {
		fmt.Printf("[%s] Round %d is done.\n", time.Now().Format("2006-01-02 03:04:05"), round)
		wg.Done()
	}()

	// 隨機等待 0-2 秒
	seconds, _ := rand.Int(rand.Reader, big.NewInt(3))
	time.Sleep(time.Duration(int(seconds.Int64())) * time.Second)
}
