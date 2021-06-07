package reporter

import (
	"fmt"
	"os"
)

type Reporter struct {
	HttpEndPoint string
	ReportBody   ReportBody
}

type ReportBody struct {
	InstanceId    string        `json:"instanceId"`
	Command       string        `json:"command"`       // Sub-command of hostctl
	SuccessReport []interface{} `json:"successReport"` // List of succeeded commands and info
	ErrorReport   []interface{} `json:"errorReport"`   // List of error commands and info
}

type Report interface {
	AddOutput(output string)
	GetReport() (report interface{})
}

// toolCmd -> eg: execT
func NewReporter(httpEndPoint string, toolCmd string, instanceId string) (reporter *Reporter) {
	reporter = &Reporter{
		HttpEndPoint: httpEndPoint,
		ReportBody: ReportBody{
			InstanceId: instanceId,
			Command:    toolCmd,
		},
	}
	return
}

func (rep *Reporter) AddSuccessReport(report interface{}) {
	// Append success report
	rep.ReportBody.SuccessReport = append(rep.ReportBody.SuccessReport, report)
}

func (rep *Reporter) AddErrorReport(report interface{}) {
	// Append error report
	rep.ReportBody.ErrorReport = append(rep.ReportBody.ErrorReport, report)
}

func (rep *Reporter) HandleErr(customMsg string, err error, reportI Report) {
	if err == nil {
		return
	}
	// Add output to report
	reportI.AddOutput(customMsg + ":" + err.Error())
	// Get report
	// Add it to report body
	rep.AddErrorReport(reportI.GetReport())

	// Submit or don't submit based on the config (not added yet)
	// Submitting by default for now and exiting
	rep.SubmitReport()
	os.Exit(1)
}

func (rep *Reporter) SubmitReport() {
	// Send the reportBody
	fmt.Println("REPORT SUBMITTED:")
	fmt.Printf("%+v\n", rep.ReportBody)
}

// Post end point
