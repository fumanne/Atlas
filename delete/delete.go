package delete

import (
	"fmt"
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

func Do(publicKey, privateKey, groupID, clusterName string) {
	baseUrl := "https://cloud.mongodb.com/api/atlas/v1.0"
	deleteUrl := fmt.Sprintf("%s/groups/%s/clusters/%s", baseUrl, groupID, clusterName)

	dr := dac.NewRequest(publicKey, privateKey, "DELETE", deleteUrl, "")
	dr.Header.Set("Content-Type", "application/json")

	resp, err := dr.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	//defer resp.Body.Close()
	//b, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(b))
}
