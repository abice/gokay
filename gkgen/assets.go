// Code generated by go-bindata.
// sources:
// gkgen/bcp47.tmpl
// gkgen/hexgen.tmpl
// gkgen/len.tmpl
// gkgen/length.tmpl
// gkgen/main.tmpl
// gkgen/notnil.tmpl
// gkgen/required.tmpl
// gkgen/uuid.tmpl
// DO NOT EDIT!

package gkgen

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

var _gkgenBcp47Tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x24\x8e\x41\x4b\x03\x31\x10\x46\xef\xf9\x15\x9f\x7b\x90\xb6\xe0\xe6\x52\x28\x54\x7a\x50\x51\xe8\x45\x7a\xf2\x9e\x36\x93\xed\x60\x9a\x94\x64\x50\xca\x30\xff\x5d\xba\x9e\x1f\xef\xf1\x54\xfd\xca\x01\xaf\x6f\x87\xf5\xe6\x2b\x64\x8e\x41\x6a\xc3\x44\x85\x5a\x10\xea\x38\xd5\x48\x90\x73\x10\xfc\x72\xce\xf8\xa1\xc6\xe9\x06\x4e\x08\x48\x4c\x39\x82\x3b\xc2\xbf\x8e\x53\xbd\x5c\x83\xf0\x31\x13\xba\x34\x2e\x93\x03\xce\x22\xd7\xbe\xf5\x5e\x6a\xcd\x7d\x64\x92\x34\xd6\x36\xf9\xb3\x5c\xb2\x3f\x9e\xae\xeb\x8d\x5b\x79\x33\xa7\x1a\x29\x71\x21\x0c\x73\x6a\x30\x73\x9c\x40\xad\x61\xbb\xc3\x54\xbf\xc3\x6d\xdc\xf7\x19\x2d\x54\x39\xa1\x54\xc1\x62\xdf\x0f\xd2\x30\x62\x69\xf6\xa8\x4a\x25\x9a\xf5\x51\x75\xfc\xb8\x8f\x7d\x86\x0b\x99\x2d\x9f\xe7\xc8\xc3\x0e\x85\x33\xd4\x01\xaa\x78\x89\xf1\xbd\xb5\x7a\x57\x07\x6a\x6d\x80\x99\xbb\x2f\x50\x89\x78\x32\x73\x7f\x01\x00\x00\xff\xff\x8f\xb3\x11\x82\x14\x01\x00\x00")

func gkgenBcp47TmplBytes() ([]byte, error) {
	return bindataRead(
		_gkgenBcp47Tmpl,
		"gkgen/bcp47.tmpl",
	)
}

func gkgenBcp47Tmpl() (*asset, error) {
	bytes, err := gkgenBcp47TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gkgen/bcp47.tmpl", size: 276, mode: os.FileMode(420), modTime: time.Unix(1477960797, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gkgenHexgenTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x1c\xcc\x31\x0a\xc2\x40\x10\x05\xd0\x7e\x4e\xf1\xdd\x42\x92\xc2\x3d\x80\x92\xc2\x42\x49\x1a\xf1\x0a\x81\x9d\xc8\x62\xdc\xc0\x6c\x8a\xc8\xf0\xef\x2e\xe6\x00\xef\x89\x7b\xd2\x29\x17\x45\xe8\x75\x0b\xa4\xe4\x09\x6a\x86\x73\x87\xd7\xf2\x1e\xbf\x71\xa8\xbd\x6e\x8d\x7b\x9e\x50\x96\x15\xcd\x50\x9f\xab\x21\xa2\x25\x8f\xee\x5a\x12\x59\xa3\x7b\xbc\x67\x9d\xd3\x63\xfc\x28\xd9\x5e\xf6\xe2\xd0\xa1\xe4\x19\x2e\x80\x3b\xae\x29\xdd\xcc\x96\x3f\x0d\x6a\x16\x40\x0a\x65\x1f\x70\x22\xe5\x17\x00\x00\xff\xff\x31\x82\x3d\x47\x8a\x00\x00\x00")

func gkgenHexgenTmplBytes() ([]byte, error) {
	return bindataRead(
		_gkgenHexgenTmpl,
		"gkgen/hexgen.tmpl",
	)
}

func gkgenHexgenTmpl() (*asset, error) {
	bytes, err := gkgenHexgenTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gkgen/hexgen.tmpl", size: 138, mode: os.FileMode(420), modTime: time.Unix(1477695249, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gkgenLenTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x91\x41\x8b\xd6\x30\x10\x86\xef\xfd\x15\xef\xf6\x20\xed\xa2\xfd\xee\xca\x77\x58\x16\x05\xc1\x5d\x16\xd4\xdb\x5e\x62\x3b\x6d\x86\x4d\x93\x3a\x99\x22\x25\xe4\xbf\x4b\x52\x44\x44\xf6\x36\x43\xe9\xfb\xce\xf3\xa4\x49\x69\xa2\x99\x3d\xa1\x75\xe4\xdb\x9c\x9b\x94\x2e\xb7\xf8\x1e\x09\x6a\x09\x3f\x76\x76\x0a\xf6\x70\xe4\x31\xef\x7e\x54\x0e\x1e\x73\x90\xfa\x55\x8f\x8d\x22\xd4\x1a\x45\xdc\xb7\x2d\x88\x82\x75\xc0\xed\xa5\xc6\x80\x67\x04\x41\x07\x8e\x0f\x66\xc3\x80\xbe\xce\x77\x22\xe6\x38\x37\xfa\x89\xe1\x13\x93\x9b\xbe\x1d\x1b\xa1\x8d\x2a\xec\x97\xb6\xcf\xb9\x01\x2e\x17\xdc\x5b\x1a\x5f\xd8\x2f\xa5\x7c\x51\xdb\xa0\x24\xde\x74\x8e\x7c\x17\x87\x94\xce\x5f\x1f\xcd\x4a\x39\xf7\xb8\x5e\x91\xd2\xf0\x64\xc4\xac\x65\x4d\x0d\x00\xa4\x84\xbb\x69\xfa\x28\x12\x04\x03\x5a\x2a\x43\x1c\x1e\xe9\x57\xf7\xdc\x7e\xa9\xa1\x78\xe0\xb8\x1a\x1d\xed\x73\xdb\xb7\xa8\xcd\xf5\x76\x72\x91\xce\xb5\xfa\xb8\x37\xce\x21\x86\x95\xb0\x90\x27\xe1\x11\x2b\xa9\x0d\xd3\x09\x3f\x1a\x0f\x62\xb5\x24\x98\x02\x84\x66\x47\xa7\x28\xa3\x90\xdd\x43\x79\xa5\x7a\x4f\x90\x1a\xa2\xb6\x50\xd5\x0a\x0d\x98\x79\xd9\x85\x10\x76\x2d\x7c\xc5\xeb\xc9\x0b\x8e\x18\x83\x08\x8d\x7f\x9c\x56\x01\x24\x82\xf7\x57\x2c\xe1\xc5\x1c\xc3\x09\xf1\xb5\x8a\xeb\xfe\xf2\xbf\x45\x4a\x3c\xc3\x07\x45\xf7\x39\x3e\x69\xa1\xef\x73\x7e\x93\x12\xf9\x29\xe7\xff\xec\x7d\xa8\xa9\x37\x57\x78\x76\xaf\xaa\xfb\x47\x8f\x9f\x50\x1f\xb9\x0c\xef\x72\x6e\x7e\x07\x00\x00\xff\xff\xd4\xa1\xf5\x54\x48\x02\x00\x00")

func gkgenLenTmplBytes() ([]byte, error) {
	return bindataRead(
		_gkgenLenTmpl,
		"gkgen/len.tmpl",
	)
}

func gkgenLenTmpl() (*asset, error) {
	bytes, err := gkgenLenTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gkgen/len.tmpl", size: 584, mode: os.FileMode(420), modTime: time.Unix(1477966155, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gkgenLengthTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x24\xcc\x41\xaa\xc2\x30\x14\x85\xe1\x79\x56\x71\x5e\x06\x8f\x16\x34\x0b\x50\x3a\x70\xa0\x20\x88\x14\x5c\x41\x20\xb7\x35\xd8\xa6\x70\x9b\x89\x5c\xce\xde\xc5\x3a\xfe\xf9\x3f\x67\x96\x64\xc8\x45\xe0\x6f\x52\xc6\xfa\xf4\xa4\xcb\x03\x44\x15\x87\x0e\xe3\xf2\x8a\xef\xf0\x2b\x8f\xaa\xb9\x8c\x8d\x59\xe8\xa3\xc6\x99\xdc\xc1\x2c\x0f\x28\x4b\x45\x73\x5d\xfb\xaa\x08\x68\xc9\x7f\x33\x29\x89\x5c\x83\x59\xb8\x64\x99\xd2\x3d\xce\x42\xb6\xc7\x4d\xfd\xeb\x50\xf2\x04\x73\x80\x19\x4e\x29\x9d\x55\x97\xef\xea\x45\xd5\x83\x74\x74\x9b\x80\x3d\xe9\x3e\x01\x00\x00\xff\xff\x40\x65\x23\xa6\xa0\x00\x00\x00")

func gkgenLengthTmplBytes() ([]byte, error) {
	return bindataRead(
		_gkgenLengthTmpl,
		"gkgen/length.tmpl",
	)
}

func gkgenLengthTmpl() (*asset, error) {
	bytes, err := gkgenLengthTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gkgen/length.tmpl", size: 160, mode: os.FileMode(420), modTime: time.Unix(1477696253, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gkgenMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\xcd\x6e\xd4\x30\x10\xbe\xfb\x29\x3e\x45\x2b\x91\xac\xba\x49\xcf\x95\xb6\x07\x50\x0f\x3d\x14\x24\x40\x5c\x10\x07\x93\xcc\x26\xd6\x3a\x76\xe4\x38\x85\x62\xf9\xdd\x91\x7f\xb2\xa4\x2d\x42\xec\xc1\xf2\x7a\xc6\xdf\x9f\x27\xce\x1d\xd0\xd1\x49\x28\x42\x31\x10\xef\xc8\x14\xde\xb3\x89\xb7\x67\xde\x13\x9c\xab\xf3\xd6\x7b\xc6\xc4\x38\x69\x63\x51\xb2\xa2\x17\x76\x58\xbe\xd7\xad\x1e\x9b\x5f\xa4\x5a\xdd\x91\x69\x7a\x7d\xe6\x4f\x69\x2d\x58\xc5\x9c\x23\xd5\xe1\x10\xee\x6d\x39\x66\x6b\x96\xd6\x06\x0e\xe7\x76\x66\x91\xf4\xc0\x27\xdc\x1c\x51\x87\xfd\x1c\xfb\x4f\x8b\x6a\x51\xce\xd8\x3b\x57\x2b\x3e\x92\xf7\x15\xbe\x70\x29\x3a\x6e\x09\x65\x85\x92\x8c\xd1\xa6\x82\x63\x80\x73\xcd\x1e\x1f\x94\x7c\xc2\xc8\xcf\x04\x3b\x10\x62\x75\xc6\xc8\x27\x88\x13\x7e\xd0\x1b\x43\xe8\xb5\x50\x3d\xac\xc6\x32\x13\x84\xad\xb1\x6f\xbc\x8f\xd7\x43\x4f\x6f\x51\x4a\x52\x59\x43\x85\x6b\xc4\x22\x8d\x41\x58\xc0\x2d\xa3\xab\xfa\x2e\x20\x3f\xf0\xa9\x4a\x37\x83\xbf\x8c\xd2\xec\x71\x6f\xc9\x04\x81\x76\x30\x7a\xe9\x87\x28\x25\xe0\x45\x21\x56\xa3\x35\x94\xca\xe9\x78\xfe\x23\xc1\x70\xd5\x13\x76\xc1\xea\x15\x76\x27\x41\xb2\x7b\x91\x08\x10\x09\x77\xe9\xe4\xe6\x98\xbb\xea\x8f\x2f\x3b\x36\x66\x76\x17\x33\x6b\xbd\x69\xf0\x96\x7a\xa1\x22\x54\x60\x83\xf7\x6b\xb0\x42\xab\x99\xc5\xae\x47\x6e\x72\x84\xdb\xb6\x4d\x00\x9f\xa4\x68\x89\x65\xc6\x03\xb2\x7a\x73\xaf\x3a\xfa\x79\x95\x68\xa3\xc4\x24\x36\x73\x47\xf6\xd5\x42\xfd\x3e\xa0\x1e\x2e\x25\xe7\xf0\x8e\x4b\xf9\x99\xc6\x49\x86\x8c\x12\x46\x5a\x2f\xde\x0e\x21\xef\x30\x4d\xe1\xaf\x38\x41\x92\x2a\x5f\xe9\xac\x70\x8b\xeb\x38\x18\xe1\x47\xe3\xd7\x62\x53\x2c\xbe\xe1\xf8\xda\x5a\xec\xcd\xb0\x4d\x83\x3b\xd5\xfd\x3b\x9f\xe7\x0f\xbf\xaa\xda\x4e\x62\x3b\x50\x7b\xfe\xeb\x28\x62\xe0\x8f\xf9\xf9\xff\x6f\x04\x57\x9f\x63\x85\xdb\xd5\x97\x21\xbb\x18\x05\x1a\x59\xd0\xfd\x5c\x44\xae\x29\x21\x99\x4f\x5f\xa0\xf7\xec\x77\x00\x00\x00\xff\xff\xa6\x69\x3f\x4e\xe3\x03\x00\x00")

func gkgenMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_gkgenMainTmpl,
		"gkgen/main.tmpl",
	)
}

func gkgenMainTmpl() (*asset, error) {
	bytes, err := gkgenMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gkgen/main.tmpl", size: 995, mode: os.FileMode(420), modTime: time.Unix(1477695218, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gkgenNotnilTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\xcd\xb1\x8a\xc3\x30\x10\x04\xd0\xde\x5f\x31\xa8\xba\x2b\x4e\xfe\x02\x17\x57\xdc\x41\x1a\x11\x48\xeb\x46\x41\x2b\x58\xb2\x91\x8c\x64\x27\xc5\xa2\x7f\x0f\xb2\xdd\xcd\x0e\xbc\x1d\xd5\x40\x91\x13\xc1\xb8\xbc\x3a\x16\xd3\xda\x00\xa8\x82\x23\x2e\xd5\x6d\x22\xfe\x2e\x04\x8b\xbd\xe7\x88\x6a\x55\xed\x3f\x93\x04\xe7\x9f\xd4\x1a\xa6\x09\x89\x05\x3a\x00\x3b\xfc\x0d\xe1\xaf\x94\x5c\x60\x61\xa8\x87\x6a\x1d\xbd\xbf\x66\xc3\x15\x8e\x65\x36\xdf\xe6\x78\x76\x0e\x91\x54\x3a\x8a\x71\xc4\xed\xc1\xcb\x42\x01\x2e\x27\x5c\x33\xa7\x95\x0a\x5e\x5e\x36\x42\xcc\xa5\x73\x50\x3f\xfd\xca\x39\x9d\x3c\x85\xae\x55\x7b\xf8\x69\x6d\xf8\x04\x00\x00\xff\xff\x19\x5a\x9a\x08\xd2\x00\x00\x00")

func gkgenNotnilTmplBytes() ([]byte, error) {
	return bindataRead(
		_gkgenNotnilTmpl,
		"gkgen/notnil.tmpl",
	)
}

func gkgenNotnilTmpl() (*asset, error) {
	bytes, err := gkgenNotnilTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gkgen/notnil.tmpl", size: 210, mode: os.FileMode(420), modTime: time.Unix(1477695187, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gkgenRequiredTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x8e\xb1\x6a\xc3\x40\x10\x44\x7b\x7d\xc5\xb0\x55\x52\xe4\xfe\x40\x45\x8a\x04\xd2\xa8\x4a\xa9\x46\xe1\x56\x61\xe1\x7c\xb2\xf6\x2c\x1b\x7b\xb9\x7f\x37\x27\x59\x8d\x31\x6a\xec\x6e\x66\x60\xe6\x8d\x99\xe7\x5e\x22\x83\x94\xc7\x49\x94\x3d\xe5\x5c\x01\x66\x90\x1e\x3f\xa9\x99\x42\xe8\xfe\x02\xc3\x61\xce\xa5\x47\x72\x66\xee\x5b\x38\xf8\xa6\xdb\x71\xce\xa8\x6b\x44\x09\xb0\x0a\x98\x8b\x9f\xde\x7f\xa9\x0e\x0a\x07\xe2\x22\x92\x6b\xf8\xf4\xd6\x92\x24\xac\x94\x96\xde\x69\x59\xbc\xd1\x38\x24\x2e\xeb\x3c\x62\x19\xff\x3d\xef\x19\x94\x0e\x2a\xf1\x9f\x36\xe1\x44\xaf\x60\xcf\xc1\xb1\x53\x5c\x58\x87\x3b\xc8\x6a\xcb\xa7\xad\x27\x8f\xaa\xcf\x3d\x8b\xbe\x78\xb3\x22\x3e\x72\xae\xae\x01\x00\x00\xff\xff\xd8\x16\x05\x41\xb0\x01\x00\x00")

func gkgenRequiredTmplBytes() ([]byte, error) {
	return bindataRead(
		_gkgenRequiredTmpl,
		"gkgen/required.tmpl",
	)
}

func gkgenRequiredTmpl() (*asset, error) {
	bytes, err := gkgenRequiredTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gkgen/required.tmpl", size: 432, mode: os.FileMode(420), modTime: time.Unix(1478006159, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gkgenUuidTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x1c\xcc\x41\xaa\xc2\x30\x14\x85\xe1\x79\x56\x71\x5e\x06\x8f\x16\x34\x0b\x50\x3a\x10\x54\xe8\x44\x3a\xe9\x02\x02\xb9\x95\x60\x9b\xc2\x4d\x27\x72\x39\x7b\x97\x76\xfc\xf3\x7f\xce\x2c\xc9\x94\x8b\xc0\x8f\x63\x7f\xf7\xa4\xcb\x13\x44\x15\x97\x0e\xef\xf5\x13\xbf\xa1\xaf\x7b\x69\xcc\xc2\x10\x35\x2e\xe4\x09\x66\x79\x42\x59\x37\x34\x7d\x1d\x36\x45\x40\x4b\xfe\x9b\x49\x49\x64\x0d\x66\xe1\x99\x65\x4e\xaf\xb8\x08\xd9\x5e\x0f\xef\xaf\x43\xc9\x33\xcc\x01\x66\xb8\xa5\xf4\x50\x5d\xf7\xd5\x8b\xaa\x07\xe9\xe8\x0e\x01\x67\xd2\xfd\x02\x00\x00\xff\xff\x91\x6f\xc1\xad\x98\x00\x00\x00")

func gkgenUuidTmplBytes() ([]byte, error) {
	return bindataRead(
		_gkgenUuidTmpl,
		"gkgen/uuid.tmpl",
	)
}

func gkgenUuidTmpl() (*asset, error) {
	bytes, err := gkgenUuidTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gkgen/uuid.tmpl", size: 152, mode: os.FileMode(420), modTime: time.Unix(1477961719, 0)}
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
	"gkgen/bcp47.tmpl": gkgenBcp47Tmpl,
	"gkgen/hexgen.tmpl": gkgenHexgenTmpl,
	"gkgen/len.tmpl": gkgenLenTmpl,
	"gkgen/length.tmpl": gkgenLengthTmpl,
	"gkgen/main.tmpl": gkgenMainTmpl,
	"gkgen/notnil.tmpl": gkgenNotnilTmpl,
	"gkgen/required.tmpl": gkgenRequiredTmpl,
	"gkgen/uuid.tmpl": gkgenUuidTmpl,
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
	"gkgen": &bintree{nil, map[string]*bintree{
		"bcp47.tmpl": &bintree{gkgenBcp47Tmpl, map[string]*bintree{}},
		"hexgen.tmpl": &bintree{gkgenHexgenTmpl, map[string]*bintree{}},
		"len.tmpl": &bintree{gkgenLenTmpl, map[string]*bintree{}},
		"length.tmpl": &bintree{gkgenLengthTmpl, map[string]*bintree{}},
		"main.tmpl": &bintree{gkgenMainTmpl, map[string]*bintree{}},
		"notnil.tmpl": &bintree{gkgenNotnilTmpl, map[string]*bintree{}},
		"required.tmpl": &bintree{gkgenRequiredTmpl, map[string]*bintree{}},
		"uuid.tmpl": &bintree{gkgenUuidTmpl, map[string]*bintree{}},
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

