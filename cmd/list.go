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
	"os"

	"github.com/chaopeng/to/bookmark"

	"github.com/spf13/cobra"
)

var (
	currFlag bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   `List saved bookmarks.`,
	Long:    `List saved bookmarks.`,
	Run: func(cmd *cobra.Command, args []string) {
		ensureConfigFileDir()

		filters := []bookmark.BookmarkFilter{}
		var dir string
		var prefix string
		var err error

		if currFlag {
			dir, err = os.Getwd()
			if err != nil {
				log.Fatalf("pwd failed: %v\n", err)
			}
			filters = append(filters, bookmark.NewChildrenDirFilter(dir))
		}
		if arg != "" {
			prefix = arg
			filters = append(filters, bookmark.NewPrefixFilter(prefix))
		}
		listWithFilters(prefix, dir, filters)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&currFlag, "curr", "c", false, "only list bookmarks under current dir")
	listCmd.Flags().StringVarP(&arg, "filter", "f", "", "list bookmarks with given prefix")
}
