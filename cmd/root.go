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

// Package cmd store flags handling and logic.
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	arg string
)

var (
	dbDir  = filepath.Join(os.Getenv("HOME"), ".config", "to")
	dbFile = filepath.Join(dbDir, "db.json")
)

var rootCmd = &cobra.Command{
	Use:   "to",
	Short: "A dir bookmark tool",
	Long:  `This is a simple bookmark tools to help me remember dir and access them with shortcut.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func ensureConfigFileDir() {
	if _, err := os.Stat(dbDir); err != nil {
		err := os.MkdirAll(dbDir, 0755)
		if err != nil {
			log.Fatalf("Failed to creat to database dir: %v\n", err)
		}
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
