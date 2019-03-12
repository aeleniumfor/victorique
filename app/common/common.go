package common

// Containers is managed
type Containers struct {
	Host      string
	Container []Container // Container List
}

// Container is managed
type Container struct {
	Name string
	ID   string
	IP   string
}
