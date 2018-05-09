package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin/json"
)

func main() {
	byteResult, err := json.Marshal(struct {
		ExtData map[string]interface{} `json:"ext_data"`
	}{
		ExtData: map[string]interface{}{
			"clickId": "8888",
		},
	})
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println(string(byteResult))
}
