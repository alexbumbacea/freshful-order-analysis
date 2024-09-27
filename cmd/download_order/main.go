package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpYXQiOjE3Mjc0NTUzOTEsImV4cCI6MTcyNzQ1NzE5MSwicm9sZXMiOlsiUk9MRV9VU0VSIl0sInVzZXJuYW1lIjoiZW1hZ0BidW1iYWNlYS5ybyJ9.Hght9rKsWqU306ZYm6nFHVeyk1aS6PHGBg0lsRXpdWxZ_CFHZVyoxMxrYtPYbZhkmLompQo0kmsGDtRGXu414mFt8VdTzCyKuv5QnKlELEavvmIoyYWg22O8giL5Jctl3EtXVmzy_0mMhB4giuBKTP1K1fqoutqYkDGEm0Yi9lcwOSkibDaFPOIRWsNee3SGb0Qn1J23VyXP4ewCkUhJGSTlWB-0E0aCi-SBP4rB1o62-gE0WaUrzphDRm4VU-sZVlzL7vm93RXvyqUfbZJrmGSChrwVFH_TvfUapvePEYMXuknHZEN6-7M5mlQHZVb2-oMlvmBHZwciOuNrtAmddbFMv4i4JSRJt8OZOMAzdLtR1ZaaRSJpH_Pd1yHm8Fwy7Z-zj_-pwXFGmm6QOx0MA-kf-XwYPMUWd8VuHVAKSAcGidfTcXotgf-343QqWG8R9BpMw4HUFbnEkGWRn-Ik76kj4QOXH9kC1zzirE5c0EeID3GZhJt2mz8tsZIQRsjJvZujlsotGxtFZbkmqMouHlhwpb8aXXshhBPoAvK6dbPZBSLBOYCIt_c8h0yFd3r-3uCk_oCaET0B9GiRtpuw5_7FAZOY0upC221biystXnoJu9m7kWyYTJFJqWy99NtRRqneW62MHEQi7L2LDAGX8Nqxs5DBd3aOo8ZoSIDTuTQ"
	i := 1
	for {
		url := fmt.Sprintf("https://www.freshful.ro/api/v2/shop/orders?page=%d&itemsPerPage=30", i)
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
				return
			}

			req, err := http.NewRequest("GET", fmt.Sprintf("https://www.freshful.ro/api/v2/shop/myaccount_orders/%s", item.TokenValue), nil)
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
