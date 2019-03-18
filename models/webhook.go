package models

// MessagesWebhook para receber os eventos do workplace
type MessagesWebhook struct {
	Object string `json:"object"`
	Entry  []struct {
		ID        string `json:"id"`
		Time      int64  `json:"time"`
		Messaging []struct {
			Sender struct {
				ID string `json:"id"`
			} `json:"sender"`
			Recipient struct {
				ID string `json:"id"`
			} `json:"recipient"`
			Delivery struct {
				Mids      []string `json:"mids"`
				Watermark int64    `json:"watermark"`
				Seq       int      `json:"seq"`
			} `json:"delivery"`
			Read struct {
				Watermark int64 `json:"watermark"`
				Seq       int   `json:"seq"`
			} `json:"read"`
		} `json:"messaging"`
	} `json:"entry"`
}
