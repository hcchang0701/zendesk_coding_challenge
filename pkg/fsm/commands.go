package fsm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"

	"github.com/joho/godotenv"
)

var (
	client   *http.Client
	baseURL  *url.URL
	username string
	password string
	tickets  []*Ticket
	meta     *Meta
	err      error
)

func init() {
	client = &http.Client{}
	baseURL, _ = url.Parse("https://zcchcc.zendesk.com")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env file:", err)
		os.Exit(1)
	}
	username = os.Getenv("ZENDESK_USERNAME")
	password = os.Getenv("ZENDESK_PASSWORD")
}

func list() {

	tickets, meta, err = getTicketsWithCursor("", "")
	if err != nil {
		fmt.Println("failed to get tickets:", err)
		os.Exit(1)
	}

	ticketString := fmt.Sprint(tickets)
	fmt.Println("\n", ticketString[1:len(ticketString)-1])
}

func prev() {

	if meta.HasMore {
		tickets, meta, err = getTicketsWithCursor(meta.BeforeCursor, "")
		if err != nil {
			fmt.Println("failed to get tickets:", err)
			os.Exit(1)
		}
	}
	ticketString := fmt.Sprint(tickets)
	fmt.Println("\n", ticketString[1:len(ticketString)-1])
}

func next() {

	if meta.HasMore {
		tickets, meta, err = getTicketsWithCursor("", meta.AfterCursor)
		if err != nil {
			fmt.Println("failed to get tickets:", err)
			os.Exit(1)
		}
	}
	ticketString := fmt.Sprint(tickets)
	fmt.Println("\n", ticketString[1:len(ticketString)-1])
}

func selc(num int) bool {

	i := sort.Search(len(tickets), func(i int) bool { return tickets[i].ID <= num })
	if i == len(tickets) || tickets[i].ID != num {
		fmt.Printf("Ticket #%d is not in the list, please select again\n", num)
		return false
	}

	view := fmt.Sprintf("\n\n#%d %v\n", tickets[i].ID, tickets[i].Subject) +
		fmt.Sprintf("Type: %v; Status: %v; Priority: %v\n\n", tickets[i].Type, tickets[i].Status, tickets[i].Priority) +
		fmt.Sprintf("%v\n\n", tickets[i].Description) +
		fmt.Sprintf("Created at: %v; Last updated at: %v", tickets[i].CreatedAt, tickets[i].UpdatedAt)

	fmt.Println(view)
	return true
}

func back() {
	ticketString := fmt.Sprint(tickets)
	fmt.Println("\n", ticketString[1:len(ticketString)-1])
}

func quit() {
	fmt.Println("Have a good one!")
	os.Exit(0)
}

func getTicketsWithCursor(before, after string) ([]*Ticket, *Meta, error) {

	baseURL.Path = "/api/v2/tickets.json"
	params := url.Values{}

	params.Add("sort", "-id")
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
		err = errors.New("failed to make a HTTP request: " + err.Error())
		return nil, nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("failed to read API response: " + err.Error())
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		return nil, nil, err
	}

	var result ticketResp
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		err = errors.New("failed to parse API response: " + err.Error())
		return nil, nil, err
	}

	return result.Tickets, &result.Meta, nil
}
