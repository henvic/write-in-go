package main

import "net/url"
import "os"

func getHostFromURL(uri string) (host string, err error) {
	var u *url.URL
	if u, err = url.Parse(uri); err != nil {
		return "", err
	}

	return u.Host, nil
}

func main() {
	if _, err := getHostFromURL("https://www.flickr.com/"); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
