package domain

import "golang.org/x/exp/slices"

type Attractive struct {
}

func NewAttractive() Disease {
	return &Attractive{}
}

func (a *Attractive) match(patient Patient, symptoms []string) bool {
	return patient.Age == 18 && patient.Gender == "female" && slices.Contains(symptoms, "sneeze")
}

func (a *Attractive) getPrescription() *Prescription {
	return NewPrescription("青春抑制劑", "有人想你了 (專業學名：Attractive)", "假鬢角、臭味", "把假鬢角黏在臉的兩側，讓自己異性緣差一點，自然就不會有人想妳了。")
}
