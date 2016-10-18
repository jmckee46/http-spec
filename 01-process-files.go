package main

import (
	"fmt"
	"sync"
)

func processFiles(https bool, hostname string, pathnames []string) {
	if logFunctions {
		fmt.Println("01 processFiles")
	}

	if logContext {
		fmt.Printf("https: %v, hostname: %v", https, hostname)
	}

	var uriScheme string

	if https {
		uriScheme = "https://"
	} else {
		uriScheme = "http://"
	}

	waitGroup := &sync.WaitGroup{}
	defer waitGroup.Wait()

	for _, pathname := range pathnames {
		waitGroup.Add(1)

		context := context{
			Pathname:  pathname,
			HTTPS:     https,
			HostName:  hostname,
			URIScheme: uriScheme,
			WaitGroup: waitGroup,
		}

		go processFile(context)
	}
}
