package reporter

type MachineInstance struct {
	Id      string `json:"id"`
	Team    string `json:"team"`
	Machine string `json:"machine"`
	Ip      string `json:"ip"`
	State   string `json:"state"`
	Region  string `json:"region"`
}

type TeamInstances struct {
	Count int `json: "count"`
	Instances []*MachineInstance `json: "instances"`
}

type MachineReport map[string]*TeamInstances
