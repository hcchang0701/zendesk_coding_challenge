package fsm

import (
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
)

var (
	client *http.Client
	baseURL *url.URL
	username string
	password string
	tickets []*Ticket
	meta    *Meta
	err error
)

func init() {
	client = &http.Client{}
	baseURL, _ = url.Parse("https://zcchcc.zendesk.com")
	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
}

func list() {

	tickets, meta, err = getTicketsWithCursor("", "")
	if err != nil {
		fmt.Println("Faile to get tickets")
		os.Exit(1)
	}

	ticketString := fmt.Sprint(tickets)
	fmt.Println("", ticketString[1:len(ticketString)-1])
}

func prev() {

	if meta.HasMore {
		tickets, meta, err = getTicketsWithCursor(meta.BeforeCursor, "")
		if err != nil {
			fmt.Println("Faile to get tickets")
			os.Exit(1)
		}
	}
	ticketString := fmt.Sprint(tickets)
	fmt.Println("", ticketString[1:len(ticketString)-1])
}

func next() {

	if meta.HasMore {
		tickets, meta, err = getTicketsWithCursor("", meta.AfterCursor)
		if err != nil {
			fmt.Println("Faile to get tickets")
			os.Exit(1)
		}
	}
	ticketString := fmt.Sprint(tickets)
	fmt.Println("", ticketString[1:len(ticketString)-1])
}

func selc() {

}

func quit() {
	fmt.Println("Have a good one!")
	os.Exit(0)
}

func getTicketsWithCursor(before, after string) ([]*Ticket, *Meta, error) {
	
	baseURL.Path = "/api/v2/tickets.json"
	params := url.Values{}

	params.Add("page[size]", "25")
	if before != "" {
		params.Add("page[before]", before)
	}
	if after != "" {
		params.Add("page[after]", after)
	}
	baseURL.RawQuery = params.Encode()

	request, _ := http.NewRequest(http.MethodGet, baseURL.String(), nil)
	request.SetBasicAuth(username, password)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Failed to make a HTTP request:", err)
		return nil, nil, nil
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read API response", err)
		return nil, nil, nil
	}

	var result ticketResp
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		fmt.Println("Failed to parse API response", err)
		return nil, nil, nil
	}

	return result.Tickets, &result.Meta, nil
}