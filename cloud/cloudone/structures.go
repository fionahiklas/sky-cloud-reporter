package cloudone

type CloudOneInstance struct {
	Id             string `json:"ID"`
	TeamName       string `json:"TeamName"`
	Machine        string `json:"Machine"`
	IpAddress      string `json:"IPAddress"`
	DeployedRegion string `json:"DeployedRegion"`
	State          string `json:"State"`
}

type CloudOne []CloudOneInstance
