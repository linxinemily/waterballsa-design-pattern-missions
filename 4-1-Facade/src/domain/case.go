package domain

type Case struct {
	CaseTime      string          `json:"case_time"`
	Symptoms      []string        `json:"symptoms"`
	Prescriptions []*Prescription `json:"prescriptions"`
}
