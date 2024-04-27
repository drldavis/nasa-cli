package models

type SolarFlare struct {
	BeginTime string `json:"beginTime"`
	PeakTime  string `json:"peakTime"`
	EndTime   string `json:"endTime"`
	Note      string `json:"note"`
}
