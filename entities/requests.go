package entities

// Field names should start with an uppercase letter
type IncomingRequest struct {
	DataIdReq    string `query:"indicator"`
	SubtitlesReq bool   `query:"subs"`
	AudioOnly    bool   `query:"ao"`
}
