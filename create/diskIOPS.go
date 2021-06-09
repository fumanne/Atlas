package create

import "log"

func computeDiskIOPS(instanceSize string, diskSizeGB int) int {
	rate := 30
	MaxIO := map[string]int{
		"M30":  3600,
		"M40":  6000,
		"M50":  8000,
		"M60":  16000,
		"M80":  32000,
		"M140": 32000,
		"M200": 32000,
		"M300": 32000,
		"M400": 32000,
		"M700": 32000,
	}
	MaxDiskSize := map[string]int{
		"M30":  512,
		"M40":  1024,
		"M50":  4096,
		"M60":  4096,
		"M80":  4096,
		"M140": 4096,
		"M200": 4096,
		"M300": 4096,
		"M400": 4096,
		"M700": 4096,
	}

	thresholdMap := make(map[string]int)
	for k, v := range MaxIO {
		threshold := v / rate
		thresholdMap[k] = threshold
	}
	//fmt.Println(thresholdMap)
	boundary := thresholdMap[instanceSize]
	if diskSizeGB > MaxDiskSize[instanceSize] {
		log.Fatalf("%s: Disk Size Bigger than max size %d", instanceSize, MaxDiskSize[instanceSize])
	}

	if diskSizeGB > boundary {
		return MaxIO[instanceSize]
	}
	return diskSizeGB * rate
}
