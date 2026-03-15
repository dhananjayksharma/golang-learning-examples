package main

import (
	"bytes"
	addresses "dkgosql-interface-gc/Addresses"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// Pool of reusable byte buffers
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func processRequest(data []byte) string {
	// Get a buffer from pool (or allocate new if pool empty)
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()            // ALWAYS reset before use!
	defer bufPool.Put(buf) // Return to pool when done

	buf.Write(data)
	buf.WriteString(" [processed]")
	return buf.String()
}

// Without pool: each call allocates a new bytes.Buffer on heap
// With pool:    objects are reused → ~0 allocations in steady state

func main() {
	addresser := addresses.NewZip("560035")

	address := addresser.Add("G005")
	fmt.Println("address:", address)
	addressByte, err := json.Marshal(address)
	if err != nil {
		log.Printf("Error:%v", err)
	}

	out := processRequest(addressByte)

	fmt.Printf("Address: %s\n", out)
}
