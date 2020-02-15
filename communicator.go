package dockerstats

import (
	"encoding/json"
	"os/exec"
	"strings"
)

const (
	defaultDockerPath        string = "/usr/bin/docker"
	defaultDockerCommand     string = "stats"
	defaultDockerNoStreamArg string = "--no-stream"
	defaultDockerFormatArg   string = "--format"
	defaultDockerFormat      string = `{"container_id":"{{.ID}}","container_name":"{{.Name}}","memory":{"raw":"{{.MemUsage}}","percent":"{{.MemPerc}}"},"cpu":"{{.CPUPerc}}","io":{"network":"{{.NetIO}}","block":"{{.BlockIO}}"},"pids":{{.PIDs}}}`
)

// Communicator provides an interface for communicating with and retrieving stats
// from Docker.
type Communicator interface {
	Stats(filter ...string) ([]Stats, error)
}

// CliCommunicator uses the Docker CLI to retrieve stats for currently running Docker
// containers.
type CliCommunicator struct {
	DockerPath string
	Command    []string
}

// DefaultCommunicator is the default way of retrieving stats from Docker.
//
// When calling `Current()`, the `DefaultCommunicator` is used, and when
// retriving a `Monitor` using `NewMonitor()`, it is initialized with the
// `DefaultCommunicator`.
var DefaultCommunicator Communicator = CliCommunicator{
	DockerPath: defaultDockerPath,
	Command:    []string{defaultDockerCommand, defaultDockerNoStreamArg, defaultDockerFormatArg, defaultDockerFormat},
}

// Stats returns Docker container statistics using the Docker CLI.
func (c CliCommunicator) Stats(filters ...string) ([]Stats, error) {
	args := c.Command
	if len(filters) > 0 {
		args = append(c.Command, filters...)
	}

	out, err := exec.Command(c.DockerPath, args...).Output()
	if err != nil {
		return nil, err
	}

	containers := strings.Split(string(out), "\n")
	stats := make([]Stats, 0)
	for _, con := range containers {
		if len(con) == 0 {
			continue
		}

		var s Stats
		if err := json.Unmarshal([]byte(con), &s); err != nil {
			return nil, err
		}

		stats = append(stats, s)
	}

	return stats, nil
}
