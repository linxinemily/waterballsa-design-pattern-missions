package domain

type Patient struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Age    int     `json:"age"`
	Height float32 `json:"height"`
	Weight float32 `json:"weight"`
}
