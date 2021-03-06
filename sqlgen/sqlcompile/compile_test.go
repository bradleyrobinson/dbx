// Copyright (C) 2017 Space Monkey, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlcompile

import (
	"testing"

	"gopkg.in/spacemonkeygo/dbx.v1/sqlgen/sqltest"
	"gopkg.in/spacemonkeygo/dbx.v1/testutil"
)

func TestCompile(t *testing.T) {
	tw := testutil.Wrap(t)
	tw.Parallel()
	tw.Runp("fuzz render", testCompileFuzzRender)
	tw.Runp("idempotent", testCompileIdempotent)
	tw.Runp("fuzz normal form", testCompileFuzzNormalForm)
}

func testCompileFuzzRender(tw *testutil.T) {
	g := sqltest.NewGenerator(tw)

	for i := 0; i < 1000; i++ {
		sql := g.Gen()
		compiled := Compile(sql)
		exp := sql.Render()
		got := compiled.Render()

		if exp != got {
			tw.Logf("sql:      %#v", sql)
			tw.Logf("compiled: %#v", compiled)
			tw.Logf("exp:      %q", exp)
			tw.Logf("got:      %q", got)
			tw.Error()
		}
	}
}

func testCompileIdempotent(tw *testutil.T) {
	g := sqltest.NewGenerator(tw)

	for i := 0; i < 1000; i++ {
		sql := g.Gen()
		first := Compile(sql)
		second := Compile(first)

		if !sqlEqual(first, second) {
			tw.Logf("sql:    %#v", sql)
			tw.Logf("first:  %#v", first)
			tw.Logf("second: %#v", second)
			tw.Error()
		}
	}
}

func testCompileFuzzNormalForm(tw *testutil.T) {
	g := sqltest.NewGenerator(tw)

	for i := 0; i < 1000; i++ {
		sql := g.Gen()
		compiled := Compile(sql)

		if !sqlNormalForm(compiled) {
			tw.Logf("sql:      %#v", sql)
			tw.Logf("compiled: %#v", compiled)
			tw.Error()
		}
	}
}
