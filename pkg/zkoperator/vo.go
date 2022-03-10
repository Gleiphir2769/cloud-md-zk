package zkoperator

import (
	"fmt"
	"github.com/fatih/structs"
	"time"
)

type ZKClusterConfig struct {
	Size   int                `json:"size"`
	CPU    string             `json:"cpu"`
	Memory string             `json:"memory"`
	Volume string             `json:"volume"`
	Config ZKClusterSubConfig `json:"config"`
}

type ZKClusterSubConfig struct {
	AutoPurgePurgeInterval   int `json:"autoPurgePurgeInterval"`
	AutoPurgeSnapRetainCount int `json:"autoPurgeSnapRetainCount"`
	CommitLogCount           int `json:"commitLogCount"`
	GlobalOutstandingLimit   int `json:"globalOutstandingLimit"`
	InitLimit                int `json:"initLimit"`
	MaxClientCnxns           int `json:"maxClientCnxns"`
	MaxSessionTimeout        int `json:"maxSessionTimeout"`
	MinSessionTimeout        int `json:"minSessionTimeout"`
	PreAllocSize             int `json:"preAllocSize"`
	SnapCount                int `json:"snapCount"`
	SnapSizeLimitInKb        int `json:"snapSizeLimitInKb"`
	SyncLimit                int `json:"syncLimit"`
	TickTime                 int `json:"tickTime"`
}

type ZKClusterPodStatus struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Generation        int       `json:"generation"`
		ManagedFields     []struct {
			Manager   string    `json:"manager"`
			Operation string    `json:"operation"`
			Time      time.Time `json:"time"`
		} `json:"managedFields"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
		UID             string `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		AdminServerService struct {
		} `json:"adminServerService"`
		ClientService struct {
		} `json:"clientService"`
		Config          ZKClusterSubConfig `json:"config"`
		HeadlessService struct {
		} `json:"headlessService"`
		Image struct {
			PullPolicy string `json:"pullPolicy"`
			Repository string `json:"repository"`
			Tag        string `json:"tag"`
		} `json:"image"`
		Labels struct {
			App     string `json:"app"`
			Release string `json:"release"`
		} `json:"labels"`
		Persistence struct {
			ReclaimPolicy string `json:"reclaimPolicy"`
			Spec          struct {
				AccessModes []string `json:"accessModes"`
				Resources   struct {
					Requests struct {
						Storage string `json:"storage"`
					} `json:"requests"`
				} `json:"resources"`
			} `json:"spec"`
		} `json:"persistence"`
		Ports []struct {
			ContainerPort int    `json:"containerPort"`
			Name          string `json:"name"`
		} `json:"ports"`
		Probes struct {
			LivenessProbe struct {
				FailureThreshold    int `json:"failureThreshold"`
				InitialDelaySeconds int `json:"initialDelaySeconds"`
				PeriodSeconds       int `json:"periodSeconds"`
				SuccessThreshold    int `json:"successThreshold"`
				TimeoutSeconds      int `json:"timeoutSeconds"`
			} `json:"livenessProbe"`
			ReadinessProbe struct {
				FailureThreshold    int `json:"failureThreshold"`
				InitialDelaySeconds int `json:"initialDelaySeconds"`
				PeriodSeconds       int `json:"periodSeconds"`
				SuccessThreshold    int `json:"successThreshold"`
				TimeoutSeconds      int `json:"timeoutSeconds"`
			} `json:"readinessProbe"`
		} `json:"probes"`
		Replicas    int    `json:"replicas"`
		StorageType string `json:"storageType"`
	} `json:"spec"`
	Status struct {
		Conditions []struct {
			LastTransitionTime time.Time `json:"lastTransitionTime,omitempty"`
			LastUpdateTime     time.Time `json:"lastUpdateTime,omitempty"`
			Status             string    `json:"status"`
			Type               string    `json:"type"`
		} `json:"conditions"`
		CurrentVersion         string `json:"currentVersion"`
		ExternalClientEndpoint string `json:"externalClientEndpoint"`
		InternalClientEndpoint string `json:"internalClientEndpoint"`
		Members                struct {
			Ready   []string `json:"ready"`
			Unready []string `json:"unready"`
		} `json:"members"`
		MetaRootCreated bool `json:"metaRootCreated"`
		ReadyReplicas   int  `json:"readyReplicas"`
		Replicas        int  `json:"replicas"`
	} `json:"status"`
}

type ZKClusterCR struct {
	APIVersion string              `json:"apiVersion"`
	Kind       string              `json:"kind"`
	Metadata   ZKClusterCRMetadata `json:"metadata"`
	Spec       ZKClusterCRSpec     `json:"spec"`
}

type ZKClusterCRMetadata struct {
	Name string `json:"name"`
}

type ZKClusterCRSpec struct {
	Replicas    int                `json:"replicas"`
	Image       ImageConfig        `json:"image"`
	StorageType string             `json:"storageType"`
	Persistence PersistentConfig   `json:"persistence"`
	Pod         PodConfig          `json:"pod"`
	Config      ZKClusterSubConfig `json:"config"`
}

type ImageConfig struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	// todo: wait to test
	PullPolicy string `json:"pullPolicy"`
}

type PersistentConfig struct {
	ReclaimPolicy string         `json:"reclaimPolicy"`
	Spec          PersistentSpec `json:"spec"`
}

type PersistentSpec struct {
	StorageClassName string              `json:"storageClassName"`
	Resources        PersistentResources `json:"resources"`
}

type PersistentResources struct {
	Requests PersistentResourceRequest `json:"requests"`
	//Limits   PersistentResourceLimit   `yaml:"limits"`
}

type PersistentResourceRequest struct {
	Storage string `json:"storage"`
}

//type PersistentResourceLimit struct {
//	Storage string `json:"storage"`
//}

type PodConfig struct {
	ServiceAccountName string       `json:"serviceAccountName"`
	Resources          PodResources `json:"resources"`
}

type PodResources struct {
	Requests PodResourceRequest `json:"requests"`
	Limits   PodResourceLimit   `json:"limits"`
}

type PodResourceRequest struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

type PodResourceLimit struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

func DefaultConfig() *ZKClusterConfig {
	return &ZKClusterConfig{
		Size:   3,
		CPU:    "200m",
		Memory: "256Mi",
		Volume: "20Gi",
		Config: ZKClusterSubConfig{1,
			3,
			500,
			1000,
			10,
			60,
			40000,
			4000,
			65536,
			10000,
			4194304,
			2,
			2000},
	}
}

// TestConfig only use for update test
func TestConfig() *ZKClusterConfig {
	return &ZKClusterConfig{
		Size:   5,
		CPU:    "",
		Memory: "",
		Volume: "",
		Config: ZKClusterSubConfig{1,
			4,
			600,
			1100,
			0,
			0,
			41000,
			0,
			0,
			0,
			0,
			0,
			2100},
	}
}

func MakeCR(clusterName string, config *ZKClusterConfig) *ZKClusterCR {
	cr := DefaultCR(clusterName)
	cr.Spec.Replicas = config.Size
	cr.Spec.Pod.Resources.Requests.CPU = config.CPU
	cr.Spec.Pod.Resources.Requests.Memory = config.Memory
	cr.Spec.Pod.Resources.Limits.CPU = config.CPU
	cr.Spec.Pod.Resources.Limits.Memory = config.Memory
	cr.Spec.Persistence.Spec.Resources.Requests.Storage = config.Volume
	cr.Spec.Config = config.Config

	return cr
}

func DefaultCR(clusterName string) *ZKClusterCR {
	return &ZKClusterCR{
		APIVersion: "zookeeper.pravega.io/v1beta1",
		Kind:       "ZookeeperCluster",
		Metadata:   ZKClusterCRMetadata{Name: clusterName},
		Spec: ZKClusterCRSpec{
			Replicas: 3,
			Image: ImageConfig{
				Repository: "pravega/zookeeper",
				Tag:        "0.2.13",
				PullPolicy: "IfNotPresent",
			},
			StorageType: "persistence",
			Persistence: PersistentConfig{
				ReclaimPolicy: "Delete",
				Spec: PersistentSpec{
					StorageClassName: "cloud-nfs",
					Resources: PersistentResources{
						Requests: PersistentResourceRequest{Storage: "20Gi"},
						//Limits:   PersistentResourceLimit{},
					},
				},
			},
			Pod: PodConfig{
				ServiceAccountName: "zookeeper",
				Resources: PodResources{
					Requests: PodResourceRequest{
						CPU:    "200m",
						Memory: "256Mi",
					},
					Limits: PodResourceLimit{
						CPU:    "200m",
						Memory: "256Mi",
					},
				},
			},
			Config: DefaultConfig().Config,
		},
	}
}

type ZKPatch struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type ZKPatches []*ZKPatch

func (zkp *ZKPatches) Display() {
	for _, patch := range *zkp {
		fmt.Println(patch)
	}
}

func MakePatches(config *ZKClusterConfig) ZKPatches {
	patches := make([]*ZKPatch, 0)
	configMap := structs.Map(config)

	for k, v := range configMap {
		if k == "Config" {
			continue
		}
		if _, ok := v.(string); ok && len(v.(string)) != 0 {
			patches = append(patches, selectPatch(k, v)...)
			continue
		}
		if _, ok := v.(int); ok && v.(int) != 0 {
			patches = append(patches, selectPatch(k, v)...)
		}
	}

	for k, v := range structs.Map(config.Config) {
		if _, ok := v.(int); ok && v.(int) != 0 {
			patches = append(patches, selectPatch(k, v)...)
			continue
		}
		if _, ok := v.(string); ok && len(v.(string)) != 0 {
			patches = append(patches, selectPatch(k, v)...)
		}
	}

	return patches
}

func selectPatch(key string, value interface{}) []*ZKPatch {
	switch key {
	case "Size":
		return []*ZKPatch{ReplicasPatch(value)}
	case "CPU":
		return []*ZKPatch{CPURequestPatch(value), CPULimitPatch(value)}
	case "Memory":
		return []*ZKPatch{MemRequestPatch(value), MemLimitPatchPatch(value)}
	case "AutoPurgePurgeInterval":
		return []*ZKPatch{AutoPurgePurgeIntervalPatch(value)}
	case "AutoPurgeSnapRetainCount":
		return []*ZKPatch{AutoPurgeSnapRetainCountPatch(value)}
	case "CommitLogCount":
		return []*ZKPatch{CommitLogCountPatch(value)}
	case "GlobalOutstandingLimit":
		return []*ZKPatch{GlobalOutstandingLimitPatch(value)}
	case "InitLimit":
		return []*ZKPatch{InitLimitPatch(value)}
	case "MaxClientCnxns":
		return []*ZKPatch{MaxClientCnxnsPatch(value)}
	case "MaxSessionTimeout":
		return []*ZKPatch{MaxSessionTimeout(value)}
	case "MinSessionTimeout":
		return []*ZKPatch{MinSessionTimeoutPatch(value)}
	case "PreAllocSize":
		return []*ZKPatch{PreAllocSizePatch(value)}
	case "SnapCount":
		return []*ZKPatch{SnapCountPatch(value)}
	case "SnapSizeLimitInKb":
		return []*ZKPatch{SnapSizeLimitInKbPatch(value)}
	case "SyncLimit":
		return []*ZKPatch{SyncLimitPatch(value)}
	case "TickTime":
		return []*ZKPatch{TickTimePatch(value)}
	default:
		return []*ZKPatch{}
	}
}

func CPURequestPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/pod/resources/requests/cpu",
		Value: value,
	}
}

func CPULimitPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/pod/resources/limits/cpu",
		Value: value,
	}
}

func MemRequestPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/pod/resources/requests/memory",
		Value: value,
	}
}

func MemLimitPatchPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/pod/resources/limits/memory",
		Value: value,
	}
}

func ReplicasPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/replicas",
		Value: value,
	}
}

func AutoPurgePurgeIntervalPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/autoPurgePurgeInterval",
		Value: value,
	}
}

func AutoPurgeSnapRetainCountPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/autoPurgeSnapRetainCount",
		Value: value,
	}
}

func CommitLogCountPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/commitLogCount",
		Value: value,
	}
}

func GlobalOutstandingLimitPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/globalOutstandingLimit",
		Value: value,
	}
}

func InitLimitPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/initLimit",
		Value: value,
	}
}

func MaxClientCnxnsPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/maxClientCnxns",
		Value: value,
	}
}

func MaxSessionTimeout(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/maxSessionTimeout",
		Value: value,
	}
}

func MinSessionTimeoutPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/minSessionTimeout",
		Value: value,
	}
}

func PreAllocSizePatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/preAllocSize",
		Value: value,
	}
}

func SnapCountPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/snapCount",
		Value: value,
	}
}

func SnapSizeLimitInKbPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/snapSizeLimitInKb",
		Value: value,
	}
}

func SyncLimitPatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/syncLimit",
		Value: value,
	}
}

func TickTimePatch(value interface{}) *ZKPatch {
	return &ZKPatch{
		Op:    "replace",
		Path:  "/spec/config/tickTime",
		Value: value,
	}
}
