package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestWorker(t *testing.T) {
	res := "{\"TaskId\":\"a5b02f55d441488e80611bebbfe5a4b5\",\"RequestId\":\"81D5CB45-E9A1-5775-826F-543412B3CFC1\",\"StatusText\":\"SUCCESS\",\"BizDuration\":3101,\"SolveTime\":167584,\"RequestTime\":1676682639710,\"StatusCode\":21050000,\"Result\":{\"Sentences\":[{\"EndTime\":2510,\"SilenceDuration\":0,\"BeginTime\":880,\"Text\":\"北京的天气。\",\"ChannelId\":0,\"Spte\":184,\"EmotionValue\":6.7}]}}"
	fmt.Println(res)
	var v map[string]interface{}
	err := json.Unmarshal([]byte(res), &v)
	if err != nil {
		return
	}
	fmt.Println(v["RequestTime"])
	fmt.Println(v["SolveTime"])
	fmt.Printf("%T", v["Result"])
	file, err := os.Create("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(file, "Hello World")
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}
