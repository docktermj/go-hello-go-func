package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func goRoutine(ctx context.Context, number int) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	sleepTime := 1
	fmt.Printf("    goRoutine: #%d - Sleep %d\n", number, sleepTime)
	time.Sleep((time.Duration(sleepTime) * time.Second))
	fmt.Printf("    goRoutine: #%d - exit\n", number)
	return nil
}

func lockedRoutine(ctx context.Context) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	fmt.Println("lockedRoutine: start")
	for _, value := range []int{1, 2, 3, 4} {
		go goRoutine(ctx, value)
	}
	sleepTime := 7
	fmt.Printf("lockedRoutine: Sleep %d\n", sleepTime)
	time.Sleep((time.Duration(sleepTime) * time.Second))
	fmt.Println("lockedRoutine: stop")
	return nil
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()
	fmt.Println("         main: start")
	go lockedRoutine(ctx)
	sleepTime := 10
	fmt.Printf("         main: Sleep %d\n", sleepTime)
	time.Sleep((time.Duration(sleepTime) * time.Second))
	fmt.Println("         main: stop")
}
