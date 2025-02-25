package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {

	// create our WaitGroup and Mutex objects
	var wg sync.WaitGroup
	var mu sync.Mutex

	numberOfPoints := 0
	numberOfRoutines := 0

	for {
		fmt.Println("----------Monte Carlo using Concurrency----------")
		fmt.Println("Enter number of points: ")
		fmt.Scanf("%d", &numberOfPoints)
		fmt.Println("Enter number of rountines to use: ")
		fmt.Scanf("%d", &numberOfRoutines)
		insideCircle := 0
		start := time.Now()

		//run the process with the given number of routines and points
		for i := 0; i < numberOfRoutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				pointsInCircle := 0
				for j := 0; j < numberOfPoints/numberOfRoutines; j++ {
					x := rand.Float64()
					y := rand.Float64()

					if (x*x + y*y) <= 1 {
						pointsInCircle++
					}
				}
				mu.Lock()
				insideCircle += pointsInCircle
				mu.Unlock()
			}()
		}

		wg.Wait()

		estimatedPi := 4.0 * float64(insideCircle) / float64(numberOfPoints)
		end := time.Now()

		elapsed := float64(end.Sub(start).Seconds())
		difference := math.Abs(estimatedPi - math.Pi)

		fmt.Printf("Estimated Pi: %.5f \n", estimatedPi)
		fmt.Printf("Delta between estimated and real Pi: %.5f \n", difference)
		fmt.Printf("Time elapsed: %.5f s\n", elapsed)
		fmt.Println("Try again? Y or N")
		var choice string
		fmt.Scanln(&choice)

		if choice == "n" || choice == "N" {
			fmt.Println("Quitting.....")
			break
		}

	}
}
