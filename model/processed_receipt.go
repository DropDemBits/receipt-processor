package model

import "github.com/google/uuid"

var processedReceiptRepo = new(Repository[string, ProcessedReceipt])

// A processed receipt, including how many points to award.
type ProcessedReceipt struct {
	// Id of the processing. Formatted as a UUID
	Id string
	// How many points were awarded.
	Points int
}

// Adds a processed receipt record.
// Returns an id to later fetch the record.
func AddProcessedReceipt(r ProcessedReceipt) string {
	id := uuid.NewString()
	r.Id = id

	// UUIDs should be unique
	processedReceiptRepo.Add(id, r)

	return id
}

// Gets a processed receipt record.
// Returns the value and true if there is a corresponding value for the key, false otherwise.
func GetProcessedReceipt(id string) (ProcessedReceipt, bool) {
	return processedReceiptRepo.GetSingle(id)
}
