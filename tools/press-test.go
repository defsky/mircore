package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	loops, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	success, failed := 0, 0
	var mu sync.Mutex

	wg := &sync.WaitGroup{}
	wg.Add(loops)

	for i := 0; i < loops; i++ {
		go func() {
			defer wg.Done()

			conn, err := net.DialTimeout("tcp", "192.168.1.20:7000", time.Second*60)
			if err != nil {
				mu.Lock()
				failed++
				mu.Unlock()

				log.Println(err)
				return
			}
			mu.Lock()
			success++
			mu.Unlock()

			//conn.Write([]byte(fmt.Sprintf("client id %d",i)))
			conn.Close()
		}()
	}

	wg.Wait()

	log.Println("Test complete...")
	log.Printf("Success: %d Failed: %d\n", success, failed)
}
