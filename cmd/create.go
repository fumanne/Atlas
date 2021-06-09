/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"Atlas/pkg/create"
	"github.com/spf13/cobra"
)

var regionName, instanceSize string
var numShards, diskSizeGB int

const (
	DefaultDiskSizeGB   int    = 10
	DefaultNumShards    int    = 2
	DefaultInstanceSize string = "M30"
	DefaultClusterName  string = "Cluster-Auto"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create MongoDB cluster",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("create called")
		create.Do(publicKey, privateKey, groupID, clusterName, regionName, instanceSize, numShards, diskSizeGB)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.Flags().StringVarP(&regionName, "regionName", "r", "", "region Name of AWS")
	createCmd.Flags().StringVarP(&instanceSize, "instanceSize", "i", DefaultInstanceSize, "Size of MongoDB Instance, M30 or greater")
	createCmd.Flags().IntVarP(&numShards, "numShards", "c", DefaultNumShards, "number of shards")
	createCmd.Flags().IntVarP(&diskSizeGB, "diskSizeGB", "d", DefaultDiskSizeGB, "Disk Size")
	createCmd.Flags().StringVarP(&clusterName, "clusterName", "n", DefaultClusterName, "Name of MongoDB Cluster")
	createCmd.MarkFlagRequired("regionName")
}
