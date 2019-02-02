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

// GetMultiHostContainer is docker ps s
func (c *Containers) GetMultiHostContainer() {
	hostname := c.host
	for i := 0; i < len(hostname); i++ {
		s := New(hostname[i])
		fmt.Println(s)
	}
	finished := make(chan bool)
	for i := 0; i < len(hostname); i++ {
		go func(i int) {
			fmt.Println(i)
			finished <- true
		}(i)
	}

	for i := 0; i < len(hostname); i++ {
		fmt.Println(<-finished)
	}

	fmt.Println("finished")

}

// 	s.CreateContainer()
// 	s.RunContainer()
// 	s.GetContainerInspect()
// 	s.StopContainer()
