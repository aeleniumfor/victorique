package container

import "fmt"

// Containers type struct
type Containers struct {
	hostList  []string
	host      string
	container []Container
}

// ContainersNew is
func ContainersNew() (c *Containers) {
	return &Containers{}
}

// SetHostList is
func (c *Containers) SetHostList(hostList []string) {
	c.hostList = hostList
	fmt.Println(c.host)
}

// GetMultiHostContainerList is docker ps s
func (c *Containers) GetMultiHostContainerList() {
	containerList := make(chan []string)
	// 並列化する処理
	for i := 0; i < len(c.hostList); i++ {
		go func(i int) {
			fmt.Println(i)
			s := New(c.hostList[i])
			fmt.Println(s)
			containerList <- s.ListContainer()
		}(i)
	}

	// 並列化したものをこっちにもってくる処理
	for i := 0; i < len(c.hostList); i++ {
		fmt.Println(<-containerList)
	}

	fmt.Println("finished")

}

// 	s.CreateContainer()
// 	s.RunContainer()
// 	s.GetContainerInspect()
// 	s.StopContainer()
