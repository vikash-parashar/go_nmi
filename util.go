package nmi

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func LogError(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}
