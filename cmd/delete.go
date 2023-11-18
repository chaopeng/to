// Copyright 2023 chaopeng@chaopeng.me
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

package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm", "del"},
	Short:   `Delete given bookmark.`,
	Long:    `Delete given bookmark.`,
	Run: func(cmd *cobra.Command, args []string) {
		ensureConfigFileDir()
		if len(args) != 1 {
			log.Fatalln("want exact 1 argument as bookmark name")
		}
		delete(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
