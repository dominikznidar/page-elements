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

var _templatesTemplateHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x58\xcf\x72\xdb\x36\x13\x3f\xdb\x4f\xb1\x9f\xbe\x8b\x3d\x23\x4a\x4d\xdc\x64\x32\x19\x45\x33\x49\xa6\xed\x29\x4d\x27\x93\x4b\x8f\x2b\x72\x49\xa2\x06\x01\x06\xbb\x14\xa9\x9c\xfa\x10\x7d\xc2\x3e\x49\x07\x00\x49\xd1\xb2\x6c\x47\x8e\x3b\xbd\x91\xc0\xfe\xdf\xdf\xfe\x00\x72\x45\x9a\x2a\x32\xb2\x2e\x09\x33\x72\xab\xe5\xf0\x7e\x7e\xbe\xca\xd4\x16\x52\x8d\xcc\x6f\x66\x85\x53\x19\x75\xb3\xf5\xf9\xd9\x74\xb5\xc6\x82\x92\xa8\xe8\xb7\xce\x56\xe5\xb3\xf5\x3b\x6b\x85\xc5\x61\x0d\x5e\x07\xa8\xc3\xaa\xd6\xc4\xab\x65\xf9\x2c\x88\xd4\x83\xb2\x26\xcc\x66\xeb\x77\xc8\x2a\x8d\xa2\x1a\x77\xb6\x11\x06\xb1\x50\x90\xc0\xce\x36\x90\x63\xa5\xb4\x42\x07\xad\x92\x12\x36\x8d\xd2\x99\x32\x45\x78\x53\x06\xa4\x24\x38\x70\xc7\x3b\x16\xaa\x16\xab\x65\xed\x43\x5d\x66\x6a\xbb\x3e\x3f\x3f\x5b\x95\x57\xeb\xcf\xa5\x23\x02\xfa\xd2\xa0\x86\xd4\xea\xa6\x32\x3e\xa6\x2b\x2f\x56\xaf\x7f\x21\x01\xd9\x0b\x24\xad\xca\xa4\x1c\xc4\x60\xc5\xe2\xac\x29\xd6\x2c\xe8\xc4\xfb\x47\x81\x8c\xf8\x5a\x6c\xcd\x80\x26\x03\x4e\x51\xfb\x75\xb1\xa0\xd1\x15\x34\xee\xae\x96\xbd\xea\x02\x3e\x1a\xa8\xec\x46\x69\xbf\xb9\x55\x29\xf1\x1c\x04\x37\x9a\x24\x9a\xd8\x90\xb6\xed\x3c\xa4\x34\xb8\x6d\x95\xd6\x80\x8d\xd8\x0a\x45\xa5\xa8\xf5\x0e\x58\x30\xbd\x1e\xb2\x9b\x34\xc2\xd9\x36\x36\x60\xb2\x96\x5a\x9d\x54\x59\xf2\xe3\x6c\xbd\x18\x1e\xfb\x82\x3c\xb9\xdc\x91\x42\x37\xe6\xc1\x52\x3f\x61\x79\xc1\xe6\xb0\x45\xa7\x6c\xe3\xcb\x96\x49\xc9\x0b\xf8\x44\x15\x55\x1b\x72\xf3\x88\x8c\xc1\x1b\x97\xb6\xd1\x19\x60\x96\x41\x53\x7b\x93\xd2\x92\xde\x12\xe4\xd6\x01\x02\x2b\x53\x68\x82\xd2\x3a\xf5\xd5\x1a\x41\x0d\x1b\x6d\xd3\xeb\x05\x7c\xb0\x8e\x40\x4a\xf4\xa8\x43\x99\x87\xb8\x46\x93\x3e\xf0\xd8\x1b\x1f\xa6\xb1\x50\xa1\x08\xb9\xd0\xcd\xad\xa2\xb6\xb6\x4e\x4e\xeb\xda\xd5\xbe\xca\x57\xf7\x75\xe3\xe5\x5e\xee\xe5\x7d\x72\xc7\xec\xdd\xec\x5a\x6b\xef\xe8\xd5\x7e\xe3\x49\x06\xe1\xa4\x32\xbc\xda\x87\xfd\xea\xbb\x40\xf9\x73\xa3\x75\x44\xc6\x7c\xe8\x71\x4c\x6a\x4c\x76\xb0\x28\xd4\x49\xd2\xa2\x33\xca\x14\xb3\xf5\xaf\xb6\x47\x8f\xdf\x24\x06\x74\x04\x86\x52\x62\x46\xb7\x0b\x98\xc9\x1b\x3d\xf0\x45\x4f\x9d\x1c\x73\xf4\x8e\xdd\x91\xf2\x46\x2e\xf3\x65\x35\xc4\x42\xd9\xed\xb2\xff\xd6\x43\x27\xb3\x69\xe3\x0d\xa2\x28\x6b\xe6\x41\xdc\x17\x57\x31\x10\xf2\xee\xef\x3f\xff\xfa\xa3\x61\x81\xba\x11\x40\x70\xb6\xf5\x43\x30\xf5\xa1\x0c\xa0\x01\xea\x54\x54\x8b\x5b\x0b\xf8\x5c\x2a\x86\x42\x6d\x89\x03\xc5\x3e\x71\x7f\xe7\x31\x3d\x34\x56\x4a\x9f\x46\x6b\xe1\x22\x32\x41\x9c\xcb\xcb\x29\x79\x07\x1b\x6e\x88\xac\x47\x46\xbd\x7e\x2b\x37\xa9\x12\x58\x7d\x3d\xe4\xcb\xcc\xb6\x26\xd0\x25\xef\x89\xc4\xaf\x4b\x49\xca\x1d\x54\x36\x32\xe9\x23\xb8\xf3\x55\xd8\x38\x1b\x21\xe8\x5f\x8e\xa9\x7e\xfb\x68\x9e\x20\xb9\xc7\xfb\xe3\x81\x3f\xe2\xef\x83\xea\x28\x7b\x3d\x94\x35\x14\x30\x76\x6e\x8f\xba\xcf\x37\x4e\xd3\xab\xe9\x79\x0a\x25\x32\xe4\xb6\x71\x20\x8a\x1c\x07\xa0\xc5\x81\x78\x0d\x1d\xc3\x45\x5d\x5a\x43\x7c\x39\x07\xae\xe0\xa2\xef\xd2\xe5\x1c\xaa\x0c\x2e\x06\x80\x5c\x46\xca\xd4\x05\x5c\xf4\x5d\x1f\x77\x16\xf0\xbb\x6d\x20\x45\x03\x0d\xfb\xe9\x42\xa7\x77\x80\x66\x07\xa9\xad\x36\xca\x04\xf4\x7b\x97\x7d\xb3\xfb\x49\x14\x0b\xa9\x23\x14\x82\xca\x53\x73\xb6\x33\x58\xa9\x34\x38\xc9\x35\x75\x6a\xa3\x69\xb8\x51\xec\x91\xf5\x13\xa6\x65\xc8\x61\x92\x42\x00\x35\x31\x34\xf5\x1c\x2a\x42\x13\x46\x2c\x0f\xb3\x51\x6b\x34\x60\x0d\x30\x49\x18\x07\x8f\x59\xc6\x8a\x7a\x2c\x87\xf9\xef\xfa\xc9\xa8\xe6\x41\xc5\x1a\xbd\x03\x43\x94\xf9\x00\xb9\xa6\x54\xe5\x3b\xe8\xf8\x04\xdc\x75\x9c\x3c\x7b\x0e\x07\x04\x18\x17\x1f\xe6\xc2\x8e\x93\x97\x70\x80\x8d\xb0\x76\x17\x4a\xbe\x35\xa4\x6f\xb4\xfa\xdf\xa8\x9e\x9a\xcb\xc4\xd8\xbd\xde\x8f\xc8\x3d\x38\x5c\x03\x4d\xcd\x8f\x4f\xd9\x63\x70\xc0\x55\x5f\x08\x5d\x1c\x03\x44\xd8\x5e\x0c\xfb\x0f\x97\x53\x17\x47\xca\xe9\x17\xbf\x0b\x1a\x5c\x1d\xb1\xea\x17\x1f\x0e\xe8\xdf\x54\x3d\xd6\xaf\xf7\xe1\x58\x80\x54\x13\x3a\x65\x8a\x49\x73\x10\x4a\x47\xf9\x9b\x59\x29\x52\xbf\x5e\x2e\x0b\x92\xcd\xc0\x88\x8b\xd4\x56\xcb\x94\x79\xf9\x7f\x4f\x8c\x89\x23\xae\xad\x61\xb5\x25\xff\x48\xc2\xb3\xf5\x7b\x6f\x0f\x72\x6d\x51\x78\xb5\xc4\xb5\x3f\x3d\x23\x03\xa8\x14\x36\x8e\xf0\xba\xb6\xca\xc4\x0f\x9c\xda\xd1\x96\x8c\x00\xb6\xd7\x2d\xba\x0c\x5a\x87\x75\x3d\x7c\xda\xf8\xcb\xf3\x96\x0c\xa4\xd6\x08\x99\x53\x6e\x8e\xd3\xba\x5c\x4d\x0e\xaf\x1b\xb5\xb9\x0a\xa7\xcb\xc6\x85\xed\x4f\xe4\x0f\x57\x4f\x5d\x6e\xbc\xaa\x82\x75\x90\x96\x94\x5e\x83\x12\xb0\x8d\x78\x0a\x0c\x02\x81\xe7\xe3\x4d\xd9\x0c\x5f\x76\x8b\xfb\x4e\xa8\x83\x78\x8e\xc4\x32\xf6\xe7\x6c\xf5\xbf\x24\x81\xb7\x59\x38\xc2\x81\x3a\x71\x18\x3b\x94\xab\x2e\xb8\x0c\xcc\xea\xf7\x1c\x7d\x69\x94\xa3\x6c\x1f\x6f\x92\xdc\x72\x3e\x68\x6e\x15\xfb\xa3\x20\xe9\x78\xb6\x9e\xf8\x7a\x6c\x9c\xdf\xa3\x7a\x0c\x8a\x1f\xf3\x9c\x3d\x57\xd4\x0d\x97\x91\x31\x6a\x7f\x53\x8d\x90\xda\x03\xf3\x93\x7f\x07\x1b\x84\x39\x4a\xfb\x0b\xd1\x20\xcf\x77\x41\xed\x04\xe8\x70\x95\xbc\x80\x83\x0b\x49\x58\x7b\xf8\xfb\x62\x54\xe5\x2a\x89\x31\x26\xe3\x01\x36\x32\x78\xbf\xf1\xc3\x2d\xd3\x53\xa5\xd1\xd7\xf8\x34\xa8\x3d\x86\x9b\x46\xde\xac\xb2\x3e\x40\x5d\x4c\x73\xdb\x7b\x79\x31\x92\xe0\xbd\x59\x1e\x58\xdb\x87\xf7\x7c\xb4\x3e\x3c\xdc\x4e\xf7\x96\xb7\xa9\xfa\xe8\x7e\x7c\xba\x23\xf1\x01\x44\xe3\x7f\x9b\xdc\x5a\xb9\xf1\xdf\xe6\x9f\x00\x00\x00\xff\xff\xa2\xea\x94\xdf\xd2\x11\x00\x00")

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

	info := bindataFileInfo{name: "templates/template.html", size: 4562, mode: os.FileMode(420), modTime: time.Unix(1478375870, 0)}
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

