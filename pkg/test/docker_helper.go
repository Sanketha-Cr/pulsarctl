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

package test

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/docker/docker/client"
)

// SkipIfDockerUnavailable skips the test if Docker is not available
// or if running on s390x architecture
func SkipIfDockerUnavailable(t *testing.T) {
	// Skip Docker tests on s390x architecture
	if runtime.GOARCH == "s390x" {
		t.Skip("Skipping Docker tests on s390x architecture")
		return
	}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Skip("Docker client could not be created, skipping test:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = cli.Ping(ctx)
	if err != nil {
		t.Skip("Docker daemon is not available, skipping test:", err)
	}
}