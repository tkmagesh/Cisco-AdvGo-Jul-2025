package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	dataCh := make(chan int)

	/*
		wg.Add(1)
		go Source(wg, "data1.dat", dataCh)

		wg.Add(1)
		go Source(wg, "data2.dat", dataCh)

		wg.Add(1)
		go Source(wg, "data3.dat", dataCh)
	*/
	fileNames := []string{"data1.dat", "data2.dat", "data3.dat"}
	for _, fileName := range fileNames {
		wg.Add(1)
		go Source(wg, fileName, dataCh)
	}

	oddCh, evenCh := Splitter(dataCh)
	oddSumCh := Sum(oddCh)
	evenSumCh := Sum(evenCh)

	doneCh := Merger(oddSumCh, evenSumCh, "result.txt")

	wg.Wait()
	close(dataCh)

	<-doneCh
	fmt.Println("Done!")
}

func Source(wg *sync.WaitGroup, fileName string, ch chan int) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if no, err := strconv.Atoi(line); err == nil {
			ch <- no
		}
	}
}

func Splitter(dataCh chan int) (<-chan int, <-chan int) {
	oddCh := make(chan int)
	evenCh := make(chan int)
	go func() {
		for no := range dataCh {
			if no%2 == 0 {
				evenCh <- no
			} else {
				oddCh <- no
			}
		}
		log.Println("[Splitter] - all data from dataCh is processed")
		close(evenCh)
		close(oddCh)
	}()
	return oddCh, evenCh

}

func Sum(noCh <-chan int) <-chan int {
	sumCh := make(chan int)
	go func() {
		total := 0
		for no := range noCh {
			total += no
		}
		sumCh <- total
	}()
	return sumCh
}

func Merger(oddSumCh <-chan int, evenSumCh <-chan int, resultFile string) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		file, err := os.Create(resultFile)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		for range 2 {
			select {
			case evenSum := <-evenSumCh:
				fmt.Fprintf(file, "Even Total : %d\n", evenSum)
			case oddSum := <-oddSumCh:
				fmt.Fprintf(file, "Odd Total : %d\n", oddSum)
			}
		}
		close(doneCh)
	}()
	return doneCh
}
