package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	converters "mable-to-facebook/internal/convertors"
	"mable-to-facebook/internal/models"
	"mable-to-facebook/internal/utils"
	"net/http"
)

func SendToFacebookHandler(w http.ResponseWriter, r *http.Request) {
	var response models.APIResponse

	// 1. Parse the input
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response = models.APIResponse{Success: false, Message: "Failed to process request", Details: "Failed to read request body"}
		utils.WriteResponse(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	var mablePayload models.MablePayload
	err = json.Unmarshal(body, &mablePayload)
	if err != nil {
		response = models.APIResponse{Success: false, Message: "Failed to process request", Details: "Invalid data format"}
		utils.WriteResponse(w, http.StatusBadRequest, response)
		return
	}

	// 2. Convert to Facebook format
	fbData := converters.ConvertToFacebookFormat(mablePayload)

	// 3. Send the data to Facebook's Graph API
	fbEndpoint := fmt.Sprintf("https://graph.facebook.com/v13.0/%s/events?access_token=%s",
		mablePayload.DestinationSettings.PixelID,
		mablePayload.DestinationSettings.AccessKey)

	fbPayload, _ := json.Marshal(fbData)
	resp, err := http.Post(fbEndpoint, "application/json", bytes.NewReader(fbPayload))
	if err != nil {
		response = models.APIResponse{Success: false, Message: "Failed to communicate with Facebook", Details: err.Error()}
		utils.WriteResponse(w, http.StatusInternalServerError, response)
		return
	}
	defer resp.Body.Close()

	fbResponseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		response = models.APIResponse{Success: false, Message: "Failed to process Facebook's response", Details: err.Error()}
		utils.WriteResponse(w, http.StatusInternalServerError, response)
		return
	}

	if resp.StatusCode != http.StatusOK {
		response = models.APIResponse{Success: false, Message: fmt.Sprintf("Received %d response from Facebook", resp.StatusCode), Details: string(fbResponseBody)}
		utils.WriteResponse(w, http.StatusInternalServerError, response)
		return
	}

	response = models.APIResponse{Success: true, Message: "Data sent successfully to Facebook", Details: json.RawMessage(fbResponseBody)}
	utils.WriteResponse(w, http.StatusOK, response)
}
