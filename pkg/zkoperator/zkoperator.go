package zkoperator

import (
	"cloud-md-zk/lib/logger"
	"cloud-md-zk/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

const api = "http://localhost:8081/apis/zookeeper.pravega.io/v1beta1/namespaces/default/zookeeperclusters"

func CreateZKCluster(clusterName string, config *ZKClusterConfig) (*ZKClusterPodStatus, error) {
	cr := MakeCR(clusterName, config)
	resp, err := util.HTTPPost(api, &cr, nil, nil)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.GetLog().Error(err.Error())
		}
	}(resp.Body)
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("craete zk cluster failed, status code: %d, reason: %s", resp.StatusCode, string(buf))
	}

	info := &ZKClusterPodStatus{}
	err = json.Unmarshal(buf, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func CreateDefaultZKCluster(clusterName string) (*ZKClusterPodStatus, error) {
	cr := DefaultCR(clusterName)
	resp, err := util.HTTPPost(api, &cr, nil, nil)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.GetLog().Error(err.Error())
		}
	}(resp.Body)
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("craete zk cluster failed, status code: %d, reason: %s", resp.StatusCode, string(buf))
	}

	info := &ZKClusterPodStatus{}
	err = json.Unmarshal(buf, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func DeleteZKCluster(clusterName string) error {
	resp, err := util.HTTPDelete(urlJoin(api, clusterName), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		return nil
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("delete zk cluster failed, status code: %d, reason: %s", resp.StatusCode, err)
	}
	return fmt.Errorf("delete zk cluster failed, status code: %d, reason: %s", resp.StatusCode, string(buf))
}

func GetZKClusterInfo(clusterName string) (*ZKClusterPodStatus, error) {
	resp, err := util.HTTPGet(urlJoin(api, clusterName), nil, nil)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.GetLog().Error(err.Error())
		}
	}(resp.Body)
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	info := &ZKClusterPodStatus{}
	err = json.Unmarshal(buf, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func UpdateZKCluster(clusterName string, config *ZKClusterConfig) (*ZKClusterPodStatus, error) {
	patches := MakePatches(config)
	resp, err := util.HTTPPatch(urlJoin(api, clusterName), &patches, nil, nil)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.GetLog().Error(err.Error())
		}
	}(resp.Body)
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("craete zk cluster failed, status code: %d, reason: %s", resp.StatusCode, string(buf))
	}

	info := &ZKClusterPodStatus{}
	err = json.Unmarshal(buf, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func urlJoin(segments ...string) (result string) {
	for i, segment := range segments {
		if i == 0 {
			result = segment
		} else {
			result += "/" + segment
		}
	}
	return result
}
