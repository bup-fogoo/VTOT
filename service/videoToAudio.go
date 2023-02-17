package service

import (
	"VTOT/utils"
	"fmt"
	"os/exec"
)

func VideoToAudioService(inputFile string) string {
	outputFile := fmt.Sprintf("%s.mp3", utils.RandSeq(8))
	cmd := fmt.Sprintf("ffmpeg -i %s -vn -c:a mp3 -b:a 192k ./tmp/%s", inputFile, outputFile)
	_, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	//fmt.Printf("Output: %s\n", output)
	return outputFile
}
