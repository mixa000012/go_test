package service

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type HMACRequest struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type HMACResponse struct {
	Signature string `json:"signature"`
}

func GenerateHMAC(w http.ResponseWriter, r *http.Request) {
	var req HMACRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	hmacKey := []byte(req.Key)
	h := hmac.New(sha512.New, hmacKey)
	h.Write([]byte(req.Text))
	signature := hex.EncodeToString(h.Sum(nil))

	resp := HMACResponse{Signature: signature}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
