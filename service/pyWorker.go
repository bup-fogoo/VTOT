package service

import (
	"fmt"
	"os/exec"
	"sync"
)

func Worker(fileLink string, outFilePath string, ch chan struct{}, wg *sync.WaitGroup) {
	// 等待 channel 可用
	ch <- struct{}{}

	defer wg.Done()

	// worker
	cmd := fmt.Sprintf("python3 ./pyWorker/work.py %s %s", fileLink, outFilePath)
	fmt.Printf(cmd)
	output, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("Output: %s\n", output)

	// 从channel 中取出信号
	<-ch
}
