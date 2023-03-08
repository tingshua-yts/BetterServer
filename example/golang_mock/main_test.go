package main

import (
	"fmt"
	"testing"
)

func get_page_mock(url string) {
	fmt.Println("in test get page")
}
func TestSum(t *testing.T) {
	d := Downloader{}
	d.get = get_page_mock
	d.download()
}
