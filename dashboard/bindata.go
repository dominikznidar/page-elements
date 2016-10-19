// Code generated by go-bindata.
// sources:
// templates/list.html
// DO NOT EDIT!

package main

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

var _templatesListHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x91\xcb\x6a\xc3\x30\x10\x45\xd7\xd6\x57\x0c\x22\xcb\x12\x91\x6d\x19\x69\xd7\x6d\x5b\x68\x7e\x40\xb1\xc7\x0f\xb0\x64\xd7\x1a\x1b\x8a\xd0\xbf\x17\xf9\x81\x09\xd9\xce\x9c\xb1\xcf\xbd\xc2\xf6\x66\x3e\xbc\x7d\xf4\x54\xc1\x68\x1b\x02\xea\xc9\x91\xe7\x80\xaa\xbd\x19\x21\xb0\x1e\x26\x07\xb6\xe4\x6e\xf0\x5a\xaa\x60\x17\x92\xe0\x88\xdb\xa1\xd2\xf2\xfb\xeb\xe7\x2e\x8d\x40\xce\xf7\x46\x14\x31\xc2\x64\x7d\x43\x70\x09\x34\x2d\x5d\x49\x9f\xd6\xd1\x1b\x5c\x02\x5b\x9e\x03\xbc\x6b\xb8\x42\x4a\xa2\x40\x9e\x8c\x28\x0a\xe4\xca\xc4\xf8\x04\x43\x4a\xa8\xb8\x3a\xb6\x18\xa8\xa7\x92\xc1\x5b\x47\x5a\xbe\xb2\x32\x83\x05\x0e\x63\xd6\x83\xc5\xf6\x33\x69\x29\xcd\x50\xd7\xa8\xb6\xe1\x0a\x9c\x5e\x6c\x9b\xac\xb1\x1b\x5d\xef\xb6\x09\xab\xd1\xf9\x91\x18\xa1\xab\x81\x7e\x37\xf6\x00\x8f\x8a\x52\xda\x8c\xa8\x8a\x11\xc8\xe7\xc1\x1a\x21\xb3\x59\xfd\xf9\xa7\x1b\x90\xb3\xa8\xed\xcc\xec\xe1\x50\xad\x0d\x9c\x08\xaa\xbd\x43\x81\x9d\x1f\x67\x06\xfe\x1b\x49\xcb\x30\x3f\x5c\xc7\x12\x54\x5e\xa8\xfc\x14\x46\xfc\x07\x00\x00\xff\xff\xc8\xb2\xd9\x71\xb2\x01\x00\x00")

func templatesListHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesListHtml,
		"templates/list.html",
	)
}

func templatesListHtml() (*asset, error) {
	bytes, err := templatesListHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/list.html", size: 434, mode: os.FileMode(420), modTime: time.Unix(1476865052, 0)}
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
	"templates/list.html": templatesListHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"list.html": &bintree{templatesListHtml, map[string]*bintree{}},
	}},
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
