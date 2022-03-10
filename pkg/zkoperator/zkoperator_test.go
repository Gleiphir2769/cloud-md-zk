package zkoperator

import (
	"fmt"
	"testing"
)

//func TestGetZKClusterInfo(t *testing.T) {
//	//err := GetZKClusterInfo("zookeeper")
//	//if err != nil {
//	//	t.Error(err)
//	//}
//}
//
//func TestUrlJoin(t *testing.T) {
//	fmt.Println(urlJoin("http://www.baidu.com", "shenjiaqi", "test"))
//}

//func TestCreateDefaultZKCluster(t *testing.T) {
//	info, err := CreateDefaultZKCluster("zookeeper-test")
//	if err != nil {
//		t.Error(err)
//	} else {
//		fmt.Println(info.Spec.Config)
//	}
//	//buf, _ := json.Marshal(DefaultCR("zookeeper-test"))
//	//fmt.Println(string(buf))
//}

func TestMakePatches(t *testing.T) {
	info, err := UpdateZKCluster("zookeeper-test", TestConfig())
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(info.Spec.Config)
	}
}

//func TestDeleteZKCluster(t *testing.T) {
//	err := DeleteZKCluster("zookeeper-the-other")
//	if err != nil {
//		t.Error(err)
//	}
//}
