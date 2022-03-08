package zkoperator

import "time"

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

type ZKClusterPodInfo struct {
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
	APIVersion string              `yaml:"apiVersion"`
	Kind       string              `yaml:"kind"`
	Metadata   ZKClusterCRMetadata `yaml:"metadata"`
	Spec       ZKClusterCRSpec     `yaml:"spec"`
}

type ZKClusterCRMetadata struct {
	Name string `yaml:"name"`
}

type ZKClusterCRSpec struct {
	Replicas    int                `yaml:"replicas"`
	Image       ImageConfig        `yaml:"image"`
	StorageType string             `yaml:"storageType"`
	Persistence PersistentConfig   `yaml:"persistence"`
	Pod         PodConfig          `yaml:"pod"`
	Config      ZKClusterSubConfig `yaml:"config"`
}

type ImageConfig struct {
	Repository string `yaml:"repository"`
	Tag        string `yaml:"tag"`
}

type PersistentConfig struct {
	ReclaimPolicy string         `yaml:"reclaimPolicy"`
	Spec          PersistentSpec `yaml:"spec"`
}

type PersistentSpec struct {
	StorageClassName string              `yaml:"storageClassName"`
	Resources        PersistentResources `yaml:"resources"`
}

type PersistentResources struct {
	Requests PersistentResourceRequest `yaml:"requests"`
	Limits   PersistentResourceLimit   `yaml:"limits"`
}

type PersistentResourceRequest struct {
	Storage string `yaml:"storage"`
}

type PersistentResourceLimit struct {
	Storage string `yaml:"storage"`
}

type PodConfig struct {
	ServiceAccountName string       `yaml:"serviceAccountName"`
	Resources          PodResources `yaml:"resources"`
}

type PodResources struct {
	Requests PodResourceRequest `yaml:"requests"`
	Limits   PodResourceLimit   `yaml:"limits"`
}

type PodResourceRequest struct {
	CPU    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

type PodResourceLimit struct {
	CPU    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
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
			},
			StorageType: "persistence",
			Persistence: PersistentConfig{
				ReclaimPolicy: "Delete",
				Spec: PersistentSpec{
					StorageClassName: "cloud-nfs",
					Resources: PersistentResources{
						Requests: PersistentResourceRequest{Storage: "20Gi"},
						Limits:   PersistentResourceLimit{},
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
					Limits: PodResourceLimit{},
				},
			},
			Config: DefaultConfig().Config,
		},
	}
}
