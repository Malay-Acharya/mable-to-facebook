package models

type GDPRConsentEnum string
type SchemaVersion string
type MableSourcePlatform string
type MableEventOrigin string
type SessionLookupKeys string
type ShopifySessionSyncMethods string

type MableEvent struct {
	EventID     string `json:"eventId"`
	EventName   string `json:"eventName"`
	EventType   string `json:"eventType"`
	SessionData struct {
		SessionID   string `json:"sessionId"`
		AnonymousID string `json:"anonymousId"`
		BrowserData struct {
			UserAgent string `json:"userAgent"`
			Etag      struct {
				ID       int64 `json:"id"`
				LastSeen int64 `json:"lastSeen"`
				Session  int   `json:"session"`
			} `json:"etag"`
			Utm struct {
				UtmCampaign string `json:"utmCampaign"`
				UtmContent  string `json:"utmContent"`
				UtmMedium   string `json:"utmMedium"`
				UtmSource   string `json:"utmSource"`
			} `json:"utm"`
			Cookies struct {
				_Fbp   string `json:"_fbp"`
				_Ttp   string `json:"_ttp"`
				_Fbc   string `json:"_fbc"`
				_Ga    string `json:"_ga"`
				_Gid   string `json:"_gid"`
				_GclAu string `json:"_gcl_au"`
				_Clck  string `json:"_clck"`
			} `json:"cookies"`
			Clid struct {
				Fbclid string `json:"fbclid"`
				Gclid  string `json:"gclid"`
				Ttclid string `json:"ttclid"`
			} `json:"clid"`
		} `json:"browserData"`
		GdprData struct {
			CustomerConsent string          `json:"customerConsent"`
			GdprAllowed     bool            `json:"gdprAllowed"`
			GdprPlugin      string          `json:"gdprPlugin"`
			GdprConsent     GDPRConsentEnum `json:"gdprConsent"`
			GdprMode        string          `json:"gdprMode"`
		} `json:"gdprData"`
		CustomerData struct {
			ID          string `json:"id"`
			Email       string `json:"email"`
			Phone       string `json:"phone"`
			FirstName   string `json:"firstName"`
			LastName    string `json:"lastName"`
			DateOfBirth string `json:"dateOfBirth"`
			Address     struct {
				Street      string `json:"street"`
				Zip         string `json:"zip"`
				City        string `json:"city"`
				Country     string `json:"country"`
				CountryCode string `json:"countryCode"`
				Region      string `json:"region"`
				RegionCode  string `json:"regionCode"`
			} `json:"address"`
		} `json:"customerData"`
		IpLocationData struct {
			IP            string `json:"ip"`
			As            string `json:"as"`
			City          string `json:"city"`
			Continent     string `json:"continent"`
			ContinentCode string `json:"continentCode"`
			Country       string `json:"country"`
			CountryCode   string `json:"countryCode"`
			Currency      string `json:"currency"`
			District      string `json:"district"`
			Isp           string `json:"isp"`
			Proxy         bool   `json:"proxy"`
			Region        string `json:"region"`
			RegionName    string `json:"regionName"`
			Timezone      string `json:"timezone"`
			Zip           string `json:"zip"`
			Org           string `json:"org"`
		} `json:"ipLocationData"`
	} `json:"sessionData"`
	GaParams struct {
		Ga    string `json:"ga"`
		Gid   string `json:"gid"`
		GclAu string `json:"gclAu"`
	} `json:"gaParams"`
	MetaData struct {
		SourceID        int                 `json:"sourceID"`
		SchemaVersion   SchemaVersion       `json:"schemaVersion"`
		UserID          string              `json:"userID"`
		SourcePlatform  MableSourcePlatform `json:"sourcePlatform"`
		EventOriginType MableEventOrigin    `json:"eventOriginType"`
		SourceUID       string              `json:"sourceUID"`
	} `json:"metaData"`
	ContextData struct {
		Page struct {
			Host        string `json:"host"`
			ReferrerUrl string `json:"referrerUrl"`
			UrlPath     string `json:"urlPath"`
			UrlTitle    string `json:"urlTitle"`
		} `json:"page"`
		SessionLookupData struct {
			CartToken         string                    `json:"cartToken"`
			SessionIDFound    bool                      `json:"sessionIdFound"`
			SessionLookupKeys []SessionLookupKeys       `json:"sessionLookupKeys"`
			SessionSyncMethod ShopifySessionSyncMethods `json:"sessionSyncMethod"`
		} `json:"sessionLookupData"`
	} `json:"contextData"`
	EventTimeData struct {
		OriginalTimestamp int64 `json:"originalTimestamp"`
	} `json:"eventTimeData"`
	SourceCustomData map[string]interface{} `json:"SourceCustomData"`
	CustomData       map[string]interface{} `json:"customData"`
}

type MablePayload struct {
	DestinationSettings struct {
		PixelID       string `json:"pixelId"`
		TestEventCode string `json:"testEventCode"`
		AccessKey     string `json:"accessKey"`
	} `json:"destinationSettings"`
	Events []MableEvent `json:"events"`
}

// The various enum values are not given in the provided JSON
// Placeholder values added for completion, but you may need to replace them with actual enum values from the documentation
const (
	ALLOWED    GDPRConsentEnum           = "ALLOWED"
	V2         SchemaVersion             = "V2"
	SHOPIFY    MableSourcePlatform       = "SHOPIFY"
	SCRIPTTAG  MableEventOrigin          = "SCRIPTTAG"
	SESSION_ID SessionLookupKeys         = "SESSION_ID"
	CART_TOKEN ShopifySessionSyncMethods = "CART_TOKEN"
)
