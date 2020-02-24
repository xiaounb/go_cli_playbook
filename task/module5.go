package main

// Module5GetExampleDotCom usese th "net/http" package to send a GET request to example.com
func Module5GetExampleDotCom() {
	resp, err := http.Get("http://example.com/")

}
