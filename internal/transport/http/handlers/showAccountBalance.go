package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sqlBankCLI/internal/models"
	"sqlBankCLI/pkg/response"
)

func (h *Handler) ShowAccountBalance(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.Account

	err := json.NewDecoder(r.Body).Decode(&inputData)

	if err != nil {
		resp = response.BadRequest
		return
	}

	name, balance, err := h.svc.ShowAccountBalance(inputData)

	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success

	resp.Payload = fmt.Sprintf("Баланс счета %v является %v", name, balance)
}
