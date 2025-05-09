package dto

type LocationInput struct {
	Cep string `json:"cep"`
}

type LocationOutput struct {
	Localidade string `json:"localidade"`

	Erro bool `json:"erro"`
}
