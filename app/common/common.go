package common

// Container is managed
type Container struct {
	Name         string
	ID           string
	IP           string
	NetNamespace string
}

// ResCotainerIDList is controll network
type ResCotainerIDList struct {
	ContainerID []string `json:"Container_ID_List"`
}
