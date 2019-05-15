package common

// Container is 
type Container struct {
	Name         string `json:"Container_Name"`
	ID           string `json:"Container_ID"`
	NetNamespace string `json:"Container_NetNamespace"`
}

// ResCotainerIDList is 
type ResCotainerIDList struct {
	ContainerIDList []string `json:"Container_ID_List"`
}

// ResCotainerNameList is 
type ResCotainerNameList struct {
	ContainerNameList []string `json:"Container_Name_List"`
}

// ReqContainerID is 
type ReqContainerID struct {
	ContainerID []string `json:"Container_ID"`
}