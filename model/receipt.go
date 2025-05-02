package model

import "strconv"

// Representation of an amount of currency.
type Currency string

// The currency as a decimal value.
// This is always in the format of "1111111.11", i.e. there are exactly 2 decimal places.
//
// As a shortcut, this is parsed as a float64, which may incur precision loss
// for large currency values. However, the scale at which this occurs is around
// 2^46 or 7.03 × 10^13, so we wouldn't run into this precision issue in
// practice.
func (c Currency) Decimal() (float64, error) {
	return strconv.ParseFloat(string(c), 64)
}

// A structured Receipt.
// Result of converting an image of a receipt into a structured format.
type Receipt struct {
	// The name of the retailer or store the receipt is from.
	// Allows alphanumeric characters, "-", "&", and spaces.
	Retailer string `json:"retailer" validate:"brand_name"`
	// The date of the purchase printed on the receipt. ISO-8601 date expected.
	PurchaseDate string `json:"purchaseDate" validate:"datetime=2006-01-02"`
	// The time of the purchase printed on the receipt. 24-hour time expected.
	// Timezone is assumed to be in Central Standard/Daylight Time.
	PurchaseTime string `json:"purchaseTime" validate:"datetime=15:04"`
	// The total amount paid on the receipt.
	Total Currency `json:"total" validate:"currency"`

	// Items that are part of the receipt. A receipt must have at least 1 item.
	Items []Item `json:"items" validate:"min=1,dive"`
}

// Validates that the receipt matches the expected schema.
// See the [Receipt] field documentation for more details.
func (r *Receipt) Validate() error {
	return validate.Struct(r)
}

// An Item on a [Receipt].
type Item struct {
	// The Short Product Description for the item.
	// Allows alphanumeric characters, "-", "&", and spaces.
	ShortDescription string `json:"shortDescription" validate:"brand_name"`
	// The total price paid for this item.
	Price Currency `json:"price" validate:"currency"`
}
