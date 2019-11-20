// Copyright 2019 The Mangos Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"testing"

	"nanomsg.org/go/mangos/v2"
)

// VerifyRaw verifies that the socket created is raw, and cannot be changed to cooked.
func VerifyRaw(t *testing.T, f func() (mangos.Socket, error)) {
	s, err := f()
	MustSucceed(t, err)
	defer s.Close()
	val, err := s.GetOption(mangos.OptionRaw)
	MustSucceed(t, err)
	if b, ok := val.(bool); ok {
		MustBeTrue(t, b)
	} else {
		t.Fatalf("Not a boolean")
	}

	err = s.SetOption(mangos.OptionRaw, false)
	MustFail(t, err)
}

// VerifyCooked verifies that the socket created is cooked, and cannot be changed to raw.
func VerifyCooked(t *testing.T, f func() (mangos.Socket, error)) {
	s, err := f()
	MustSucceed(t, err)
	defer s.Close()
	val, err := s.GetOption(mangos.OptionRaw)
	MustSucceed(t, err)
	if b, ok := val.(bool); ok {
		MustBeFalse(t, b)
	} else {
		t.Fatalf("Not a boolean")
	}

	err = s.SetOption(mangos.OptionRaw, true)
	MustFail(t, err)
}
