package billbee

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiClient struct {
	user     string
	password string
	apiKey   string
}

func NewApiClient(user string, password string, apiKey string) ApiClient {
	return ApiClient{
		user:     user,
		password: password,
		apiKey:   apiKey,
	}
}

func (api ApiClient) get(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("request failed (1)")
	}

	api.authenticateRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("request failed (2)")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("request failed (3)")
	}

	return string(body)
}

func (api ApiClient) authenticateRequest(req *http.Request) {
	b64auth := base64.StdEncoding.EncodeToString([]byte(api.user + ":" + api.password))
	req.Header.Add("Authorization", "Basic "+b64auth)
	req.Header.Add("X-Billbee-Api-Key", api.apiKey)
}

func (api ApiClient) GetLastOrders(minOrderId int) []Order {
	currentTime := time.Now().Add(1 * 24 * time.Hour * -1)

	response := &OrdersResponse{}
	url := fmt.Sprintf(
		"https://api.billbee.io/api/v1/orders?pageSize=200&page=1&minOrderDate=%d-%d-%d&minimumBillBeeOrderId=%d",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		minOrderId,
	)
	fmt.Println(url)
	err := json.Unmarshal(
		[]byte(
			api.get(
				url,
			),
		),
		response,
	)
	if err != nil {
		panic("couldn't parse billbee response")
	}

	return response.Data
}
