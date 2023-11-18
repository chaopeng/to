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
	"fmt"
	"log"
)

type BookmarkErrType int

const (
	NoErr BookmarkErrType = iota
	NotFound
	PrefixNotFound
	AlreadyExists
	MoreThanOneMatch
)

// Err for bookmark
type Err struct {
	errType BookmarkErrType
	message string
}

func (e *Err) Error() string {
	return e.message
}

func IsErrType(e error, t BookmarkErrType) bool {
	er, ok := e.(*Err)
	if !ok {
		log.Fatalf("Error is not BookmarkErr")
	}
	return er.errType == t
}

func notFoundErr(name string) *Err {
	return &Err{
		errType: NotFound,
		message: fmt.Sprintf("bookmark %v not found", name),
	}
}

func prefixNotFoundErr(name string) *Err {
	return &Err{
		errType: PrefixNotFound,
		message: fmt.Sprintf("bookmark with prefix %q not found", name),
	}
}

func alreadyExistsErr(name string) *Err {
	return &Err{
		errType: AlreadyExists,
		message: fmt.Sprintf("bookmark %v already exists", name),
	}
}

func moreThanOneMatchErr(name string) *Err {
	return &Err{
		errType: MoreThanOneMatch,
		message: fmt.Sprintf("bookmark with %q prefix has more than 1 matches", name),
	}
}
