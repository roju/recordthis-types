package types

// SubscriptionsMessage is sent by client to subscribe to streamers
// Example:
// {
// 	"type": "subscriptions",
// 	"streamers": {
// 		"platformA": [
// 			"streamer1",
// 			"streamer2"
// 		],
// 		"platformB": [
// 			"streamer3",
// 			"streamer4"
// 		]
// 	}
// }
type SubscriptionsMessage struct {
	Type      string              `json:"type"`      // "subscriptions"
	Streamers map[string][]string `json:"streamers"` // Platform -> list of usernames
}

// ReauthMessage is sent by client to reauthenticate with a new access token
// Example:
// {
// 	"type": "reauth",
// 	"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
// }
type ReauthMessage struct {
	Type        string `json:"type"`         // "reauth"
	AccessToken string `json:"access_token"` // New access token
}

// ReauthResponseMessage is sent by server after reauth attempt
// Example:
// {
// 	"type": "reauth_response",
// 	"success": true,
// 	"message": "Reauthentication successful"
// }
type ReauthResponseMessage struct {
	Type    string `json:"type"`    // "reauth_response"
	Success bool   `json:"success"` // Whether reauth succeeded
	Message string `json:"message"` // Error message if failed
}

// LiveStatusMessage is sent by server to notify clients of streamer live status
// Example:
// {
// 	"type": "live_status",
// 	"is_live": {
// 		"platformA": {
// 			"streamer1": true,
// 			"streamer2": false
// 		},
// 		"platformB": {
// 			"streamer3": false,
// 			"streamer4": false
// 		}
// 	}
// }
type LiveStatusMessage struct {
	Type   string                     `json:"type"`    // "live_status"
	IsLive map[string]map[string]bool `json:"is_live"` // Platform -> Username -> IsLive
}
