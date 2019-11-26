// Code generated for package mustgather by go-bindata DO NOT EDIT. (@generated)
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

var _templatesJobTemplateYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\x4f\x6f\xdb\x3e\x0c\xbd\xf7\x53\x10\x45\x8f\x3f\x27\xf8\x01\x3b\x09\xd8\x61\x68\xd1\xa1\xdb\xd2\x16\x2d\xb0\x6b\x41\xcb\x4c\xa2\xd5\xfa\x33\x89\xce\x66\x18\xfe\xee\x83\xfc\x27\x51\x1c\xb7\xc8\x65\xcb\xc9\x21\x1f\x1f\x1f\x45\x3d\xa1\x53\xdf\xc9\x07\x65\x8d\x80\x1c\x59\x6e\x97\xbb\xff\x2f\x5e\x95\x29\x04\x7c\xb1\xf9\x85\x26\xc6\x02\x19\xc5\x05\x80\x41\x4d\x02\x9a\x06\x16\x0f\xf9\x0f\x92\xbc\x22\xc6\xc5\x3d\x6a\x82\xb6\x1d\xd2\xc1\xa1\x9c\xc7\x74\x99\x08\x0c\x8e\x64\x64\x63\xd2\xae\x44\xa6\xf8\x0d\x30\x46\xe3\xcf\x53\x60\xf4\xfc\x68\x4b\x25\x6b\x01\x0f\xe6\x16\x55\x59\x79\x1a\xd2\xd2\x1a\x46\x65\xc8\x87\xb1\x20\x03\x32\xec\x6b\x67\x95\xe1\x31\x16\xa3\xcb\x5c\x99\x65\x8e\x61\x9b\xc4\x32\x99\x02\xaa\xe0\x97\xa5\x95\x58\x76\xd0\xca\x95\x16\x8b\x7d\x5e\x69\xdc\x90\x80\x9f\x15\xd6\x0b\x65\x97\x9e\x8a\x2d\x72\x26\xad\x5b\xea\x2a\x70\xb6\x41\xde\x92\xcf\xac\x23\x8f\x6c\xbd\x88\xc3\x04\xde\x57\xf7\xa7\x35\xa1\xdc\xd9\xb2\xd2\xb4\xb2\x95\xe1\x90\x2a\xd5\x31\xf2\x88\xbc\x15\x70\x4c\x5e\xb1\xab\x0e\x9c\x23\xeb\x3b\x90\xb7\xb9\x26\x52\xce\xe0\x22\xb3\x4b\x45\x0e\x13\x05\xf2\xf1\x2b\x21\xda\x61\x59\xd1\xad\xb7\x5a\x24\x41\x80\x40\xd2\x13\x7f\xa5\xfa\x89\xd6\xc7\x99\xb1\xb7\xc4\x40\x99\x46\x83\x1b\xd2\x64\x38\xeb\x2b\x26\xd0\x57\xaa\x67\xba\x8e\x72\x1c\x86\xf0\xcb\xfa\xe2\xdf\xca\x39\xe9\x9a\x25\x14\xea\x44\x8c\x80\xcb\x68\x89\x67\x47\x72\x71\x8d\x81\xee\x6e\xa0\x6d\x2f\x4f\x8a\xe3\x2a\x5e\xfa\x55\xbc\x9c\x6c\x7e\x20\x7a\xef\x7a\xcc\xf1\x9c\x6c\x7d\x8e\xa7\x07\xc1\xcc\x6f\x28\x54\x46\xf1\xf5\xc1\x77\x43\xae\x69\x32\xf0\x68\x36\x04\x57\xca\x14\xf4\xfb\x3f\xb8\xa2\xb2\x3b\x3a\x10\x1f\x87\x71\x57\x55\xe0\xcf\x5d\x97\xbb\x68\xa7\x00\x6d\x7b\xc4\x9c\x81\xb4\x5a\xa3\x29\xc4\xd4\x97\xd1\x91\xbd\xbe\xa9\x23\x9b\xe6\xd0\xa8\x7b\x7a\xd2\x25\x0e\x23\x45\x4c\x27\x2a\x45\xfc\x3d\xfb\xc5\xa3\x20\x53\x1c\x9a\x05\xf2\x3b\x25\xe9\x93\x94\x91\xfb\x7e\xff\x74\x76\x87\xf2\x7c\x94\x7c\xa2\x75\xf2\x8c\x1e\x74\xa6\xef\x9b\x76\x5c\xdf\x28\x2f\xa0\x99\xce\xfb\xa6\xde\xb3\xab\x8e\xae\x48\x76\x9e\x17\xfa\xbf\xa9\x8d\xfa\xc8\x64\xd0\x78\xd9\x57\x7b\x92\x61\xde\xe7\x0e\x99\x4c\x3d\x73\xe9\x72\x94\xaf\x76\xbd\xfe\xa6\xb4\x62\x01\x1f\xfe\x04\x00\x00\xff\xff\x78\x81\x59\x2a\xa6\x06\x00\x00")

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

	info := bindataFileInfo{name: "templates/job.template.yaml", size: 1702, mode: os.FileMode(420), modTime: time.Unix(1574785482, 0)}
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
