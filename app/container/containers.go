package container

import "fmt"

// Containers type struct
type Containers struct {
	hostList []string
}

// HostContainer type struct
type HostContainer struct {
	host          string
	containerName []string
}

// ContainersNew is
func ContainersNew() (c *Containers) {
	return &Containers{}
}

// SetHostList is
func (c *Containers) SetHostList(hostList []string) {
	c.hostList = hostList
	fmt.Println(c.hostList)
}

// GetMultiHostContainerList is docker ps s
func (c *Containers) GetMultiHostContainerList() (ContainerList []HostContainer) {
	containerList := make(chan HostContainer)
	// 並列化する処理
	for i := 0; i < len(c.hostList); i++ {
		go func(i int) {
			hostContainerList := HostContainer{}
			s := New(c.hostList[i])
			hostContainerList.host = c.hostList[i]
			hostContainerList.containerName = s.ListContainer()
			containerList <- hostContainerList
		}(i)
	}

	// 構造体のスライスを作成する
	HostList := []HostContainer{}
	// 並列化したものをこっちにもってくる処理
	for i := 0; i < len(c.hostList); i++ {
		// TODO: これの処理は果たしてベストなのか?
		HostList = append(HostList, <-containerList)
	}
	fmt.Println(HostList)
	return HostList
}

// 	s.CreateContainer()
// 	s.RunContainer()
// 	s.GetContainerInspect()
// 	s.StopContainer()
