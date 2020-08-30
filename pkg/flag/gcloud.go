package flag

import "os/exec"

func gcloud() (which string, err error) {
	which, err = exec.LookPath("gcloud")
	return
}

func project() (result *exec.Cmd, err error) {
	file, err := gcloud()
	if err != nil {
		return
	}
	result = exec.Command(file, "config", "get-value", "project")
	return
}
