package zkoperator

import "time"

type ZKClusterConfig struct {
	Size   int    `json:"size"`
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
	Volume string `json:"volume"`
	Config struct {
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
	} `json:"config"`
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
		Config struct {
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
		} `json:"config"`
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
