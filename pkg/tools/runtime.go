package tools

import (
	"log"
	"runtime"
	"time"
)

//看goroutine数量可以在main里用，仅针对开发环境
func System() {
	mem := &runtime.MemStats{}

	for {
		cpu := runtime.NumCPU()
		log.Println("CPU:", cpu)

		rot := runtime.NumGoroutine()
		log.Println("Goroutine:", rot)

		// Byte
		runtime.ReadMemStats(mem)
		log.Println("Memory:", mem.Alloc)

		time.Sleep(2 * time.Second)
		log.Println("-------")
	}
}
