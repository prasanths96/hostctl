package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

var commonFilUmask = os.FileMode(0777)

func writeBytesToFile(fileName string, pathS string, dataB []byte) {
	// f, err := os.Create(path.Join(pathS, fileName))
	err := ioutil.WriteFile(path.Join(pathS, fileName), dataB, commonFilUmask)
	cobra.CheckErr(err)
}

func readFileBytes(fullPath string) (dataB []byte) {
	dataB, err := ioutil.ReadFile(fullPath)
	cobra.CheckErr(err)
	return
}

func execAndGetOutput(dir string, comdS string, args ...string) (out []byte, err error) {
	comd := exec.Command(comdS, args...)
	if dir != "" {
		comd.Dir = dir
	}
	out, err = comd.Output()
	return
}
