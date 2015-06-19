// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// Package plugin contains the code that interfaces with plugins for workload
// process technologies such as Docker, Rocket, or systemd.
//
// Plugins of this type are expected to handle three commands: launch, status,
// and stop.  See the functions of the same name for more information about each
// command.
//
// If the plugin command completes successfully, the plugin should exit with a
// 0 exit code. If there is a problem completing the command, the plugin should
// print the error details to stdout and return a non-zero exit code.
//
// Any information written to stderr will be piped to the unit log.
package plugin

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"github.com/juju/deputy"
	"github.com/juju/errors"
	"github.com/juju/loggo"
	"gopkg.in/juju/charm.v5"
)

var logger = loggo.GetLogger("juju.process.plugin")

// validationErr represents an error signifying an object with an invalid value.
type validationErr struct {
	*errors.Err
}

// IsInvalid returns whether the given error indicates an invalid value.
func IsInvalid(e error) bool {
	_, ok := e.(validationErr)
	return ok
}

// Find returns the plugin for the given name.
func Find(name string) (*Plugin, error) {
	path, err := exec.LookPath("juju-process-" + name)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &Plugin{Name: name, Path: path}, nil
}

// ProcDetails represents information about a process launched by a plugin.
type ProcDetails struct {
	// ID is a unique string identifying the process to the plugin.
	ID string `json:"id"`
	// Status represents the human-readable string returned by the plugin for
	// the launched process.
	Status string `json:"status"`
}

// ProcStatus represents the data returned from the Status call.
type ProcStatus struct {
	// Status represents the human-readable string returned by the plugin for
	// the process.
	Status string `json:"status"`
}

// Validate returns nil if this value is valid, and an error if it is not.
func (p ProcDetails) Validate() error {
	if p.ID == "" {
		e := errors.NewErr("ID cannot be empty")
		return validationErr{&e}
	}
	return nil
}

// Plugin represents a provider for launching, destroying, and introspecting
// workload processes via a specific technology such as Docker or systemd.
type Plugin struct {
	// Name is the name of the plugin.
	Name string
	// Path is the filename disk where the plugin executable resides.
	Path string
}

// Launch runs the given plugin, passing it the "launch" command, with the path
// to the image to launch as an argument.
//
//		launch <image>
//
// The plugin is expected to start the image as a new process, and write json
// output to stdout.  The form of the output is expected to be:
//		{
//			"id" : "some-id", # unique id of the process
//			"status" : "details" # plugin-specific metadata about the started process
//		}
//
// The id should be a unique identifier of the process that the plugin can use
// later to introspect the process and/or stop it. The contents of status can
// be whatever information the plugin thinks might be relevant to see in the
// service's status output.
func (p Plugin) Launch(proc charm.Process) (ProcDetails, error) {
	b, err := json.Marshal(proc)
	if err != nil {
		return ProcDetails{}, errors.Annotate(err, "can't convert charm.Process to json")
	}
	details := ProcDetails{}
	out, err := p.run("launch", string(b))
	if err != nil {
		return details, errors.Trace(err)
	}
	if err := json.Unmarshal(out, &details); err != nil {
		return details, errors.Annotatef(err, "error parsing data returned from %q", p.Name)
	}
	if err := details.Validate(); err != nil {
		return details, errors.Annotatef(err, "invalid details returned by plugin %q", p.Name)
	}
	return details, nil
}

// Destroy runs the given plugin, passing it the "destroy" command, with the id of the
// process to destroy as an argument.
//
// 		destroy <id>
func (p Plugin) Destroy(id string) error {
	_, err := p.run("destroy", id)
	return errors.Trace(err)
}

// Status runs the given plugin, passing it the "status" command, with the id of
// the process to get status about.
//
//		status <id>
//
// The plugin is expected to write raw-string status output to stdout if
// successful.
func (p Plugin) Status(id string) (ProcStatus, error) {
	out, err := p.run("status", id)
	status := ProcStatus{}
	if err != nil {
		return status, errors.Trace(err)
	}
	if err := json.Unmarshal(out, &status); err != nil {
		return status, errors.Annotatef(err, "error parsing data returned from %q", p.Name)
	}

	return status, nil
}

// run runs the given subcommand of the plugin with the given args.
func (p Plugin) run(subcommand string, args ...string) ([]byte, error) {
	return runCmd(p.Name, p.Path, subcommand, args...)
}

// runCmd runs the executable at path with the subcommand as the first argument
// and any args in args as further arguments.  It logs to loggo using the name
// as a namespace.
var runCmd = func(name, path, subcommand string, args ...string) ([]byte, error) {
	cmd := exec.Command(path, append([]string{subcommand}, args...)...)
	log := getLogger("juju.process.plugin." + name)
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	err := deputy.Deputy{
		Errors:    deputy.FromStdout,
		StderrLog: func(s string) { log.Infof(s) },
	}.Run(cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return stdout.Bytes(), nil
}

type infoLogger interface {
	Infof(s string, args ...interface{})
}

var getLogger = func(name string) infoLogger {
	return loggo.GetLogger(name)
}
