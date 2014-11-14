// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package diskmanager

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/juju/errors"

	"github.com/juju/juju/storage"
)

var pairsRE = regexp.MustCompile(`([A-Z]+)=(?:"(.*?)")`)

func listBlockDevices() ([]storage.BlockDevice, error) {
	columns := []string{
		"KNAME", // kernel name
		"SIZE",  // size
		"LABEL", // filesystem label
		"UUID",  // filesystem UUID
	}

	logger.Debugf("executing lsblk")
	output, err := exec.Command(
		"lsblk",
		"-b", // output size in bytes
		"-P", // output fields as key=value pairs
		"-o", strings.Join(columns, ","),
	).Output()
	if err != nil {
		return nil, errors.Annotate(
			err, "cannot list block devices: lsblk failed",
		)
	}

	blockDeviceMap := make(map[string]storage.BlockDevice)
	s := bufio.NewScanner(bytes.NewReader(output))
	for s.Scan() {
		pairs := pairsRE.FindAllStringSubmatch(s.Text(), -1)
		var dev storage.BlockDevice
		for _, pair := range pairs {
			switch pair[1] {
			case "KNAME":
				dev.DeviceName = pair[2]
			case "SIZE":
				size, err := strconv.ParseUint(pair[2], 10, 64)
				if err != nil {
					logger.Errorf(
						"invalid size %q from lsblk: %v", pair[2], err,
					)
				} else {
					dev.Size = size / bytesInMiB
				}
			case "LABEL":
				dev.Label = pair[2]
			case "UUID":
				dev.UUID = pair[2]
			default:
				logger.Debugf("unexpected field from lsblk: %q", pair[1])
			}
		}

		// Check if the block device is in use. We need to know this so we can
		// issue an error if the user attempts to allocate an in-use disk to a
		// unit.
		dev.InUse, err = blockDeviceInUse(dev)
		if os.IsNotExist(err) {
			// In LXC containers, lsblk will show the block devices of the
			// host, but the devices will typically not be present.
			continue
		} else if err != nil {
			logger.Errorf(
				"error checking if %q is in use: %v", dev.DeviceName, err,
			)
			// We cannot detect, so err on the side of caution and default to
			// "in use" so the device cannot be used.
			dev.InUse = true
		}
		blockDeviceMap[dev.DeviceName] = dev
	}
	if err := s.Err(); err != nil {
		return nil, errors.Annotate(err, "cannot parse lsblk output")
	}

	blockDevices := make([]storage.BlockDevice, 0, len(blockDeviceMap))
	for _, dev := range blockDeviceMap {
		blockDevices = append(blockDevices, dev)
	}
	return blockDevices, nil
}
