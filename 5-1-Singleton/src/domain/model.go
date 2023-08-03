package domain

type Model interface {
	Count(input [size]float32) [size]float32
}

type concreteModel struct {
	data [size][size]float32
}

func (c concreteModel) Count(input [size]float32) [size]float32 {
	var output [size]float32

	// 線性轉換的函數：output[i] = c.data[0][i] * input[0] + c.data[1][i] * input[1] + ... + c.data[999][i] * input[999]
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			output[i] += c.data[j][i] * input[j]
		}
	}

	return output
}

func newConcreteModel(data [size][size]float32) *concreteModel {
	return &concreteModel{
		data: data,
	}
}
