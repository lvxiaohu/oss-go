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
	"github.com/oss-go/pkg/project"
	"github.com/spf13/cobra"
)

var localFilePath string
var bucketName string
var remoteFilePath string

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "upload file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(bucketName) == 0 {
			cmd.Help()
			return
		}
		project.PutObject(bucketName,localFilePath,remoteFilePath)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	putCmd.Flags().StringVarP(&localFilePath,"file","f","","file's name | eg: -f /tmp/a.tar.gz")
	putCmd.Flags().StringVarP(&bucketName,"bucket","b","","bucket's name | eg: -b ossbucket")
	putCmd.Flags().StringVarP(&remoteFilePath,"remoteFilePath","r","","remote file path | eg: -r version1.1/filename")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
