package datamodels

type Service struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Versions    []string `json:"versions"`
}
