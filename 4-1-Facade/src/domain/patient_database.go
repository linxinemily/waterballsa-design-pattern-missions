package domain

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"
)

type PatientDatabase struct {
	patients []Patient
}

func NewPatientDatabase() *PatientDatabase {
	return &PatientDatabase{}
}

func (db *PatientDatabase) saveCase(patient Patient, prescriptions []*Prescription, symptoms []string, filename string, filetype string) {
	if _, err := os.Stat("data/" + patient.Id); os.IsNotExist(err) {
		err := os.Mkdir("data/"+patient.Id, 0750)
		if err != nil {
			panic(err)
		}
	}
	filename = "data/" + patient.Id + "/" + filename + "." + filetype
	_case := &Case{
		time.Now().String(),
		symptoms,
		prescriptions,
	}

	switch filetype {
	case "json":
		file, err := json.MarshalIndent(_case, "", " ")
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(filename, file, 0666)
		if err != nil {
			panic(err)
		}
	case "csv":
		csvFile, err := os.Create(filename)

		if err != nil {
			panic(fmt.Sprintf("failed creating file: %s\n", err))
		}

		data := make([][]string, len(prescriptions)+1)

		rowCount := 0
		t := reflect.TypeOf(*prescriptions[0])
		for i := 0; i < t.NumField(); i++ {
			data[rowCount] = append(data[rowCount], t.Field(i).Name)
		}
		rowCount += 1

		for _, prescription := range prescriptions {
			r := reflect.ValueOf(*prescription)
			for j := 0; j < t.NumField(); j++ {
				data[rowCount] = append(data[rowCount], reflect.Indirect(r).FieldByName(t.Field(j).Name).String())
			}
			rowCount += 1
		}

		csvwriter := csv.NewWriter(csvFile)

		for _, row := range data {
			_ = csvwriter.Write(row)
		}

		csvwriter.Flush()
	}
}

func (db *PatientDatabase) find(patientId string) Patient {
	var patient Patient
	for _, p := range db.patients {
		if p.Id == patientId {
			patient = p
			break
		}
	}

	return patient
}

func (db *PatientDatabase) loadPatientsData(filename string) {
	fileContent, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		panic("parse file error")
	}

	var data []Patient

	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		fmt.Println(err)
		panic("unmarshal json error")
	}

	db.patients = data
}
