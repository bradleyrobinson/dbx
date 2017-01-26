// Code generated by go-bindata.
// sources:
// golang.create.tmpl
// golang.delete-all.tmpl
// golang.delete-world.tmpl
// golang.delete.tmpl
// golang.footer.tmpl
// golang.get-all.tmpl
// golang.get-count.tmpl
// golang.get-has.tmpl
// golang.get-last.tmpl
// golang.get-limitoffset.tmpl
// golang.get-paged.tmpl
// golang.get.tmpl
// golang.header.tmpl
// golang.misc.tmpl
// golang.update.tmpl
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _golangCreateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x92\x41\x6b\xdc\x30\x10\x85\xcf\xd6\xaf\x98\x2e\x14\x6c\xd8\xd5\x0f\x08\xec\xa1\x84\x14\x02\x61\x21\xd9\x94\x1e\x8d\x62\x8d\x8c\x5a\xaf\xe4\x8e\x66\xeb\x5d\x84\xfe\x7b\x91\x55\x17\xe7\xd0\x1c\xf6\x60\xb0\x2d\xf1\xbe\xf7\xde\x4c\x8c\x3b\xd0\x68\xac\x43\xd8\x04\xdb\x3b\xc5\x67\xc2\x0d\xec\x52\x12\xf7\x84\x8a\x31\x46\x90\xc7\xb3\x31\xf6\x02\x29\xd5\x31\x42\xc7\x97\x51\x91\x3a\x81\xfc\x42\x7d\x80\x94\x1a\xa8\x45\x15\x23\xfc\xfd\xfb\x82\x7c\x26\x07\x29\x6d\x01\x89\xf2\xe3\xa9\x11\x99\x83\x4e\xcf\xc2\x62\x0d\x7d\xf3\xfa\x3a\xf3\x62\x16\xd9\x81\x35\x20\x0f\x88\x3a\x1c\xfc\x04\x29\x89\xaa\x6d\x9d\x9f\xe0\x6e\x0f\x07\x3f\xd5\x8d\xfc\xf6\x7a\x5f\x37\x33\x6f\x91\xcb\xef\xd6\x59\x76\x38\x81\xfc\x6a\x71\xd0\xd9\x95\x10\x55\xe7\x5d\x60\x68\xdb\xc0\x27\xde\x67\x83\x64\x1d\x1b\xd8\x7c\xfe\xb5\x01\x79\x7c\x7e\x9a\xe5\xfd\xdb\x0f\x39\xf8\xfe\xc8\x27\xae\xcb\xd5\x2d\xc4\x08\x8a\xfa\x95\x58\x23\x0a\xc5\xe4\x2e\xc6\xd1\x13\x87\x12\xd3\xba\x1e\x56\x16\x56\xe9\x45\x95\xd3\xef\x21\xeb\x6b\xb2\xbf\x91\xe4\xf3\x19\xe9\xfa\xe2\xa7\x0f\x38\xf2\xd8\x29\x97\x5b\x56\x5a\x93\x37\x50\x9b\x41\x31\xa3\x5b\x84\x9b\xd9\x4c\x65\xcd\xdc\xed\xa7\x3d\x38\x3b\x40\x14\x55\x45\x85\xeb\xec\xb0\x85\xef\xa4\xc6\x07\xa2\x1a\x89\x1a\x51\x25\xb1\x1c\x2e\xb8\xd5\x84\x9c\x1d\x4a\x97\x43\xc0\x52\x66\xdb\x12\x86\x32\xba\xbb\x77\xee\x1f\x2e\xd8\xfd\xc7\xf9\x2d\x9e\xda\x76\xfc\xf9\x0f\x33\x43\xe5\x93\x0a\xfc\xe8\x02\x12\x3f\xea\xfa\xe6\x94\xd9\x72\x8f\x9c\xc5\xf2\xee\x96\xb0\xf2\xf5\x3a\x62\x5e\xe0\x8e\x2f\x5b\xc8\xec\xf7\x3b\x94\xc4\xea\xe3\x4f\x00\x00\x00\xff\xff\xad\x71\xb5\x3f\x14\x03\x00\x00")

func golangCreateTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangCreateTmpl,
		"golang.create.tmpl",
	)
}

func golangCreateTmpl() (*asset, error) {
	bytes, err := golangCreateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.create.tmpl", size: 788, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangDeleteAllTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\x31\x6b\xfb\x30\x10\x47\x67\xdd\xa7\xf8\xfd\x0d\x7f\x70\x20\x31\x1d\x4a\x87\x82\x87\x40\xb3\x75\x69\x33\x74\x34\x8e\x7d\x0a\x2a\x8e\x9c\x9e\xcf\x69\xca\xa1\xef\x5e\xec\x84\xb6\x5b\xe9\xa0\x49\x7a\xef\x9d\xce\x6c\x85\x96\x7d\x88\x8c\x6c\x08\xfb\x58\xeb\x28\x9c\x61\x95\x12\x3d\x70\xc7\xca\x66\x28\xb6\xa3\xf7\xe1\x8c\x94\x72\x33\x34\x7a\x3e\xd6\x52\x1f\x50\xac\x65\x3f\x20\xa5\x05\x72\x72\x4d\x3f\x46\x45\x88\x7a\x77\xbb\x04\x8b\x4c\xa7\x97\x05\x4d\x7e\x8e\xed\x2c\xa4\x9f\xb1\x5d\xdf\x7e\xcc\x1d\x9b\xe0\x38\x28\xaa\x6a\xd0\x83\x96\x66\x38\x4a\x88\xea\x91\xfd\x7f\xcb\x50\x6c\x9f\x1e\x91\x12\xb9\x7e\xf7\x5a\x74\xfd\x7e\xab\x07\xcd\x2f\x4f\x97\x30\x83\x0f\xdc\xb5\xa7\xba\x1b\xf9\x7b\x20\x22\x57\x55\xc2\xc3\x65\x92\xfb\x12\x13\xdb\x4a\x38\xb1\x14\x9b\x33\x37\xbf\xf2\x2e\xf8\x19\xfd\x57\x22\x86\x0e\x46\xce\x09\xeb\x28\x11\x37\x4b\xbc\x48\x7d\xdc\x88\xe4\x2c\xb2\x20\x97\xe8\xfa\xf9\x4b\xac\xc4\x5c\x2e\x9e\xfb\xf7\x61\xed\x3d\x37\xca\x6d\xfe\x67\xe1\xf5\xee\xea\x8d\xa1\xa3\x44\x66\x5f\x8b\xfc\x0c\x00\x00\xff\xff\xdd\x30\x4b\x51\xb5\x01\x00\x00")

func golangDeleteAllTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangDeleteAllTmpl,
		"golang.delete-all.tmpl",
	)
}

func golangDeleteAllTmpl() (*asset, error) {
	bytes, err := golangDeleteAllTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.delete-all.tmpl", size: 437, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangDeleteWorldTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\x31\x6b\xf3\x30\x10\x40\x67\xdd\xaf\xb8\xcf\xf0\x81\x4d\x13\xd1\xa1\x74\xf3\x10\xda\x6c\x5d\x9a\x0e\x1d\x83\x63\x9f\x83\x8a\x90\x92\xd3\x39\x75\x11\xfa\xef\xc5\x72\x1c\xb2\x76\x10\x02\xa1\x7b\x4f\x7a\x31\xae\xb1\xa3\xde\x38\xc2\x22\x98\xa3\x6b\x64\x60\x2a\x70\x9d\x12\xbc\x92\x25\xa1\x8d\xb5\x65\x2b\x23\xb6\xde\x09\x8d\xa2\x5f\xe6\xbd\xc2\x12\x54\xeb\x07\x27\x68\x9c\x3c\x3f\xad\x90\x98\xa7\xe5\xb9\x82\x09\x4a\xae\xcb\x14\xb8\x37\x1c\x7c\xf7\x93\xe1\x11\xd4\xa5\x61\xdc\xef\x99\x02\x86\xb3\xd5\x3b\x0a\x83\x95\xe5\xf4\x0e\x3c\x03\xb8\x71\x47\x42\xfd\xf1\xfe\x16\x30\x25\x50\x79\x70\x76\xd6\xe8\x0f\x5f\xba\x63\x73\x21\xd6\xdb\x91\xda\x32\x46\x3c\xb1\x71\xd2\x63\xf1\xff\x5c\xa0\xc6\x94\x2a\x50\xa6\xcf\xd7\xff\xd5\xe8\x8c\xc5\x08\x4a\x31\xc9\xc0\x0e\x1f\x57\xf8\xc9\xcd\x69\xcb\x5c\x12\x73\x05\x2a\xc1\x24\xc8\x6f\x58\x14\xd9\xa7\x77\xfe\x3b\x6c\xfa\x9e\x5a\xa1\xae\xfc\x23\xf2\x1a\xeb\xa1\x5e\xbe\x77\xab\x34\x45\x5a\x06\xaf\x52\x67\x2c\x24\x88\xf1\x56\xf1\x37\x00\x00\xff\xff\x1d\x0d\x4b\x82\xa7\x01\x00\x00")

func golangDeleteWorldTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangDeleteWorldTmpl,
		"golang.delete-world.tmpl",
	)
}

func golangDeleteWorldTmpl() (*asset, error) {
	bytes, err := golangDeleteWorldTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.delete-world.tmpl", size: 423, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangDeleteTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x90\xb1\x6a\xf3\x30\x14\x85\x67\xe9\x29\xce\x6f\xf8\xc1\x01\xc7\x74\x2e\xa4\x10\x68\xb6\x2e\x6d\x86\x8e\xc6\xb1\xae\x82\x8a\x22\xa5\x57\x72\x9a\x72\xd1\xbb\x17\x3b\x21\xed\x56\xe8\xa0\x49\xe7\x3b\xdf\xe1\x8a\x2c\x61\xc8\xba\x40\xa8\x92\xdb\x87\x3e\x8f\x4c\x15\x96\xa5\xe8\x47\xf2\x94\x49\x04\xed\x76\xb4\xd6\x9d\x51\x4a\x2d\x82\x21\x9f\x8f\x3d\xf7\x07\xb4\x6b\xde\x27\x94\xb2\x40\xad\x95\x99\xc3\x06\xbb\x18\x7d\x03\x62\x9e\x5e\xe4\x85\x9e\x04\x14\xcc\xdc\xa8\x7f\xda\x76\xd1\x7c\xce\x22\xd1\x6a\x88\x21\x65\x74\x5d\xca\x87\xbc\x12\xc1\x91\x5d\xc8\x16\xd5\xff\xf7\x0a\xed\xf6\xf9\x09\xa5\x68\x15\x77\x6f\xad\x8f\xfb\x6d\x3e\xe4\xfa\x12\x6d\x20\x02\xeb\xc8\x9b\x53\xef\x47\xfa\x5e\xa4\xb5\xea\x3a\xa6\x74\x59\x72\xbf\xc2\xc4\x1a\x76\x27\xe2\x76\x73\xa6\xe1\x57\x5e\x39\x3b\xa3\xff\x56\x08\xce\x43\xb4\x52\x4c\x79\xe4\x00\xdb\xfb\x44\x0d\x5e\xb9\x3f\x6e\x98\x6b\x62\x5e\x68\x55\x66\xe1\x10\xc7\x90\x6f\xca\x79\x40\xfb\x12\x3f\xd2\xda\x5a\x1a\x32\x99\xfa\x4f\xbd\xd7\xff\x6b\x3d\x1e\x70\xd7\x4c\xac\x2e\x5a\xe4\x76\xd9\xaf\x00\x00\x00\xff\xff\xec\x7c\x37\xba\xc7\x01\x00\x00")

func golangDeleteTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangDeleteTmpl,
		"golang.delete.tmpl",
	)
}

func golangDeleteTmpl() (*asset, error) {
	bytes, err := golangDeleteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.delete.tmpl", size: 455, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xce\x31\x0b\xc2\x30\x10\x05\xe0\xb9\xf7\x2b\x6e\x6c\x05\xf3\x23\x44\xdc\x9c\xcc\xe0\x1a\x93\xb3\x2d\xe8\x59\x93\x13\x23\x47\xfe\xbb\x14\x21\x2e\xd9\xde\xf0\xde\xc7\x03\xf9\x2c\x84\xf6\x7c\x24\x99\x1e\x21\xe1\xcc\x42\xf1\xea\x3c\xa1\x82\xea\x16\xa3\xe3\x91\xd0\x1c\x5e\xec\x13\x96\x02\x9d\x2a\x9a\x35\xa8\x22\x71\x58\x53\x81\x1f\x22\xb9\x85\x74\xd5\xae\xc5\xfd\xae\x59\x3c\xf9\x89\xee\xae\x1f\x30\x49\x9c\x79\x6c\x2d\xc3\xa5\xb9\xac\x20\x40\xf7\x8e\x6e\xb1\xb9\x97\x8c\x9b\xf4\xbc\x19\x9b\x87\xff\x31\x28\xf0\x0d\x00\x00\xff\xff\xec\xf5\x69\xba\xf0\x00\x00\x00")

func golangFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangFooterTmpl,
		"golang.footer.tmpl",
	)
}

func golangFooterTmpl() (*asset, error) {
	bytes, err := golangFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.footer.tmpl", size: 240, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetAllTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x92\x4d\xeb\xdc\x20\x10\xc6\xcf\xfa\x29\xa6\x81\x82\x81\xfd\xe7\x03\x6c\xc9\xa1\x94\xd2\x4b\x29\x6c\x73\xe8\x31\xb8\x71\x0c\x16\x57\xd3\x89\xd9\x17\x06\xbf\x7b\x31\xd9\xdd\xf6\x50\x16\x7a\xf0\xa0\xce\x3c\xcf\x6f\x1e\x65\x7e\x03\x83\xd6\x05\x84\x6a\x76\x63\xd0\x69\x21\xac\xe0\x2d\x67\xf9\x05\x13\x33\x34\xdd\x62\xad\xbb\x42\xce\x8a\x19\x86\x74\x9d\x34\xe9\x13\x34\x1f\x69\x9c\x21\xe7\x1a\x94\x14\x14\x2f\x33\x30\xc3\xec\xdd\x80\xd1\x42\xf3\x3d\x5e\x20\xe7\x1d\x20\x51\x59\x91\x6a\x59\x8c\x30\x98\x55\x59\xfe\xed\x7a\x8c\xe6\x56\x41\xce\x2c\xc5\x10\xc3\x9c\xa0\xef\xe7\x74\x4a\x2d\x33\x4c\xe4\x42\xb2\x50\xbd\xff\x55\x41\xd3\x1d\xbe\x42\xce\x52\xc4\xe3\xcf\xc6\xc7\xb1\x4b\xa7\xa4\xb6\xd2\x5d\xf1\xb6\x0e\xbd\x39\x6b\xbf\xe0\x1f\x36\x29\x45\xdf\x17\xb8\x8d\x64\xdf\x42\x69\x36\xe4\xce\x48\xcd\x61\x41\xba\xbd\x52\x28\x02\xc2\xd9\xb5\xf5\x5d\x0b\xc1\x79\x60\x29\x04\x61\x5a\x28\x94\xed\x0e\x7e\x90\x9e\x3e\x13\x29\x24\xaa\xa5\xc8\x52\x18\xb4\x48\xb0\x99\x36\x9f\x7c\x9c\x51\x15\x0a\x1b\x9f\x87\xdf\xf0\x9a\x54\xbd\x2a\x31\x83\x0b\x2e\x05\xbc\x3c\x12\x93\x42\x14\xb7\xf6\x51\xdc\x0d\x3a\x94\xd8\xb5\x31\x14\x2d\x28\xeb\x75\x4a\x18\xd6\xf2\x7a\x9d\x50\xfc\x83\xf0\x05\x62\x61\xdc\x9e\xab\x05\x3d\x4d\x18\x8c\xda\xf2\x29\x26\x34\x3e\x38\xb6\x61\xee\xca\xfb\x27\x4e\xd1\xa9\x3f\xfc\x57\x20\xf7\xcb\xcd\x24\x38\x2f\xb3\x64\x7e\x7e\x84\xdf\x01\x00\x00\xff\xff\x69\xb2\x42\xa6\x7e\x02\x00\x00")

func golangGetAllTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetAllTmpl,
		"golang.get-all.tmpl",
	)
}

func golangGetAllTmpl() (*asset, error) {
	bytes, err := golangGetAllTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-all.tmpl", size: 638, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetCountTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x8f\x3f\x4f\xc3\x30\x10\xc5\x67\xdf\xa7\x78\x44\x02\x25\x52\x1a\x31\x20\xb6\x0e\x08\xb1\xb1\x94\x0c\x8c\x55\xda\x38\x95\x51\x6a\x87\x8b\x5d\x5a\x9d\xee\xbb\x23\x07\xf1\x67\x63\xf0\xe4\xf7\x7b\xf7\x7e\x22\x2b\xf4\x76\x70\xde\xa2\x98\xdd\xc1\x77\x31\xb1\x2d\xb0\x52\xa5\xc7\x90\x7c\x14\x41\xd3\xa6\x61\x70\x67\xa8\x96\x22\xd8\xc7\xf3\xd4\x71\x77\x44\xf3\xc0\x87\x19\xaa\x15\x4a\x32\xfb\x9c\x85\xf3\xf1\xfe\xae\x86\x65\xce\x2f\x70\x45\xb9\xde\xfa\x7e\xe9\xa3\xbf\xb7\x76\xa1\xbf\x14\x50\x95\xcc\xfa\x39\x62\xbb\x9d\xe3\x31\xae\x45\x30\xb1\xf3\x71\x40\x71\xfd\x5e\xa0\x69\x37\xcf\x50\x25\x13\x76\x6f\xcd\x18\x0e\x6d\x3c\xc6\xf2\x2b\x5a\x43\x04\x83\xb3\x63\x7f\xea\xc6\x64\x7f\xf7\x10\x99\xbc\x60\x8d\xcc\xf4\xec\x4e\x96\x9b\x4d\xb2\x7c\x79\x09\x1f\xff\xb1\x4d\xbb\xef\x7c\x79\xb3\xe8\x54\x64\xdc\xb0\xc8\x5c\xad\xe1\xdd\x08\x21\x63\xd8\xc6\xc4\x1e\xb7\x35\x5e\xb9\x9b\x9e\x98\x4b\xcb\x5c\x91\x51\xa2\xef\xbf\x05\xae\x33\x41\x4a\x22\x3f\xfe\x9f\x01\x00\x00\xff\xff\x73\xce\xbf\x3a\x6b\x01\x00\x00")

func golangGetCountTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetCountTmpl,
		"golang.get-count.tmpl",
	)
}

func golangGetCountTmpl() (*asset, error) {
	bytes, err := golangGetCountTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-count.tmpl", size: 363, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetHasTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x8f\xbd\x6e\xeb\x30\x0c\x85\x67\xf1\x29\xce\x35\x70\x0b\x1b\x70\xfc\x06\x19\x3a\x14\xe8\xd0\x25\xf5\xd0\x31\x50\x62\x29\x51\xe1\x48\x2e\x25\xa7\x09\x08\xbe\x7b\xe1\x14\xfd\xd9\x3a\x70\x22\xbf\xc3\xf3\x89\xac\x30\x38\x1f\xa2\x43\x95\xc3\x21\xda\x32\xb3\xab\xb0\x52\xa5\x47\x9b\x45\xd0\xf5\xb3\xf7\xe1\x02\xd5\x5a\x04\xfb\x72\x99\x2c\xdb\x13\xba\x7b\x3e\x64\xa8\x36\xa8\xc9\x1c\x6d\xc6\x2e\xa5\xb1\x85\x63\x5e\x26\x71\x43\x4b\xb2\x8b\xc3\x2d\x8a\x7e\xbf\xd9\xa5\xe1\x5a\x41\x55\xc8\xec\x53\xcc\x05\xdb\x6d\x2e\xa7\xb2\x16\xc1\xc4\x21\x16\x8f\xea\xff\x5b\x85\xae\xdf\x3c\x41\x95\x4c\xda\xbd\x76\x63\x3a\xf4\xe5\x54\xea\xcf\xd3\x16\x22\xf0\xc1\x8d\xc3\xd9\x8e\xb3\xfb\x29\x43\x64\x96\x06\x6b\x2c\xcc\xc0\xe1\xec\xb8\xdb\xcc\x8e\xaf\xcf\xe9\xfd\x2f\xb6\xeb\xf7\x36\xd6\x77\x47\x9b\x1b\x32\xc1\xdf\x54\xfe\xad\x11\xc3\x08\x21\x63\xd8\x95\x99\x23\xbc\x1d\xb3\x6b\xf1\xc2\x76\x7a\x60\xae\x1d\x73\x43\x46\xe9\x6b\x7d\xb4\xb9\x5d\x10\x52\x12\xf9\xd6\xff\x08\x00\x00\xff\xff\x88\xaa\xbc\x96\x65\x01\x00\x00")

func golangGetHasTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetHasTmpl,
		"golang.get-has.tmpl",
	)
}

func golangGetHasTmpl() (*asset, error) {
	bytes, err := golangGetHasTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-has.tmpl", size: 357, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetLastTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x90\x31\x6f\xf2\x40\x0c\x86\xe7\xf3\xaf\xf0\x17\xe9\x93\x12\x09\x32\x55\xdd\x98\xaa\x6e\x2c\x40\xa5\x8e\xe8\x48\x9c\xe8\x4a\xf0\xa5\x8e\x69\xa1\xd6\xfd\xf7\x2a\x41\x54\x61\x38\xdd\xf2\xbe\x7e\x1e\xdb\x6c\x89\x35\x35\x81\x09\xb3\x21\xb4\xec\xf5\x2c\x94\xe1\x32\x25\x68\x49\xd7\x7e\x50\x33\x2c\xb7\xa4\x67\xe1\xf2\xed\xda\x13\xa6\x94\x57\x7a\xc1\x2a\xb2\xd2\x45\xcb\x97\xdb\xbf\x00\xd7\x1f\x31\xb0\x3e\x3f\x15\x98\x83\x33\xc3\xde\x8b\x3f\xdd\xbb\x98\xd2\x02\x49\x64\x7c\x51\x0a\x18\xb9\xc4\xf5\x04\x82\xb9\xc4\x21\xd6\xd7\x89\x6f\xe0\xaa\xc8\x83\xe2\x7e\x3f\xe8\x49\x57\xe3\x44\x09\xac\x0d\x66\xff\x3f\x33\x2c\x77\x9b\x35\xa6\x04\x2e\x1e\x3e\xca\x2e\xb6\x3b\x3d\x69\x7e\x8b\x2e\xb0\x3f\x16\x30\x39\x04\x0e\x3a\x53\x00\x37\x2a\xac\x70\xec\xd4\x12\xbe\x48\xca\xcd\x99\xe4\xba\x8d\xdf\xf3\x6e\xb9\xab\x3c\xe7\x66\xe8\xeb\x5a\x62\x83\x79\xd3\x79\x55\xe2\xfb\xa4\x02\x53\x2a\xc0\x85\x66\xda\xe8\xdf\x0a\x39\x74\x68\xe0\x9c\xdc\x40\x66\xf8\x43\x12\x1f\x76\x7f\x17\xdf\xbf\x8a\xe4\x24\x52\x80\x4b\x30\xcb\x7a\x69\x1f\xa2\x1c\x3a\x48\x60\xf6\x77\xa0\xdf\x00\x00\x00\xff\xff\x77\x00\x17\x29\xa5\x01\x00\x00")

func golangGetLastTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetLastTmpl,
		"golang.get-last.tmpl",
	)
}

func golangGetLastTmpl() (*asset, error) {
	bytes, err := golangGetLastTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-last.tmpl", size: 421, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetLimitoffsetTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x92\x41\x8b\xd4\x30\x14\xc7\xcf\xc9\xa7\xf8\x5b\x10\x5a\xe8\xf6\x24\x1e\x56\x7a\x10\x11\x2f\x22\xac\x73\xf0\x38\x64\xdb\x97\x21\x92\x26\xf5\x35\xdd\x99\x25\xbe\xef\x2e\x69\x77\x06\x0f\x22\xe8\xa1\xd0\x24\x8f\xdf\xff\x97\xf7\x92\xf3\x1d\x46\xb2\x2e\x10\xaa\xc5\x9d\x82\x49\x2b\x53\x85\x3b\x11\xfd\x89\x52\xce\xe8\x0e\xab\xb5\xee\x02\x91\x3a\x67\x0c\xe9\x32\x1b\x36\x13\xba\xf7\x7c\x5a\x44\x5a\xad\xbc\x9b\x5c\x82\x0b\xa9\x45\xb4\x76\xa1\xed\xff\xed\x9b\x06\xb5\x56\x1c\xcf\x0b\x72\xc6\xe2\xdd\x40\xd1\xa2\xfb\x1a\xcf\x10\x69\x41\xcc\xe5\x8b\xdc\xe8\x62\x40\x61\xdc\x22\xf5\xef\x3a\x8f\x71\x7c\xae\x20\x92\xb5\x1a\x62\x58\x12\x8e\xc7\x25\x4d\xa9\xcf\x19\x33\xbb\x90\x2c\xaa\xd7\x3f\x2a\x74\x87\x87\xcf\x10\xd1\x2a\x3e\x7e\xef\x7c\x3c\x1d\xd2\x94\xea\xbd\xb4\x2d\xd9\xd6\x91\x1f\x9f\x8c\x5f\x69\x97\xc6\x4f\x0c\x71\x9a\x0c\x44\x36\xf5\xab\x76\xa3\xb5\x3a\x1e\x8b\xf1\xae\x77\xdf\xa3\x10\x47\x76\x4f\xc4\xdd\xc3\x4a\xfc\xfc\x7f\x58\xe5\xec\x06\x7c\xd5\x23\x38\x8f\xac\x95\x62\x4a\x2b\x87\xb2\x6c\xf1\x8d\xcd\xfc\x91\xb9\x26\xe6\x46\x2b\xd1\x6a\x24\x4b\x8c\x5d\xa5\xfb\xe0\xe3\x42\x75\x71\xb3\xf1\xb6\xf9\x85\x2e\xa9\x6e\x36\x52\xce\x70\xc1\xa5\x40\xe7\x6b\x73\xb5\x52\x25\xad\xbf\x16\x1f\x06\x13\xca\xe8\xcc\x38\x72\xb4\xa8\xad\x37\x29\x51\xd8\xca\x1b\x88\x34\x5a\xfd\xc1\xf0\x2f\x8a\xc5\x71\x9f\x6c\x0f\x33\xcf\x14\xc6\x7a\xef\x5a\x09\xe1\xd3\xd5\x63\xbf\xcc\x0b\xf9\xfe\xa6\x53\x38\xcd\xbb\x7f\x6a\xc8\xcb\xe1\x1e\x12\x9c\xd7\xa2\x73\xbe\xbd\x99\x5f\x01\x00\x00\xff\xff\x7c\x87\x5b\x55\xc2\x02\x00\x00")

func golangGetLimitoffsetTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetLimitoffsetTmpl,
		"golang.get-limitoffset.tmpl",
	)
}

func golangGetLimitoffsetTmpl() (*asset, error) {
	bytes, err := golangGetLimitoffsetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-limitoffset.tmpl", size: 706, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetPagedTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x93\x4f\x8b\xdb\x30\x10\xc5\xcf\xd2\xa7\x98\x0a\x0a\x36\x78\xcd\x9e\xb7\xb8\x50\x4a\xe9\x65\x29\xdd\xe6\xd0\x63\xd0\xda\xa3\xa0\xae\x2d\xa5\xe3\xf1\x26\x8b\xaa\xef\x5e\x24\x25\x4e\x4a\x0b\x6d\x0f\x81\xfc\x79\xf3\xe6\xf7\x26\xcf\x21\xdc\xc0\x80\xc6\x3a\x04\x35\xdb\x9d\xd3\xbc\x10\x2a\xb8\x89\x51\x7e\x44\x0e\x01\xda\xcd\x62\x8c\x3d\x42\x8c\x55\x08\xd0\xf3\x71\xaf\x49\x4f\xd0\xbe\xa3\xdd\x0c\x31\x36\x52\xf4\xec\x9f\xd0\xc1\xcc\x64\xdd\xae\x81\xd1\x4e\x96\xc1\x3a\xae\xa1\x92\x82\xfc\x61\x86\x10\x60\x1e\x6d\x8f\xde\x40\xfb\xc5\x1f\xd2\x18\x94\x29\xbf\xf0\x3a\x88\x44\xe9\xe5\xa9\x96\x89\x0a\xdd\x90\x31\xe4\x35\xe2\xa3\x1f\x5e\x14\xc4\x18\xa4\xb0\xe6\xe4\x01\x5d\x07\x4a\x41\x90\xe2\x8c\xd2\x81\xba\x55\x52\x44\x29\x45\xef\xdd\xcc\xb0\xdd\xce\x3c\x71\x17\x02\xec\xc9\x3a\x36\xa0\x5e\x7f\x57\xd0\x6e\x1e\xee\x21\x46\x29\xfc\xe3\xb7\x76\xf4\xbb\x0d\x4f\x5c\x15\x69\x93\x98\x8d\xc5\x71\x78\xd6\xe3\x82\xa7\xb8\x3f\xa0\xf7\xd3\xa4\x21\xc6\xb2\xe8\x14\xb6\x96\x52\x6c\xb7\x29\x69\x09\x71\xd7\x41\x72\x1c\xc8\x3e\x23\xb5\x0f\x0b\xd2\xcb\xbf\xd8\xfe\xe6\x9a\x22\x26\xbf\x57\x1d\x38\x3b\xe6\x80\x84\xbc\x90\x4b\x1f\x1b\x50\xaa\x81\xaf\xa4\xf7\x1f\x88\x2a\x24\xaa\x53\x60\x31\xa0\x41\x82\x42\xd3\xbe\x1f\xfd\x8c\x55\xc2\x0b\x01\xac\xb3\xec\xf0\x00\xed\xbd\x9e\xf9\xf3\x53\x0e\x6e\xfc\xaa\xfd\x84\x47\xae\xea\xbc\xe4\x5a\x5c\xfe\x2e\x29\x44\x02\xe9\xce\xe2\x4d\xaf\x5d\xaa\x83\x1e\x06\xf2\x06\x2a\x33\x6a\x66\x74\x59\x5e\x5f\x9d\xe9\x22\xb9\x6c\xad\xa5\xf8\x43\xb0\xbf\x24\x4b\xd1\x4a\x97\x3a\xd0\xfb\x3d\xba\xa1\x2a\xf7\x4e\x1b\x68\x77\xe6\x2c\x37\x38\xb9\xdf\xad\xb8\xc9\xa7\x7e\xf3\xbf\xb7\xcc\x46\xa5\xcd\x6f\xe1\x36\x8f\xa4\x2f\xd0\xe5\xd5\x75\xaa\x5d\xf9\x35\xf3\x5f\xfa\xdc\x81\x99\xb8\xdd\xe4\xa6\x55\x67\xbe\x5f\xe2\x47\x29\x22\xe0\x38\xe3\x55\x69\xcb\x64\x79\x5f\xb6\x9f\xf8\x4a\xce\x55\xd4\x24\x60\x19\x65\x08\xeb\x13\xf2\x33\x00\x00\xff\xff\xdb\x75\x07\x28\xc4\x03\x00\x00")

func golangGetPagedTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetPagedTmpl,
		"golang.get-paged.tmpl",
	)
}

func golangGetPagedTmpl() (*asset, error) {
	bytes, err := golangGetPagedTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-paged.tmpl", size: 964, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x90\x3f\x6f\xe3\x30\x0c\xc5\x67\xe9\x53\xf0\x0c\x1c\x60\x03\x89\xbe\x81\x87\x1b\x82\x5b\x0e\x07\x24\x1e\x3a\x06\x4a\x44\x19\x2a\x1c\x29\xa1\xe5\xfc\x29\xc1\xef\x5e\xc8\x6d\x93\x76\x69\x07\x2e\x04\xdf\xef\x3d\x3e\xe6\x25\x38\xf4\x21\x22\x54\x63\xe8\xa3\xcd\x13\x61\x05\x4b\x11\xfd\x17\x33\x33\x98\x6e\xf2\x3e\x5c\x41\xa4\x66\x86\x7d\xbe\x1e\x2d\xd9\x03\x98\x3f\xd4\x8f\x20\xd2\x40\xad\x15\x33\xbc\x6f\x37\xe9\x02\x22\x0b\x40\xa2\x32\x89\x1a\x5d\x1c\x30\xba\x19\xa9\x3f\xdb\xed\x92\xbb\x55\x20\xc2\x5a\xed\x53\x1c\x33\x6c\xb7\x63\x3e\xe4\xb6\xd0\x28\xc4\xec\xa1\xfa\x7d\xaa\xc0\x74\xeb\x7f\x20\xa2\x55\xda\x3d\x9b\x21\xf5\x5d\x3e\xe4\xfa\xed\x74\x01\xcc\xe0\x03\x0e\xee\x6c\x87\x09\x1f\xa1\xf4\x9c\x29\xc4\x90\x3f\x22\x69\x55\x22\xb5\x50\x20\x8e\xc2\x19\xc9\xac\x27\xa4\xdb\x26\x5d\x7e\x82\x99\x6e\x6f\x63\x79\xde\x3a\x47\xc9\x43\xed\x07\x9b\x33\xc6\x19\xdd\xcc\x76\x2a\xf8\xf9\xe5\xb6\x85\xf1\x34\x98\x15\xd1\xff\xb4\x49\x97\x11\x58\x2b\x45\x98\x27\x8a\x85\xfe\x82\x94\x1e\x1d\xc5\x30\x68\x25\x77\xed\xaf\xb6\x6c\xbe\x55\x3c\x91\x3d\xae\x88\x6a\x24\x6a\x66\xe9\xe3\xd0\x52\xff\x95\x2c\x9a\xf9\xde\xfb\x6b\x00\x00\x00\xff\xff\xf0\xdb\xef\xb8\xe6\x01\x00\x00")

func golangGetTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetTmpl,
		"golang.get.tmpl",
	)
}

func golangGetTmpl() (*asset, error) {
	bytes, err := golangGetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get.tmpl", size: 486, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x57\x5d\x8f\x9b\x38\x14\x7d\xc6\xbf\xe2\x2e\x6a\x2b\x18\xa5\x44\xfb\x9a\x55\x1e\x9a\x49\x76\x35\x52\x77\xd2\xce\x50\xad\x56\xaa\x34\x75\xc0\x30\x74\xc0\x26\xb6\x49\xa8\x58\xfe\xfb\xca\xc6\x80\xc9\x24\xb3\xd3\x7d\x8a\xb1\xef\x3d\xf7\xdc\x8f\x63\xc8\x7c\x0e\x1f\xbe\x84\xdb\x3f\x36\xb7\x9b\xbb\x0f\xe1\x66\x0d\xab\xbf\x21\x65\xe5\x53\x1a\x64\x74\x2e\x4a\x1c\x91\x82\xd1\x27\xf2\x23\x65\xf3\x78\x57\x07\x87\x5f\xd1\x7c\x0e\xeb\x2d\xdc\x6e\x43\xd8\xac\x6f\xc2\x00\xa1\x12\x47\x4f\x38\x25\xd0\x34\x10\x7c\x32\xeb\xb6\x45\x28\x2b\x4a\xc6\x25\x78\xc8\x71\x23\x46\x25\xa9\xa5\x8b\x1c\x37\xc6\x12\xef\xb0\x20\x73\xb1\xcf\xd5\x73\x52\xe8\x6d\x99\x15\x44\xfd\x0a\xc9\x23\x46\x0f\x2e\x42\x4d\xf3\x1e\x38\xa6\x29\x81\x60\x53\x4b\x8e\x6f\x34\x9c\x50\xd0\x8e\x0a\xf5\x41\x2d\xc1\x9d\x46\x75\xb5\x1b\xa1\xb1\x32\xf3\x91\x22\xfb\x89\x93\x03\xa1\x12\x22\x46\xe3\x4c\x66\x8c\xe2\x1c\x32\x83\x95\x70\x56\x40\x84\x2b\x91\xd1\x14\x76\x55\x96\xc7\x90\xe0\x2c\xaf\x38\x11\xe8\x80\x39\x3c\xc0\x12\x0c\xa3\xe0\x46\x32\x8c\xf4\xee\x5f\x1c\x97\x1b\xce\x61\x09\x49\x45\x23\x8f\x70\x0e\x84\x73\xc6\xfd\xee\x07\x1a\x4e\x64\xc5\xa9\x7a\x6a\xb5\xc3\x47\x96\xa6\x84\x77\xd6\x09\xe3\x05\x96\x0a\x35\xa3\xe9\x0c\x30\x4f\x05\x04\x41\x90\x51\x49\x78\x82\x23\xd2\xb4\xbe\xf6\xb9\x65\x47\x58\x82\x2a\x4b\x70\xcb\x8e\x08\xc9\x1f\x25\x81\x98\x67\x07\xc2\x61\x30\x86\x06\x39\x9b\x9a\x44\xde\xbe\x22\xfc\xc7\x8b\xa0\xe0\x89\x7d\x1e\xdc\x11\x51\xe5\x72\x66\xf8\x22\xe7\xb3\xf2\x7b\x95\xf7\x95\x76\x67\x47\x71\xe2\x7c\xc7\x8e\xaf\xf1\xef\xdd\x51\x6b\x52\x59\xaf\x94\x43\x15\x49\x95\x83\x3e\x5d\xaf\x90\x13\xef\xfe\x24\xf2\x91\xc5\x42\xd9\xa9\x7a\xc1\xb6\x24\xd4\xeb\xf2\x9e\x81\x60\x15\x8f\x88\x89\xe4\x83\x17\xef\xe0\x6a\xbd\xd2\x8c\xfa\x16\x34\xc8\x11\xfb\xfc\x21\xde\x75\xbb\x8b\x25\x28\xec\x33\x28\x3e\x72\xb2\x44\xdb\xfc\xb2\x04\x9a\xe5\xca\xd3\x31\x9d\xa3\x59\x3e\xeb\xdb\xac\xfa\xeb\x23\xa7\x45\x4e\x4c\x92\xbe\x8b\x5d\x08\x30\xbc\x75\xd4\x33\x68\x86\x49\x70\x9d\x33\x41\x3c\xdf\x71\x90\xa3\x70\x5a\xe3\xee\xa3\x81\x42\x47\x53\xd9\x7e\xca\x68\xea\xf9\xbf\xfd\x14\x31\xe4\xa8\x89\xc9\x8a\x32\x87\xb1\x80\x8e\x38\x66\x32\x7a\xec\x67\xa6\xb1\x05\xb5\xce\x70\x4e\x22\x23\xa6\x08\x8b\x4e\xbc\xb7\xb8\x20\xf0\x0f\x94\x3c\xa3\x32\x01\xf7\xed\xde\x85\xb6\x5d\xa8\xcc\x14\xf2\x12\x28\x39\x0e\x66\xed\x98\xc4\x44\xa8\x24\xda\x52\x55\x6c\x8d\xac\xb2\x7b\x98\xc1\x49\x86\x7a\x60\x9b\x66\x12\x27\x80\xb6\x7d\x9e\xf4\x0b\x59\x3b\xad\xad\x74\x6b\xa9\xba\x84\xab\x5c\x2e\x2e\x14\x2d\x29\x64\xb0\x51\xa3\x92\x78\x6e\x45\x45\x55\xaa\x9b\x80\xc4\x7d\x99\xde\x0a\x77\x66\xd6\xbe\x2e\x2e\xea\x51\xde\xad\x57\x8a\xd3\x7a\xb5\x30\x89\xcc\x90\x33\xce\xeb\x42\x97\x7f\x86\x9c\x76\xa6\xe2\x0d\xe3\xeb\xb1\xdd\x77\x35\xa3\x3e\x98\x21\x00\x6f\x3a\xac\x06\xbd\xa7\xc7\x76\xdf\x83\xf5\xaa\x9f\x18\xff\x0c\x8e\x1e\x65\x25\xc8\xb0\x9e\x59\x38\xb2\x1e\x06\xde\x60\xac\x48\x9a\x51\xef\xa7\xc7\x7c\x4c\x38\xac\x95\xad\xac\x17\x20\xeb\x99\x5e\x0d\xc9\xaa\x10\x47\x8e\xcb\xb0\xf6\x64\xed\x4f\xd2\xd6\xea\x0e\x6b\x4b\xdd\xb2\xee\x84\x12\xd6\x68\xc4\x18\x33\x53\xc7\x61\xed\xc3\x35\x2b\x8a\x4c\xfe\x67\x85\x64\x1d\xc8\x3a\xe8\x8d\xfd\xe7\x38\x77\x2c\xcf\x77\x38\x7a\x7a\x25\xd2\x68\xae\xb1\x2e\xa8\x04\x35\x0d\xbc\x89\x77\x3a\xb9\xc5\xf2\xb9\x56\xc4\x7a\xe5\x76\x93\x08\x6f\x64\x7d\xd9\x2c\xac\x07\x33\x35\x30\x97\x0d\x6f\x8a\x32\xd7\xa6\x5d\x41\x27\x0e\x6d\x6b\x55\xd7\x0c\x6e\xf7\x73\x32\x2f\x27\x5e\x3e\xe4\x2c\xbd\x97\x85\xf4\x84\x2c\xa6\xef\xa0\x20\x08\x60\x72\x61\x37\xdd\x5b\xd6\x88\xfd\xa3\xe5\x37\x38\xf8\xc8\x66\x67\x8a\x33\xe5\x36\x5c\x91\xc8\x39\x25\x33\x50\x3d\xb9\x56\xec\x6b\xf5\x6a\x0a\x3c\xb6\xf0\xdd\xe4\xa0\xd1\x4a\x5c\x40\xa7\xc9\x93\x38\x8b\xce\xd8\xda\xd1\x57\x4b\x57\xaf\xde\xa7\x55\x23\x7c\xa6\x78\x43\x08\x1f\xee\xa3\x47\x52\x60\xcf\x37\x65\xb3\xc8\x7c\x53\xf4\xbb\xe3\xfb\xcf\x1f\xa1\x6d\xbf\xbd\x8c\x34\x08\xa7\x97\x85\x0f\x83\x2c\x4e\x73\x34\xb3\xd4\x91\x1e\xa5\xf8\x73\x39\x76\x3e\x43\x8e\x43\xcb\x06\xf0\x0b\x62\xbd\xd8\xb2\x0b\x93\xf1\xf2\x27\x40\x83\x9c\xf9\x1c\xc2\xed\x7a\xbb\x00\x4e\x68\x4c\x38\x94\x39\x8e\xc8\x23\xcb\x63\xc2\x85\xbe\xa7\xcc\x97\x92\x75\x55\x75\x3b\x9e\x2b\xf6\xf9\xe2\x2b\x7d\x2b\xbe\x52\x05\xae\x96\x7b\x77\x06\xe3\x3c\xfa\x26\x39\xeb\x55\xa0\x54\x66\xa4\x7c\xaf\xf3\x13\xbd\xf4\x4c\xba\x83\xf0\x26\x32\x33\x87\x76\x55\xac\x4b\xe1\xf7\x8c\xe4\xf1\xf8\x15\x6a\xdc\x75\x45\x42\x53\x25\x8b\x82\x21\x94\x25\x10\x6c\xcb\xee\xe3\xf3\x86\x0a\xc2\xe5\x08\x33\x04\x0e\xae\x39\xc1\x92\x74\x54\x7b\xdc\x73\x14\x2e\x21\x4d\x08\x5d\xe9\xa1\xb4\xb1\xa6\xbc\xec\x3a\x0d\x0c\xbe\x94\xf1\xab\x18\x68\x3b\xbc\xcb\xc9\xff\x0c\x7e\xbe\xa0\x6a\xf3\x4d\x62\xf5\x66\x8a\x30\x76\x28\x79\xde\x22\x38\xd8\x2d\x80\x76\x98\xd3\xb1\x9d\x0f\xf6\x2d\x33\x31\xf7\x4f\x51\x1b\x30\x1a\x9c\xee\x37\x87\x05\x1c\x06\x70\xaf\x3a\x39\xf6\x21\x62\x79\x55\x50\x15\xc3\xba\x27\x7a\x2c\xfd\x57\xe5\x5a\x5b\xa8\x7f\x2a\x16\xcc\xd5\x29\xce\x01\xe7\x95\x82\xb0\xf4\x03\x0d\x68\x85\x54\xb0\x3c\xf7\x1e\xd7\xdf\x81\xe6\xb1\x0a\x0e\xe8\x32\xc7\x3b\x92\x78\xfe\x69\xc8\x91\xe5\xbb\x0a\x2c\x15\xbd\x9f\x76\xee\xdf\x00\x00\x00\xff\xff\xb4\xbf\x97\xed\x2e\x0e\x00\x00")

func golangHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangHeaderTmpl,
		"golang.header.tmpl",
	)
}

func golangHeaderTmpl() (*asset, error) {
	bytes, err := golangHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.header.tmpl", size: 3630, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangMiscTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x41\x4b\xc3\x40\x10\x85\xcf\xcd\xaf\x78\x84\x08\x2a\xa6\x3f\xa0\xe0\xa5\x07\x41\x04\x0f\x5a\x3c\x77\xcd\x4e\xc2\x48\xba\x8d\x9b\x8d\x50\xc6\xfd\xef\x32\xbb\xb2\x56\x7a\x5b\xde\x7c\xef\xbd\x99\x15\x69\x61\xa9\x67\x47\xa8\x2d\x75\x63\x8d\x18\xab\x7e\x71\x1d\xae\x8f\xef\x1f\xb8\x15\xc1\xfa\x85\x3a\xe2\x2f\xf2\x5b\x33\x13\x62\x7c\x3c\x4c\xe3\x0d\x74\xf0\xca\x83\x33\x61\xf1\xaa\x26\x61\x7b\xb4\x27\x0d\x10\x01\x39\x8b\x36\xc6\xaa\x3a\x6f\x98\x83\x5f\xba\x90\x3a\xc2\x69\xa2\xe4\x79\x36\x87\xe4\xcf\x33\x88\x9a\xbd\x71\x03\x61\xfd\xc0\x34\xda\x59\xe9\x95\x08\xb8\x2f\xf0\xb9\xef\xb7\x2b\x8b\x3b\x4d\x4d\x4f\xa5\x77\x66\x98\xd3\x12\xab\x7d\x09\x6d\xf8\x0e\x4d\xc0\xe6\xfe\x6f\x9c\xe9\x86\x2f\xd2\x9e\x48\xaf\xd9\xe8\xf3\xcd\x8c\x0b\xe1\x1b\x93\x67\x17\x7a\xd4\x57\x9f\x75\x86\x32\xbd\xd7\x0d\xdb\x72\x74\xd1\xab\x7f\x7f\xf1\x13\x00\x00\xff\xff\x1b\x36\x80\x57\x6d\x01\x00\x00")

func golangMiscTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangMiscTmpl,
		"golang.misc.tmpl",
	)
}

func golangMiscTmpl() (*asset, error) {
	bytes, err := golangMiscTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.misc.tmpl", size: 365, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangUpdateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x55\xc1\x6e\xe3\x36\x10\x3d\x4b\x5f\x31\x15\x50\x40\x42\x1c\xa2\xe8\x71\x01\x1d\x82\x6d\x5a\x04\x48\x8d\x24\xee\xb6\x87\xc5\xc2\x60\xac\xa1\xc0\x40\x26\x6d\x72\x14\x3b\x10\xf4\xef\xc5\x90\xb2\x23\xc5\x4e\xda\xa2\x97\x3d\xf8\x40\x73\xe6\xcd\x9b\xf7\x86\xa3\xae\xbb\x84\x0a\x95\x36\x08\x99\xd7\xb5\x91\xd4\x3a\xcc\xe0\xb2\xef\xd3\x2f\x9b\x4a\x12\x76\x1d\x88\x45\xab\x94\xde\x43\xdf\xe7\x5d\x07\x2b\xda\x6f\xa4\x93\x6b\x10\x57\xae\xf6\xd0\xf7\xb3\x34\x69\x43\x28\x84\x58\x72\xed\x8a\x44\x4c\x8e\x87\xb9\x5c\x23\xf4\x7d\x01\x79\x9a\x74\x1d\x0c\xd9\x0f\x48\xad\x33\x9c\x0f\xe8\x1c\xff\xac\x2b\x52\xe6\x83\xa6\x0a\x04\xd2\x31\xb9\x47\x5b\xbd\x04\x5e\x5d\x9a\xac\xac\xf1\x04\xcb\xa5\xa7\x35\x2d\x37\x0e\x95\xde\x97\x0c\xec\xb4\x21\x05\xd9\x8f\xdb\x0c\xc4\xe2\xfe\xf6\x2e\xdc\x40\xdf\xbf\xc9\xf0\xa1\x9d\x73\x19\xc7\x46\xd3\x34\x79\x96\x0e\xb6\x1e\xbe\x7e\x7b\x7c\x21\x8c\xc7\x67\xd9\xb4\xc8\x7f\x69\x43\xe8\x94\x5c\x61\xc7\x91\x5d\x07\x5a\x81\xb8\xb3\x5e\x93\xb6\x46\x36\x57\xae\x6e\xd7\x68\xc8\x87\xda\x4f\xf0\xa9\x64\x69\x1a\x34\x47\xcd\x18\x2f\x60\x95\x20\x37\x1b\x34\x55\x1e\xcf\x33\x0e\x54\x1a\x9b\x2a\x9c\x8f\xf1\x45\xa8\x72\x14\x26\x91\x55\x75\x28\xc2\xe8\xaa\x35\xab\xdc\xb0\xcc\x9e\x9c\x36\x75\x01\x5d\x9a\x24\xdb\x11\xfc\xd6\xcf\x80\x03\x84\x10\x45\x9a\x24\xac\xec\x47\x9c\x93\xa7\x8b\x8b\x33\x08\x19\x84\x4e\xc4\x21\xf8\xa8\x70\x36\xe0\xbe\x4d\xf0\xe4\x56\xd6\x3c\x8b\x1b\xb2\x32\x7f\x2a\xde\x89\xca\x66\x90\x8d\x88\x61\xe3\x31\xb2\xf8\xd7\xf5\xdf\x00\x98\x2a\xe4\xf7\x69\xb2\x64\x80\x57\xb1\xa2\x5b\x4e\x9a\x1a\xa7\xa3\x2a\x1f\x1b\xfc\x95\x75\x8f\xfd\x6b\x05\x71\xa6\x05\x97\x1b\x06\x18\x7e\x28\xc1\xe8\x26\x68\x3b\xc2\xcc\x33\x8e\xf9\x6c\x9b\x76\xcd\xe3\x9c\x31\x8b\xf7\xdc\x3d\x05\x15\xe1\x2a\x2f\x8a\x40\x77\x30\xb9\x8f\x63\x15\x3d\x9a\x23\x56\x7e\x6e\x77\x81\xd8\x72\x69\xec\x8e\x2d\x9f\xdb\x5d\x5e\x88\x2f\x7f\x7c\xce\xa7\xb3\x31\xea\xef\xaa\x25\x3b\xea\xe9\x84\xf2\xc0\x80\x09\x7f\x30\x8d\xe2\xc6\x68\xfa\x53\x36\x93\x29\xec\x5f\xe7\xde\x58\x3a\x29\xa5\x15\x4f\x7b\xbe\xf5\x05\x94\x25\xfc\x14\x14\x73\xf1\xbd\x1b\xdd\xcc\xe0\x2f\x27\x37\xd7\xce\xe5\x6a\x4d\xe2\x9a\x1f\xbe\xca\x33\x83\x58\x01\xd9\x41\x22\x90\x04\x0d\x4a\x4f\x60\x0d\xc6\x17\x91\x9d\xd5\xe8\x48\xe1\xdc\x28\x5f\xfe\x9f\x97\x16\x0a\xc4\x95\xc1\x82\x4f\xd6\x0d\x5c\x0c\x2f\x2d\xdf\xfa\xaf\x9f\x86\x5e\x2f\x7f\xfe\x56\xc0\xc5\x74\xcb\xa4\x89\x7d\x7c\x12\x8d\xad\x17\xb4\xa6\x3c\x5e\xcd\x86\x45\x12\x26\x36\xf6\x60\x34\x8d\x36\xe2\x71\xa5\x2c\xda\xcd\xc6\x3a\xf2\xf1\x46\x9b\x3a\x5c\xf2\xba\x2c\x81\x71\x2b\xa7\x9f\xd1\x89\xfb\x16\xdd\xcb\x83\xdd\x9d\xc1\x17\x8b\x95\x34\xbc\xb6\x65\x55\x39\xab\x20\x57\x8d\x24\xe2\x4d\x14\x31\x8b\xd8\xb2\x56\x61\x09\x97\x25\xf8\x6d\xc3\x96\xcc\xed\x83\xdd\xf9\x13\xe3\x8c\x6e\x82\x07\x43\xfc\xe8\x41\x9c\xb3\x17\x9d\x7b\xb5\x8c\x1f\x75\xf0\x63\x19\x17\xfe\xa4\x83\xeb\x3d\xae\xce\xa9\xf3\x5f\x0b\x1d\xf6\x7c\x8d\x14\xc0\x4e\x96\xfc\x6f\x48\x8b\xfb\xdb\x20\xe3\xd8\x98\x2c\xbf\xf9\xfd\xee\xf6\xe6\xfa\x97\x02\xb2\x60\xe1\x01\xe0\xfd\x31\xf9\xd8\x87\x7f\xcc\xff\x7e\x8c\x39\x2c\x8e\x21\x90\x29\xb9\x7a\xf2\x7d\x66\xf0\x7e\xf2\x69\xfe\x3b\x00\x00\xff\xff\xd1\xe8\xeb\x3e\x37\x08\x00\x00")

func golangUpdateTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangUpdateTmpl,
		"golang.update.tmpl",
	)
}

func golangUpdateTmpl() (*asset, error) {
	bytes, err := golangUpdateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.update.tmpl", size: 2103, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"golang.create.tmpl": golangCreateTmpl,
	"golang.delete-all.tmpl": golangDeleteAllTmpl,
	"golang.delete-world.tmpl": golangDeleteWorldTmpl,
	"golang.delete.tmpl": golangDeleteTmpl,
	"golang.footer.tmpl": golangFooterTmpl,
	"golang.get-all.tmpl": golangGetAllTmpl,
	"golang.get-count.tmpl": golangGetCountTmpl,
	"golang.get-has.tmpl": golangGetHasTmpl,
	"golang.get-last.tmpl": golangGetLastTmpl,
	"golang.get-limitoffset.tmpl": golangGetLimitoffsetTmpl,
	"golang.get-paged.tmpl": golangGetPagedTmpl,
	"golang.get.tmpl": golangGetTmpl,
	"golang.header.tmpl": golangHeaderTmpl,
	"golang.misc.tmpl": golangMiscTmpl,
	"golang.update.tmpl": golangUpdateTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"golang.create.tmpl": &bintree{golangCreateTmpl, map[string]*bintree{}},
	"golang.delete-all.tmpl": &bintree{golangDeleteAllTmpl, map[string]*bintree{}},
	"golang.delete-world.tmpl": &bintree{golangDeleteWorldTmpl, map[string]*bintree{}},
	"golang.delete.tmpl": &bintree{golangDeleteTmpl, map[string]*bintree{}},
	"golang.footer.tmpl": &bintree{golangFooterTmpl, map[string]*bintree{}},
	"golang.get-all.tmpl": &bintree{golangGetAllTmpl, map[string]*bintree{}},
	"golang.get-count.tmpl": &bintree{golangGetCountTmpl, map[string]*bintree{}},
	"golang.get-has.tmpl": &bintree{golangGetHasTmpl, map[string]*bintree{}},
	"golang.get-last.tmpl": &bintree{golangGetLastTmpl, map[string]*bintree{}},
	"golang.get-limitoffset.tmpl": &bintree{golangGetLimitoffsetTmpl, map[string]*bintree{}},
	"golang.get-paged.tmpl": &bintree{golangGetPagedTmpl, map[string]*bintree{}},
	"golang.get.tmpl": &bintree{golangGetTmpl, map[string]*bintree{}},
	"golang.header.tmpl": &bintree{golangHeaderTmpl, map[string]*bintree{}},
	"golang.misc.tmpl": &bintree{golangMiscTmpl, map[string]*bintree{}},
	"golang.update.tmpl": &bintree{golangUpdateTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

