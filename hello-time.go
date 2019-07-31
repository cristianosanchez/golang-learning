package main

import (
	"fmt"
	"time"
)

func howTimeWorks() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())

	now := time.Now()
	fmt.Printf("and now is %s\n", now.Local())
}
