package model

//Facture : la structure de nos factures
type Facture struct {
	Username string `json:"Username"`
	Adress   string `json:"Adress"`
	Birthday string `json:"Birthday"`
	Work     string `json:"Work"`
	Price    int    `json:"Price"`
}
