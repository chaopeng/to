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

package bookmark

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadFromFileFileNotExists(t *testing.T) {
	got := ReadFromFile(filepath.Join(os.TempDir(), "not-exists"))
	want := &Bookmarks{data: map[string]string{}}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Bookmarks{})); diff != "" {
		t.Errorf("-want +got: %v", diff)
	}
}

func TestSaveToFile(t *testing.T) {
	want := &Bookmarks{
		data: map[string]string{
			"aaa": "bbb",
		},
	}

	f, err := os.CreateTemp("", "bmtest")
	if err != nil {
		t.Fatalf("CreateTemp: %v", err)
	}

	file := f.Name()
	f.Close()

	defer os.Remove(file)

	want.SaveToFile(file)

	got := ReadFromFile(file)
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Bookmarks{})); diff != "" {
		t.Errorf("-want +got: %v", diff)
	}
}

func TestListAll(t *testing.T) {
	b := &Bookmarks{
		data: map[string]string{
			"aaa":  "bbb",
			"aaa1": "bbb1",
		},
	}

	got := b.ListWithFilters([]BookmarkFilter{})
	want := []Bookmark{
		{Name: "aaa", Path: "bbb"},
		{Name: "aaa1", Path: "bbb1"},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("-want +got: %v", diff)
	}
}

func TestListWithPrefixFilter(t *testing.T) {
	b := &Bookmarks{
		data: map[string]string{
			"aaa":  "bbb",
			"aaa1": "bbb1",
		},
	}

	got := b.ListWithFilters([]BookmarkFilter{
		NewPrefixFilter("a"),
	})
	want := []Bookmark{
		{Name: "aaa", Path: "bbb"},
		{Name: "aaa1", Path: "bbb1"},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("-want +got: %v", diff)
	}
}

func TestListWithChildrenFilter(t *testing.T) {
	b := &Bookmarks{
		data: map[string]string{
			"aaa":  "bbb",
			"aaa1": "bbb1",
			"aaa2": "abb",
		},
	}

	got := b.ListWithFilters([]BookmarkFilter{
		NewChildrenDirFilter("b"),
	})
	want := []Bookmark{
		{Name: "aaa", Path: "bbb"},
		{Name: "aaa1", Path: "bbb1"},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("-want +got: %v", diff)
	}
}

func TestListWithPrefixAndChildrenFilter(t *testing.T) {
	b := &Bookmarks{
		data: map[string]string{
			"aaa":  "bbb",
			"aaa1": "bbb1",
			"aaa2": "abb",
		},
	}

	got := b.ListWithFilters([]BookmarkFilter{
		NewPrefixFilter("a"),
		NewChildrenDirFilter("a"),
	})
	want := []Bookmark{
		{Name: "aaa2", Path: "abb"},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("-want +got: %v", diff)
	}
}

func TestAddFailed(t *testing.T) {
	b := &Bookmarks{
		data: map[string]string{
			"aaa": "bbb",
		},
	}

	if err := b.Add("aaa", "ccc"); err == nil || !IsErrType(err, AlreadyExists) {
		t.Errorf("want already exists error")
	}
}

func TestAdd(t *testing.T) {
	b := &Bookmarks{
		data: map[string]string{
			"aaa": "bbb",
		},
	}

	if err := b.Add("aaa1", "ccc"); err != nil {
		t.Fatalf("Add failed: %v", err)
	}

	want := &Bookmarks{
		data: map[string]string{
			"aaa":  "bbb",
			"aaa1": "ccc",
		},
	}
	if diff := cmp.Diff(want, b, cmp.AllowUnexported(Bookmarks{})); diff != "" {
		t.Errorf("-want +got: %v", diff)
	}
}

func TestDeleteFailed(t *testing.T) {
	b := &Bookmarks{
		data: map[string]string{
			"aaa": "bbb",
		},
	}

	if err := b.Delete("aaa1"); err == nil || !IsErrType(err, NotFound) {
		t.Errorf("want not found error")
	}
}

func TestDelete(t *testing.T) {
	b := &Bookmarks{
		data: map[string]string{
			"aaa": "bbb",
		},
	}

	if err := b.Delete("aaa"); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	want := &Bookmarks{
		data: map[string]string{},
	}
	if diff := cmp.Diff(want, b, cmp.AllowUnexported(Bookmarks{})); diff != "" {
		t.Errorf("-want +got: %v", diff)
	}
}

func TestMatch(t *testing.T) {
	b := &Bookmarks{
		data: map[string]string{
			"aaa":  "bbb",
			"aab1": "ccc",
			"aab2": "ddd",
			"abc":  "eee",
		},
	}

	tests := []struct {
		n       string
		input   string
		result1 *Bookmark
		result2 []Bookmark
		et      BookmarkErrType
	}{
		{n: "exact match", input: "aaa", result1: &Bookmark{Name: "aaa", Path: "bbb"}, et: NoErr},
		{n: "prefix not found", input: "b", et: PrefixNotFound},
		{
			n: "match shortest", input: "aa",
			result1: &Bookmark{Name: "aaa", Path: "bbb"},
			result2: []Bookmark{
				{Name: "aaa", Path: "bbb"}, {Name: "aab1", Path: "ccc"}, {Name: "aab2", Path: "ddd"}},
			et: NoErr},
		{
			n:       "found only 1 prefix match",
			input:   "ab",
			result1: &Bookmark{Name: "abc", Path: "eee"},
			result2: []Bookmark{{Name: "abc", Path: "eee"}},
			et:      NoErr,
		},
		{
			n:       "match more than 1",
			input:   "aab",
			result2: []Bookmark{{Name: "aab1", Path: "ccc"}, {Name: "aab2", Path: "ddd"}},
			et:      MoreThanOneMatch,
		},
	}

	for _, tc := range tests {
		t.Run(tc.n, func(t *testing.T) {
			got1, got2, err := b.Match(tc.input)
			if diff := cmp.Diff(tc.result1, got1); diff != "" {
				t.Errorf("-want +got: %v", diff)
			}
			if diff := cmp.Diff(tc.result2, got2); diff != "" {
				t.Errorf("-want +got: %v", diff)
			}
			if err != nil && tc.et == NoErr {
				t.Errorf("want no err, got %v", err)
			}
			if err == nil && tc.et != NoErr {
				t.Errorf("want err type %v, but no err", tc.et)
			}
			if err != nil {
				et := err.(*Err).errType
				if et != tc.et {
					t.Errorf("want err type %v, got err type %v", tc.et, et)
				}
			}
		})
	}
}
