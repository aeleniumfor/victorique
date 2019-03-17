package common

// Containers is managed
type Containers struct {
	Host      string
	ContainerStructs []Container // Container List
}

// Container is managed
type Container struct {
	Name string
	ID   string
	IP   string
}


