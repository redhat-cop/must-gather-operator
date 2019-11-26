// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/job.template.yaml
package mustgather

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesJobTemplateYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\x4d\x8f\xd3\x30\x10\xbd\xef\xaf\x98\xc3\x1e\x49\x23\x24\x4e\x96\x38\xa0\x5d\x81\x16\xe8\x82\xb6\x12\xd7\xd5\xc4\x99\xb6\xa6\xf1\x07\xf6\xa4\x50\x45\xfd\xef\xc8\xf9\x68\x9d\x34\xbb\xea\x05\x72\x8a\x66\x9e\xdf\xbc\x19\xcf\x33\x3a\xf5\x83\x7c\x50\xd6\x08\x28\x90\xe5\x36\xdf\xbf\xbd\xd9\x29\x53\x0a\xf8\x6c\x8b\x1b\x4d\x8c\x25\x32\x8a\x1b\x00\x83\x9a\x04\x34\x0d\x2c\xbe\x15\x3f\x49\xf2\x92\x18\x17\x8f\xa8\x09\x8e\xc7\x3e\x1d\x1c\xca\x79\x4c\x9b\x89\xc0\xe0\x48\x46\x36\x26\xed\x2a\x64\x8a\xff\x00\x43\x34\x7e\xd2\x1a\x46\x65\xc8\x87\x21\x92\x01\x19\xf6\x07\x67\x95\xe1\x21\x16\xa3\x79\xa1\x4c\x5e\x60\xd8\x26\xb1\x4c\xa6\x80\x3a\xf8\xbc\xb2\x12\xab\x16\x5a\xbb\xca\x62\x79\xca\x2b\x8d\x1b\x12\xf0\xab\xc6\xc3\x42\xd9\xdc\x53\xb9\x45\xce\xa4\x75\xb9\xae\x03\x67\x1b\xe4\x2d\xf9\xcc\x3a\xf2\xc8\xd6\x8b\xa8\x36\xf0\xe9\x74\x37\x8e\x09\xe5\xde\x56\xb5\xa6\xa5\xad\x0d\x87\x54\xa9\x8e\x91\xef\xc8\x5b\x01\x63\xf2\x9a\x5d\x7d\xe6\x1c\x58\x5f\x81\xbc\xcc\x35\x91\x72\x05\x17\x99\x7d\x2a\xb2\xef\x28\x90\x8f\x7f\x09\xd1\x1e\xab\x9a\x3e\x7a\xab\x45\x12\x04\x08\x24\x3d\xf1\x17\x3a\x3c\xd1\x7a\x9c\x19\x6a\x4b\x0c\xb4\x44\x83\x1b\xd2\x64\x78\xd5\xe2\x27\xc0\x1d\x1d\x66\x6a\x0e\x62\x1c\x86\xf0\xdb\xfa\xf2\x7f\x8a\xb9\xa8\x99\x25\x04\xea\x42\x4a\xb7\xee\x2b\x47\x72\x71\x87\x81\x1e\xee\x3b\x3b\x8c\x8f\xc6\x4b\x78\xee\x2e\xe1\xf9\xe2\xce\x7b\x9a\xd7\x16\x63\x8e\xe7\xe2\xbe\xe7\x78\x3a\x10\xcc\x7c\xfd\x41\x65\x14\xdf\x9d\x1d\xd7\xe7\x9a\x26\x03\x8f\x66\x43\x70\xab\x4c\x49\x7f\xde\xc0\x2d\x55\xed\xdc\x40\xbc\xef\xbb\x5d\xd6\x81\x3f\xb5\x55\x1e\xa2\x91\x02\x1c\x8f\x23\xe6\x0c\xa4\xd5\x1a\x4d\x29\xa6\x8e\x8c\x5e\xec\xf4\x4d\xbd\xd8\x34\xe7\x42\xc9\x18\xbb\xe6\xfb\x96\x22\xa6\x15\x95\x22\xfe\x9d\xf1\xe2\x28\xc8\x94\xe7\x62\x81\xfc\x5e\x49\xfa\x20\x65\xe4\x7e\x3c\xbd\x8a\xed\x50\x56\xa3\xe4\x13\xad\x93\x17\xf2\xac\x33\x7d\xd9\xb4\xe3\xc3\xbd\xf2\x02\x9a\x69\xbf\x2f\xea\xbd\xfa\xd4\x68\x45\xb2\x6b\x7c\xd0\xd9\x28\x35\x50\x17\x99\xb4\x79\x37\xe2\xe8\xbb\xed\xa8\x92\x9e\x67\x56\xae\x40\xb9\xb3\xeb\xf5\x57\xa5\x15\x0b\x78\xf7\x37\x00\x00\xff\xff\x4d\x45\x1d\x4f\x7f\x06\x00\x00")

func templatesJobTemplateYamlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJobTemplateYaml,
		"templates/job.template.yaml",
	)
}

func templatesJobTemplateYaml() (*asset, error) {
	bytes, err := templatesJobTemplateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/job.template.yaml", size: 1663, mode: os.FileMode(420), modTime: time.Unix(1574777678, 0)}
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
	"templates/job.template.yaml": templatesJobTemplateYaml,
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
		"job.template.yaml": &bintree{templatesJobTemplateYaml, map[string]*bintree{}},
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
