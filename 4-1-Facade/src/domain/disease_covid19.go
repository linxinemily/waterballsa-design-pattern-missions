package domain

import "golang.org/x/exp/slices"

type COVID19 struct {
}

func NewCOVID19() Disease {
	return &COVID19{}
}

func (c *COVID19) match(patient Patient, symptoms []string) bool {
	return slices.Contains(symptoms, "sneeze") && slices.Contains(symptoms, "headache") && slices.Contains(symptoms, "cough")
}

func (c *COVID19) getPrescription() *Prescription {
	return NewPrescription("清冠一號", "新冠肺炎（專業學名：COVID-19）", "清冠一號", "將相關藥材裝入茶包裡，使用500 mL 溫、熱水沖泡悶煮1~3 分鐘後即可飲用。")
}
