/*
 * Flow CLI
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package project

import (
	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/flow/project/commands/deploy_contracts"
	"github.com/onflow/flow-cli/flow/project/commands/initialize"
	"github.com/onflow/flow-cli/flow/project/commands/start_emulator"
)

var Cmd = &cobra.Command{
	Use:              "project",
	Short:            "Manage your Cadence project",
	TraverseChildren: true,
}

func init() {
	Cmd.AddCommand(initialize.Cmd)
	Cmd.AddCommand(start_emulator.Cmd)
	Cmd.AddCommand(deploy_contracts.Cmd)
}
