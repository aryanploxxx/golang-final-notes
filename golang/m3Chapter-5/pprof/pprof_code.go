package pprof

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
)

var profile = flag.String("cpuprofile", "", "output pprof data to file")

// go run pprof.go -cpuprofile=cpu.prof
// go tool pprof cpu.prof

func main() {

	flag.Parse()
	if *profile != "" {
		flag, err := os.Create(*profile)
		if err != nil {
			fmt.Println("Could not create profile", err)
		}
		pprof.StartCPUProfile(flag)
		defer pprof.StopCPUProfile()

	}

	// add rest of the code here

}
