package zkoperator

import (
	"testing"
)

func TestGetZKClusterInfo(t *testing.T) {
	err := GetZKClusterInfo("zookeeper")
	if err != nil {
		t.Error(err)
	}
}
