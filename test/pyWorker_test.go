package test

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestName(t *testing.T) {
	fileLink := "https://gw.alipayobjects.com/os/bmw-prod/0574ee2e-f494-45a5-820f-63aee583045a.wav"
	outFilePath := " ./tmp/abc.json"
	cmd := fmt.Sprintf("python3 ../pyWorker/work.py %s %s", fileLink, outFilePath)
	fmt.Printf(cmd)
	output, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("Output: %s\n", output)
}
