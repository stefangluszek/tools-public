package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: unixtime <unixtimestamp>")
	}

	timestamp, err := strconv.ParseInt(os.Args[1], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	t := time.Unix(timestamp, 0)
	fmt.Println(t.Format(time.UnixDate))
}
