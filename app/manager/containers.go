package manager

import (
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Containers is managed
type Containers struct {
	host      string
	client    *client.Client
	Container []Container
}

// Container is managed
type Container struct {
	name    string
	id      string
	inspect types.ContainerJSON
}

// New is manager
func New() (c *Containers) {
	log.Println(1)
	tr := &http.Transport{}
	hostname := "http://localhost:2375"
	cli, err := client.NewClient(hostname, client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		log.Println(err)
	}
	return &Containers{client: cli}
}

// GetContainerList is multi host get container list
func (c *Containers)GetContainerList(){
	
}

// SetHostList is
// func (c *Containers) SetHostList(hostList []string) {
// 	c.hostList = hostList
// 	fmt.Println(c.hostList)
// }

// // GetMultiHostContainerList is docker ps s
// func (c *Containers) GetMultiHostContainerList() (ContainerList []HostContainer) {
// 	containerList := make(chan HostContainer)
// 	// 並列化する処理
// 	for i := 0; i < len(c.hostList); i++ {
// 		go func(i int) {
// 			hostContainerList := HostContainer{}
// 			s := New(c.hostList[i])
// 			hostContainerList.host = c.hostList[i]
// 			hostContainerList.containerName = s.ListContainer()
// 			containerList <- hostContainerList
// 		}(i)
// 	}

// 	// 構造体のスライスを作成する
// 	HostList := []HostContainer{}
// 	// 並列化したものをこっちにもってくる処理
// 	for i := 0; i < len(c.hostList); i++ {
// 		// TODO: これの処理は果たしてベストなのか?
// 		HostList = append(HostList, <-containerList)
// 	}
// 	fmt.Println(HostList)
// 	c.hostContainers = HostList
// 	return HostList
// }

// // DecisionPriority is docker create
// func (c *Containers) DecisionPriority() {
// 	// TODO 優先順位をどうやって決めるか考え中
// 	// for i := 0; i < len(c.hostContainers); i++ {
// 	// 	list := c.hostContainers[i].containerName
// 	// }
// 	c.hostContainers = c.hostContainers
// 	fmt.Println(c.hostContainers)
// }

// 	s.CreateContainer()
// 	s.RunContainer()
// 	s.GetContainerInspect()
// 	s.StopContainer()
