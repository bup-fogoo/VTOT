package test

import (
	"VTOT/utils"
	"fmt"
	"os/exec"
	"testing"
)

func TestVideoToAudioService(t *testing.T) {
	path := "../tmp/"
	inputFile := path + "atest.MP4"
	fmt.Printf(inputFile)
	outputFile := fmt.Sprintf(path + utils.RandSeq(8) + ".mp3")
	cmd := fmt.Sprintf("ffmpeg -i %s -vn -c:a mp3 -b:a 192k %s", inputFile, outputFile)
	print(cmd)
	output, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("Output: %s\n", output)
}
