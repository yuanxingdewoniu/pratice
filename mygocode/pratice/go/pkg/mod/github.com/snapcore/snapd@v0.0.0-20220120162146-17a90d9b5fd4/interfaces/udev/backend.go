// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2016-2018 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

// Package udev implements integration between snapd, udev and
// snap-confine around tagging character and block devices so that they
// can be accessed by applications.
//
// TODO: Document this better
package udev

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/snapcore/snapd/dirs"
	"github.com/snapcore/snapd/interfaces"
	"github.com/snapcore/snapd/osutil"
	"github.com/snapcore/snapd/sandbox/cgroup"
	"github.com/snapcore/snapd/snap"
	"github.com/snapcore/snapd/timings"
)

// Backend is responsible for maintaining udev rules.
type Backend struct {
	preseed bool
}

// Initialize does nothing.
func (b *Backend) Initialize(opts *interfaces.SecurityBackendOptions) error {
	if opts != nil && opts.Preseed {
		b.preseed = true
	}
	return nil
}

// Name returns the name of the backend.
func (b *Backend) Name() interfaces.SecuritySystem {
	return interfaces.SecurityUDev
}

// snapRulesFileName returns the path of the snap udev rules file.
func snapRulesFilePath(snapName string) string {
	rulesFileName := fmt.Sprintf("70-%s.rules", snap.SecurityTag(snapName))
	return filepath.Join(dirs.SnapUdevRulesDir, rulesFileName)
}

// Setup creates udev rules specific to a given snap.
// If any of the rules are changed or removed then udev database is reloaded.
//
// UDev has no concept of a complain mode so confinement options are ignored.
//
// If the method fails it should be re-tried (with a sensible strategy) by the caller.
func (b *Backend) Setup(snapInfo *snap.Info, opts interfaces.ConfinementOptions, repo *interfaces.Repository, tm timings.Measurer) error {
	snapName := snapInfo.InstanceName()
	spec, err := repo.SnapSpecification(b.Name(), snapName)
	if err != nil {
		return fmt.Errorf("cannot obtain udev specification for snap %q: %s", snapName, err)
	}
	content := b.deriveContent(spec.(*Specification), snapInfo)
	subsystemTriggers := spec.(*Specification).TriggeredSubsystems()

	dir := dirs.SnapUdevRulesDir
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("cannot create directory for udev rules %q: %s", dir, err)
	}

	rulesFilePath := snapRulesFilePath(snapInfo.InstanceName())

	if len(content) == 0 {
		// Make sure that the rules file gets removed when we don't have any
		// content and exists.
		err = os.Remove(rulesFilePath)
		if err != nil && !os.IsNotExist(err) {
			return err
		} else if err == nil {
			// FIXME: somehow detect the interfaces that were
			// disconnected and set subsystemTriggers appropriately.
			// ATM, it is always going to be empty on disconnect.
			return b.reloadRules(subsystemTriggers)
		}
		return nil
	}

	var buffer bytes.Buffer
	buffer.WriteString("# This file is automatically generated.\n")
	if (opts.DevMode || opts.Classic) && !opts.JailMode {
		buffer.WriteString("# udev tagging/device cgroups disabled with non-strict mode snaps\n")
	}
	for _, snippet := range content {
		if (opts.DevMode || opts.Classic) && !opts.JailMode {
			buffer.WriteRune('#')
			snippet = strings.Replace(snippet, "\n", "\n#", -1)
		}
		buffer.WriteString(snippet)
		buffer.WriteByte('\n')
	}

	rulesFileState := &osutil.MemoryFileState{
		Content: buffer.Bytes(),
		Mode:    0644,
	}

	// EnsureFileState will make sure the file will be only updated when its content
	// has changed and will otherwise return an error which prevents us from reloading
	// udev rules when not needed.
	err = osutil.EnsureFileState(rulesFilePath, rulesFileState)
	if err == osutil.ErrSameState {
		return nil
	} else if err != nil {
		return err
	}

	// FIXME: somehow detect the interfaces that were disconnected and set
	// subsystemTriggers appropriately. ATM, it is always going to be empty
	// on disconnect.
	return b.reloadRules(subsystemTriggers)
}

// Remove removes udev rules specific to a given snap.
// If any of the rules are removed then udev database is reloaded.
//
// This method should be called after removing a snap.
//
// If the method fails it should be re-tried (with a sensible strategy) by the caller.
func (b *Backend) Remove(snapName string) error {
	rulesFilePath := snapRulesFilePath(snapName)
	err := os.Remove(rulesFilePath)
	if os.IsNotExist(err) {
		// If file doesn't exist we avoid reloading the udev rules when we return here
		return nil
	} else if err != nil {
		return err
	}

	// FIXME: somehow detect the interfaces that were disconnected and set
	// subsystemTriggers appropriately. ATM, it is always going to be empty
	// on disconnect.
	return b.reloadRules(nil)
}

func (b *Backend) deriveContent(spec *Specification, snapInfo *snap.Info) (content []string) {
	content = append(content, spec.Snippets()...)
	return content
}

func (b *Backend) NewSpecification() interfaces.Specification {
	return &Specification{}
}

// SandboxFeatures returns the list of features supported by snapd for mediating access to kernel devices.
func (b *Backend) SandboxFeatures() []string {
	commonFeatures := []string{
		"tagging", /* Tagging dynamically associates new devices with specific snaps */
	}
	cgroupv1Features := []string{
		"device-filtering", /* Snapd can limit device access for each snap */
		"device-cgroup-v1", /* Snapd creates a device group (v1) for each snap */
	}

	if cgroup.IsUnified() {
		// TODO: update v2 device cgroup is supported
		return commonFeatures
	}

	features := append(cgroupv1Features, commonFeatures...)
	return features
}