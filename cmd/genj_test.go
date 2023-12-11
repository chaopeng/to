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
	"testing"

	"github.com/chaopeng/to/bookmark"

	"github.com/google/go-cmp/cmp"
)

func TestGenFishAutoForJ(t *testing.T) {
	b := bookmark.NewBookMarkForTesting()
	b.Add("a", "111")
	b.Add("b", "(home)/222")

	homeDir = "(home)"

	got := genFishAutoForJ(b)
	want := "a\t111\nb\t~/222"
	if d := cmp.Diff(want, got); d != "" {
		t.Errorf("-want, +got:\n%v", d)
	}
}

func TestGenFishAutoForJWithEmptyBookmark(t *testing.T) {
	b := bookmark.NewBookMarkForTesting()

	homeDir = "(home)"

	got := genFishAutoForJ(b)
	want := ""
	if d := cmp.Diff(want, got); d != "" {
		t.Errorf("-want, +got:\n%v", d)
	}
}
