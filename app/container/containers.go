package container

import "fmt"

// Containers type struct
type Containers struct {
	host      []string
	container Container
}

// ContainersNew is
func ContainersNew() (c *Containers) {
	return &Containers{}
}

// SetHostList is
func (c *Containers) SetHostList(host []string) {
	c.host = host
	fmt.Println(c.host)
}

// GetMultiHostContainerlist is docker ps s
func (c *Containers) GetMultiHostContainerList() {
	hostname := c.host
	containerList := make(chan []string)
	for i := 0; i < len(hostname); i++ {
		go func(i int) {
			fmt.Println(i)
			s := New(hostname[i])
			fmt.Println(s)
			containerList <- s.ListContainer()
		}(i)
	}

	for i := 0; i < len(hostname); i++ {
		fmt.Println(<-containerList)
	}

	fmt.Println("finished")

}

// 	s.CreateContainer()
// 	s.RunContainer()
// 	s.GetContainerInspect()
// 	s.StopContainer()
