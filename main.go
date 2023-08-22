package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}

	//req, err := http.NewRequest("GET", "http://128.131.132.179:8080/lpims/rest/v1/reactors?running=true", nil)
	//req, err := http.NewRequest("GET", "http://128.131.132.179:8080/lpims/rest/v1/processes?name=FMU_23_03", nil)
	//req, err := http.NewRequest("GET", "http://128.131.132.179:8080/lpims/rest/v1/processes?processId=11572", nil)
	//req, err := http.NewRequest("GET", "http://128.131.132.179:8080/lpims/rest/v1/reactors/2?currentValues=20", nil)
	req, err := http.NewRequest("GET", "http://128.131.132.179:8080/lpims/rest/v1/signals/1838523", nil)
	//calling_str = rest_url * "reactors/" * string(reactor_id) * "?currentValues=" * port_str

	if err != nil {
		fmt.Println("ERR!:", err)
	}
	req.SetBasicAuth("student", "bioVT")
	//req.Header.Add("Authorization", "Basic "+basicAuth(USERNAME, PASSWORD))
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	//req.Header.Add("Transfer-Encoding", "chunked")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	//Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP Error:", resp.Status)
	}

	// Read the HTTP response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	fmt.Println("Response Body:", string(body))
}
