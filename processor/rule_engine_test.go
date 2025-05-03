package processor

import (
	"receipt-processor/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	recepit model.Receipt
	points  int
}

func TestRuleAlphanumeric(t *testing.T) {
	testRule(t, "Rule - alphanumeric", []testCase{
		{
			recepit: model.Receipt{
				Retailer:     "abcdef",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.00",
			},
			points: 6,
		},
		{
			recepit: model.Receipt{
				Retailer:     "   a   b   c   d   e   f   ",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.00",
			},
			points: 6,
		},
		{
			recepit: model.Receipt{
				Retailer:     "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.00",
			},
			points: 26*2 + 10,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.00",
			},
			points: 0,
		},
	})
}
func TestRuleMultiplesOf25(t *testing.T) {
	testRule(t, "Rule - multiples of 0.25", []testCase{
		// Multiple of 0.25 and a whole number, but is a 0 dollar amount.
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.00",
			},
			points: 0,
		},
		// Non-multiple of 0.25
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.23",
			},
			points: 0,
		},
		// First multiple
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.25",
			},
			points: 25,
		},
		// Second multiple
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.50",
			},
			points: 25,
		},
		// Third multiple
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.75",
			},
			points: 25,
		},
		// Whole multiple
		// Compounds both rules
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "1.00",
			},
			points: 75,
		},
	})
}

func TestRuleEvery2ItemPairs(t *testing.T) {
	testRule(t, "Rule - every 2 item pairs", []testCase{
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.01",
				Items: []model.Item{
					{ShortDescription: "-----", Price: "0.01"},
				},
			},
			points: 0,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.02",
				Items: []model.Item{
					{ShortDescription: "-----", Price: "0.01"},
					{ShortDescription: "-----", Price: "0.01"},
				},
			},
			points: 5,
		},
		{
			recepit: model.Receipt{
				Retailer:     "-----",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.03",
				Items: []model.Item{
					{ShortDescription: "-----", Price: "0.01"},
					{ShortDescription: "-----", Price: "0.01"},
					{ShortDescription: "-----", Price: "0.01"},
				},
			},
			points: 5,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.04",
				Items: []model.Item{
					{ShortDescription: "-----", Price: "0.01"},
					{ShortDescription: "-----", Price: "0.01"},
					{ShortDescription: "-----", Price: "0.01"},
					{ShortDescription: "-----", Price: "0.01"},
				},
			},
			points: 10,
		},
	})
}

func TestRuleTrimmedShortDescription(t *testing.T) {
	testRule(t, "Rule - trimmed short description", []testCase{
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.03",
				Items: []model.Item{
					{ShortDescription: "aaa", Price: "0.01"},
					{ShortDescription: "bbb", Price: "0.01"},
					{ShortDescription: "ccc", Price: "0.01"},
				},
			},
			points: 3 + 5,
		},
		// Total length of 3 but not a trimmed length of 3
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.05",
				Items: []model.Item{
					{ShortDescription: "  a", Price: "0.01"},
					{ShortDescription: " b ", Price: "0.01"},
					{ShortDescription: "c  ", Price: "0.01"},
					{ShortDescription: " aa", Price: "0.01"},
					{ShortDescription: "cc ", Price: "0.01"},
				},
			},
			points: 0 + 5*2,
		},
		// Trimmed length of 3 (since we don't care about interstitial spaces)
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.03",
				Items: []model.Item{
					{ShortDescription: "  a a", Price: "0.01"},
					{ShortDescription: " b b ", Price: "0.01"},
					{ShortDescription: "c c", Price: "0.01"},
				},
			},
			points: 3 + 5,
		},
		// Below, above and exactly at a multiple of 5
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.03",
				Items: []model.Item{
					{ShortDescription: "aaa", Price: "19.99"},
					{ShortDescription: "bbb", Price: "20.00"},
					{ShortDescription: "ccc", Price: "20.01"},
				},
			},
			points: 13 + 5,
		},
	})
}

func TestRuleOddDate(t *testing.T) {
	testRule(t, "Rule - between times", []testCase{
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "09:58",
				Total:        "0.01",
			},
			points: 0,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-01",
				PurchaseTime: "09:58",
				Total:        "0.01",
			},
			points: 6,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-03",
				PurchaseTime: "09:58",
				Total:        "0.01",
			},
			points: 6,
		},
	})
}

func TestRuleBetweenTimes(t *testing.T) {
	testRule(t, "Rule - between times", []testCase{
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "13:59",
				Total:        "0.01",
			},
			points: 0,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "14:00",
				Total:        "0.01",
			},
			points: 10,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "14:01",
				Total:        "0.01",
			},
			points: 10,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "14:30",
				Total:        "0.01",
			},
			points: 10,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "15:59",
				Total:        "0.01",
			},
			points: 10,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "16:00",
				Total:        "0.01",
			},
			points: 0,
		},
		{
			recepit: model.Receipt{
				Retailer:     "------",
				PurchaseDate: "2027-06-02",
				PurchaseTime: "16:01",
				Total:        "0.01",
			},
			points: 0,
		},
	})
}

func testRule(t *testing.T, group string, tests []testCase) {
	rules := DefaultRules()

	for _, tc := range tests {
		t.Run(group, func(t *testing.T) {
			// t.Parallel()

			points, err := rules.Process(&tc.recepit)
			if assert.NoError(t, err) {
				assert.Equal(t, tc.points, points)
			}
		})
	}
}
