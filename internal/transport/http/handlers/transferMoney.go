package handlers

import (
	"encoding/json"
	"net/http"
	"sqlBankCLI/internal/models"
	"sqlBankCLI/pkg/response"
)

func (h *Handler) TransferMoney(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.TransferResponse

	err := json.NewDecoder(r.Body).Decode(&inputData)
	
	if err != nil {
		resp = response.BadRequest
		return
	}


	err = h.svc.TransferMoney(inputData.SenderPhone,inputData.RecipientPhone, inputData.Amount)

	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
