package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	flag "github.com/spf13/pflag"
)

func main() {
	max_loop := flag.IntP("max_loop", "m", 8500, "sampling size of monte carlo")
	flag.Parse()

	quadrant := 0
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println("Approximating pi using sample size", *max_loop)
	for i := 0; i < *max_loop; i++ {
		x := r1.Float64()
		y := r1.Float64()
		d := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
		if d <= 1.0 {
			quadrant += 1
		}
	}
	estimate_pi := float64(quadrant) / float64(*max_loop) * 4.0
	fmt.Println("Approximating pi is:")
	fmt.Println(estimate_pi)
}
