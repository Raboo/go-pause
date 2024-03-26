package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // Create a channel to receive OS signals
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt, os.Kill)

    // Print a log message indicating that the pause container has started
    log.Println("Pause container started.")

    // Block until a termination signal is received
    sig := <-signals

    // Print a log message indicating the received signal
    log.Printf("Received signal: %v. Exiting.", sig)
}

