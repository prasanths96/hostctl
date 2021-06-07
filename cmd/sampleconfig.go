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
	"encoding/json"

	"github.com/spf13/cobra"
)

var execTSampleConfigFlags struct {
	Path string
}

type Object map[string]interface{}

var sampleJson = Object{
	"bimt-host-endpoint": "http://localhost:8090/deployer/status",
	"commands": []Object{
		{
			"dir":     "",
			"command": []string{"echo", "commands involving | & >> should be avoided"},
		},
		{
			"dir":     "<dir to execute cmd from>",
			"command": []string{"sudo", "./hlfd", "install", "prereqs"},
		},
		{
			"dir":     "<dir to execute cmd from>",
			"command": []string{"sudo", "./hlfd", "deploy", "ca", "-n", "first-ca", "-p", "8040"},
		},
	},
}

// sampleconfigCmd represents the sampleconfig command
var sampleconfigCmd = &cobra.Command{
	Use:   "sampleconfig",
	Short: "Generates sample config.",
	Long:  `Generates sample config.`,
	Args: func(cmd *cobra.Command, args []string) (err error) {
		return
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// Path validation
	},
	Run: func(cmd *cobra.Command, args []string) {
		generateSample()
	},
}

func init() {
	execTCmd.AddCommand(sampleconfigCmd)
	sampleconfigCmd.Flags().StringVarP(&execTSampleConfigFlags.Path, "path", "p", "", `Path of to generate sample config file (required).`)

	sampleconfigCmd.MarkFlagRequired("path")
}

func generateSample() {
	mc, err := json.Marshal(sampleJson)
	cobra.CheckErr(err)
	writeBytesToFile("sample-config.json", execTSampleConfigFlags.Path, mc)
}
