# myMessagePack
Trying to implement a converter that can convert from JSON to MessagePack and from MessagePack to JSON.

# Support type
### Encode
`bool`, `int` family, `uint` family, `string`, `float32`, `float64`, map, array, slice, timestamp

### Decode
`bool`, `uint` family,`int` family,  `string`, `float32`, `float64`

# How to use

```Go
func main() {
    // Declare a map as a JSON object
    m := map[string]interface{}{
        "compact": true,
        "schema" 0,
    }

    b, err := encoder.JsonToMsgPack(json)
	if err != nil {
		log.Fatalf("Convert JSON to msgpack occured error: %s", err.Error())
	}

	fmt.Printf("%x\n", b) // output: 82a7c3a600
}

```

# Run test
### Test encoding
```Shell
$ cd internal/encoder
$ go test -v .
```

### Test decoding
```Shell
$ cd internal/dncoder
$ go test -v .
```