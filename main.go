package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const (
	DefaultCount                = 20
	DefaultSleep                = 1
	DefaultPidFile              = "./go-sandbox.pid"
	DefaultFileMode os.FileMode = 0644
)

var (
	count   int
	sleep   time.Duration
	pidFile string
)

func main() {
	flag.IntVar(&count, "count", DefaultCount, "loop count")
	flag.DurationVar(&sleep, "sleep", DefaultSleep, "sleep second")
	flag.StringVar(&pidFile, "pid-file", DefaultPidFile, "pid file path")
	flag.Parse()

	fmt.Println("loop started.")

	pid := os.Getpid()
	pidStr := strconv.Itoa(pid)
	err := ioutil.WriteFile(pidFile, []byte(pidStr), DefaultFileMode)
	if err != nil {
		fmt.Println(fmt.Sprintf("write pid file failed.\n%s", err.Error()))
	}

	for i := 0; i < count; i++ {
		fmt.Println(fmt.Sprintf("loop: %d, left: %d", i+1, count-i-1))
		time.Sleep(sleep * time.Second)
	}

	fmt.Println("loop completed.")
}
