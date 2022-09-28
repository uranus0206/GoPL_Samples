// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//<- Ex 1.8
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		//-> End of Ex 1.8

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

		//<- Ex 1.9
		fmt.Printf("Response: %s\n\n", resp.Status)
		//-> Endof Ex 1.9

		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)

		//<- Ex1.7
		n, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\nCopy: %d", n)
		//-> Endof Ex1.7

	}
}
