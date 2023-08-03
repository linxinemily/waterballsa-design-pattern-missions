package domain

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Prescriber struct {
	supportedDisease []Disease
	queue            chan *PrescribeJob
}

func NewPrescriber() *Prescriber {
	p := &Prescriber{}
	p.queue = make(chan *PrescribeJob)
	return p
}

func (p *Prescriber) run() {
	for {
		select {
		case job := <-p.queue:
			fmt.Printf("do prescribing, patient name: %s\n", job.patient.Name)
			job.handle()
		default:
			// do nothing
		}
	}
}

func (p *Prescriber) prescribe(patient Patient, symptoms []string, afterPrescribe func([]*Prescription)) {

	job := &PrescribeJob{
		patient: patient,
		handle: func() {
			// 診斷需耗時 3 秒
			time.Sleep(time.Second * 3)
			// 找到 match 的 Disease 並得到診斷結果
			var prescriptions []*Prescription
			for _, disease := range p.supportedDisease {
				if disease.match(patient, symptoms) {
					prescriptions = append(prescriptions, disease.getPrescription())
				}
			}

			if len(prescriptions) > 0 {
				for _, prescription := range prescriptions {
					fmt.Printf("disease: %s\n", prescription.Name)
				}
			}
			afterPrescribe(prescriptions)
			fmt.Println("prescribe done.")
		},
	}
	p.queue <- job

}

func (p *Prescriber) loadSupportedDisease(filename string) {
	content, err := os.Open(filename)

	if err != nil {
		panic("parse file error")
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)

	//var diseaseNames []string
	diseaseFactory := map[string]func() Disease{
		"COVID-19":           NewCOVID19,
		"Attractive":         NewAttractive,
		"SleepApneaSyndrome": NewSleepApneaSyndrome,
	}

	for scanner.Scan() {
		p.supportedDisease = append(p.supportedDisease, diseaseFactory[scanner.Text()]())
	}

	if err := scanner.Err(); err != nil {
		panic("scan error")
	}
}
