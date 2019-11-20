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

package star

import (
	"testing"
	"time"

	"nanomsg.org/go/mangos/v2"
	. "nanomsg.org/go/mangos/v2/internal/test"
	_ "nanomsg.org/go/mangos/v2/transport/inproc"
)

func TestStarNonBlock(t *testing.T) {
	maxqlen := 2

	rp, err := NewSocket()
	MustSucceed(t, err)
	MustNotBeNil(t, rp)
	defer rp.Close()

	MustSucceed(t, rp.SetOption(mangos.OptionWriteQLen, maxqlen))
	MustSucceed(t, rp.Listen(AddrTestInp()))

	msg := []byte{'A', 'B', 'C'}

	start := time.Now()
	for i := 0; i < maxqlen*10; i++ {
		MustSucceed(t, rp.Send(msg))
	}
	end := time.Now()
	MustBeTrue(t, end.Sub(start) < time.Second/10)
}
