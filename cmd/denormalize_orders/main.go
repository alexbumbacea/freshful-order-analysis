package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Order struct {
	CheckoutNoticeV2 []interface{} `json:"checkoutNoticeV2"`
	CartNoticeV2     []interface{} `json:"cartNoticeV2"`
	OrderNotice      []interface{} `json:"orderNotice"`
	ShippingAddress  struct {
		ContactPhonenumber string `json:"contactPhonenumber"`
		Province           string `json:"province"`
		ContactName        string `json:"contactName"`
		StreetAddress      string `json:"streetAddress"`
		City               string `json:"city"`
		FirstName          string `json:"firstName"`
		CountryCode        string `json:"countryCode"`
		Postcode           string `json:"postcode"`
	} `json:"shippingAddress"`
	BillingAddress struct {
		ContactPhonenumber string `json:"contactPhonenumber"`
		Province           string `json:"province"`
		ContactName        string `json:"contactName"`
		StreetAddress      string `json:"streetAddress"`
		City               string `json:"city"`
		FirstName          string `json:"firstName"`
		CountryCode        string `json:"countryCode"`
		Postcode           string `json:"postcode"`
	} `json:"billingAddress"`
	Schedule struct {
		Id          int  `json:"id"`
		IsAvailable bool `json:"isAvailable"`
	} `json:"schedule"`
	CouponInfo struct {
		Code          string  `json:"code"`
		PromotionName string  `json:"promotionName"`
		TotalDiscount float64 `json:"totalDiscount"`
	} `json:"couponInfo"`
	AgeCheckRequired bool `json:"ageCheckRequired"`
	TipOptions       []struct {
		Key   int    `json:"key"`
		Label string `json:"label"`
	} `json:"tipOptions"`
	Tip     int `json:"tip"`
	Summary []struct {
		Type string `json:"type"`
		Size string `json:"size"`
		Key  struct {
			PopoverContent string `json:"popoverContent"`
			Text           string `json:"text"`
			Bold           bool   `json:"bold"`
			Style          string `json:"style"`
		} `json:"key,omitempty"`
		Value struct {
			Text  string `json:"text"`
			Bold  bool   `json:"bold"`
			Style string `json:"style"`
		} `json:"value,omitempty"`
		Modal struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			ButtonLabel string `json:"buttonLabel"`
			Summary     []struct {
				Type string `json:"type"`
				Size string `json:"size"`
				Key  struct {
					PopoverContent string `json:"popoverContent"`
					Text           string `json:"text"`
					Bold           bool   `json:"bold"`
					Style          string `json:"style"`
				} `json:"key,omitempty"`
				Value struct {
					Text  string `json:"text"`
					Bold  bool   `json:"bold"`
					Style string `json:"style"`
				} `json:"value,omitempty"`
			} `json:"summary"`
		} `json:"modal,omitempty"`
		Description struct {
			Text  string `json:"text"`
			Bold  bool   `json:"bold"`
			Style string `json:"style"`
		} `json:"description,omitempty"`
	} `json:"summary"`
	Customer struct {
		Id                    int    `json:"id"`
		Email                 string `json:"email"`
		FirstName             string `json:"firstName"`
		PhoneNumber           string `json:"phoneNumber"`
		ShouldShowGeniusPrice bool   `json:"shouldShowGeniusPrice"`
	} `json:"customer"`
	Channel  string `json:"channel"`
	Payments []struct {
		Id     int `json:"id"`
		Method struct {
			CheckoutEnabled bool   `json:"checkoutEnabled"`
			Id              int    `json:"id"`
			Code            string `json:"code"`
			Translations    struct {
				RoRO struct {
					Id           int    `json:"id"`
					Name         string `json:"name"`
					Description  string `json:"description"`
					Instructions string `json:"instructions"`
				} `json:"ro_RO"`
			} `json:"translations"`
			Type         string `json:"type"`
			Name         string `json:"name"`
			Description  string `json:"description"`
			Instructions string `json:"instructions"`
		} `json:"method"`
		State    string  `json:"state"`
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
		Type     string  `json:"type"`
	} `json:"payments"`
	Shipments []struct {
		Id     int    `json:"id"`
		Method string `json:"method"`
	} `json:"shipments"`
	CurrencyCode           string        `json:"currencyCode"`
	LocaleCode             string        `json:"localeCode"`
	CheckoutState          string        `json:"checkoutState"`
	PaymentState           string        `json:"paymentState"`
	ShippingState          string        `json:"shippingState"`
	TokenValue             string        `json:"tokenValue"`
	Id                     int           `json:"id"`
	Number                 string        `json:"number"`
	Items                  []interface{} `json:"items"`
	ItemsHash              string        `json:"itemsHash"`
	ItemsTotalInt          int           `json:"itemsTotalInt"`
	TotalInt               int           `json:"totalInt"`
	State                  string        `json:"state"`
	CreatedDate            string        `json:"createdDate"`
	DisplayState           string        `json:"displayState"`
	Currency               string        `json:"currency"`
	TaxTotal               float64       `json:"taxTotal"`
	TaxTotalInt            int           `json:"taxTotalInt"`
	Total                  float64       `json:"total"`
	ShippingTotal          float64       `json:"shippingTotal"`
	ShippingTotalIgnore    int           `json:"shippingTotalIgnore"`
	ItemsTotal             float64       `json:"itemsTotal"`
	OrderPromotionTotalInt int           `json:"orderPromotionTotalInt"`
	OrderPromotionTotal    float64       `json:"orderPromotionTotal"`
	Invoices               []struct {
		Label string `json:"label"`
		Url   string `json:"url"`
	} `json:"invoices"`
	CheckoutCompletedAt       string        `json:"checkoutCompletedAt"`
	CanReportProblem          bool          `json:"canReportProblem"`
	ReportProblemTypes        []interface{} `json:"reportProblemTypes"`
	CanReturnProducts         bool          `json:"canReturnProducts"`
	DeliveryDate              string        `json:"deliveryDate"`
	TipTotal                  int           `json:"tipTotal"`
	SurchargeTotal            float64       `json:"surchargeTotal"`
	SurchargeLabel            string        `json:"surchargeLabel"`
	ShouldRetryPayment        bool          `json:"shouldRetryPayment"`
	IncompletePromotionsCount int           `json:"incompletePromotionsCount"`
	AdditionalDiscounts       []interface{} `json:"additionalDiscounts"`
	MealVoucherTotal          int           `json:"mealVoucherTotal"`
	NonMealVoucherItemsTotal  float64       `json:"nonMealVoucherItemsTotal"`
	SummaryItems              []struct {
		Replacements []interface{} `json:"replacements"`
		ProductName  string        `json:"productName"`
		VariantName  string        `json:"variantName"`
		Id           int           `json:"id"`
		Quantity     int           `json:"quantity"`
		ProductSlug  string        `json:"productSlug"`
		Subtotal     float64       `json:"subtotal"`
		Sku          string        `json:"sku"`
		VariantCode  string        `json:"variantCode"`
		Image        struct {
			Default string `json:"default"`
		} `json:"image"`
		Brand       string `json:"brand"`
		State       string `json:"state"`
		Message     string `json:"message"`
		Breadcrumbs []struct {
			Code string `json:"code"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"breadcrumbs"`
		IsGift     bool   `json:"isGift"`
		PriceLabel string `json:"priceLabel"`
	} `json:"summaryItems"`
	CanBeCancelled         bool   `json:"canBeCancelled"`
	IsZone0                bool   `json:"isZone0"`
	ItemsTotalDiscount     string `json:"itemsTotalDiscount"`
	FeedbackUrl            string `json:"feedbackUrl"`
	Notes                  string `json:"notes"`
	MergeVersion           int    `json:"mergeVersion"`
	CanChangePaymentMethod bool   `json:"canChangePaymentMethod"`
}

func main() {
	// Path to the folder containing JSON files
	dataFolder := "./data"

	// Read all files in the data folder
	files, err := os.ReadDir(dataFolder)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	outFile, err := os.Create("denormalized.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer outFile.Close()
	csvWriter := csv.NewWriter(outFile)
	csvWriter.Write([]string{
		"data.Number",
		"data.month",
		"data.CheckoutCompletedAt",
		"data.DeliveryDate",
		"data.State",
		"data.CouponInfo.Code",
		"item.Sku",
		"taxon 1",
		"taxon 2",
		"taxon 3",
		"item.Quantity",
		"item.Subtotal",
		"item.ProductName",
		"item.IsGift",
		"item.UnitPrice",
	})

	// Loop through each file in the folder
	for _, file := range files {
		// Only process .json files
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(dataFolder, file.Name())

			// Open and read the JSON file
			jsonFile, err := os.Open(filePath)
			if err != nil {
				log.Printf("Failed to open file %s: %v", file.Name(), err)
				continue
			}
			defer jsonFile.Close()

			// Read the file content
			byteValue, err := ioutil.ReadAll(jsonFile)
			if err != nil {
				log.Printf("Failed to read file %s: %v", file.Name(), err)
				continue
			}

			// Decode the JSON data into a struct
			var data Order
			err = json.Unmarshal(byteValue, &data)
			if err != nil {
				log.Printf("Failed to decode JSON file %s: %v", file.Name(), err)
				continue
			}

			for _, item := range data.SummaryItems {
				if len(item.Replacements) == 0 {
					split := strings.Split(data.CheckoutCompletedAt[3:10], ".")
					slices.Reverse(split)
					csvWriter.Write([]string{
						data.Number,
						strings.Join(split, "."),
						data.CheckoutCompletedAt,
						data.DeliveryDate,
						data.State,
						data.CouponInfo.Code,
						item.Sku,
						item.Breadcrumbs[0].Slug,
						item.Breadcrumbs[1].Slug,
						item.Breadcrumbs[2].Slug,
						fmt.Sprintf("%d", item.Quantity),
						fmt.Sprintf("%f", item.Subtotal),
						item.ProductName,
						fmt.Sprintf("%v", item.IsGift),
						fmt.Sprintf("%.2f", item.Subtotal/float64(item.Quantity)),
					})
				}
			}
		}
	}
}
