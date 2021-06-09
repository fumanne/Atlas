package show

import (
	"fmt"
	dac "github.com/xinsnake/go-http-digest-auth-client"
	"io/ioutil"
)

func Do(publicKey, privateKey, groupID, clusterName string) {
	baseUrl := "https://cloud.mongodb.com/api/atlas/v1.5"
	queryUrl := fmt.Sprintf("%s/groups/%s/clusters/%s", baseUrl, groupID, clusterName)

	dr := dac.NewRequest(publicKey, privateKey, "GET", queryUrl, "")
	dr.Header.Set("Content-Type", "application/json")

	resp, err := dr.Execute()
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
