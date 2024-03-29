package handlers

import (
	"encoding/json"
	"net/http"
	"sqlBankCLI/internal/models"
	"sqlBankCLI/pkg/response"
)

func (h *Handler) TopUpClientsAccount(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.Response

	var account models.Account

	err := json.NewDecoder(r.Body).Decode(&inputData)
	
	if err != nil {
		resp = response.BadRequest
		return
	}

	account.PhoneNumber = inputData.PhoneNumber

	err = h.svc.TopUpClientsAccount(account, inputData.Amount)

	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
}