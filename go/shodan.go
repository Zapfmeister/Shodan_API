package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendHostInformation() {
	// Host Information (GET https://api.shodan.io/shodan/host/1.1.1.1?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/host/1.1.1.1?key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendSearchShodanWithoutResults() {
	// Search Shodan without Results (GET https://api.shodan.io/shodan/host/count?query=port:22&facets=org,%20os&key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/host/count?query=port:22&facets=org,%20os&key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendSearchShodan() {
	// Search Shodan (GET https://api.shodan.io/shodan/host/search?query=product:nginx&facets=org,%20os&key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/host/search?query=product:nginx&facets=org,%20os&key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendListAllSearchFacets() {
	// List all search facets (GET https://api.shodan.io/shodan/host/search/facets?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/host/search/facets?key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendListAllFiltersThatCanBeUsedWhenSearching() {
	// List all filters that can be used when searching (GET https://api.shodan.io/shodan/host/search/filters?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/host/search/filters?key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendBreakTheSearchQueryIntoTokens() {
	// Break the search query into tokens (GET https://api.shodan.io/shodan/host/search/tokens?key=your_api_key&query=product:nginx)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/host/search/tokens?key=your_api_key&query=product:nginx", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendListAllPortsThatShodanIsCrawlingOnTheInternet() {
	// List all ports that Shodan is crawling on the Internet (GET https://api.shodan.io/shodan/ports?query=product:nginx&facets=org,%20os&key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/ports?query=product:nginx&facets=org,%20os&key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendListAllProtocolsThatCanBeUsedWhenPerformingOnDemandInternetScansViaShodan() {
	// List all protocols that can be used when performing on-demand Internet scans via Shodan (GET https://api.shodan.io/shodan/protocols?query=product:nginx&facets=org,%20os&key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/protocols?query=product:nginx&facets=org,%20os&key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"bytes"
)

func sendRequestShodanToCrawlAnIpNetblockInternetScansViaShodanDuplicate() {
	// Request Shodan to crawl an IP/ netblock Internet scans via Shodan Duplicate (POST https://api.shodan.io/shodan/scan?key=your_api_key)

	params := url.Values{}
	params.Set("ips", "1.1.1.1,8.8.8.8")
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://api.shodan.io/shodan/scan?key=your_api_key", body)

	// Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"bytes"
)

func sendCrawlTheInternetForASpecificPortAndProtocolUsingShodanDuplicate2() {
	// Crawl the Internet for a specific port and protocol using Shodan Duplicate (2) (POST https://api.shodan.io/shodan/scan/internet?key=your_api_key)

	params := url.Values{}
	params.Set("port", "22")
	params.Set("protocol", "http")
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://api.shodan.io/shodan/scan/internet?key=your_api_key", body)

	// Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendGetListOfAllTheCreatedScans() {
	// Get list of all the created scans (GET https://api.shodan.io/shodan/scans?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/scans?key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendGetTheStatusOfAScanRequest() {
	// Get the status of a scan request (GET https://api.shodan.io/shodan/scan/Lxu8uAEahnlMEm9r?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/scan/Lxu8uAEahnlMEm9r?key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func sendCreateAnAlertToMonitorANetworkRange() {
	// Create an alert to monitor a network range (POST https://api.shodan.io/shodan/alert?key=your_api_key)

	json := []byte(`{"name": "DNS Alert","filters": {"ip": ["8.8.8.8","1.1.1.1"]},"expires": "0"}`)
	body := bytes.NewBuffer(json)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://api.shodan.io/shodan/alert?key=your_api_key", body)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func sendGetAListOfAllTheCreatedAlerts() {
	// Get a list of all the created alerts (GET https://api.shodan.io/shodan/alert/info?key=your_api_key)

	json := []byte(`{"name": "DNS Alert","filters": {"ip": ["8.8.8.8","1.1.1.1"]},"expires": "0"}`)
	body := bytes.NewBuffer(json)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/alert/info?key=your_api_key", body)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendGetTheDetailsForANetworkAlert() {
	// Get the details for a network alert (GET https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/info?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/info?key=your_api_key", nil)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendDeleteAnAlert() {
	// Delete an alert (DELETE https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3?key=your_api_key", nil)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func sendEditTheNetworksMonitoredInAnAlert() {
	// Edit the networks monitored in an alert (POST https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3?key=your_api_key)

	json := []byte(`{"filters": {"ip": "1.1.1.1"}}`)
	body := bytes.NewBuffer(json)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3?key=your_api_key", body)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendGetAListOfAvailableTriggers() {
	// Get a list of available triggers (GET https://api.shodan.io/shodan/alert/triggers?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/shodan/alert/triggers?key=your_api_key", nil)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendEnableATrigger() {
	// Enable a trigger (PUT https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("PUT", "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware?key=your_api_key", nil)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendDisableATrigger() {
	// Disable a trigger (PUT https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("PUT", "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware?key=your_api_key", nil)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendAddToWhitelist() {
	// Add to Whitelist (PUT https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware/ignore/1.1.1.1:53?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("PUT", "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware/ignore/1.1.1.1:53?key=your_api_key", nil)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRemoveFromWhitelist() {
	// Remove from Whitelist (DELETE https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware/ignore/1.1.1.1:53?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware/ignore/1.1.1.1:53?key=your_api_key", nil)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendAddTheNotifierToTheAlert() {
	// Add the notifier to the alert (PUT https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/notifier/default?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("PUT", "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/notifier/default?key=your_api_key", nil)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRemoveTheNotifierFromTheAlert() {
	// Remove the notifier from the alert (DELETE https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/notifier/default?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/notifier/default?key=your_api_key", nil)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendListAllUserCreatedNotifiers() {
	// List all user-created notifiers (GET https://api.shodan.io/notifier?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/notifier?key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendListOfAvailableNotificationProviders() {
	// List of available notification providers (GET https://api.shodan.io/notifier/provider?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/notifier/provider?key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"bytes"
)

func sendCreateANewNotificationServiceForTheUser() {
	// Create a new notification service for the user (POST https://api.shodan.io/notifier?key=your_api_key)

	params := url.Values{}
	params.Set("provider", "email")
	params.Set("to", "test@google.de")
	params.Set("description", "\"email notification\"")
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://api.shodan.io/notifier?key=your_api_key", body)

	// Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"bytes"
)

func sendDeleteANotificationService() {
	// Delete a notification service (DELETE https://api.shodan.io/notifier/ZKJpffUnarJushBf?key=your_api_key)

	params := url.Values{}
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", "https://api.shodan.io/notifier/ZKJpffUnarJushBf?key=your_api_key", body)

	// Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"bytes"
)

func sendGetInformationAboutANotifier() {
	// Get information about a notifier (GET https://api.shodan.io/notifier/ZKJpffUnarJushBf?key=your_api_key)

	params := url.Values{}
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/notifier/ZKJpffUnarJushBf?key=your_api_key", body)

	// Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"bytes"
)

func sendEditANotifier() {
	// Edit a notifier (GET https://api.shodan.io/notifier/ZKJpffUnarJushBf?key=your_api_key)

	params := url.Values{}
	params.Set("to", "test@google.de")
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/notifier/ZKJpffUnarJushBf?key=your_api_key", body)

	// Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendDomainInformation() {
	// Domain Information (GET https://api.shodan.io/dns/domain/google.de?key=your_api_key&history=False&page=1)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/dns/domain/google.de?key=your_api_key&history=False&page=1", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendDnsLookup() {
	// DNS Lookup (GET https://api.shodan.io/dns/resolve?key=your_api_key&hostnames=google.com,bing.com)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/dns/resolve?key=your_api_key&hostnames=google.com,bing.com", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendReverseDnsLookup() {
	// Reverse DNS Lookup (GET https://api.shodan.io/dns/reverse?key=your_api_key&ips=1.1.1.1,8.8.8.8)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/dns/reverse?key=your_api_key&ips=1.1.1.1,8.8.8.8", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendHttpHeaders() {
	// HTTP Headers (GET https://api.shodan.io/tools/httpheaders?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/tools/httpheaders?key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendMyIpAddress() {
	// My IP Address (GET https://api.shodan.io/tools/myip?key=your_api_key)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://api.shodan.io/tools/myip?key=your_api_key", nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
	  fmt.Println(parseFormErr)    
	}

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


