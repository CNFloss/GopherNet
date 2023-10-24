package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/CNFloss/GopherNet/api/data"
	"github.com/CNFloss/GopherNet/api/handlers"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	
	if *debug {
		fmt.Println("\n   @@@@@@@@@@@@@@@@@@\n       Debug mode\n   @@@@@@@@@@@@@@@@@@\n ")
	}

	fmt.Println(">>>>>> Gopher Net <<<<<<")
	fmt.Println("\nlistening on port:", ":8080")

	l := log.New(os.Stdout, "Gopher Net ", log.LstdFlags)

	usersCache := data.NewCache()
	err := usersCache.Init("users.json", &data.User{})
	if err != nil {
		l.Println("cache intialization failed")
		l.Println(err)
	}

	hh := handlers.NewHello(l)
	uh := handlers.NewUsers(usersCache,l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/user", uh)

	s := &http.Server{
		Addr: ":8080",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}