package main

import (
	"fmt"
	"log"

	"github.com/penril0326/myMessagePack/internal/encoder"
)

func main() {
	// JSON to MessagePack
	json := map[string]interface{}{
		"compact": true,
		"schema":  0,
	}

	b, err := encoder.JsonToMsgPack(json)
	if err != nil {
		log.Fatalf("Convert JSON to msgpack occured error: %s", err.Error())
	}

	fmt.Printf("%x\n", b)
}
