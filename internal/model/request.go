package model

type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type TopSecretRequest struct {
	Satellites []Satellite `json:"satellites"`
}
