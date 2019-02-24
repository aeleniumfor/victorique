package container

import (
	"testing"
)

func TestContainer(t *testing.T) {
	s := ContainersNew()
	if s != nil {
		s.SetHostList([]string{"localhost"})
	}

}
