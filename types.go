package types

// Stream represents a single stream variant with quality information
type Stream struct {
	URL     string `json:"url"`     // The URL for this stream variant
	Quality string `json:"quality"` // Quality identifier (e.g., "1080p", "720p", "best", "worst")
}

// Streamer represents a streamer being monitored for live streams
type Streamer struct {
	Platform    string   `json:"platform"`
	Username    string   `json:"username"`
	LastChecked int64    `json:"last_checked"` // Unix timestamp
	Checking    bool     `json:"checking"`
	Paused      bool     `json:"paused"`
	AutoRecord  bool     `json:"auto_record"`
	IsLive      bool     `json:"is_live"`
	LastLive    string   `json:"last_live"` // ISO 8601 string or "unknown"
	VODs        int      `json:"vods"`      // Number of recorded VODs
	VodPath     string   `json:"vod_path"`  // Path to the folder where recordings are saved
	StreamUrls  []Stream `json:"stream_urls"`
	IsPrivate   bool     `json:"is_private"`
}

type StreamerId struct {
	Platform string `json:"platform"`
	Username string `json:"username"`
}
