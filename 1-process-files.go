package main

import "fmt"

func processFiles(https bool, hostname string, pathnames []string) {
	if logFunctions {
		fmt.Println("1 processFiles")
	}

	if logContext {
		fmt.Printf("https: %v, hostname: %v", https, hostname)
	}

	var uriScheme string

	if https {
		uriScheme = "https://" + hostname
	} else {
		uriScheme = "http://" + hostname
	}

	for _, pathname := range pathnames {
		context := &context{
			Pathname:  pathname,
			URIScheme: uriScheme,
		}

		processFile(context)
	}
}
