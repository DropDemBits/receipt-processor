package processor

import (
	"math"
	"receipt-processor/model"
	"regexp"
	"strings"
	"time"
)

// Any potential rule that can award points.
//
// Matches criteria on a receipt, and may give a certain amount of points.
// Should an error be encountered at any point, processing stops and no points
// are awarded.
type RuleFunc func(*model.Receipt) (int, error)

// A set of rules to award points by for a given [model.Receipt].
type RuleEngine struct {
	rules []RuleFunc
}

// Creates a new rule engine instance to process receipts.
func NewProcessor() *RuleEngine {
	return new(RuleEngine)
}

// Adds one or more rules to awards points by.
// This is not thread-safe, and it is expected that a rule engine be finalized
// before using it for processing.
func (e *RuleEngine) AddRule(fs ...RuleFunc) {
	e.rules = append(e.rules, fs...)
}

// Processes a receipt and computes the awarded points.
// Returns the sum of the points awarded by each rule, or the error encountered during evaluation.
func (e *RuleEngine) Process(receipt *model.Receipt) (int, error) {
	points := 0

	for _, rule := range e.rules {
		awarded, err := rule(receipt)
		if err != nil {
			return 0, err
		}

		points += awarded
	}

	return points, nil
}

var (
	alphanumericMatcher = regexp.MustCompile("[a-zA-Z0-9]")
)

// Default set of rules to award points by.
func DefaultRules() *RuleEngine {
	r := NewProcessor()

	// Based on the challenge requirements:
	//
	//   - +1 point for every alphanumeric character in the retailer name.
	r.AddRule(func(r *model.Receipt) (int, error) {
		alnums := alphanumericMatcher.FindAll([]byte(r.Retailer), -1)
		return len(alnums), nil
	})
	//   - +50 points if the total is a round dollar amount with no cents.
	//   - +25 points if the total is a multiple of 0.25.
	r.AddRule(func(r *model.Receipt) (int, error) {
		// We fuse the rules as they both operate on cents.
		cents, err := r.Total.Cents()
		if err != nil {
			return 0, err
		}
		// Sanity check: ensure that the total value is at least greater than 0
		total, err := r.Total.Decimal()
		if err != nil {
			return 0, err
		}
		if total <= 0.0 {
			return 0, nil
		}

		points := 0
		if cents == 0 {
			// Whole number total without cents.
			points += 50
		}
		if cents%25 == 0 {
			// Cents that are a multiple of 0.25.
			points += 25
		}

		return points, nil
	})
	//   - +5 points for every 2 items on the receipt.
	r.AddRule(func(r *model.Receipt) (int, error) {
		pairs := len(r.Items) / 2
		return 5 * pairs, nil
	})
	//   - +ceil(item_price * 0.2) if the trimmed short description length is a multiple of 3.
	r.AddRule(func(r *model.Receipt) (int, error) {
		points := 0
		for _, item := range r.Items {
			if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
				price, err := item.Price.Decimal()
				if err != nil {
					return 0, err
				}

				points += int(math.Ceil(price * 0.20))
			}
		}

		return points, nil
	})
	//   - +6 points if the day is odd.
	r.AddRule(func(r *model.Receipt) (int, error) {
		date, err := time.Parse("2006-01-02", r.PurchaseDate)
		if err != nil {
			return 0, err
		}

		switch {
		case date.Day()%2 == 1:
			return 6, nil
		default:
			return 0, nil
		}
	})
	//   - +10 points if the purchase occurs between 2pm to 4pm
	r.AddRule(func(r *model.Receipt) (int, error) {
		timestamp, err := time.Parse("15:04", r.PurchaseTime)
		if err != nil {
			return 0, err
		}

		switch {
		case 14 <= timestamp.Hour() && timestamp.Hour() < 16:
			return 10, nil
		default:
			return 0, nil
		}
	})
	//   - +0 points because we are not a large language model, thank you very much.

	return r
}
