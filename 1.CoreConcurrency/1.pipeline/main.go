package main

/*
┌───────────┐    ┌───────────┐    ┌───────────┐    ┌───────────┐
│  Source   │ →  │ Stage 1   │ →  │ Stage 2   │ →  │  Sink     │
└───────────┘    └───────────┘    └───────────┘    └───────────┘
   Data Gen        Process-1        Process-2         Output

- Implementation
Source : Generate function would generate the numbers and write to the next stage
Stage 1: Process function would process the input and square this and write to the next stage
Stage 2: Is again a process function say multiply this by const C and write the output to next stage
Sink: Would render the output to display or a file etc.
*/

import (
	"fmt"
	"sync"
	"time"
)

const (
	ConstMultiplier = 2
)

// source generating numbers
func generate(wg *sync.WaitGroup) <-chan int {
	ch := make(chan int) //Creating unbuffered channel
	wg.Add(1)
	go func() { // goroutine with a closure

		defer wg.Done() //signaling wait group, that go routine is exiting

		defer close(ch) // closing channel as we are done with generating nymber
		// It is important to close channel,
		// otherwise it will lead to memory leak
		// goroutine generating numbers from 0 - 999 and writing this to channel
		for i := range 10 {
			ch <- i
		}
	}()
	return ch
}

func process_stage1(wg *sync.WaitGroup, data <-chan int) <-chan int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for input := range data {
			ch <- input * input
		}
	}()
	return ch
}

func process_stage2(wg *sync.WaitGroup, data <-chan int) <-chan int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for input := range data {
			time.Sleep(10 * time.Millisecond)
			ch <- input * ConstMultiplier
		}
	}()
	return ch
}

func sink(wg *sync.WaitGroup, data <-chan int) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for input := range data {
			fmt.Println("Processed output:", input)
		}
	}()
}

/**
* Pipeline Orachastration type A
**/
func orchestratePipeline(wg *sync.WaitGroup) {
	source := generate(wg) //Creating source

	ps1 := process_stage1(wg, source) // Creatining stage1

	ps2 := process_stage2(wg, ps1) //Creating stage2

	sink(wg, ps2) //wrting to synk
}

/*
Orchestration type 2
*/

func orchestratePipelineType2(wg *sync.WaitGroup) {
	sink(wg, process_stage1(wg, generate(wg)))
}

func main() {
	fmt.Println("Demo:Pipeline...")
	wg := new(sync.WaitGroup)
	orchestratePipeline(wg)      // type 1
	orchestratePipelineType2(wg) //type2
	wg.Wait()
}
