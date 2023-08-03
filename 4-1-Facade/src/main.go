package main

import "C4M1H1-facade/domain"

func main() {
	facade := domain.NewPrescriberFacade()
	facade.LoadPatientsData("data/patients.json")
	facade.LoadSupportedDisease("data/supported_disease.txt")

	go func() {
		facade.Prescribe("A103456789", []string{"snore"}, "test", "csv")
		facade.Prescribe("A123456789", []string{"sneeze", "headache", "cough"}, "test", "json")
		facade.Prescribe("B223456789", []string{"sneeze", "headache", "cough"}, "test", "cs")
	}()

	facade.Run()
}
