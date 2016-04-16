// Code generated by go-bindata.
// sources:
// server/static/css/index_style.css
// server/static/css/overview_style.css
// server/static/export_list.json
// server/static/index.html
// server/static/js/d3.min.js
// server/static/js/d3utils.js
// server/static/js/jquery-1.8.3.min.js
// server/static/js/overview.js
// server/static/js/render_settings.js
// server/static/js/render_settings_example.js
// server/static/js/render_utils.js
// server/static/js/spin.min.js
// server/static/js/ulogme.js
// server/static/js/ulogme_common.js
// server/static/js/underscore.min.js
// server/static/overview.html
// server/static/rules.json
// DO NOT EDIT!

package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// serverStaticCssIndex_styleCss reads file data from disk. It returns an error on failure.
func serverStaticCssIndex_styleCss() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/css/index_style.css"
	name := "server/static/css/index_style.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticCssOverview_styleCss reads file data from disk. It returns an error on failure.
func serverStaticCssOverview_styleCss() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/css/overview_style.css"
	name := "server/static/css/overview_style.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticExport_listJson reads file data from disk. It returns an error on failure.
func serverStaticExport_listJson() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/export_list.json"
	name := "server/static/export_list.json"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticIndexHtml reads file data from disk. It returns an error on failure.
func serverStaticIndexHtml() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/index.html"
	name := "server/static/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsD3MinJs reads file data from disk. It returns an error on failure.
func serverStaticJsD3MinJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/d3.min.js"
	name := "server/static/js/d3.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsD3utilsJs reads file data from disk. It returns an error on failure.
func serverStaticJsD3utilsJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/d3utils.js"
	name := "server/static/js/d3utils.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsJquery183MinJs reads file data from disk. It returns an error on failure.
func serverStaticJsJquery183MinJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/jquery-1.8.3.min.js"
	name := "server/static/js/jquery-1.8.3.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsOverviewJs reads file data from disk. It returns an error on failure.
func serverStaticJsOverviewJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/overview.js"
	name := "server/static/js/overview.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsRender_settingsJs reads file data from disk. It returns an error on failure.
func serverStaticJsRender_settingsJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/render_settings.js"
	name := "server/static/js/render_settings.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsRender_settings_exampleJs reads file data from disk. It returns an error on failure.
func serverStaticJsRender_settings_exampleJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/render_settings_example.js"
	name := "server/static/js/render_settings_example.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsRender_utilsJs reads file data from disk. It returns an error on failure.
func serverStaticJsRender_utilsJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/render_utils.js"
	name := "server/static/js/render_utils.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsSpinMinJs reads file data from disk. It returns an error on failure.
func serverStaticJsSpinMinJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/spin.min.js"
	name := "server/static/js/spin.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsUlogmeJs reads file data from disk. It returns an error on failure.
func serverStaticJsUlogmeJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/ulogme.js"
	name := "server/static/js/ulogme.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsUlogme_commonJs reads file data from disk. It returns an error on failure.
func serverStaticJsUlogme_commonJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/ulogme_common.js"
	name := "server/static/js/ulogme_common.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticJsUnderscoreMinJs reads file data from disk. It returns an error on failure.
func serverStaticJsUnderscoreMinJs() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/js/underscore.min.js"
	name := "server/static/js/underscore.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticOverviewHtml reads file data from disk. It returns an error on failure.
func serverStaticOverviewHtml() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/overview.html"
	name := "server/static/overview.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// serverStaticRulesJson reads file data from disk. It returns an error on failure.
func serverStaticRulesJson() (*asset, error) {
	path := "/home/hxr/work/go/src/github.com/erasche/gologme/server/static/rules.json"
	name := "server/static/rules.json"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
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
	"server/static/css/index_style.css":           serverStaticCssIndex_styleCss,
	"server/static/css/overview_style.css":        serverStaticCssOverview_styleCss,
	"server/static/export_list.json":              serverStaticExport_listJson,
	"server/static/index.html":                    serverStaticIndexHtml,
	"server/static/js/d3.min.js":                  serverStaticJsD3MinJs,
	"server/static/js/d3utils.js":                 serverStaticJsD3utilsJs,
	"server/static/js/jquery-1.8.3.min.js":        serverStaticJsJquery183MinJs,
	"server/static/js/overview.js":                serverStaticJsOverviewJs,
	"server/static/js/render_settings.js":         serverStaticJsRender_settingsJs,
	"server/static/js/render_settings_example.js": serverStaticJsRender_settings_exampleJs,
	"server/static/js/render_utils.js":            serverStaticJsRender_utilsJs,
	"server/static/js/spin.min.js":                serverStaticJsSpinMinJs,
	"server/static/js/ulogme.js":                  serverStaticJsUlogmeJs,
	"server/static/js/ulogme_common.js":           serverStaticJsUlogme_commonJs,
	"server/static/js/underscore.min.js":          serverStaticJsUnderscoreMinJs,
	"server/static/overview.html":                 serverStaticOverviewHtml,
	"server/static/rules.json":                    serverStaticRulesJson,
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
	"server": &bintree{nil, map[string]*bintree{
		"static": &bintree{nil, map[string]*bintree{
			"css": &bintree{nil, map[string]*bintree{
				"index_style.css":    &bintree{serverStaticCssIndex_styleCss, map[string]*bintree{}},
				"overview_style.css": &bintree{serverStaticCssOverview_styleCss, map[string]*bintree{}},
			}},
			"export_list.json": &bintree{serverStaticExport_listJson, map[string]*bintree{}},
			"index.html":       &bintree{serverStaticIndexHtml, map[string]*bintree{}},
			"js": &bintree{nil, map[string]*bintree{
				"d3.min.js":                  &bintree{serverStaticJsD3MinJs, map[string]*bintree{}},
				"d3utils.js":                 &bintree{serverStaticJsD3utilsJs, map[string]*bintree{}},
				"jquery-1.8.3.min.js":        &bintree{serverStaticJsJquery183MinJs, map[string]*bintree{}},
				"overview.js":                &bintree{serverStaticJsOverviewJs, map[string]*bintree{}},
				"render_settings.js":         &bintree{serverStaticJsRender_settingsJs, map[string]*bintree{}},
				"render_settings_example.js": &bintree{serverStaticJsRender_settings_exampleJs, map[string]*bintree{}},
				"render_utils.js":            &bintree{serverStaticJsRender_utilsJs, map[string]*bintree{}},
				"spin.min.js":                &bintree{serverStaticJsSpinMinJs, map[string]*bintree{}},
				"ulogme.js":                  &bintree{serverStaticJsUlogmeJs, map[string]*bintree{}},
				"ulogme_common.js":           &bintree{serverStaticJsUlogme_commonJs, map[string]*bintree{}},
				"underscore.min.js":          &bintree{serverStaticJsUnderscoreMinJs, map[string]*bintree{}},
			}},
			"overview.html": &bintree{serverStaticOverviewHtml, map[string]*bintree{}},
			"rules.json":    &bintree{serverStaticRulesJson, map[string]*bintree{}},
		}},
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
