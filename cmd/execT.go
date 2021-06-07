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
	"fmt"

	"hostctl/cmd/reporter"

	"github.com/spf13/cobra"
)

var execTFlags struct {
	// Type string // sc - single command; sf - shell file
	ConfigPath string
}

type ExecTConfig struct {
	ReportEndpoint string    `json:"report-endpoint"`
	Commands       []Command `json:"commands"`
}

type Command struct {
	Dir     string   `json:"dir"`
	Command []string `json:"command"`
}

var execTConfig ExecTConfig

// execTCmd represents the execT command
var execTCmd = &cobra.Command{
	Use:   "execT",
	Short: "Executes a task given and reports back status to given HTTP end point.",
	Long:  `Executes a task given and reports back status to given HTTP end point.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		loadConfig()
		configValidation()
		newReporter()
	},
	Run: func(cmd *cobra.Command, args []string) {
		runExecT()
	},
}

func init() {
	rootCmd.AddCommand(execTCmd)

	// deployCaCmd.Flags().StringVarP(&depCaFlags.Type, "type", "t", "sc", `Type of the execution ("sc" - single command |"sf" - shell file)`)
	execTCmd.Flags().StringVarP(&execTFlags.ConfigPath, "file", "f", "", `Path of config file.`)
}

func loadConfig() {
	dataB := readFileBytes(execTFlags.ConfigPath)
	err := json.Unmarshal(dataB, &execTConfig)
	cobra.CheckErr(err)
}

func configValidation() {
	// // Check if reporter end point is valid

	// // Check if commands are valid
	// for i := 0; i < len(execTConfig.Commands); i++ {
	// 	thisCmd := execTConfig.Commands[i]
	// 	// Dir path is legit
	// 	_, err := os.Stat(thisCmd.Dir)

	// 	// atleast one string
	// }

}

type ExecTReport struct {
	CommandId int     // Taken from order in array : 0,1,2, etc.
	Command   Command // Actual Command that ran
	Output    string  // Success or error output // If error, fill this and use reportBody.AddError()
}

func (execTRep *ExecTReport) AddOutput(output string) {
	execTRep.Output = output
}

func (execTRep *ExecTReport) GetReport() (execTReport interface{}) {
	execTReport = *execTRep
	return
}

var breporter *reporter.Reporter

func newReporter() {
	breporter = reporter.NewReporter(execTConfig.ReportEndpoint, "execT")
}

func runExecT() {
	for i := 0; i < len(execTConfig.Commands); i++ {
		thisCmd := execTConfig.Commands[i]

		// Create a report for every command
		execTRep := &ExecTReport{
			CommandId: i,
			Command:   thisCmd,
			Output:    "", // To be filled later
		}

		// Validate command
		if len(thisCmd.Command) < 1 {
			breporter.HandleErr("", fmt.Errorf("invalid command"), execTRep)
		}

		firstCmd := thisCmd.Command[0]
		restCmd := thisCmd.Command[1:]
		outB, err := execAndGetOutput(thisCmd.Dir, firstCmd, restCmd...)
		// Error out
		breporter.HandleErr(string(outB), err, execTRep)

		// Success out
		if err == nil {
			execTRep.Output = string(outB)
			breporter.AddSuccessReport(execTRep.GetReport())
		}
	}

	// Submit report
	breporter.SubmitReport()
}
