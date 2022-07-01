package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	wordlist := os.Args[2]

	fmt.Printf("Url: %s\n", url)
	fmt.Printf("Wordlist: %s\n\n", wordlist)

	// readFile, err := os.Open("/usr/share/wordlists/awesome-wordlists/dirbuster/directory-list-lowercase-2.3-medium.txt")
	readFile, err := os.Open(wordlist)
	if err != nil {
		fmt.Print(err)
	}

	// counter := 1
	scanner := bufio.NewScanner(readFile)

	for scanner.Scan() {
		line := scanner.Text()

		// Checking if url has / in the end
		var newUrl string
		if string(url[len(url)-1]) == "/" {
			newUrl = url + line
			// fmt.Printf("%s\n", newUrl)
		} else {
			newUrl = url + "/" + line
		}

		// Sending Get Request to see if the endpoint exists or not
		resp, err := http.Get(newUrl)
		if err != nil {
			fmt.Print(err)
			// continue
		}
		if resp.StatusCode == 200 || resp.StatusCode == 403 || resp.StatusCode == 301 || resp.StatusCode == 302 {
			fmt.Println(newUrl, resp.StatusCode)
		}

		// counter++

	}
	readFile.Close()

}
