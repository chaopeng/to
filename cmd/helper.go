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
	"os"
	"strings"
)

var (
	homeDir = os.Getenv("HOME")
)

func dirShorten(dir string, color bool) string {
	if strings.HasPrefix(dir, homeDir) {
		sb := strings.Builder{}
		if color {
			sb.WriteString(cyanBold.Sprintf("~"))
		} else {
			sb.WriteString("~")
		}
		sb.WriteString(strings.TrimPrefix(dir, homeDir))
		return sb.String()
	}
	return dir
}
