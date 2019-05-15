package common

// Container is 
type Container struct {
	Name         string
	ID           string
	NetNamespace string
}

// ResCotainerIDList is 
type ResCotainerIDList struct {
	ContainerID []string `json:"Container_ID_List"`
}

// ResCotainerNameList is 
type ResCotainerNameList struct {
	ContainerName []string `json:"Container_Name_List"`
}

// ReqContainerID is 
type ReqContainerID struct {
	ContainerID []string `json:"Container_ID"`
}