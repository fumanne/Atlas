package create

import "time"

type ResElectableSpecs struct {
	InstanceSize  string `json:"instanceSize"`
	DiskIOPS      int    `json:"diskIOPS"`
	EbsVolumeType string `json:"ebsVolumeType"`
	NodeCount     int    `json:"nodeCount"`
}

type ResRegionConfig struct {
	ElectableSpecs *ResElectableSpecs `json:"electableSpecs"`
	Priority       int                `json:"priority"`
	ProviderName   string             `json:"providerName"`
	RegionName     string             `json:"regionName"`
}

type ResReplicationSpecs struct {
	Id            string             `json:"id"`
	NumShards     int                `json:"numShards"`
	ZoneName      string             `json:zoneName`
	RegionConfigs []*ResRegionConfig `json:"regionConfigs"`
}

type ConnectionStrings struct {
	Standard    string `json:"standard"`
	StandardSrv string `json:"standardSrv"`
}

type CreateClusterResponseElements struct {
	BackupEnabled            bool                   `json:"backupEnabled"`
	ClusterType              string                 `json:"clusterType"`
	ConnectionStrings        *ConnectionStrings     `json:"connectionStrings"`
	CreateDate               time.Time              `json:"createDate"`
	DiskSizeGB               int                    `json:"diskSizeGB"`
	EncryptionAtRestProvider string                 `json:"encryptionAtRestProvider"`
	GroupId                  string                 `json:"groupId"`
	Id                       string                 `json:"id"`
	Labels                   []map[string]string    `json:"labels"`
	MongoDBMajorVersion      string                 `json:"mongoDBMajorVersion"`
	MongoDBVersion           string                 `json:"mongoDBVersion"`
	Name                     string                 `json:"name"`
	Paused                   bool                   `json:"paused"`
	PitEnabled               bool                   `json:"pitEnabled"`
	ReplicationSpecs         []*ResReplicationSpecs `json:"replicationSpecs"`
	RootCertType             string                 `json:"rootCertType"`
	StateName                string                 `json:"stateName"`
}

func NewCreateClusterResponseElements() *CreateClusterResponseElements {
	ele := &ResElectableSpecs{}
	reg := &ResRegionConfig{
		ElectableSpecs: ele,
	}

	cs := &ConnectionStrings{}

	replica := &ResReplicationSpecs{
		RegionConfigs: []*ResRegionConfig{reg},
	}

	return &CreateClusterResponseElements{
		ConnectionStrings: cs,
		ReplicationSpecs:  []*ResReplicationSpecs{replica},
	}

}
