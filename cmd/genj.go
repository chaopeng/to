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
	"bytes"
	"fmt"
	"log"
	"text/template"

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

func genFishAutoForJ(b *bookmark.Bookmarks) string {
	t, err := template.New("fishj").Parse(fishJTemplate)
	if err != nil {
		log.Fatalf("template parse failed: %v", err)
	}

	list := b.ListWithFilters(nil)
	for i, b := range list {
		list[i].Path = dirShorten(b.Path, false)
	}

	var out bytes.Buffer
	t.Execute(&out, list)
	return out.String()
}

var fishJTemplate = `#!/usr/bin/env fish
{{if .}}
set -l bookmark_keys{{ range . }} \
  {{ .Name }}{{ end }}
{{end}}

# cleanup current autocomplete
complete -c j -e

# list all bookmarks
{{ range . }}
complete -f -c j -n "not __fish_seen_subcommand_from $bookmark_keys" -a '{{ .Name }}' -d '{{ .Path }}'
{{ end }}
`
