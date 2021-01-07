// Code generated by go-bindata.
// sources:
// data/en_US-web.dic
// data/en_US-web.aff
// DO NOT EDIT!

package data

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


func dataEn_usWebDicBytes() ([]byte, error) {
	return bindataRead(
		_dataEn_usWebDic,
		"data/en_US-web.dic",
	)
}

func dataEn_usWebDic() (*asset, error) {
	bytes, err := dataEn_usWebDicBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/en_US-web.dic", size: 891233, mode: os.FileMode(420), modTime: time.Unix(1567821227, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataEn_usWebAff = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x56\xdd\x52\xe3\x38\x13\xbd\xf7\x53\xf4\x57\x5c\xc0\x4c\x0d\x54\x1c\x98\x1f\xe6\x0e\x42\x12\xc2\x4f\x92\xca\x0f\x93\x40\x85\xaf\x8c\xdd\xb1\x35\x04\x3b\x63\x59\xc3\x9a\xe2\x62\x1f\x63\xf7\xf5\xf6\x49\xb6\x6c\x49\x2d\xd9\x98\xe5\x62\x7c\x74\xfa\xf4\xd1\x51\x5b\x71\xcd\xb4\x3b\x83\xf9\xac\xb7\xff\xcd\x99\x4d\x96\x80\x9c\x79\x71\x9a\x25\x1b\x3f\x10\xe1\xd3\x36\x7a\xc8\xd7\xbf\x1f\x9f\x5f\xba\xd3\xc1\xc9\x70\x32\x1b\x5d\x75\xce\xe6\xfd\xeb\xf1\xf9\xe9\xb2\x77\x73\xf9\xe3\x76\xd7\x19\x74\x46\xc3\x1b\x68\xab\xe7\x3f\x7f\xfe\x0d\xbb\x84\xff\x82\x5d\x67\x38\x9a\xce\xfb\xfd\xee\x74\x06\xff\x73\x9c\x1d\x48\xd2\x80\xc5\xde\x06\x62\xf1\xf4\x80\x29\x77\x3a\xa3\xeb\xf1\x68\x3e\x3c\xbb\x1e\x0c\xc1\x2d\xea\xf1\x26\x07\x16\x83\x9f\x3c\x6d\x13\x11\x07\xfc\x3b\xb8\x59\xf4\x09\xda\xc5\x3f\x87\x59\xe4\x8c\x86\x57\xcb\xc1\x50\xb7\x81\xef\xec\x90\x16\x52\xb1\x41\xfe\xdd\xd9\x01\xf7\x00\xee\x5a\xfb\xc7\xab\x8f\x6e\xf9\xc8\x22\xd8\x73\x5b\x85\x83\x5b\x9a\xb9\xa5\xdb\xe7\x2f\x5f\xdd\xa3\x02\x60\xe6\x1f\x7c\x70\x76\xa0\xad\xbb\xee\x5a\xed\xfd\xe3\xd5\x9e\xcb\xb3\xd7\x76\x1c\xbc\x1e\xa6\xc1\xeb\xdd\x51\xe9\xf3\x01\xf6\xda\x2e\xcf\x3e\x41\xbb\x1d\x07\x85\xd1\x61\x2a\x1f\x96\x91\xce\x36\x99\x5f\x75\xa1\x5d\x5d\xc6\x1f\xdd\xac\xce\x3c\x6d\x9d\x1f\xa3\xc9\x59\xe7\xfc\x64\x32\x85\x56\xe1\xf5\xf9\xcb\xd7\x6f\xc7\x8e\x33\xee\x2d\xe0\x04\x96\xe0\x2a\x04\xd0\x82\xe2\x2f\x45\xd0\x7f\x07\x52\x35\x20\xd5\x80\x54\x2c\xae\xab\xe6\xa4\x9a\x93\x4a\xbc\x51\x75\x48\xd5\x21\x55\x60\x76\xd4\xb2\x2e\xc9\xba\x46\xc6\x78\x5d\xd6\x23\x59\x8f\x64\x7e\xf2\x66\xd3\x4b\x92\x5d\x92\x6c\x9b\x26\xb6\x6c\xda\x5b\xc0\x0d\x0c\xa1\xad\x10\x80\x0c\xc5\x7e\x53\x38\xa4\x52\xab\x5e\xba\xbb\xc7\x95\xf4\x18\xc2\x12\x0e\x15\x22\x0f\x93\x08\xa9\x94\xcb\x92\xef\x65\xaa\x9c\x53\x49\xda\xa3\x39\xc6\xdd\x3d\xe6\xca\x7f\x41\xfe\x0b\xdb\x9f\x57\xfc\x17\x75\x7f\x4e\xfe\x0b\xcb\x9f\x37\xf8\x9f\xd3\x0c\xce\x8d\x09\x66\x91\x52\xe6\x54\x92\x26\x54\x28\x4c\xb4\xc7\xb2\x1c\xb7\x44\x5a\xb8\xc9\xeb\xd3\xee\xc3\x52\xed\xd4\x37\x27\x89\xc3\xea\xa4\xfa\xd6\x85\x0b\xdf\x4e\xfb\x82\x3c\x2e\x6c\x8f\xea\x34\x2e\x6c\x0f\xfe\xc6\xe3\x0c\x96\x70\xa4\x10\x5d\x35\x73\x21\x91\x4a\x7a\x18\x81\xc9\xe1\x21\x4b\xc4\x2a\xaf\x75\x1b\x85\x9a\xec\xfb\x65\x6d\x50\x4a\x66\x30\x54\x49\x66\x24\xe6\x99\x11\x63\xb8\xa2\xaa\x0e\x43\xf5\x6a\x18\x63\x60\x14\x66\xb7\xf7\x15\xe6\x22\x4c\x68\x2c\x13\xf3\x69\xa8\x8f\x65\x62\x25\x49\x9b\xc7\x62\xba\x8d\xa2\x9a\xa4\x59\x61\x92\xdc\x52\x92\x5b\x93\x84\xd7\x92\xdc\xda\x49\x78\x63\x92\x5b\x6b\x1f\xde\x98\xa4\x59\x61\x92\x4c\x29\xc9\xd4\x7e\x05\xcd\x07\x9f\x9a\x77\x08\xcd\x07\x9f\x5a\xaf\xc0\x28\xf8\x1f\x2f\xd1\xea\x3f\x1c\xee\x0b\x81\x4e\x34\xa6\xcf\xc1\xd8\x24\x8a\x91\xf3\xa6\x44\x63\xf2\x33\x8a\x6a\xa2\x66\x05\xfd\xb0\xaf\xe9\x87\x7d\x4d\xc2\xdd\xca\x47\xb9\xa8\x9d\x52\xa6\x53\x52\x79\x0f\x1b\xac\x4d\xe9\x7d\x05\x22\xd5\xf0\xbd\x6e\x94\x7b\x5d\x51\xa2\x2b\x72\x7a\xc2\x38\x33\x89\x26\xdd\x31\x1c\xb7\xca\x87\x07\xc8\x4a\x80\x0c\x3c\xcd\xe4\x92\xc9\x35\xc3\x80\x61\x89\x18\x82\x27\xe5\xde\x26\xc9\xc0\xfb\xff\x26\xc9\xe4\x32\x2d\x2a\x29\x61\xf4\x2c\xac\x79\x96\x16\x6b\xc2\xa8\x30\x5a\xbd\x68\xf5\xa2\xd5\x8b\x9e\xe9\x2d\xb1\xc5\x93\x8f\xe5\x8f\x96\xbf\x1f\x41\x26\x51\x86\xe0\x47\xc4\xc9\x83\x64\xcc\xe6\x84\xe4\x84\xc5\xf1\x12\x70\x8b\x79\x2c\xc1\xa3\x66\xd6\xb0\x95\x60\x1b\xc1\xba\x04\xa1\x06\x6b\x08\x65\x89\x01\xd3\x28\x8c\x80\x29\x4e\xc8\x41\x8b\x9c\x18\x54\xf1\x51\x31\x3f\x21\x90\x20\x60\xf0\x53\x31\x61\x28\x37\x09\x0d\x23\xbb\x42\x54\x0c\xb7\x0e\x47\xf1\x25\xe3\x6b\xe6\x11\x7c\x5f\x32\x3e\x1d\xe8\x97\x3c\xfe\x2f\xa1\x99\x67\x4d\x25\x80\x9e\x50\x13\x17\x90\x68\xee\x59\x52\xcf\x9a\xa9\x50\x89\x46\x42\x6e\x2d\x98\xae\x26\x89\xc5\x59\x3a\x49\xd9\x2a\xc5\x28\x8d\x80\x44\x9e\x34\x41\x2a\x31\x14\xea\x6a\x0a\xcd\xa1\x95\x42\xc8\x06\xb1\x5e\x43\x22\xd4\x2b\x48\x92\x4a\x97\x32\x2f\xbe\xd6\x74\xf5\xbc\xb4\x58\xbf\xb9\x6e\xe5\xb5\x55\x9a\x67\x6b\x5c\x72\xbb\x17\xe0\x6a\xda\x1c\x5e\x24\x88\x44\x0c\xc5\x7f\x36\xcc\x8a\x57\x56\x3e\xad\xd8\x0b\x82\xcf\x38\x3a\xff\x06\x00\x00\xff\xff\x03\xa4\xf6\x92\x1d\x0c\x00\x00")

func dataEn_usWebAffBytes() ([]byte, error) {
	return bindataRead(
		_dataEn_usWebAff,
		"data/en_US-web.aff",
	)
}

func dataEn_usWebAff() (*asset, error) {
	bytes, err := dataEn_usWebAffBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/en_US-web.aff", size: 3101, mode: os.FileMode(420), modTime: time.Unix(1567968621, 0)}
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
	"data/en_US-web.dic": dataEn_usWebDic,
	"data/en_US-web.aff": dataEn_usWebAff,
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
	"data": &bintree{nil, map[string]*bintree{
		"en_US-web.aff": &bintree{dataEn_usWebAff, map[string]*bintree{}},
		"en_US-web.dic": &bintree{dataEn_usWebDic, map[string]*bintree{}},
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
