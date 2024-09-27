package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Order struct {
	Tip      int `json:"tip"`
	Customer struct {
		ShouldShowGeniusPrice bool `json:"shouldShowGeniusPrice"`
	} `json:"customer"`
	CurrencyCode           string   `json:"currencyCode"`
	CheckoutState          string   `json:"checkoutState"`
	ShippingState          string   `json:"shippingState"`
	TokenValue             string   `json:"tokenValue"`
	Id                     int      `json:"id"`
	Number                 string   `json:"number"`
	TotalInt               int      `json:"totalInt"`
	State                  string   `json:"state"`
	CreatedDate            string   `json:"createdDate"`
	DisplayState           string   `json:"displayState"`
	Currency               string   `json:"currency"`
	Total                  float64  `json:"total"`
	CheckoutCompletedAt    string   `json:"checkoutCompletedAt"`
	CanReportProblem       bool     `json:"canReportProblem"`
	ReportProblemTypes     []string `json:"reportProblemTypes"`
	DeliveryDate           string   `json:"deliveryDate"`
	CanBeCancelled         bool     `json:"canBeCancelled"`
	FeedbackUrl            string   `json:"feedbackUrl"`
	Notes                  string   `json:"notes"`
	MergeVersion           int      `json:"mergeVersion"`
	CanChangePaymentMethod bool     `json:"canChangePaymentMethod"`
}
type OrderListResponse struct {
	Page              int         `json:"page"`
	ItemsPerPage      int         `json:"itemsPerPage"`
	Limit             int         `json:"limit"`
	Pages             int         `json:"pages"`
	Total             int         `json:"total"`
	Count             int         `json:"count"`
	Items             []Order     `json:"items"`
	AnalyticsListId   interface{} `json:"analyticsListId"`
	AnalyticsListName interface{} `json:"analyticsListName"`
}

func main() {
	token := os.Getenv("TOKEN")
	i := 1
	for {
		url := fmt.Sprintf("%s/orders?page=%d&itemsPerPage=30", os.Getenv("URL"), i)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		bdy, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			panic("received status code no 200:" + resp.Status)
		}
		var data OrderListResponse
		err = json.Unmarshal(bdy, &data)
		if len(data.Items) == 0 {
			return
		}
		for _, item := range data.Items {
			filePath := fmt.Sprintf("./data/%s.json", item.Number)
			//
			// Check if file already exists
			if _, err := os.Stat(filePath); err == nil {
				log.Printf("File %s already exists. Skipping download.", filePath)
				continue
				return
			}

			req, err := http.NewRequest("GET", fmt.Sprintf("%s/myaccount_orders/%s", os.Getenv("URL"), item.TokenValue), nil)
			if err != nil {
				panic(err)
			}
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatalf("Failed to download file: %v", err)
			}
			defer resp.Body.Close()

			// Check for non-200 status code
			if resp.StatusCode != http.StatusOK {
				log.Fatalf("Failed to download file: received status code %d", resp.StatusCode)
			}

			// Create the file
			outFile, err := os.Create(filePath)
			if err != nil {
				log.Fatalf("Failed to create file: %v", err)
			}
			defer outFile.Close()

			// Write the body to file
			_, err = io.Copy(outFile, resp.Body)
			if err != nil {
				log.Fatalf("Failed to write file: %v", err)
			}

			fmt.Printf("File %s downloaded successfully.\n", filePath)
			fmt.Printf("%s:%s\n", item.Number, item.CheckoutCompletedAt)
		}
		i++
	}

}
