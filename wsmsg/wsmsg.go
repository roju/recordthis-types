// Websocket messages sent between client and server
package wsmsg

// MessageType represents the type of websocket message
type MessageType string

const (
	TypeSubscriptions  MessageType = "subscriptions"
	TypeReauthRequest  MessageType = "reauth_request"
	TypeReauthResponse MessageType = "reauth_response"
	TypeLiveStatus     MessageType = "live_status"
)

// Subscriptions is sent by client to subscribe to streamers. Example:
//
//	{
//		"type": "subscriptions",
//		"streamers": {
//			"platformA": [
//				"streamer1",
//				"streamer2"
//			],
//			"platformB": [
//				"streamer3",
//				"streamer4"
//			]
//		}
//	}
type Subscriptions struct {
	Type      MessageType         `json:"type"`      // "subscriptions"
	Streamers map[string][]string `json:"streamers"` // Platform -> list of usernames
}

// ReauthRequest is sent by client to reauthenticate with a new access token. Example:
//
//	{
//		"type": "reauth_request",
//		"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
//	}
type ReauthRequest struct {
	Type        MessageType `json:"type"`         // "reauth_request"
	AccessToken string      `json:"access_token"` // New access token
}

// ReauthResponse is sent by server after reauth attempt. Example:
//
//	{
//		"type": "reauth_response",
//		"success": true,
//		"message": "Reauthentication successful"
//	}
type ReauthResponse struct {
	Type    MessageType `json:"type"`    // "reauth_response"
	Success bool        `json:"success"` // Whether reauth succeeded
	Message string      `json:"message"` // Error message if failed
}

// LiveStatus is sent by server to notify clients of streamer live status. Example:
//
//	{
//		"type": "live_status",
//		"is_live": {
//			"platformA": {
//				"streamer1": true,
//				"streamer2": false
//			},
//			"platformB": {
//				"streamer3": false,
//				"streamer4": false
//			}
//		}
//	}
type LiveStatus struct {
	Type   MessageType                `json:"type"`    // "live_status"
	IsLive map[string]map[string]bool `json:"is_live"` // Platform -> Username -> IsLive
}
