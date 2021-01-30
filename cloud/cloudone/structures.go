package cloudone

// TODO: I wonder if there is an option to just lower-case all JSON
// TODO: Actually, is this even needed since we're just reading data
type CloudOneInstance struct {
	Id             string `json:"ID"`
	TeamName       string `json:"TeamName"`
	Machine        string `json:"Machine"`
	IpAddress      string `json:"IPAddress"`
	DeployedRegion string `json:"DeployedRegion"`
	State          string `json:"State"`
}

type CloudOne []CloudOneInstance
