package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

const (
	minTemperatureValue = 40
	maxTemperatureValue = 50
)

var (
	minTemperatureName = fmt.Sprintf("Handler <= %dº", minTemperatureValue)
	maxTemperatureName = fmt.Sprintf("Handler  > %dº", maxTemperatureValue)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	temperatureChan := make(chan int)
	minTemperatureChan := make(chan int)
	maxTemperatureChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go generateTemperatures(ctx, temperatureChan)
	go handleTemperature(ctx, minTemperatureName, minTemperatureChan, &wg)
	go handleTemperature(ctx, maxTemperatureName, maxTemperatureChan, &wg)
	go receiveTemperatures(ctx, temperatureChan, minTemperatureChan, maxTemperatureChan)

	wg.Wait()
}

func generateTemperatures(ctx context.Context, temperatureChan chan int) {
	defer close(temperatureChan)
	for {
		select {
		case <-ctx.Done():
			return
		case temperatureChan <- rand.Int() % 100:
		}
	}
}

func handleTemperature(ctx context.Context, handlerName string, temperatureChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	const skip = 10

	for {
		select {
		case <-ctx.Done():
			return
		case _, ok := <-temperatureChan:
			if !ok {
				return
			}
			count++
			if count%skip == 0 {
				fmt.Printf("%s received %d temperatures\n", handlerName, skip)
			}
		}
	}
}

func receiveTemperatures(ctx context.Context, temperatureChan, minTemperatureChan, maxTemperatureChan chan int) {
	defer close(minTemperatureChan)
	defer close(maxTemperatureChan)

	for {
		select {
		case <-ctx.Done():
			return
		case temperature, ok := <-temperatureChan:
			if !ok {
				return
			}
			if temperature <= minTemperatureValue {
				minTemperatureChan <- temperature
			} else if temperature > maxTemperatureValue {
				maxTemperatureChan <- temperature
			}
		}
	}
}

//var minTemperatureValue = 40
//var maxTemperatureValue = 50
//var minTemperatureName = fmt.Sprintf("Handler <= %dº", minTemperatureValue)
//var maxTemperatureName = fmt.Sprintf("Handler  > %dº", maxTemperatureValue)
//
//func main() {
//	temperaturesChan := make(chan int)
//	minTemperatureHandlerChan := make(chan int)
//	maxTemperatureHandlerChan := make(chan int)
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	go GenerateTemperatures(temperaturesChan)
//	go HandleTemperature(minTemperatureName, minTemperatureHandlerChan, &wg)
//	go HandleTemperature(maxTemperatureName, maxTemperatureHandlerChan, &wg)
//	go ReceiveTemperatures(temperaturesChan, minTemperatureHandlerChan, maxTemperatureHandlerChan)
//
//	wg.Wait()
//}
//
//func GenerateTemperatures(temperaturesChan chan int) {
//	for i := 0; i < 10_000_000; i++ {
//		temperature := rand.Int() % 100
//		temperaturesChan <- temperature
//	}
//	close(temperaturesChan)
//}
//
//func HandleTemperature(handlerName string, temperaturesChan chan int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	count := 0
//	skip := 10
//
//	for range temperaturesChan {
//		count++
//		if count%skip == 0 {
//			fmt.Printf("%s received %d temperatures\n", handlerName, skip)
//		}
//	}
//}
//
//func ReceiveTemperatures(temperaturesChan chan int, minTemperatureHandlerChan chan int, maxTemperatureHandlerChan chan int) {
//	for temperature := range temperaturesChan {
//		if temperature <= minTemperatureValue {
//			minTemperatureHandlerChan <- temperature
//		} else if temperature > maxTemperatureValue {
//			maxTemperatureHandlerChan <- temperature
//		}
//	}
//	close(minTemperatureHandlerChan)
//	close(maxTemperatureHandlerChan)
//}
