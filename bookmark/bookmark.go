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

// Package bookmark places bookmark function.
package bookmark

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// Bookmarks contains list, list-with-filter, save, delete and file feature.
// func will crash if error.
type Bookmarks struct {
	data map[string]string
}

func NewBookMarkForTesting() *Bookmarks {
	return &Bookmarks{
		data: map[string]string{},
	}
}

// ReadFromFile reads bookmark from file.
func ReadFromFile(file string) *Bookmarks {
	// Read the JSON file.
	jsonFile, err := os.Open(file)
	if err != nil {
		return &Bookmarks{
			data: map[string]string{},
		}
	}

	// Close the file when we're done.
	defer jsonFile.Close()

	// Read the contents of the file into a byte slice.
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v\n", err)
	}

	// Create a new hashmap to store the JSON data.
	data := make(map[string]string)

	// Unmarshal the JSON data into the hashmap.
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal the db file: %v\n", err)
	}

	return &Bookmarks{data}
}

// SaveToFile save the bookmark to file.
func (b *Bookmarks) SaveToFile(file string) {
	os.Remove(file)
	// Open the output file.
	outputFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}

	// Close the file when we're done.
	defer outputFile.Close()

	// Marshal the hashmap to JSON.
	jsonData, err := json.Marshal(b.data)
	if err != nil {
		log.Fatalf("Failed to marshal: %v\n", err)
	}

	// Write the JSON data to the output file.
	_, err = outputFile.Write(jsonData)
	if err != nil {
		log.Fatalf("Failed to write file: %v\n", err)
	}

	// Flush the output file.
	err = outputFile.Sync()
	if err != nil {
		log.Fatalf("Failed to flush file: %v\n", err)
	}
}

// Bookmark use as result in ListAll() and ListWithFilter()
type Bookmark struct {
	Name, Path string
}

// ListAll lists all saved bookmarks.
func (b *Bookmarks) ListWithFilters(filters []BookmarkFilter) []Bookmark {
	res := []Bookmark{}
	for k, v := range b.data {
		rejected := false
		bm := Bookmark{
			Name: k, Path: v,
		}
		for _, f := range filters {
			if !f.Filter(&bm) {
				rejected = true
				break
			}
		}
		if !rejected {
			res = append(res, bm)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Name < res[j].Name
	})
	return res
}

// Add a bookmark
func (b *Bookmarks) Add(name, path string) error {
	if _, exists := b.data[name]; exists {
		return alreadyExistsErr(name)
	}
	b.data[name] = path
	return nil
}

// Delete a bookmark
func (b *Bookmarks) Delete(name string) error {
	if _, exists := b.data[name]; !exists {
		return notFoundErr(name)
	}
	delete(b.data, name)
	return nil
}

// Match finds the matched bookmark with order.
// 1. exact match
// 2. shortest bookmark name with given as prefix, return error if more than 1.
func (b *Bookmarks) Match(name string) (*Bookmark, []Bookmark, error) {
	if path, exists := b.data[name]; exists {
		return &Bookmark{Name: name, Path: path}, nil, nil
	}

	res := []Bookmark{}
	for k, v := range b.data {
		if strings.HasPrefix(k, name) {
			res = append(res, Bookmark{Name: k, Path: v})
		}
	}

	if len(res) == 0 {
		return nil, nil, prefixNotFoundErr(name)
	}

	if len(res) == 1 {
		return &res[0], res, nil
	}

	sort.Slice(res, func(i, j int) bool {
		li := len(res[i].Name)
		lj := len(res[j].Name)
		if li == lj {
			return res[i].Name < res[j].Name
		}
		return li < lj
	})

	if len(res[0].Name) == len(res[1].Name) {
		return nil, res, moreThanOneMatchErr(name)
	}

	return &res[0], res, nil
}
