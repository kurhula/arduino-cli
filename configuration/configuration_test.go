// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func tmpDirOrDie() string {
	dir, err := ioutil.TempDir(os.TempDir(), "cli_test")
	if err != nil {
		panic(fmt.Sprintf("error creating tmp dir: %v", err))
	}
	return dir
}

func TestSearchConfigTreeNotFound(t *testing.T) {
	tmp := tmpDirOrDie()
	require.Empty(t, searchConfigTree(tmp))
}

func TestSearchConfigTreeSameFolder(t *testing.T) {
	tmp := tmpDirOrDie()
	defer os.RemoveAll(tmp)
	_, err := os.Create(filepath.Join(tmp, "arduino-cli.yaml"))
	require.Nil(t, err)
	require.Equal(t, tmp, searchConfigTree(tmp))
}

func TestSearchConfigTreeInParent(t *testing.T) {
	tmp := tmpDirOrDie()
	defer os.RemoveAll(tmp)
	target := filepath.Join(tmp, "foo", "bar")
	err := os.MkdirAll(target, os.ModePerm)
	require.Nil(t, err)
	_, err = os.Create(filepath.Join(tmp, "arduino-cli.yaml"))
	require.Nil(t, err)
	require.Equal(t, tmp, searchConfigTree(target))
}

var result string

func BenchmarkSearchConfigTree(b *testing.B) {
	tmp := tmpDirOrDie()
	defer os.RemoveAll(tmp)
	target := filepath.Join(tmp, "foo", "bar", "baz")
	os.MkdirAll(target, os.ModePerm)

	var s string
	for n := 0; n < b.N; n++ {
		s = searchConfigTree(target)
	}
	result = s
}
