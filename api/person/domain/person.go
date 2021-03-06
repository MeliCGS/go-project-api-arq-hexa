package person_domain

type Person struct {
	ID        string     `json:"id"`
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Edad      int        `json:"edad"`
	Direccion *Direccion `json:"direccion"`
	Text      string     `json:"text"`
	Done      bool       `json:"done"`
}

type Direccion struct {
	Ciudad    string `json:"ciudad"`
	Municipio string `json:"municipio"`
}
