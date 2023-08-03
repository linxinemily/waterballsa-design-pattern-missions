package domain

type PrescribeJob struct {
	patient Patient
	handle  func()
}
