package main

import (
    "os"
    "log"
    "net"
    "sync"
    "time"
    "strconv"
)

func main(){
    loops, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Fatalln(err)
    }
    
    success,failed := 0,0
    
    wg := &sync.WaitGroup{}
    for i:=0;i<loops;i++ {
        go func(){
            wg.Add(1)
            defer wg.Done()
            
            conn, err := net.DialTimeout("tcp", "login.afkplayer.com:7000", time.Second * 60)
            if err != nil {
                failed ++
                log.Println(err)
                return
            }
            success ++
            conn.Close()
        }()
        
    }
    
    wg.Wait()
    log.Println("Test complete...")
    log.Printf("Success: %d Failed: %d\n", success, failed)
}
