// Code generated by go-bindata.
// sources:
// templates/template.html
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

var _templatesTemplateHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x56\x4b\x73\xdb\x36\x10\xbe\xfb\x57\x6c\x99\x4b\xe2\x9a\xa2\xa5\xc8\x8f\x32\x92\xa6\x99\xd4\xd3\xe6\xe0\x26\x63\xbb\xa7\x4e\x0f\x20\xb1\x22\x51\x93\x00\x06\x58\xca\x72\x3c\xfe\xef\x1d\x10\x7a\x50\x7c\xc8\x87\xd2\x07\xdb\x58\xec\xee\x87\x6f\x9f\xb3\x9f\xb8\x4a\xe9\x59\x23\xe4\x54\x16\x8b\x93\x99\xfb\x05\x05\x93\xd9\x3c\x40\x19\xb8\x03\x64\x7c\x71\x02\x00\x30\x2b\x91\x18\xa4\x39\x33\x16\x69\x1e\x54\xb4\x0c\xaf\x83\xc5\x89\x97\x91\xa0\x02\x17\x2f\x2f\x30\xfa\x1d\x09\x02\xcd\x32\x7c\x70\x47\x01\xbc\xbe\xce\x22\x2f\x6d\x58\x91\xac\xc4\x79\xc0\xd1\xa6\x46\x68\x12\x4a\x06\x90\x2a\x49\x28\x69\x1e\x84\xe1\xed\xcd\xc3\xe7\xf0\xb7\x9b\xfb\x2f\x77\x5f\xbf\x3f\x7c\xfd\xf6\x67\x18\x06\x5d\x65\x56\x51\xae\x4c\x57\xef\xf3\x5f\x0f\x7f\x7c\xbb\xab\x55\xbc\x4e\x21\xe4\x23\x18\x2c\xe6\x81\xa5\xe7\x02\x6d\x8e\x48\x01\xe4\x06\x97\xf3\x20\x27\xd2\x71\x14\x65\x48\x89\x52\x64\xc9\x30\x3d\x4a\x55\x19\x71\x61\x29\x4a\xad\x8d\xf6\xc7\xa5\x90\xa3\xd4\xda\x2d\x92\xda\x96\xff\xdb\x7d\xd1\x29\xdc\x6b\x96\x22\xa8\x8a\xb6\x90\x80\x41\x22\x08\x4e\xa3\xdd\xad\x44\xf1\x67\x78\xd9\xfd\xeb\x3e\xcd\x38\x17\x32\x0b\x49\xe9\x18\x26\xe7\x7a\xfd\xa9\x57\x9c\x28\x22\x55\xb6\x6f\xbc\x9e\x34\x01\xdc\xac\xd0\x3c\x53\x2e\x64\x06\x49\x45\x40\x39\xc2\xbf\x55\x99\x28\x32\x4a\x42\x86\x64\xc1\x0a\x8e\x60\x35\x4b\xdd\x9d\xa5\x32\x50\xaa\x44\x14\x08\x4b\x61\x2c\xc1\x4a\xe0\x93\x6d\xc2\x1d\xb9\xe0\xa3\x39\xdb\x1f\x94\xcc\x3c\x22\x09\x99\x35\xce\x96\x4a\x11\x9a\x81\x67\x19\x91\xe5\x14\xc3\xf8\x62\xe8\x61\x05\x2e\x3b\xf2\xc3\x67\x7d\xa9\x2c\xa9\x12\x5c\x4a\x81\x07\xd4\x83\x71\xc0\x7d\x3f\x6d\x3e\x14\x86\xa3\xd9\x5d\x18\xeb\x35\x58\x55\x08\x0e\xef\xf0\xc2\xfd\x34\xd1\x34\xc0\xdc\xb2\x47\xac\x99\x2d\x99\x25\xe7\xba\x86\xe4\xe8\x74\x87\x96\x95\x0e\xa3\x7b\x33\x30\x5b\x1f\x49\xb6\x12\x19\x73\x29\xde\x87\x3a\xff\xd8\x02\x5e\x32\x93\x09\xe9\xb3\xe1\xfc\x53\x9f\x68\x8b\xb8\x25\x2d\x84\xc4\xd0\xbb\x8e\x61\x7a\x24\x4b\x9a\x74\x6e\x42\xd7\x04\x76\x34\x9a\x35\xac\xf1\x2f\x6d\x2e\x53\x55\x28\x13\xc3\xbb\xab\xab\xab\x5e\x92\xbd\xda\x31\x86\xbb\x00\xc5\x0f\xac\xeb\x88\x09\x79\x88\xf0\xd7\x12\xb9\x60\xf0\xbe\x14\x32\x7c\x12\x9c\xf2\x18\xae\x2e\xaf\xf5\xfa\x43\x0b\xf3\x68\xaf\x7d\x28\xf0\x5c\xae\x77\xca\x1f\x3b\xb9\xf1\xda\x13\xfb\xbd\xb9\x50\x32\x63\xd4\x13\x2c\x20\x6f\x9b\xf6\x21\x8a\xc1\x99\x6c\x06\xe8\xf0\x81\xb7\x4c\x48\xd8\x55\x12\x94\x68\xad\x0b\x06\x93\x1c\xac\xc8\x24\x54\xda\xd5\x2f\xb5\x32\x66\x5f\xca\x87\x3e\x09\xd7\x14\xb2\x42\x64\x32\x86\x14\x25\xa1\xf9\xff\x89\xde\x70\x36\x4a\xa8\xed\x71\x93\x0d\x31\x8c\xa7\x7a\x0d\x93\x69\x9b\xbe\xa5\x92\x14\x5a\xf1\x03\x63\x98\x8c\x87\x13\xf1\xbe\xd2\x5a\x19\x4f\xc1\x8e\x8c\x6d\xeb\x6c\xbe\x7c\x2f\xed\x67\x7b\xda\x61\xbb\x47\x55\xc3\xcf\x90\x4f\x8f\x54\xdb\xe4\x7a\x18\xea\x1d\x5a\xad\xa4\x15\x2b\x8c\xe1\xbb\x32\x64\x98\x20\x20\x96\x14\xae\xa9\xba\xb8\x55\xba\x27\x45\x6d\x6a\x10\x65\x2d\x7f\x33\x5b\x6b\x27\xa5\x5a\xf9\xe6\xb2\x61\x18\x9e\x10\x2c\x12\x20\x33\x85\x38\xac\x02\xe8\x6b\xcf\x30\xd0\xa2\x61\xb0\xb0\xa1\xdb\xaa\x5b\x7d\x05\x3a\xbd\xfa\x7c\xa8\x58\xa0\x3d\x07\x0f\xfa\x64\x3f\xf8\xde\xda\x3c\xe8\x73\xc7\xca\xb3\xcb\x9c\xd7\xda\x24\x3d\x28\xd9\x9a\x82\x6e\xe6\xad\x84\xad\x58\x01\xb8\x5c\x62\x4a\x1d\x58\x43\x65\x06\xdd\x52\x1a\x24\xa2\x91\x80\x99\x11\x1c\xd7\x30\x72\x0d\xa3\x37\xf7\x86\x06\x7b\xaf\x85\x61\x33\xbe\xc5\x76\xa8\x1a\x9e\x19\x5d\x0f\x7f\xa7\x05\xb3\xf6\x74\x1e\xa4\xaa\x08\x83\x7f\x8e\x8e\x80\xc1\x71\xbe\x6b\x33\x9d\x1b\x09\x4b\x1f\x33\xa3\x2a\xc9\xc3\xed\xbc\x40\xc4\xb7\xee\x98\x2c\x61\xef\xaf\x2f\xcf\x2e\xc7\x67\xe3\xc9\xf4\x6c\x34\xbe\xf8\xd0\xd7\xde\x0e\xfa\x1a\xe7\xfc\xad\x3b\x6d\xb3\x93\x0f\x9d\xd2\x9f\x45\x9b\x15\x6f\x16\xf9\x1d\xf8\x64\xe6\xd6\xb7\xcd\xfa\xc7\xc5\x0a\x6a\xbe\x1c\x5d\x9b\xc1\x10\xec\xd7\xc1\x19\x16\x58\xa2\xa4\xfd\x46\xbc\x39\xf0\xfb\xf0\x56\xba\x71\xc4\xc5\xca\xb9\xf1\xe6\x67\x91\x5f\xc5\xff\x0b\x00\x00\xff\xff\xe5\xcb\xc5\x7c\x9b\x0b\x00\x00")

func templatesTemplateHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTemplateHtml,
		"templates/template.html",
	)
}

func templatesTemplateHtml() (*asset, error) {
	bytes, err := templatesTemplateHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/template.html", size: 2971, mode: os.FileMode(420), modTime: time.Unix(1477690431, 0)}
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
	"templates/template.html": templatesTemplateHtml,
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
		"template.html": &bintree{templatesTemplateHtml, map[string]*bintree{}},
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

