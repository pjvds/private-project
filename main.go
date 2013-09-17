package main

import (
	"github.com/pjvds/private-package"
)

func main() {
	for counter.Value() <= 10 {
		println(counter.Next())
	}
}
