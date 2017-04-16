package main

/*
usage : command.exe -uri http://wasabisyrup.com/archives/455582
*/

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"path/filepath"

	"github.com/jhoonb/archivex"
)

func readLinks(lines []string) []string {
	imgList := []string{}

	re := regexp.MustCompile(`<img class="lz-lazyload" src="/template/images/transparent.png" data-src="(.+)">`)
	for _, line := range lines {
		founds := re.FindStringSubmatch(line)
		if len(founds) > 0 {
			host := "http://wasabisyrup.com"
			path := founds[1]
			imgURL := host + path
			imgList = append(imgList, imgURL)
		}
	}

	return imgList
}

func readTitle(lines []string) string {
	re := regexp.MustCompile(`<div class="article-title" title="(.+)">`)
	for _, line := range lines {
		founds := re.FindStringSubmatch(line)
		if len(founds) > 0 {
			return founds[1]
		}
	}
	return ""
}

func downloadHTML(uri string) string {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		panic(err)
	}
	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	text := buf.String()
	return text
}

type ImageResponse struct {
	resp *http.Response
	idx  int
}

func downloadImage(idx int, uri string, ch chan *ImageResponse) {
	resp, _ := http.Get(uri)
	ch <- &ImageResponse{
		resp: resp,
		idx:  idx,
	}
}

func makeImageFileName(uri string, idx int) string {
	base := filepath.Base(uri)
	prefix := fmt.Sprintf("%003d-", idx+1)
	return prefix + base
}

var uri string

func init() {
	flag.StringVar(&uri, "uri", "", "marumaru uri")
}

func main() {
	flag.Parse()
	if len(uri) == 0 {
		fmt.Println("invalid uri")
		return
	}

	text := downloadHTML(uri)

	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.Trim(line, " ")
	}

	title := readTitle(lines)
	fmt.Println(title)
	zipfilename := title + ".zip"

	links := readLinks(lines)

	zip := new(archivex.ZipFile)
	zip.Create(zipfilename)

	// request download
	ch := make(chan *ImageResponse)

	for i, link := range links {
		go downloadImage(i, link, ch)
	}

	for i := 0; i < len(links); i++ {
		r := <-ch
		resp := r.resp
		idx := r.idx

		var buf bytes.Buffer
		io.Copy(&buf, resp.Body)
		resp.Body.Close()

		filename := makeImageFileName(links[idx], idx)
		zip.Add(filename, buf.Bytes())

		fmt.Printf("download (%d/%d) %s\n", i+1, len(links), filename)
	}

	zip.Close()
	fmt.Printf("%s complete\n", zipfilename)
}
