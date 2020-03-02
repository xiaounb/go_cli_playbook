package main

// Module5GetExampleDotCom uses the "net/http" package to send a GET request to example.com
func Module5GetExampleDotCom() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("something went wrong")
	}

	defer resp.Body.Close()
}
