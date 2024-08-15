package goroutiness

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func print(till int, message string, startTime time.Time, wege *sync.WaitGroup) {
	for i := 0; i < till; i++ {
		fmt.Println(i+1, message)
		since := time.Since(startTime)
		log.Log().Float64("processTime ", float64(since.Nanoseconds())).Msg("Success print " + message)

	}
	wege.Done()
}

func StartGorou() {
	runtime.GOMAXPROCS(2)
	wg := sync.WaitGroup{}
	now := time.Now()

	wg.Add(1)
	go print(10, "anjay", now, &wg)
	print(10, "anjay lah anjay", now, &wg)
	wg.Wait()
}
