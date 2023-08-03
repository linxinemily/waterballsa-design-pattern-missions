package domain

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

const size = 1000

var (
	models Models
	mutex  sync.Mutex
)

type Models interface {
	CreateModel(modelName string) (Model, error)
	getData(modelName string) ([size][size]float32, error)
}

type concreteModels struct {
	data  map[string][size][size]float32
	mutex sync.Mutex
}

func newConcreteModels() *concreteModels {
	return &concreteModels{
		data: make(map[string][size][size]float32),
	}
}

func (m *concreteModels) CreateModel(modelName string) (Model, error) {
	data, err := m.getData(modelName)
	if err != nil {
		return nil, err
	}
	return newConcreteModel(data), nil
}

func (m *concreteModels) getData(modelName string) ([size][size]float32, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.data[modelName]; !ok {
		d, err := retrieveDataFromFile(modelName)
		if err != nil {
			return [size][size]float32{}, err
		}
		m.data[modelName] = d
	}

	return m.data[modelName], nil
}

func retrieveDataFromFile(modelName string) ([size][size]float32, error) {
	// 開啟檔案
	file, err := os.Open(modelName + ".mat")
	if err != nil {
		return [size][size]float32{}, err
	}
	defer file.Close()

	// 使用 bufio.Scanner 逐行讀取檔案內容
	scanner := bufio.NewScanner(file)
	var result [size][size]float32
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		// 將每行的字串以空格分割成字串切片
		numStrings := strings.Fields(line)
		var row [size]float32

		// 將字串切片中的元素轉換為浮點數並添加到 row
		for i, numString := range numStrings {
			num, err := strconv.ParseFloat(numString, 32)
			if err != nil {
				return [size][size]float32{}, err
			}
			row[i] = float32(num)
		}

		// 將 row 添加到結果變數中
		result[count] = row

		count += 1
	}

	if err := scanner.Err(); err != nil {
		return [size][size]float32{}, err
	}

	return result, nil
}

func GetModels() Models {
	mutex.Lock()
	defer mutex.Unlock()

	if models == nil {
		models = newConcreteModels()
	}
	return models
}
