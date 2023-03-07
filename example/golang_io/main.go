package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Usage struct {
	cpu int
	gpu int
}

func main() {

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()
	usages := make([]Usage, 0)
	usages = append(usages, Usage{cpu: 1, gpu: 2})
	usages = append(usages, Usage{cpu: 2, gpu: 3})

	for _, usage := range usages {
		str := fmt.Sprintf("result: %v", usage)
		_, err = f.WriteString(str)
		check(err)
	}
	f.Sync()

}
