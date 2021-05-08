package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	flag "github.com/spf13/pflag"
	"github.come/gulien/fizz-buzz/pkg/server"
)

var version = "snapshot"

func main() {
	fmt.Printf("[SYSTEM] version %s\n", version)

	fs := flag.NewFlagSet("fizzbuzz", flag.ExitOnError)
	fs.Int("port", 80, "Set the port on which the fizz-buzz server should listen")
	fs.Int("timeout", 30, "Set the maximum duration in seconds before timing out execution of a request")

	// Parses the flags...
	err := fs.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// and gets their values.
	port, err := fs.GetInt("port")
	if err != nil {
		fmt.Printf("[FATAL] %s\n", err)
		os.Exit(1)
	}

	timeout, err := fs.GetInt("timeout")
	if err != nil {
		fmt.Printf("[FATAL] %s\n", err)
		os.Exit(1)
	}

	srv := server.New(time.Duration(timeout) * time.Millisecond)

	go func() {
		err := srv.Start(fmt.Sprintf(":%d", port))

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("[FATAL] %s\n", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C) SIGKILL,
	// SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(quit, os.Interrupt)

	// Block until we receive our signal.
	<-quit

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the
	// timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("[FATAL] %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
