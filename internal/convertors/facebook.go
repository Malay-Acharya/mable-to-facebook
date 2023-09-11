package converters

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"mable-to-facebook/internal/models"
	"time"
)

func ConvertToFacebookFormat(payload models.MablePayload) models.FacebookData {
	var fb models.FacebookData
	currentTime := time.Now().Unix()

	fb.TestEventCode = payload.DestinationSettings.TestEventCode

	for _, mableEvent := range payload.Events {
		h := sha256.New()
		h.Write([]byte(mableEvent.SessionData.CustomerData.Email))
		hashedEmail := hex.EncodeToString(h.Sum(nil))

		// SHA256 hash the phone number.
		h.Reset()
		h.Write([]byte(mableEvent.SessionData.CustomerData.Phone))
		hashedPhone := hex.EncodeToString(h.Sum(nil))

		fbEvent := models.FacebookEvent{
			EventName:    mableEvent.EventName,
			EventTime:    currentTime, // Assuming you still want to use the current time for all events.
			ActionSource: "website",
		}
		fbEvent.UserData.Em = append(fbEvent.UserData.Em, hashedEmail)
		fbEvent.UserData.Ph = append(fbEvent.UserData.Ph, hashedPhone)

		// Map custom data from MablePayload to FacebookEvent.

		// Map currency from IpLocationData.
		fbEvent.CustomData.Currency = mableEvent.SessionData.IpLocationData.Currency

		// Assuming 'value' in CustomData is of type float64.
		if val, ok := mableEvent.CustomData["value"].(float64); ok {
			fbEvent.CustomData.Value = fmt.Sprintf("%.2f", val)
		}

		fbEvent.CustomData.Value = "12"

		fb.Data = append(fb.Data, fbEvent)
	}

	return fb
}
