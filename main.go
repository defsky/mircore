package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"mircore/realm"
	"mircore/world"
)

const banner = `
                                  ###
                                 ####                #####
              ###        ####    ###               #######
              ###       ####                      #### ###
             ####      #####                    ####   ###
             #####    #####         ###  ####  ####   ####   #####    ###  ####   #####
            ######   ######   ###  ### ###### ####    ###   #######  ### ######  ######
            ######  ######   ####  #######   ####         ####  ###  #######    ## ###
           ####### ### ###   ###  ######    ####          ###   ### ######     ## ####
          ########### ###   ###   #####     ###          ###    ### #####     ######
          ### ######  ###   ###  #####     ###          ###    ### #####     #####
         ###   ###    ###  ###   ####      ###          ###   #### ####     ####
         ###         ###   ###   ####     ###          ###   ####  ####     ###    ###
        ###          ###  ###   ####      ###     ###  ###  ####  ####      ###  ####
        ###          ###  ###   ###       ####  ####   #######    ###       ########
       ###                ###   ###       #########     #####     ###       ######
                                           ######
`

func main() {
	var loginPort int
	var worldPort int
	var wg sync.WaitGroup
	var loginServer *realm.LoginServer
	var worldServer *world.WorldServer
	var err error

	flag.IntVar(&loginPort, "loginPort", 7000, "login server port")
	flag.IntVar(&worldPort, "worldPort", 7400, "world server port")
	flag.Parse()

	fmt.Println(banner)

	if loginServer, err = realm.NewLoginServer(loginPort); err != nil {
		panic(err)
	} else {
		go func() {
			loginServer.Start()
			wg.Done()
		}()
	}

	if worldServer, err = world.NewWorldServer(worldPort); err != nil {
		panic(err)
	} else {
		go func() {
			worldServer.Start()
			wg.Done()
		}()
	}

	c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, os.Interrupt, os.Kill)
		_ = <-c
		log.Println("Received Interrupt signal, shutdown servers ...")

		loginServer.Stop()
		worldServer.Stop()

		wg.Done()
	}()

	wg.Add(3)
	wg.Wait()
}
