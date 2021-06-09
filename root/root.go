package root

import "os"

const (
	DefaultPublicKey  string = "abc123"
	DefaultPrivateKey string = "xajh1-asbkjdb1-ba0125a0-1"
	DefaultGroupId    string = "60af70efd0107237d1d3e647"
)

func GetPublicKey() string {
	if v, ok := os.LookupEnv("PublicKey"); ok {
		return v
	}
	return DefaultPublicKey
}

func GetPrivateKey() string {
	if v, ok := os.LookupEnv("PrivateKey"); ok {
		return v
	}
	return DefaultPrivateKey
}

func GetGroupID() string {
	if v, ok := os.LookupEnv("GroupID"); ok {
		return v
	}
	return DefaultGroupId
}
