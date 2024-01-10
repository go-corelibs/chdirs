// Copyright (c) 2023  The Go-Curses Authors
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

package chdirs

import (
	"os"
	"sync"
)

var (
	gPushPop = struct {
		stack []string
		sync.RWMutex
	}{}
	gMockWD = struct {
		stack []string
		sync.RWMutex
	}{}
)

// Push notes the current working directory and changes directory to the given
// path, use Pop to return to the previous working directory
func Push(path string) (err error) {
	gPushPop.Lock()
	defer gPushPop.Unlock()
	var cwd string
	if cwd, err = os.Getwd(); err == nil {
		if err = os.Chdir(path); err == nil {
			gPushPop.stack = append(gPushPop.stack, cwd)
		}
	}
	return
}

// Pop removes the last working directory from the stack and changes directory
// to it
func Pop() (err error) {
	gPushPop.Lock()
	defer gPushPop.Unlock()
	if last := len(gPushPop.stack) - 1; last >= 0 {
		path := gPushPop.stack[last]
		if last == 0 {
			gPushPop.stack = make([]string, 0)
		} else {
			gPushPop.stack = gPushPop.stack[:last]
		}
		err = os.Chdir(path)
	}
	return
}

// Stack returns the current Push stack
func Stack() (stack []string) {
	gPushPop.RLock()
	defer gPushPop.RUnlock()
	stack = gPushPop.stack[:]
	return
}

// MockBadWD is intended to be used during unit testing and not something
// generally useful. MockBadWD changes to a new temp directory and then
// deletes the directory. This leaves the running code in a case where
// calling os.Getwd will throw and error, which in turn causes filepath.Abs
// to throw an error
//
// Call UnMockBadWD to restore the working directory to the location the
// runtime was at when MockBadWD was called
func MockBadWD() (err error) {
	gMockWD.Lock()
	defer gMockWD.Unlock()
	if len(gMockWD.stack) == 0 {
		// working directory not mocked
		var cwd, tmpDir string
		if cwd, err = os.Getwd(); err == nil {
			if tmpDir, err = os.MkdirTemp("", "corelibs-chdirs-mock-bad-wd.*.d"); err == nil {
				gMockWD.stack = append(gMockWD.stack, cwd)
				err = os.RemoveAll(tmpDir)
			}
		}
	}
	return
}

// UnMockBadWD restores the runtime to the original working directory when
// MockBadWD was called
func UnMockBadWD() (err error) {
	gMockWD.Lock()
	defer gMockWD.Unlock()
	if len(gMockWD.stack) > 0 {
		// should only have length of 1
		owd := gMockWD.stack[0]
		gMockWD.stack = []string{}
		err = os.Chdir(owd)
	}
	return
}
