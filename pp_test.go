// Copyright 2019 PushPanel Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestReplaceAtIndex(t *testing.T) {
	got := replaceAtIndex("testme", 'a', 1)
	if got != "tastme" {
		t.Errorf("Replace does not work: %s", got)
	}
}
