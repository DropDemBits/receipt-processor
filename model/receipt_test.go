package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidReceipt(t *testing.T) {
	var tests = []Receipt{
		{
			Retailer:     "Retailer",
			PurchaseDate: "2025-05-02",
			PurchaseTime: "13:05",
			Total:        "1.23",
			Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
		},
		// Multiple items
		{
			Retailer:     "Retailer",
			PurchaseDate: "2025-05-02",
			PurchaseTime: "13:05",
			Total:        "1.23",
			Items: []Item{
				{ShortDescription: "Something", Price: "1.23"},
				{ShortDescription: "Something Else", Price: "4.56"},
			},
		},
		// Leading & trailing spaces in brand names
		{
			Retailer:     "              Retailer              ",
			PurchaseDate: "2025-05-02",
			PurchaseTime: "13:05",
			Total:        "1.23",
			Items: []Item{
				{ShortDescription: "               Something            ", Price: "1.23"},
				{ShortDescription: "               Something Else       ", Price: "4.56"},
			},
		},
	}

	for _, tc := range tests {
		receipt := tc

		t.Run("Valid receipt", func(t *testing.T) {
			t.Parallel()

			assert.NoError(t, receipt.Validate())
		})
	}
}

func TestInvalidReceipts(t *testing.T) {
	var tests = []struct {
		name    string
		receipt Receipt
	}{
		{
			"Retailer Name - Empty",
			Receipt{
				Retailer:     "",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Retailer Name - Invalid Chars",
			Receipt{
				Retailer:     "Retailer!",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Purchase Date - Empty",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Purchase Date - Bad Format",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Purchase Date - Unexpected Time",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02 13:05",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Purchase Time - Empty",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Purchase Time - Bad Format",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05:02",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Total - Empty",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Total - More decimal digits",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.000",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Total - Less decimal digits",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.0",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Total - No decimal digits",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Total - No decimal place",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Total - No leading decimal digits",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        ".11",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Total - Not a decimal place",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "11?11",
				Items:        []Item{{ShortDescription: "Something", Price: "1.23"}},
			},
		},
		{
			"Items - Empty",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{},
			},
		},
		// Item-level
		{
			"Item.ShortDescription - Empty",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "", Price: "1.23"}},
			},
		},
		{
			"Item.ShortDescription - Invalid Chars",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something!", Price: "1.23"}},
			},
		},
		{
			"Item.Price - Empty",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: ""}},
			},
		},
		{
			"Item.Price - More decimal digits",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1.000"}},
			},
		},
		{
			"Item.Price - Less decimal digits",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1.0"}},
			},
		},
		{
			"Item.Price - No decimal digits",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1."}},
			},
		},
		{
			"Item.Price - No decimal place",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "1"}},
			},
		},
		{
			"Item.Price - No leading decimal digits",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: ".23"}},
			},
		},
		{
			"Item.Price - Not a decimal place",
			Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2025-05-02",
				PurchaseTime: "13:05",
				Total:        "1.23",
				Items:        []Item{{ShortDescription: "Something", Price: "12?34"}},
			},
		},
	}

	for _, tc := range tests {
		receipt := tc.receipt

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Error(t, receipt.Validate())
		})
	}
}
