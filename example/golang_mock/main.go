package main

import "fmt"

type PageGetter func(url string)

type Downloader struct {
	get PageGetter
}

func NewDownloader(pg PageGetter) *Downloader {
	return &Downloader{get: pg}
}

func (d *Downloader) download() {
	//...
	fmt.Println("in download")
	d.get("aa")
	//...
}

func get_page(url string) {
	fmt.Println("in main get page")
}

func main() {
	d := Downloader{}
	d.get = get_page
	d.download()
}
