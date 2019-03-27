package netverify

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Retrievaldata(scanReference, flag string) string {
	BaseURL := "https://netverify.com/api/netverify/v2/scans/"
	//assemble the url with different router via different demands, e.g. data, transaction only, .etc
	var url string

	switch{
	//if nothing input, then justify the status of the scan reference
	case flag == "":
		url = BaseURL + scanReference

	//if "data" input, then query all the data information
	case flag == "data":
		url = BaseURL + scanReference + "/" + flag

	//only want document data
	case flag == "document":
		url = BaseURL + scanReference + "/" + "data" + "/" + flag

	//only want transaction data
	case flag == "transaction":
		url = BaseURL + scanReference + "/" + "data" + "/" + flag

	//only want verification data
	case flag == "verification":
		url = BaseURL + scanReference + "/" + "data" + "/" + flag

	default:
		return fmt.Sprintf("No such corresponding fields, please input the correct flag!")
	}

	//initiate a http request obj with GET method
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	//application/json or image/jpeg, image/png for "Retrieving specific image"
	req.Header.Set("Accept", "application/json")

	//add the user agent for trouble shooting
	req.Header.Add("User-Agent", "Digital Wallet QSTOApp/v1.0")

	//note the specific Authorization code with jumio API credential, updated accordingly!
	req.Header.Add("Authorization","Basic OWJjZmFhM2QtNThkMy00MjhlLWE5ZTUtYzM3YTc4NDZjMjUwOkFLVXppVjFlNGo2WndYQ2d2SDR4d0o2dGlnUVFxc2Fi")


	//get the origin codes from the standard net/http package
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(respBody)
}


