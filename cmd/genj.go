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
	"fmt"
	"log"
	"strings"

	"github.com/chaopeng/to/bookmark"

	"github.com/spf13/cobra"
)

// genjCmd represents the genj command
var genjCmd = &cobra.Command{
	Use:   "genj",
	Short: "Gen the autocompletion script for the specified shell for J (Fish only)",
	Long:  `Gen the autocompletion script for the specified shell for J (Fish only)`,
	Run: func(cmd *cobra.Command, args []string) {
		ensureConfigFileDir()
		if len(args) != 1 || args[0] != "fish" {
			log.Fatalln("only support fish for now")
		}
		b := bookmark.ReadFromFile(dbFile)
		fmt.Println(genFishAutoForJ(b))
	},
}

func init() {
	rootCmd.AddCommand(genjCmd)
}

// genFishAutoForJ gen the list for `complete -f -c j -a {result}`
func genFishAutoForJ(b *bookmark.Bookmarks) string {
	var sb strings.Builder

	list := b.ListWithFilters(nil)
	for i, b := range list {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(b.Name)
		sb.WriteString("\t")
		sb.WriteString(dirShorten(b.Path, false))
	}

	return sb.String()
}
