package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt-processor/model"
	"receipt-processor/routes"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProcessRoute(t *testing.T) {
	router := routes.Router()

	receipt := model.Receipt{
		Retailer:     "!",
		PurchaseDate: "2025-05-02",
		PurchaseTime: "13:52",
		Total:        "0.01",
		Items: []model.Item{
			{
				ShortDescription: "Thing", Price: "0.01",
			},
		},
	}
	receiptJson, _ := json.Marshal(receipt)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(string(receiptJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Body must be `{"id":"<uuid>"}`
	var body struct {
		Id string `json:"id"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if assert.NoError(t, err) {
		_, err := uuid.Parse(body.Id)
		assert.NoError(t, err, "id is not a valid uuid")
	}
}

func TestProcessRouteInvalid(t *testing.T) {
	router := routes.Router()

	receipt := model.Receipt{
		// "!" is not accepted as a retailer name.
		Retailer:     "!",
		PurchaseDate: "2025-05-02",
		PurchaseTime: "13:52",
		Total:        "0.01",
		Items: []model.Item{
			{
				ShortDescription: "Thing", Price: "0.01",
			},
		},
	}
	receiptJson, _ := json.Marshal(receipt)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(string(receiptJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPointsRoute(t *testing.T) {
	router := routes.Router()

	receipt := model.Receipt{
		// +11 points (11 alphanumeric chars)
		Retailer: "Supermarket 1",
		// +6 points (day is odd)
		PurchaseDate: "2025-05-03",
		// +10 points (after 2pm but before 4pm)
		PurchaseTime: "14:00",
		// +0 points (neither .00, .25, .50, nor .75)
		Total: "0.01",
		// +0 points (0 pairs of 2)
		Items: []model.Item{
			{
				ShortDescription: "Thing", Price: "0.01",
			},
		},
	}
	receiptJson, _ := json.Marshal(receipt)

	var body struct {
		Id string `json:"id"`
	}

	{
		w := httptest.NewRecorder()
		postReq, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(string(receiptJson)))
		router.ServeHTTP(w, postReq)

		assert.Equal(t, http.StatusOK, w.Code)

		err := json.Unmarshal(w.Body.Bytes(), &body)
		if !assert.NoError(t, err) {
			return
		}
	}

	{
		w := httptest.NewRecorder()
		getReq, _ := http.NewRequest("GET", "/receipts/"+body.Id+"/points", nil)
		router.ServeHTTP(w, getReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var points struct {
			Points int `json:"points"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &points)
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, points.Points, 27)
	}
}

func TestGetPointsRouteInvalid(t *testing.T) {
	router := routes.Router()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/receipts/062e96a5-3dcf-4a42-b565-06ea3bc3d2e9/points", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
