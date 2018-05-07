package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	MAX         = 105
	EACH        = 20
	OUTPUT_PATH = "json_by_chain"
)

func main() {
	encodeToFile()
	fmt.Printf("Encode to file over \n")
	time.Sleep(time.Second)

	decodeFromFile()
}

func decodeFromFile() {
	f, err := os.Open(OUTPUT_PATH)
	if err != nil {
		fmt.Printf("Failed Create file: %+v \n", err)
		return
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	allEle := make([]int, 0, MAX)

	for {
		var chain []int

		err = dec.Decode(&chain)
		if err != nil {
			if err == io.EOF { // no more to read in file
				break
			}
			fmt.Printf("Failed Decode: %+v \n", err)
			return
		}

		allEle = append(allEle, chain...)
	}

	fmt.Printf("Total %d elements: %v \n", len(allEle), allEle)
}

func encodeToFile() {
	f, err := os.Create(OUTPUT_PATH)
	if err != nil {
		fmt.Printf("Failed Create file: %+v \n", err)
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)

	// init all elements
	allEle := make([]int, 0, MAX)
	for i := 0; i < MAX; i++ {
		allEle = append(allEle, i)
	}

	start := 0
	end := start + EACH
	for {
		chain := allEle[start:end]

		err = enc.Encode(chain)
		if err != nil {
			fmt.Printf("Failed Encode %v: %+v \n", chain, err)
			return
		}
		fmt.Printf("%d-%d encoded \n", start, end)

		// no more
		if end == MAX {
			break
		}

		// next chain
		start = start + EACH
		end = start + EACH
		if end > MAX {
			end = MAX
		}
	}
}
