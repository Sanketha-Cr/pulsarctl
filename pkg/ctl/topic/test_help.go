// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package topic

import (
	"bytes"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
)

// Override the NewPulsarClient function for testing
var origNewPulsarClient = cmdutils.NewPulsarClient
var mockClient cmdutils.Client

func TestTopicCommands(newVerb func(cmd *cmdutils.VerbCmd), args []string) (out *bytes.Buffer,
	execErr, nameErr, err error) {
	// Use mock client for tests
	cmdutils.NewPulsarClient = func() cmdutils.Client {
		if mockClient == nil {
			mockClient = NewMockClient()
		}
		return mockClient
	}
	var execError error
	cmdutils.ExecErrorHandler = func(err error) {
		execError = err
	}

	var nameError error
	cmdutils.CheckNameArgError = func(err error) {
		nameError = err
	}

	var rootCmd = &cobra.Command{
		Use:   "pulsarctl [command]",
		Short: "a CLI for Apache Pulsar",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs(append([]string{"topics"}, args...))

	resourceCmd := cmdutils.NewResourceCmd(
		"topics",
		"Operations about topics(s)",
		"",
		"topic")
	flagGrouping := cmdutils.NewGrouping()
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, newVerb)
	rootCmd.AddCommand(resourceCmd)
	err = rootCmd.Execute()

	// Reset the original function after test
cmdutils.NewPulsarClient = origNewPulsarClient
return buf, execError, nameError, err
}
