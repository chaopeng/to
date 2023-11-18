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
	"os"
	"regexp"
	"strings"

	"github.com/chaopeng/to/bookmark"

	"github.com/fatih/color"
)

var (
	splitLine = "=================================================="
	bold      = color.New(color.Bold)
	blueBold  = color.New(color.FgBlue, color.Bold)
	cyanBold  = color.New(color.FgCyan, color.Bold)
)

func listWithFilters(prefix string, dir string, filters []bookmark.BookmarkFilter) {
	b := bookmark.ReadFromFile(dbFile)
	res := b.ListWithFilters(filters)
	bold.Printf("Found %v saved bookmarks", len(res))
	if prefix != "" {
		bold.Printf(" with prefix %q", prefix)
	}
	if dir != "" {
		bold.Printf(" under dir %q", dirShorten(dir, false))
	}
	bold.Println()
	fmt.Println(splitLine)
	printBookmarks(res, func(b *bookmark.Bookmark) string {
		sb := strings.Builder{}
		sb.WriteString(blueBold.Sprint(prefix))
		sb.WriteString(strings.TrimPrefix(b.Name, prefix))
		sb.WriteString(": ")
		sb.WriteString(dirShorten(b.Path, true))
		sb.WriteString("\n")
		return sb.String()
	})
}

func printBookmarks(l []bookmark.Bookmark, formatter func(b *bookmark.Bookmark) string) {
	sb := strings.Builder{}
	for _, b := range l {
		sb.WriteString(formatter(&b))
	}
	fmt.Println(sb.String())
}

func save(name string) {
	validateBookmarkName(name)
	curr, err := os.Getwd()
	if err != nil {
		log.Fatalf("pwd failed: %v\n", err)
	}
	b := bookmark.ReadFromFile(dbFile)
	if err := b.Add(name, curr); err != nil {
		log.Fatalf("Add bookmark failed: %v\n", err)
	}
	b.SaveToFile(dbFile)
}

func delete(name string) {
	validateBookmarkName(name)
	b := bookmark.ReadFromFile(dbFile)
	if err := b.Delete(name); err != nil {
		log.Fatalf("Delete bookmark failed: %v\n", err)
	}
	b.SaveToFile(dbFile)
}

func findMatchedDir(name string) string {
	validateBookmarkName(name)
	b := bookmark.ReadFromFile(dbFile)
	r1, _, err := b.Match(name)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return r1.Path
}

var bookmarkRE = regexp.MustCompile("^[a-z][a-z0-9]*$")

func validateBookmarkName(name string) {
	if !bookmarkRE.MatchString(name) {
		log.Fatalf("Given bookmark name %v is invalid\n", name)
	}
}
