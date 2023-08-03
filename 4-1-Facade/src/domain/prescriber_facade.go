package domain

type PrescriberFacade struct {
	prescriber *Prescriber
	patientDB  *PatientDatabase
}

func NewPrescriberFacade() *PrescriberFacade {
	facade := &PrescriberFacade{}
	facade.patientDB = NewPatientDatabase()
	facade.prescriber = NewPrescriber()
	return facade
}

func (p *PrescriberFacade) Run() {
	p.prescriber.run()
}

func (p *PrescriberFacade) LoadPatientsData(filename string) {
	p.patientDB.loadPatientsData(filename)
}

func (p *PrescriberFacade) LoadSupportedDisease(filename string) {
	p.prescriber.loadSupportedDisease(filename)
}

func (p *PrescriberFacade) Prescribe(patientId string, symptoms []string, exportFilename string, filetype string) {
	patient := p.patientDB.find(patientId)

	p.prescriber.prescribe(patient, symptoms, func(prescriptions []*Prescription) {
		p.patientDB.saveCase(patient, prescriptions, symptoms, exportFilename, filetype)
	})
}
