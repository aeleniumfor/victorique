package container

import (
	"testing"
)

func TestNew(t *testing.T) {
	New()
}

func TestGetContainer(t *testing.T) {
	c := New()
	c.GetContainerList()
}
