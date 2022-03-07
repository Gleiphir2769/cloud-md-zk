package zkoperator

import (
	"cloud-md-zk/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ZKClusterStatus struct {
}

func CreateZKCluster(clusterName string, config *ZKClusterConfig) (*ZKClusterStatus, error) {
	return &ZKClusterStatus{}, nil
}

func DeleteZKCluster(clusterName string) error {
	return nil
}

func GetZKClusterInfo(clusterName string) error {
	resp, err := util.HTTPGet(fmt.Sprintf("http://localhost:8081/apis/zookeeper.pravega.io/v1beta1/namespaces/default/zookeeperclusters/%s", clusterName), nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var info *ZKClusterPodInfo
	err = json.Unmarshal(buf, info)
	if err != nil {
		return err
	}
	fmt.Println(info)
	return nil
}

func UpdateZKCluster(clusterName string, config *ZKClusterConfig) (*ZKClusterStatus, error) {
	return &ZKClusterStatus{}, nil
}
