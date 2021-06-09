package create

type DiskGB struct {
	Enabled bool `json:"enabled"`
}

type Compute struct {
	Enabled          bool `json:"enabled"`
	ScaleDownEnabled bool `json:"scaleDownEnabled"`
}

type AutoScaling struct {
	Compute *Compute `json:"compute"`
	DiskGB  *DiskGB  `json:"diskGB"`
}

type ElectableSpecs struct {
	DiskIOPS      int    `json:"diskIOPS"`
	EbsVolumeType string `json:"ebsVolumeType"`
	InstanceSize  string `json:"instanceSize"`
	NodeCount     int    `json:"nodeCount"`
}

type RegionConfig struct {
	Priority       int             `json:"priority"`
	ElectableSpecs *ElectableSpecs `json:"electableSpecs"`
	ProviderName   string          `json:"providerName"`
	RegionName     string          `json:"regionName"`
	AutoScaling    *AutoScaling    `json:"autoScaling"`
}

type ReplicationSpecs struct {
	NumShards     int             `json:"numShards"`
	RegionConfigs []*RegionConfig `json:"regionConfigs"`
}

type CreateClusterParams struct {
	ClusterType         string              `json:"clusterType"`
	Name                string              `json:"name"`
	DiskSizeGB          int                 `json:"diskSizeGB"`
	ReplicationSpecs    []*ReplicationSpecs `json:"replicationSpecs"`
	MongoDBMajorVersion string              `json:"mongoDBMajorVersion"`
}

func NewCreateClusterParams(clusterType, name, provider, region, instanceSize string, numShards, diskSizeGB int) *CreateClusterParams {

	elec := &ElectableSpecs{
		InstanceSize:  instanceSize,
		NodeCount:     3,
		EbsVolumeType: "PROVISIONED",
		DiskIOPS:      computeDiskIOPS(instanceSize, diskSizeGB),
	}
	reg := &RegionConfig{
		Priority:       7,
		ElectableSpecs: elec,
		ProviderName:   provider,
		RegionName:     region,
		AutoScaling: &AutoScaling{
			Compute: &Compute{
				Enabled:          false,
				ScaleDownEnabled: false,
			},
			DiskGB: &DiskGB{
				Enabled: true,
			},
		},
	}

	replica := &ReplicationSpecs{
		NumShards:     numShards,
		RegionConfigs: []*RegionConfig{reg},
	}

	return &CreateClusterParams{
		ClusterType:         clusterType,
		Name:                name,
		DiskSizeGB:          diskSizeGB,
		ReplicationSpecs:    []*ReplicationSpecs{replica},
		MongoDBMajorVersion: "4.2",
	}

}
