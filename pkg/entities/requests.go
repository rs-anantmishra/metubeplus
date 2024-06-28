package entities

// Field names should start with an uppercase letter
type IncomingRequest struct {
	DataIdReq    string `json:"indicator" form:"indicator"`
	SubtitlesReq bool   `json:"subs" form:"subs"`
	AudioOnly    bool   `json:"ao" form:"ao"`
}
