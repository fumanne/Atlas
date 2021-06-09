package create

import (
	"encoding/json"
	"fmt"
	dac "github.com/xinsnake/go-http-digest-auth-client"
	"io/ioutil"
	"strings"
)

const (
	DefaultClusterType string = "SHARDED"
	DefaultProvider    string = "AWS"
)

func Do(publicKey, privateKey, groupID, clusterName, regionName, instanceSize string, numShards, diskSizeGB int) {
	baseUrl := "https://cloud.mongodb.com/api/atlas/v1.5"
	queryUrl := fmt.Sprintf("%s/groups/%s/clusters", baseUrl, groupID)

	regionName = convertRegion(regionName)
	params := NewCreateClusterParams(DefaultClusterType, clusterName, DefaultProvider, regionName, instanceSize, numShards, diskSizeGB)
	bs, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(bs))
	dr := dac.NewRequest(publicKey, privateKey, "POST", queryUrl, string(bs))
	dr.Header.Set("Content-Type", "application/json")
	resp, err := dr.Execute()
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	result := NewCreateClusterResponseElements()
	json.Unmarshal(b, &result)
	fmt.Println(string(b))

}

func convertRegion(regionName string) string {
	return strings.Replace(strings.ToUpper(regionName), "-", "_", -1)
}
