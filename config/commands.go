// Copyright © 2020 Krishna Iyer Easwaran
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version   = ""
	gitCommit = ""
	buildDate = ""
)

// VersionCommand prints version information to the output stream.
func (mgr *Manager) VersionCommand(root *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", root.Name())
			fmt.Println("----------------")
			fmt.Printf("Version: %s\n", version)
			fmt.Printf("Git Commit: %s\n", gitCommit)
			fmt.Printf("Build Date: %s\n", buildDate)
			fmt.Printf("Go version: %s\n", runtime.Version())
			fmt.Printf("OS/Arch: %s\n", runtime.GOOS+"/"+runtime.GOARCH)
		},
	}
}

// ConfigCommand prints the read config values to the output stream.
func (mgr *Manager) ConfigCommand(root *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Display raw config values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s: Raw config values\n", root.Name())
			fmt.Println("----------------")
			fmt.Println(mgr.AllSettings())
		},
	}
}
