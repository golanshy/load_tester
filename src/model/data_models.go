package data_models

import "fmt"

// TestType: GetEntireMenu / Order / GetEntireMenuAndOrder

type LoadTestRequest struct {
	TestType            string `json:"test_type"`
	Repetitions         int    `json:"repetitions"`
	DelayInMilliseconds int    `json:"delay_in_milliseconds"`
}

type Location struct {
	LocationId string `json:"location_id"`
}

type Order struct {
	Data OrderData `json:"data"`
}

type OrderData struct {
	ID                       interface{}   `json:"id"`
	LocationID               string        `json:"location_id"`
	TableID                  string        `json:"table_id"`
	RoomNumber               string        `json:"room_number"`
	RoomSurname              string        `json:"room_surname"`
	CustomerName             string        `json:"customer_name"`
	PhoneNumber              string        `json:"phone_number"`
	OrderConfirmationMessage string        `json:"order_confirmation_message"`
	ReceiptURL               interface{}   `json:"receipt_url"`
	GroupID                  interface{}   `json:"group_id"`
	GroupName                interface{}   `json:"group_name"`
	OrderNumber              interface{}   `json:"order_number"`
	MerchantID               string        `json:"merchant_id"`
	Currency                 string        `json:"currency"`
	OrderType                string        `json:"order_type"`
	OrderItems               []OrderItem   `json:"order_items"`
	SignatureInBytes         []interface{} `json:"signature_in_bytes"`
	Amount                   float64       `json:"amount"`
	UIVersion                string        `json:"ui_version"`
}

type Names struct {
	En string `json:"en"`
}

type Descriptions struct {
	En string `json:"en"`
}

type AllergyInformation struct {
	En string `json:"en"`
}

type Availability struct {
	DayOfWeek         string `json:"dayOfWeek"`
	AvailabilityIndex int    `json:"availability_index"`
	StartTime         string `json:"startTime"`
	EndTime           string `json:"endTime"`
}

type OrderItem struct {
	Quantity              int           `json:"quantity"`
	ItemNotes             string        `json:"item_notes"`
	MenuItem              MenuItem      `json:"menu_item"`
	SelectedConfigurables []interface{} `json:"selected_configurables"`
}

type MenuItem struct {
	ID                 string             `json:"id"`
	CategoryID         string             `json:"category_id"`
	LocationID         string             `json:"location_id"`
	MenuID             string             `json:"menu_id"`
	Order              int                `json:"order"`
	Portrait           interface{}        `json:"portrait"`
	Price              float64            `json:"price"`
	CoverImageURL      string             `json:"cover_image_url"`
	ConfigurablesIntl  []interface{}      `json:"configurables_intl"`
	Names              Names              `json:"names"`
	Descriptions       Descriptions       `json:"descriptions"`
	UpdatedAt          interface{}        `json:"updated_at"`
	AllergyInformation AllergyInformation `json:"allergy_information"`
	Enabled            bool               `json:"enabled"`
	Timezone           interface{}        `json:"timezone"`
	Availability       []Availability     `json:"availability"`
}

func NewLocation() *Location {
	return &Location{
		LocationId: "eD4HVlwlfj88Ik0MNupE",
	}
}

func NewOrder(itemNumber int) *Order {

	orderItems := make([]OrderItem, 2)

	orderItems = append(orderItems, OrderItem{
		Quantity:  1,
		ItemNotes: fmt.Sprintf("Order counter: %d", itemNumber),
		MenuItem: MenuItem{
			ID:                "hvsYCnsAESN0UmXQkfQT",
			CategoryID:        "O6bRD44JUz7o9Phc0iLw",
			LocationID:        "eD4HVlwlfj88Ik0MNupE",
			MenuID:            "BAA77j0ViRq5foca9vtN",
			Order:             0,
			Portrait:          nil,
			Price:             6.50000000022,
			CoverImageURL:     "https://firebasestorage.googleapis.com/v0/b/leslie-tech.appspot.com/o/product_images%2Fa9a4decd-95a9-465a-81d3-3894bf75be33.jpg?alt=media&token=2265101e-e604-4e66-a7e4-ed539d97e748",
			ConfigurablesIntl: nil,
			Names: Names{
				En: "Feta and Red Onion Tart",
			},
			Descriptions: Descriptions{
				En: "It's ok I guess",
			},
			UpdatedAt: nil,
			AllergyInformation: AllergyInformation{
				En: "Contains pine nuts",
			},
			Enabled:      true,
			Timezone:     nil,
			Availability: nil,
		},
		SelectedConfigurables: nil,
	})

	orderItems = append(orderItems, OrderItem{
		Quantity:  1,
		ItemNotes: fmt.Sprintf("Order counter: %d", itemNumber),
		MenuItem: MenuItem{
			Price:      0.125,
			ID:         "service_charge",
			CategoryID: "service_charge",
			LocationID: "service_charge",
			MenuID:     "service_charge",
		},
		SelectedConfigurables: nil,
	})

	return &Order{
		Data: OrderData{
			ID:                       nil,
			LocationID:               "eD4HVlwlfj88Ik0MNupE",
			TableID:                  "29",
			RoomNumber:               "",
			RoomSurname:              "",
			CustomerName:             "",
			PhoneNumber:              "+447386238276",
			OrderConfirmationMessage: "Custom confirmation message",
			ReceiptURL:               nil,
			GroupID:                  nil,
			GroupName:                nil,
			OrderNumber:              nil,
			MerchantID:               "acct_1JbnN82e2TueidJl",
			Currency:                 "GBP",
			OrderType:                "cash",
			OrderItems:               orderItems,
			SignatureInBytes:         nil,
			Amount:                   0,
			UIVersion:                "",
		},
	}
}

type OrderResult struct {
	ID          interface{} `json:"id"`
	OrderNumber int         `json:"order_number"`
}
