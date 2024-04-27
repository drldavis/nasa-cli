package models

type Sol struct {
	Temperature Details `json:"AT"`
	WindSpeed   Details `json:"HWS"`
	Pressure    Details `json:"PRE"`
	Season      string  `json:"Season"`
}

type Details struct {
	Average float32 `json:"av"`
}
