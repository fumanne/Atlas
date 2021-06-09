### Usage

    For example:
    
            Atlas show --clusterName xxxx 
            Atlas delete --clusterName xxxx
            Atlas create --clusterName xxxx --region ap-east-1 --numShards 2 --instanceSize M30
    
    Usage:
      Atlas [command]
    
    Available Commands:
      create      Create MongoDB cluster
      delete      Delete One MongoDB Cluster
      help        Help about any command
      show        Show MongoDB Cluster

    
    Flags:
      -g, --groupID string      Atlas Project ID (default "")
      -h, --help                help for Atlas
      -s, --privateKey string   Atlas Access Private Key (default "")
      -p, --publicKey string    Atlas Access Public Key (default "")

    
    
### Create  Command
    go run main.go create -r ap-east-1 -c 3 -d 111 -n cluster-Test -i M30

    {"backupEnabled":false,"biConnector":{"enabled":false,"readPreference":"secondary"},"clusterType":"SHARDED","connectionStrings":{},"createDate":"2021-05-31T09:09:44Z","diskSizeGB":111.0,"encryptionAtRestProvider":"NONE","groupId":"60af70efd0107237d1d3e647","id":"60b4a7d8336dc74441c089ef","labels":[],"mongoDBMajorVersion":"4.2","mongoDBVersion":"4.2.14","name":"cluster-Test","paused":false,"pitEnabled":false,"replicationSpecs":[{"id":"60b4a7d8336dc74441c089ea","numShards":3,"regionConfigs":[{"analyticsSpecs":{"instanceSize":"M30","diskIOPS":3330,"ebsVolumeType":"PROVISIONED","nodeCount":0},"autoScaling":{"compute":{"enabled":false,"scaleDownEnabled":false},"diskGB":{"enabled":true}},"electableSpecs":{"instanceSize":"M30","diskIOPS":3330,"ebsVolumeType":"PROVISIONED","nodeCount":3},"priority":7,"providerName":"AWS","readOnlySpecs":{"instanceSize":"M30","diskIOPS":3330,"ebsVolumeType":"PROVISIONED","nodeCount":0},"regionName":"AP_EAST_1"}],"zoneName":"Zone 1"}],"rootCertType":"ISRGROOTX1","stateName":"CREATING"}
    
    
### Show Command    
    go run main.go show -n cluster-Test
    
    {"backupEnabled":false,"biConnector":{"enabled":false,"readPreference":"secondary"},"clusterType":"SHARDED","connectionStrings":{"standard":"mongodb://techcenter-mongodb-shard-00-00.vncux.mongodb.net:27016,techcenter-mongodb-shard-00-01.vncux.mongodb.net:27016,techcenter-mongodb-shard-00-02.vncux.mongodb.net:27016,techcenter-mongodb-shard-01-00.vncux.mongodb.net:27016,techcenter-mongodb-shard-01-01.vncux.mongodb.net:27016,techcenter-mongodb-shard-01-02.vncux.mongodb.net:27016/?ssl=true&authSource=admin","standardSrv":"mongodb+srv://techcenter-mongodb.vncux.mongodb.net"},"createDate":"2021-05-27T11:27:28Z","diskSizeGB":40.0,"encryptionAtRestProvider":"NONE","groupId":"60af70efd0107237d1d3e647","id":"60af82207449116cae2bf206","labels":[],"mongoDBMajorVersion":"4.2","mongoDBVersion":"4.2.14","name":"TechCenter-MongoDB","paused":false,"pitEnabled":false,"replicationSpecs":[{"id":"60af8183690aac43a336b679","numShards":2,"regionConfigs":[{"analyticsSpecs":{"instanceSize":"M30","diskIOPS":3000,"ebsVolumeType":"STANDARD","nodeCount":0},"autoScaling":{"compute":{"enabled":false,"scaleDownEnabled":false},"diskGB":{"enabled":false}},"electableSpecs":{"instanceSize":"M30","diskIOPS":3000,"ebsVolumeType":"STANDARD","nodeCount":3},"priority":7,"providerName":"AWS","readOnlySpecs":{"instanceSize":"M30","diskIOPS":3000,"ebsVolumeType":"STANDARD","nodeCount":0},"regionName":"AP_EAST_1"}],"zoneName":"Zone 1"}],"rootCertType":"ISRGROOTX1","stateName":"IDLE"}

    
    
