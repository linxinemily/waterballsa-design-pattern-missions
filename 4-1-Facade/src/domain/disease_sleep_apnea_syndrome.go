package domain

import "golang.org/x/exp/slices"

type SleepApneaSyndrome struct {
}

func NewSleepApneaSyndrome() Disease {
	return &SleepApneaSyndrome{}
}

func (s *SleepApneaSyndrome) match(patient Patient, symptoms []string) bool {
	return patient.Weight/(patient.Height*1/100)*(patient.Height*1/100) >= 26 && slices.Contains(symptoms, "snore")
}

func (s *SleepApneaSyndrome) getPrescription() *Prescription {
	return NewPrescription("打呼抑制劑", "睡眠呼吸中止症（專業學名：SleepApneaSyndrome）", "一捲膠帶", "睡覺時，撕下兩塊膠帶，將兩塊膠帶交錯黏在關閉的嘴巴上，就不會打呼了。")
}
