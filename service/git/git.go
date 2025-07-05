package git

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/HJyup/forge/config"
)

func GetLastCommitTime(path string) (time.Time, error) {
	var zeroTime time.Time

	cmd := exec.Command(config.GitCommand, config.GitChangeDir, path, config.GitLogCommand, config.GitLogOneCommit, config.GitLogFormat)
	output, err := cmd.Output()
	if err != nil {
		return zeroTime, err
	}

	timestamp := strings.TrimSpace(string(output))
	if timestamp == "" {
		return zeroTime, nil
	}

	var unixTime int64
	_, err = fmt.Sscanf(timestamp, "%d", &unixTime)
	if err != nil {
		return zeroTime, err
	}

	return time.Unix(unixTime, 0), nil
}

func IsGitRepository(path string) bool {
	cmd := exec.Command(config.GitCommand, config.GitChangeDir, path, config.GitRevParseCommand, config.GitRevParseGitDir)
	err := cmd.Run()
	return err == nil
}
