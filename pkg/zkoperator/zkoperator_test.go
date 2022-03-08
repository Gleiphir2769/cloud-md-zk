package zkoperator

import (
	"fmt"
	"testing"
)

func TestGetZKClusterInfo(t *testing.T) {
	//err := GetZKClusterInfo("zookeeper")
	//if err != nil {
	//	t.Error(err)
	//}
}

func TestUrlJoin(t *testing.T) {
	fmt.Println(urlJoin("http://www.baidu.com", "shenjiaqi", "test"))
}

func TestDeleteZKCluster(t *testing.T) {
	err := DeleteZKCluster("zookeeper-the-other")
	if err != nil {
		t.Error(err)
	}
}
