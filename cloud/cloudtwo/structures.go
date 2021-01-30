package cloudtwo

// TODO: I wonder if there is an option to just lower-case all JSON
// TODO: Actually, is this even needed since we're just reading data
type CloudTwoInstance struct {
	Instance_id    string `json:"instance_id"`
	Team           string `json:"team"`
	Instance_type  string `json:"instance_type"`
	Ip_address     string `json:"ip_address"`
	Region         string `json:"region"`
	Instance_state string `json:"instance_state"`
}

type CloudTwo struct {
	Total     int                `json:"total"`
	Count     int                `json:"Count""`
	Instances []CloudTwoInstance `json:"instances"`
}
