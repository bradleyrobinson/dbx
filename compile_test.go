// Copyright (C) 2017 Space Monkey, Inc.

package main

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/spacemonkeygo/dbx.v1/testutil"
)

func TestCompilation(t *testing.T) {
	tw := testutil.Wrap(t)
	fileinfos, err := ioutil.ReadDir("testdata")
	tw.AssertNoError(err)

	for _, fileinfo := range fileinfos {
		tw.Run(fileinfo.Name(), func(t *testing.T) {
			path := filepath.Join("testdata", fileinfo.Name())
			testFile(testutil.Wrap(t), path)
		})
	}
}

func testFile(t *testutil.T, file string) {
	defer func() {
		if val := recover(); val != nil {
			t.Fatal(val)
		}
	}()

	dir, err := ioutil.TempDir("", "dbx")
	t.AssertNoError(err)
	defer os.RemoveAll(dir)

	dbx_source, err := ioutil.ReadFile(file)
	t.AssertNoError(err)
	t.Context("dbx", linedSource(dbx_source))
	d := loadDirectives(t, dbx_source)

	dialects := []string{"postgres", "sqlite3"}
	if other := d.lookup("dialects"); other != nil {
		dialects = other
		t.Logf("using dialects: %q", dialects)
	}

	err = golangCmd("", dialects, "", file, dir)
	t.AssertNoError(err)

	go_file := filepath.Join(dir, filepath.Base(file)+".go")
	go_source, err := ioutil.ReadFile(go_file)
	t.AssertNoError(err)
	t.Context("go", linedSource(go_source))

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, go_file, go_source, parser.AllErrors)
	t.AssertNoError(err)

	config := types.Config{
		Importer: importer.Default(),
	}
	_, err = config.Check(dir, fset, []*ast.File{f}, nil)

	if d.has("fail") {
		t.AssertError(err, d.get("fail"))
	} else {
		t.AssertNoError(err)
	}
}