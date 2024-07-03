package protocol

type Protocol struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func StatusOk(data interface{}) Protocol {
	return Protocol{
		Status: "ok",
		Data:   data,
	}
}
