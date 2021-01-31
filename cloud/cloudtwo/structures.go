package cloudtwo

type CloudTwoInstance struct {
	InstanceId    string `json:"instance_id"`
	Team           string `json:"team"`
	InstanceType  string `json:"instance_type"`
	IpAddress     string `json:"ip_address"`
	Region         string `json:"region"`
	InstanceState string `json:"instance_state"`
}

type CloudTwo struct {
	Total     int                `json:"total"`
	Count     int                `json:"Count""`
	Instances []CloudTwoInstance `json:"instances"`
}
