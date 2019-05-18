package common

// Container is
type Container struct {
	Name         string `json:"Container_Name"`
	ID           string `json:"Container_ID"`
	NetNamespace string `json:"Container_NetNamespace"`
	Status       string `json:"Container_Status"`
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

// ResContainerDelMsg is delete container message
type ResContainerDelMsg struct {
	Msg string `json:"Container_Delete_Msg"`
}
