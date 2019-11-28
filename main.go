package main

import (
	"flag"
	"fmt"
	"log"
	"mircore/db"
	"mircore/game/proto"
	"mircore/mird"
	"mircore/realmd"
	"os"
	"os/signal"
	"sync"
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
	var loginServer *realmd.Realm
	var worldServer *mird.WorldServer
	var err error

	flag.IntVar(&loginPort, "loginPort", 7000, "login server port")
	flag.IntVar(&worldPort, "worldPort", 7400, "world server port")
	flag.Parse()

	fmt.Println(banner)

	db.Migrate()

	loginServer, err = realmd.NewRealmServer(loginPort, &proto.Protocol{})
	if err != nil {
		panic(err)
	}

	go func() {
		loginServer.Start()
		wg.Done()
	}()

	worldServer, err = mird.NewWorldServer(worldPort, &proto.Protocol{})
	if err != nil {
		panic(err)
	}

	go func() {
		worldServer.Start()
		wg.Done()
	}()

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
