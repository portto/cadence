// Code generated by go-bindata. DO NOT EDIT.
// sources:
// examples/arrays.cdc (795B)
// examples/arrays.cdc.abi.json (2.344kB)
// examples/car.cdc (733B)
// examples/car.cdc.abi.json (1.386kB)
// examples/dictionares.cdc (292B)
// examples/dictionares.cdc.abi.json (2.242kB)
// examples/events.cdc (484B)
// examples/events.cdc.abi.json (216B)
// examples/functions.cdc (409B)
// examples/functions.cdc.abi.json (1.232kB)
// examples/resources.cdc (494B)
// examples/resources.cdc.abi.json (1.416kB)

package abi

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _arraysCdc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcb\xae\xdb\x20\x18\x84\xf7\x7e\x8a\x59\x62\x1d\xeb\xc8\x97\x9d\x5d\x74\x94\x45\x17\x59\xf5\x01\x10\x0b\x42\xa0\x41\x22\xc4\xc2\xb8\x55\x1a\xe5\xdd\x2b\x30\xce\x4d\xaa\x9a\x8d\x65\x7e\x8f\xff\x99\x8f\x19\xe7\x1d\xf4\xec\x20\xf6\x7b\xe5\x89\xe8\xc1\xb6\x2e\xf0\xb2\xc7\xd6\x85\x2f\x5c\x8a\x02\x80\xd1\x10\x9f\x56\xb9\x9f\xe1\x00\x4a\x51\xe3\x12\xa7\xf0\x2a\xcc\xde\xc1\x19\x1b\x8f\xd7\x24\xfd\x25\x3c\x0c\x28\xea\xf5\x30\xcd\xc7\xb4\x2b\xcd\xe2\xf0\xf7\xc1\x58\x05\x83\x6f\xf7\x9d\xcb\xba\x69\x3e\x82\xa6\xe7\x07\x04\x33\x7c\x48\xd3\xb8\xcc\xe0\xa3\x59\x2c\xee\xae\xd3\x7c\x2c\xae\x45\x51\xac\xf9\xff\x98\xf1\x87\xee\x12\xc0\xc6\x9d\x87\x8e\x57\xd8\xf5\x91\x65\xe8\x22\x0d\x4b\xd3\x96\x0f\x1d\xcf\x50\xb0\x2a\x40\x36\x59\xdf\x72\x50\x30\xc1\xea\xf8\x1f\xab\x39\xbf\x4b\xda\x17\x49\x93\x24\xcd\xa3\xa4\x7b\x91\xb4\x49\xd2\x72\xfe\xe0\xf4\x9c\x81\x82\xc9\xa6\x82\x6c\x2b\xc8\x2e\xcb\x32\x99\x8c\x5c\x2b\x96\x14\xd6\x6e\xac\x25\x7a\x76\x32\x98\x93\x9b\xa0\xa7\x1e\x8c\x90\xb2\xdf\xb8\xf3\x57\x99\xd8\xe2\x1b\xcf\xb7\xf8\x54\x40\x72\xf6\x2a\xdc\x34\x14\x2c\x9b\x2d\x35\x90\xd8\x83\x9e\x72\x11\x65\xde\x91\xc3\x7c\x8a\x71\x54\x6e\x4f\xf4\xc4\x0c\x27\x65\x79\xfb\x66\xa8\x59\x0a\xc9\xa5\xaf\xc9\xbd\x0a\x8f\xd9\xf5\xc9\x7f\x17\xf2\x40\x84\xf7\xe2\xbc\x46\xa8\xa0\x7b\x10\x92\xc2\x2f\x08\x65\xff\x0f\x80\xf7\x09\x92\xc3\x7f\x20\x96\x18\xcc\xf0\xb7\x41\xfe\x06\x00\x00\xff\xff\x04\x3c\x8d\xe0\x1b\x03\x00\x00")

func arraysCdcBytes() ([]byte, error) {
	return bindataRead(
		_arraysCdc,
		"arrays.cdc",
	)
}

func arraysCdc() (*asset, error) {
	bytes, err := arraysCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "arrays.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x58, 0xdc, 0x75, 0x69, 0x42, 0x96, 0xfb, 0x1, 0xf4, 0xcd, 0xf3, 0x90, 0x7, 0xf5, 0x35, 0xd5, 0x96, 0x73, 0x99, 0x4e, 0xa0, 0xbf, 0xb8, 0x2a, 0xb4, 0x2a, 0x8b, 0x8, 0xc1, 0x93, 0xe1, 0x9d}}
	return a, nil
}

var _arraysCdcAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\xbb\x6e\xc3\x30\x0c\xdc\xfd\x15\x04\xe7\x4c\xcd\x96\x2d\x43\x87\x4e\x5d\xba\x15\x19\x98\x58\x42\x0d\x28\xb2\x21\x2b\x83\x53\xe8\xdf\x0b\xd9\x4d\x62\x4b\xf4\x23\x85\x82\x66\x30\x6c\x1d\x69\x92\x47\x1e\xf5\x9d\x01\x60\x2e\x64\xa1\x0b\x5b\x94\xba\xc6\x0d\x78\x08\x00\x29\xcf\x85\xb9\x1e\x01\x50\x9e\xf4\xc1\xfb\xf4\x30\x00\xac\xc8\xd0\x51\x58\x61\xfc\xaf\x9f\x57\x1c\x7a\x3e\xad\x9f\xa6\xa3\xc0\x0d\x20\xe1\x6a\x68\xb0\x4d\x25\x06\x21\x7f\x71\x32\x86\x1a\xc6\x00\x80\xa5\xf4\x91\xde\xb4\xc5\xc0\xe6\xb2\xb1\xd3\xed\x7b\x77\xcb\x8f\x46\xd8\x93\xd1\x1f\x71\x05\x58\x56\x9e\x2a\xa9\x28\xd1\x25\x50\xf7\x76\x5d\x30\x3c\x17\xd5\xbb\x5c\x3f\x73\xb7\xb6\xba\x09\x62\xb5\xc6\xba\x38\xfb\x78\xeb\xe5\x8d\x5c\xcd\x57\xbd\x4f\x3a\xe3\x24\x55\xdf\x39\x7e\xbe\xb0\xae\xa8\xd4\xbd\x7f\x99\x66\x11\xb4\x92\xa1\xee\x66\xc4\x79\x20\xa5\xb6\xca\x4b\x39\xb5\x3c\x65\x1d\x4e\x5a\xd1\x5e\xb4\x4b\x73\x49\x10\x79\xfc\x55\x0b\x31\x3e\x46\x63\xe0\x31\x3a\xe3\x61\x86\xde\xba\xfb\x69\xb1\x7e\x8e\x41\x63\x2c\x44\xfe\x49\x94\x93\x8c\xf8\x2a\x78\xf5\xc8\xd2\xbc\xd2\xe1\xeb\x01\xea\xe9\xf8\x3c\x52\x1e\xb3\x63\xbd\x63\x58\x0b\xee\x3d\xb9\x94\xcc\xa4\x6c\xc7\xfb\xc8\xa7\x5f\xce\x96\x93\xeb\x8e\xb9\x98\x66\x77\x26\x69\x63\x9f\x7b\x0b\x32\xff\xb8\xec\x27\x00\x00\xff\xff\x73\x87\x27\xbb\x28\x09\x00\x00")

func arraysCdcAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_arraysCdcAbiJson,
		"arrays.cdc.abi.json",
	)
}

func arraysCdcAbiJson() (*asset, error) {
	bytes, err := arraysCdcAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "arrays.cdc.abi.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x98, 0x1d, 0x97, 0xb6, 0xbe, 0x7c, 0x44, 0xa3, 0x2b, 0x86, 0x51, 0x70, 0xbf, 0xbe, 0x11, 0x4f, 0x77, 0xa3, 0xcb, 0x20, 0xf5, 0x8d, 0xc5, 0xee, 0xdc, 0x39, 0xc, 0xa9, 0x21, 0xf8, 0xed, 0xcf}}
	return a, nil
}

var _carCdc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x31\x4f\xc3\x30\x10\x85\x77\xff\x8a\xa3\x53\x5b\x55\x84\xb6\x5b\x50\x85\x44\xa5\x4a\x0c\x4c\x8c\x55\x06\x93\x5c\x21\xaa\xed\x58\xe7\x4b\x01\xa1\xfe\x77\x74\x76\x1a\x5a\x08\x22\x83\x7d\xb9\x67\xe7\xbe\x3c\x3d\xdf\x3e\x83\x41\x86\xe0\x11\x2b\x58\xc1\xc8\xea\xf7\xda\xb6\x76\xa4\x94\x48\x81\xa9\x2d\x19\xd6\x9a\xe0\x53\x01\x00\x48\xf3\xa0\x09\x6c\x53\xa1\xc9\xe1\x89\xa9\x76\x2f\x97\x8a\xde\xe3\xa0\xc0\x54\xdb\x5e\x88\x4a\xed\x6a\x1e\xef\x5a\x63\x9c\xb6\xfd\x9d\x49\x37\x49\x9e\x80\x66\x77\x1d\x47\x09\xda\xe8\x47\x5f\xef\x71\xa0\x2d\x63\xbe\xdb\x47\xa5\xb2\x69\xac\xb2\x0c\x1e\x5b\xc3\xb5\x37\x08\x65\xe3\xd2\x9f\x35\x14\x40\x13\x82\x6b\x18\x42\xeb\x7d\x43\x8c\x15\x7c\x20\x5f\x9d\x11\x7a\x4d\xda\x86\x1c\xb6\x09\xf0\x76\x59\xfc\xc9\x98\x8e\x6e\x6f\x8a\x41\xd4\x4e\x9d\x17\x83\xc4\x9d\xba\x28\x4e\xe0\x3d\x40\x32\x3b\x4d\x9f\x25\x83\x4f\x2f\xd1\xd4\x7f\x8c\x8b\xfb\x20\x90\x6c\x83\x2c\xb2\x75\x18\xb2\x4e\x33\x75\xbc\x08\xc4\xc6\x20\xf2\x59\x24\x24\x42\xa5\x26\x31\x69\xad\xa9\x38\x63\x2f\x35\xcd\x73\x09\xd0\x4c\x0e\x2c\x62\x79\x17\xeb\x65\xaa\x7f\x61\xcb\x77\x60\x05\x5b\xb9\x79\xf2\xe2\xa8\x14\x1e\xd0\x31\x6c\xd0\x55\x48\xf7\x71\x1d\xbf\xbd\x22\x21\x78\xa3\xcb\x3e\x3d\x33\x28\x9b\xc0\x39\x3c\x38\x9e\xa8\xaf\x00\x00\x00\xff\xff\x97\xed\x9f\x35\xdd\x02\x00\x00")

func carCdcBytes() ([]byte, error) {
	return bindataRead(
		_carCdc,
		"car.cdc",
	)
}

func carCdc() (*asset, error) {
	bytes, err := carCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "car.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1a, 0x65, 0x83, 0x19, 0xa2, 0x5c, 0x45, 0xb7, 0x6e, 0xe7, 0x81, 0x1b, 0x2f, 0x4f, 0xb3, 0xee, 0x31, 0xcb, 0xa1, 0x71, 0x6b, 0x91, 0x9a, 0x68, 0xbf, 0xad, 0xfd, 0x26, 0x3, 0x6b, 0xb9, 0x91}}
	return a, nil
}

var _carCdcAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\xc1\x4e\xc3\x30\x0c\x86\xef\x7d\x0a\xcb\xe7\x5d\x80\xdb\x8e\x20\x21\x71\xe6\x88\x38\x78\xad\x0b\x11\x69\x5a\xb9\xde\xd0\x40\x7d\x77\xd4\x8c\x0d\x27\x6b\x05\x63\x42\xe2\x32\x25\xff\x1f\xc7\x9f\xbd\xb8\xef\x05\x00\x56\x5c\xbb\xe0\xd4\xb5\xa1\xc7\x25\x8c\x12\x00\xde\x90\x1c\x36\x00\xd8\xab\xac\x4b\x35\x0a\x00\xd6\x8e\x7d\xd5\x27\x1a\x00\x36\x6d\xc5\x1e\x97\x80\xf7\x2a\x2e\x3c\xe1\x22\x31\xe9\x85\xe7\x3c\x15\xd7\x18\xef\x60\x0d\x5f\xa7\x30\x72\x92\x77\x6f\x2c\x63\xe2\x07\x13\x6f\xd7\x90\x20\xc5\xc8\x40\x4d\xcc\x5c\xaf\xbd\x8f\xeb\x45\x7e\x42\xb7\x1d\x4f\xe5\x8f\x0c\x66\xf7\x58\xe4\xab\x9d\xfb\xc9\x89\xb7\x9e\x59\x7f\xd7\xbb\x92\x24\xd7\x00\x90\x44\x68\x7b\x24\x03\x60\x5b\x4f\xa8\x36\x61\xfc\x17\x33\x7f\x98\x2d\x6c\xf8\xa3\x96\x97\x24\x17\xb3\xed\x3e\x13\x7f\xf1\x63\x84\xcb\x53\x10\xda\x6e\x1c\x07\xf2\x93\xee\x77\x88\x39\xe4\x19\xd0\x57\xff\x07\xfa\x94\x09\xe0\x50\xb1\x5c\xc7\x5f\x3b\x08\xbc\xe1\xa0\xc9\x1b\x4a\x5e\xff\xbe\xee\xce\x53\x99\xce\x27\x7a\x5a\xed\x3e\x2a\xaf\xcf\x2c\x99\x37\x37\xb7\xa6\xcd\x93\x69\xca\xb6\xd7\xe9\x9b\xee\x82\x9a\x6b\x0a\x5b\xea\xbe\xc4\xbe\x63\xae\x6c\x6d\x1b\x12\x47\x2b\x7f\x44\x32\x86\x0f\xc5\x50\x7c\x04\x00\x00\xff\xff\x01\x79\x63\xff\x6a\x05\x00\x00")

func carCdcAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_carCdcAbiJson,
		"car.cdc.abi.json",
	)
}

func carCdcAbiJson() (*asset, error) {
	bytes, err := carCdcAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "car.cdc.abi.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x53, 0x8c, 0x65, 0x7c, 0xf0, 0xfd, 0xce, 0x82, 0xe3, 0x97, 0xe8, 0x10, 0x8e, 0xfa, 0x61, 0xc7, 0x9b, 0x95, 0x0, 0x9f, 0x37, 0xf8, 0x51, 0x4a, 0x77, 0x62, 0x8b, 0x46, 0x5f, 0x99, 0x93, 0x5f}}
	return a, nil
}

var _dictionaresCdc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x31\x0b\xc2\x30\x10\x85\xf7\xfc\x8a\x37\x36\xd0\x41\xd7\x03\x29\x0e\x0e\x9d\x1d\x4b\x87\xb4\x26\x2a\x84\xa4\x24\x97\x41\x4a\xff\xbb\x34\x11\xa9\xd8\xe9\xdd\x7b\x77\xf7\xdd\x4d\x69\x80\xd5\x8c\xc1\x7b\x1b\x71\xc2\x2c\x00\xe0\x40\x30\xca\x46\x5d\x67\x77\x24\x70\x48\x5a\x2c\x42\xac\xe3\x26\x39\x18\x1f\x2e\x6a\x7c\x54\x37\xc2\x7c\xe5\xf0\x74\x77\x42\xd1\xa5\x86\x21\x54\x55\x71\x75\x11\x49\x67\xf7\x6a\xa4\x24\x74\x6b\xd1\x7f\xce\x04\xcd\x29\x38\x74\xfd\x16\xcd\x41\xb9\x38\xf9\xa8\xf7\xe0\xf2\x3f\xfa\x65\xcd\xcb\x96\x35\x2a\x6b\x33\xa6\x75\xbc\x3e\xd5\x3a\x96\xd4\x95\xc5\xa6\x97\x19\x97\x5b\xdf\x6c\x0f\xf7\x0e\x00\x00\xff\xff\x49\x2d\x55\xab\x24\x01\x00\x00")

func dictionaresCdcBytes() ([]byte, error) {
	return bindataRead(
		_dictionaresCdc,
		"dictionares.cdc",
	)
}

func dictionaresCdc() (*asset, error) {
	bytes, err := dictionaresCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dictionares.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xaa, 0x9, 0xb6, 0x73, 0xc, 0xb, 0x56, 0xc0, 0x73, 0xd1, 0x37, 0x67, 0x98, 0x9b, 0x88, 0xb7, 0x30, 0x48, 0x2c, 0x83, 0xfc, 0xd1, 0x9c, 0x93, 0xbd, 0xcf, 0x31, 0x41, 0xa1, 0x87, 0xfd, 0x8a}}
	return a, nil
}

var _dictionaresCdcAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x55\xb1\x4e\xc4\x30\x0c\xdd\xf3\x15\x96\xe7\xfb\x02\x36\x90\x18\x98\x61\x43\x0c\xbe\x36\x85\x88\x5c\x52\xa5\xe9\x49\x15\xca\xbf\xa3\x14\x28\x4d\xea\xb4\x05\xc4\x80\x38\xe9\x74\xaa\xed\xc6\xef\xbd\xd8\xef\x5e\x04\x00\xd6\xb2\x51\x46\x79\x65\x4d\x87\x17\x10\x43\x00\x78\xb4\x56\x7f\x3e\x02\xe0\x99\x9c\xa2\xa3\x96\xb3\xd8\x18\xaf\x55\x15\x5f\x25\x37\x64\x19\x00\x7c\x96\x43\x3c\x03\x6f\x8c\xc7\x43\x9a\x3a\x93\xee\xe5\x98\xbc\xb2\x56\xe3\x94\x0c\x62\xfe\x1b\xde\xde\xc2\xc6\xba\x6b\xaa\x9e\xe6\x78\x9a\xde\x8c\x8d\x93\xae\xd8\x92\xa3\x93\xf4\xd2\xc5\xa3\xef\x67\x2d\x33\x64\x86\x4e\x91\x09\xd6\x39\x2e\x3f\xb4\x39\xc5\x2d\x9a\x09\xd5\x5b\xef\x94\x79\xcc\x4e\xcd\x18\xbf\xd7\x64\x25\x41\x94\x9e\xc2\x61\x9b\x48\xb3\x97\x08\xab\xdb\x94\x2d\xea\x37\x55\x14\x09\x42\x89\x18\xc0\x03\x23\x87\x93\xbe\x77\xe6\x8e\x47\x39\x56\xd8\x76\x54\x5c\x47\x7a\x97\x66\x58\x1e\x1b\xf6\x2b\x28\x18\x28\x65\x08\x48\xce\x11\x33\xcf\xb6\xe1\x14\x5d\xc5\xc9\xa3\xe0\xa7\xdc\x3b\x32\x5d\x6b\x3b\xf9\x5f\xe7\xfc\x8b\xb7\xb4\xc7\x7a\x58\x9c\xab\x18\xb7\x6e\xa9\x22\xad\xff\xc0\x05\x2d\x3d\x37\x65\xce\xae\xdc\xaa\x39\xac\xf3\x4a\xaa\x62\x73\x36\xc7\x38\x01\xec\x72\x03\x28\xaf\x65\x52\xc2\xae\x68\x5a\x31\x5b\xd7\x92\x61\x7d\x7c\x72\x87\x59\x8f\x73\xd1\x65\xec\xf7\x5c\xeb\xa7\x7f\xc5\x8b\x59\x2b\xeb\x5d\xd6\x79\x97\xbe\xdf\xd1\x20\xdb\x46\x11\xbf\x41\xbc\x06\x00\x00\xff\xff\x2e\xd8\x18\x69\xc2\x08\x00\x00")

func dictionaresCdcAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_dictionaresCdcAbiJson,
		"dictionares.cdc.abi.json",
	)
}

func dictionaresCdcAbiJson() (*asset, error) {
	bytes, err := dictionaresCdcAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dictionares.cdc.abi.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x42, 0x3a, 0xd8, 0xe6, 0x4c, 0x78, 0x6a, 0xf0, 0xa, 0x41, 0x99, 0x8f, 0xef, 0xc6, 0xb0, 0x9d, 0x6a, 0xf7, 0xba, 0xdd, 0x14, 0xe2, 0x61, 0xfd, 0xf0, 0x22, 0xc0, 0x53, 0xe, 0xc3, 0x24, 0xec}}
	return a, nil
}

var _eventsCdc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xbd\x6e\xc2\x40\x10\x84\xfb\x7b\x8a\x29\x01\x91\x58\x69\x52\x9c\x82\x22\x85\x34\x54\x14\x90\x0a\x21\x64\x7c\x6b\xbc\x8a\xd9\x43\x77\xeb\x58\x11\xf2\xbb\x47\xfe\x23\x7f\x2e\xc7\xf3\xcd\x67\x6f\x32\x33\x26\x49\xb0\x5d\xbf\xae\xb1\xac\x42\x20\xd1\xf2\x13\xf4\x41\xa2\x11\x59\x2a\xe2\x15\x99\x17\x4d\x59\x10\x28\xfa\x2a\x64\x14\x91\x46\xe4\x4c\xa5\x9b\xe3\x58\x29\xb4\xe0\x08\x8e\x38\x12\xcb\x09\xb5\x0f\xef\xe4\xe0\xc5\x98\xa8\xa1\xca\x14\x1b\x22\x87\xab\x01\x80\x4b\x75\x44\x49\x0a\x76\x16\x2b\x51\xd3\x85\x2c\xac\x93\x21\x99\x0e\xc5\xf6\x89\x54\xe6\xf7\xec\xb0\x00\xbb\x2e\x6c\x4c\x63\x8c\x19\x3f\x03\x5b\x7f\x4e\xd5\xff\x59\x8e\x44\x2e\x5a\xec\x5a\xe9\xfe\xc7\xfe\xaf\xfc\x9f\xa5\x7b\x8b\x45\x4f\x7f\xbb\xba\x3b\x60\x59\xf8\xcb\x85\xdc\x44\x3b\x9f\x1d\xbc\xf3\xbe\x7c\x28\x29\xd7\xdb\xee\x88\x6c\x8b\xe0\xeb\x49\x5d\x50\x20\x8c\xd8\xd3\xdd\x08\x16\xbe\x3e\xe4\x69\x80\xe3\xa8\xa9\x64\x64\xf1\xb6\x12\x7d\x78\xbc\xe1\x2f\xbe\x3a\x15\x3a\x08\x29\x5a\x5c\x37\x1a\x58\x4e\x16\xbb\x7e\xe3\x79\xdf\x4c\xcd\x2c\x19\xfb\x1b\x5f\x77\x3f\x68\xd1\xf7\xe6\x50\x3e\xb7\x5c\x7b\x51\xf3\x15\x00\x00\xff\xff\xa1\x7a\xf9\x2a\xe4\x01\x00\x00")

func eventsCdcBytes() ([]byte, error) {
	return bindataRead(
		_eventsCdc,
		"events.cdc",
	)
}

func eventsCdc() (*asset, error) {
	bytes, err := eventsCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "events.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x41, 0x5e, 0x18, 0xa5, 0x1f, 0xf8, 0x37, 0x60, 0x82, 0x2, 0xd3, 0x70, 0x41, 0xe6, 0x3c, 0x91, 0x3a, 0x41, 0x6b, 0x59, 0xab, 0x74, 0x14, 0x65, 0xbd, 0x96, 0xc4, 0xaf, 0xe8, 0x9, 0x4a, 0x76}}
	return a, nil
}

var _eventsCdcAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x4a\x49\x4d\xcb\xcc\xcb\x2c\xc9\xcc\xcf\x2b\x56\xb2\x52\x00\x09\x29\x28\x28\x05\xe7\x97\xc3\x39\x0a\x0a\x4a\xa9\x65\xa9\x79\x25\x4a\x56\x0a\xd1\x50\x01\x05\xb8\x14\x58\x3a\x2f\x31\x37\x55\xc9\x4a\x41\xa9\x38\x35\x35\x45\x49\x07\x59\xa6\xa4\xb2\x00\x2c\x13\x5c\x52\x94\x99\x97\xae\x04\x97\xaa\xd5\xc1\x6f\x52\x49\x66\x6e\x6a\x31\x76\xa3\x3c\xf3\x4a\x90\xcc\x81\xb2\x62\xb9\x60\xbc\x5a\xae\x5a\x2e\x40\x00\x00\x00\xff\xff\xec\xdd\xac\x8f\xd8\x00\x00\x00")

func eventsCdcAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_eventsCdcAbiJson,
		"events.cdc.abi.json",
	)
}

func eventsCdcAbiJson() (*asset, error) {
	bytes, err := eventsCdcAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "events.cdc.abi.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x39, 0xe9, 0xc5, 0xc8, 0x6e, 0x92, 0x72, 0x67, 0x6f, 0xa9, 0x36, 0xf2, 0x45, 0x11, 0xe2, 0xf1, 0xa2, 0x32, 0x60, 0x42, 0x49, 0x5, 0xfc, 0x53, 0xad, 0x33, 0x2d, 0x74, 0xe4, 0x17, 0x88, 0x19}}
	return a, nil
}

var _functionsCdc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x31\x4f\x03\x31\x0c\x85\xf7\xfb\x15\x6f\x74\x50\x25\xf6\xa2\x8a\xb9\x1b\x13\x7b\xee\x62\xa3\x48\xe0\x80\xeb\xa0\x93\xd0\xfd\x77\x94\xe3\x7a\x4a\x69\x06\x27\x7e\x71\xbe\xf7\xf2\x59\x47\x48\x55\x5c\xbe\xcc\x69\xc6\x7c\xc4\x59\x3d\xac\xf5\x67\x00\x00\x63\xaf\xa6\x98\x1f\xe6\x61\x19\x86\xeb\x78\xca\xdf\x39\xb1\xbd\x58\x49\x75\x62\xa3\xd6\x5f\x8a\x21\x5d\xdf\x13\xad\x7b\x2b\xb8\x01\x49\x55\x8a\xf6\x56\x3f\x58\x7d\xb7\x3b\xab\x6f\x53\xbd\xe5\x63\x7a\x5a\xb5\xe5\xc6\x99\xa7\x62\xd1\x8b\x91\x54\x9d\x3c\x17\x85\x34\xbb\x70\x7c\x2d\x39\x85\x03\x46\x96\x62\xdc\x49\xcf\x07\x44\x71\xb6\x5e\x0a\x5d\x73\x97\x0f\x14\xba\x34\x59\xf0\xce\x8e\x11\xa7\x0d\xdd\xdd\xb5\x35\x52\xd8\xfb\x65\x3f\x49\xa7\x6e\x84\x88\xd3\x5f\x90\x7f\x80\x78\x07\x68\x1f\xfe\x0d\x00\x00\xff\xff\x73\x4e\xca\x61\x99\x01\x00\x00")

func functionsCdcBytes() ([]byte, error) {
	return bindataRead(
		_functionsCdc,
		"functions.cdc",
	)
}

func functionsCdc() (*asset, error) {
	bytes, err := functionsCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "functions.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x35, 0x7, 0xdc, 0xa3, 0x25, 0xd5, 0x27, 0xcf, 0x2d, 0x30, 0xc3, 0x50, 0x31, 0xe5, 0x8f, 0xf3, 0x35, 0x43, 0x83, 0x4a, 0x7a, 0xef, 0x42, 0x24, 0x6b, 0xb9, 0x88, 0xb2, 0xaa, 0x14, 0x1c, 0x71}}
	return a, nil
}

var _functionsCdcAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x52\x41\x4e\xc4\x30\x0c\xbc\xf7\x15\x96\xcf\xfb\x82\xfd\x01\x37\x0e\xdc\x10\x87\x6c\xe3\x48\x91\xba\x49\x71\x5d\x04\x5a\xe5\xef\x28\xad\x14\x92\x92\x14\x10\x85\x43\xa5\xda\x33\xce\x8c\x47\xbe\x75\x00\xa8\xc9\x58\x67\xc5\x7a\x37\xe1\x19\x62\x0b\x00\xa7\x67\x96\x54\x01\xa0\x99\x5d\x1f\x29\x59\x0f\x00\x47\xc5\xea\x4a\x42\x1c\x27\x1f\x53\x1f\x32\xce\xc2\x73\xea\x4a\x78\x06\x7c\xc5\x53\x09\x0c\xea\x42\x43\x15\x91\xb7\x71\x19\xb9\x73\x82\x19\x12\xd2\xff\xd3\xc7\x00\x32\xc9\xcc\xee\xe1\xf3\xc8\x4a\x0f\x2b\x15\xb5\x7d\xb1\x9a\xf8\x9e\xbd\x9e\x7b\xe2\xe3\xf7\xd3\xcd\xfd\xa2\xf4\xe4\xf9\xb8\x2d\x73\x07\x75\xf7\xfb\x1b\x2c\xe8\x46\xb6\x90\xdb\x0f\xb6\x74\x19\xea\x71\x53\xef\x59\x89\xff\x83\xa0\x4d\x33\xe8\xf4\x7e\x3d\xe9\xdb\x36\x82\xdc\x4f\x28\xc0\xbc\x0a\xa7\xaf\x3d\x5d\xc8\x78\xa6\xef\xea\xfa\x31\xaa\xaa\xa1\x82\xed\xba\x2a\x7d\xfd\xdc\xa5\x32\x42\xad\x33\xfc\x1f\x93\xbf\xb8\xed\xe6\xc5\x75\xf1\x0b\xdd\x7b\x00\x00\x00\xff\xff\x43\xef\xaf\x0b\xd0\x04\x00\x00")

func functionsCdcAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_functionsCdcAbiJson,
		"functions.cdc.abi.json",
	)
}

func functionsCdcAbiJson() (*asset, error) {
	bytes, err := functionsCdcAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "functions.cdc.abi.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9, 0x2f, 0x7c, 0x7b, 0xa0, 0xc4, 0x65, 0xf3, 0xe0, 0x8a, 0x7f, 0x8, 0xa2, 0x2d, 0x98, 0x59, 0xfe, 0xb, 0xb6, 0x44, 0xf2, 0xe8, 0x54, 0xf4, 0x91, 0x38, 0x97, 0xef, 0x46, 0x4f, 0x61, 0xe9}}
	return a, nil
}

var _resourcesCdc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x6b\x84\x30\x14\x84\xef\xf9\x15\x73\xb4\xa0\x3d\x17\x31\x14\xda\x53\x4f\x3d\xf5\x54\x3c\x98\x34\x5a\x41\x92\xe5\x25\x39\x2c\x8b\xff\x7d\x31\x89\x62\x74\x37\x97\xc0\xf7\xde\xbc\x99\xb9\x78\x01\xeb\xc8\x4b\x87\x4f\x33\x19\x4f\xb8\x31\x00\x58\xf8\xa4\x1c\xa8\xc6\xcf\x97\x76\x6f\x19\x1c\x1e\x41\xb1\xc2\x40\x47\x3d\xba\x62\x15\x97\x9b\xa2\xdc\xd6\x5e\x92\xd1\xf2\xac\x9a\xfa\x57\x02\x07\xe5\x68\x00\xc7\x90\x23\x01\x0e\x11\xd0\xcc\x66\xc6\x16\x73\x52\xd6\x78\x92\x0a\x1f\x9d\xee\x74\x77\x28\x20\x43\xab\x3a\xb5\xdb\x85\xcb\x07\xa7\x38\x71\x0c\x9e\x0e\x3c\xb3\xf4\x5a\xfe\x7f\xf7\xd1\xd9\x1e\xac\x45\xa4\x35\x9a\xea\x37\x6e\xbc\xb7\xbb\x00\xe7\x71\x7b\x4a\x91\x76\xd0\x54\xeb\xb5\x14\x24\x7c\x7f\xca\x3a\x32\xd7\x62\x2f\x4b\x2c\x93\x6f\xe1\xef\x01\x00\x00\xff\xff\x68\xfd\x5b\xa2\xee\x01\x00\x00")

func resourcesCdcBytes() ([]byte, error) {
	return bindataRead(
		_resourcesCdc,
		"resources.cdc",
	)
}

func resourcesCdc() (*asset, error) {
	bytes, err := resourcesCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5b, 0x7a, 0xb0, 0x1f, 0x9c, 0x8, 0x85, 0x42, 0x77, 0x80, 0x93, 0x3d, 0xfa, 0xb, 0xbc, 0x20, 0x34, 0x6a, 0x11, 0xeb, 0x95, 0xd8, 0x4d, 0xc1, 0xe9, 0xc3, 0xda, 0x4, 0x5, 0x52, 0x6b, 0xfd}}
	return a, nil
}

var _resourcesCdcAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\xb1\x4e\xc4\x30\x0c\x86\xf7\x3c\x85\xe5\x99\x07\x40\x8c\x30\x31\x31\x31\x21\x06\xb7\x97\x1e\x91\x4a\x72\x72\xd2\xe1\x40\x7d\x77\x94\x5e\xdb\x4b\xd3\xe4\x68\x10\x1d\x4e\x6a\x7e\xc7\xf6\x9f\xcf\xbe\x6f\x01\x80\x07\xd9\x28\xad\x9c\x32\xda\xe2\x03\x78\x09\x00\x9f\x4c\x6b\x3a\x9e\xcf\x00\x68\x1d\x77\xb5\x0b\x14\x00\x6c\x94\x6c\x0f\x76\xa1\x01\xa0\x4f\xc3\xd7\x67\xed\xee\xf1\x2e\xd4\x8f\x19\xbd\xba\xea\xb3\xdc\x5f\x6f\xe0\xe0\x8e\x5a\xf5\x25\xd9\xf7\x7a\x0b\x72\xc3\x6f\x58\xb8\x18\x32\x35\x7d\x4a\x5f\x9c\x17\x0d\x87\x90\x3b\x9f\x64\xa2\x6f\xd4\xfb\x66\xd5\xe3\x2e\x55\xab\xb2\xaa\xc1\xe9\x5d\xc4\x5f\x97\xe8\xd8\x19\x1f\x49\x93\xa6\x70\xa6\x2c\xad\xe9\xb8\x96\x5b\xa6\x5a\xc7\x1b\x31\xea\xf3\x5e\x4c\x3b\x23\x52\xee\xfe\x77\x9e\xa3\x97\x1c\xa8\x38\xef\x17\x97\x31\xc7\x42\xaa\x9d\xae\x3f\x5e\x9a\x0b\x5c\xfb\x57\xba\xd5\x2a\x7d\x0c\x10\x33\x9d\x13\x6f\x42\xd3\xa4\x5f\x6a\x4e\xfe\xaf\x4c\x6d\x32\xba\x74\x35\x6d\xc4\xea\x5a\xbf\x99\xce\x5e\x13\x9e\x78\x14\x8c\x38\x47\x0a\x6e\xd0\x82\xad\x44\xd6\x4c\x4a\x28\x65\x77\x48\xf8\x5f\x2f\x7e\x02\x00\x00\xff\xff\x87\x20\x62\xe9\x88\x05\x00\x00")

func resourcesCdcAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_resourcesCdcAbiJson,
		"resources.cdc.abi.json",
	)
}

func resourcesCdcAbiJson() (*asset, error) {
	bytes, err := resourcesCdcAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources.cdc.abi.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9a, 0x4, 0xd0, 0x4b, 0xd, 0xa2, 0xf9, 0x41, 0x82, 0xa1, 0x4a, 0xd, 0xf7, 0xe7, 0xc7, 0xf2, 0xac, 0x56, 0x89, 0xa, 0xf8, 0x10, 0x77, 0x87, 0xff, 0xb2, 0xa9, 0x6a, 0x5, 0xd2, 0x57, 0x85}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"arrays.cdc": arraysCdc,

	"arrays.cdc.abi.json": arraysCdcAbiJson,

	"car.cdc": carCdc,

	"car.cdc.abi.json": carCdcAbiJson,

	"dictionares.cdc": dictionaresCdc,

	"dictionares.cdc.abi.json": dictionaresCdcAbiJson,

	"events.cdc": eventsCdc,

	"events.cdc.abi.json": eventsCdcAbiJson,

	"functions.cdc": functionsCdc,

	"functions.cdc.abi.json": functionsCdcAbiJson,

	"resources.cdc": resourcesCdc,

	"resources.cdc.abi.json": resourcesCdcAbiJson,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"arrays.cdc":               &bintree{arraysCdc, map[string]*bintree{}},
	"arrays.cdc.abi.json":      &bintree{arraysCdcAbiJson, map[string]*bintree{}},
	"car.cdc":                  &bintree{carCdc, map[string]*bintree{}},
	"car.cdc.abi.json":         &bintree{carCdcAbiJson, map[string]*bintree{}},
	"dictionares.cdc":          &bintree{dictionaresCdc, map[string]*bintree{}},
	"dictionares.cdc.abi.json": &bintree{dictionaresCdcAbiJson, map[string]*bintree{}},
	"events.cdc":               &bintree{eventsCdc, map[string]*bintree{}},
	"events.cdc.abi.json":      &bintree{eventsCdcAbiJson, map[string]*bintree{}},
	"functions.cdc":            &bintree{functionsCdc, map[string]*bintree{}},
	"functions.cdc.abi.json":   &bintree{functionsCdcAbiJson, map[string]*bintree{}},
	"resources.cdc":            &bintree{resourcesCdc, map[string]*bintree{}},
	"resources.cdc.abi.json":   &bintree{resourcesCdcAbiJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
