package domain

type Disease interface {
	match(patient Patient, symptoms []string) bool
	getPrescription() *Prescription
}
