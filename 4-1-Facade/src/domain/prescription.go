package domain

type Prescription struct {
	Name             string `json:"name"`
	PotentialDisease string `json:"potential_disease"`
	Medicines        string `json:"medicines"`
	Usage            string `json:"usage"`
}

func NewPrescription(name, potentialDisease, medicines, usage string) *Prescription {
	return &Prescription{name, potentialDisease, medicines, usage}
}
