// Code generated for package widgets by go-bindata DO NOT EDIT. (@generated)
// sources:
// header.html
// sidebar.html
// sidebar_level.html
// timeseries.html
// timeseries_init.html
package widgets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
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

var _headerHtml = []byte(`{{ define "header" }}
<header>
    <nav>
        <a id="brand" href="/">
            <span class="logo">
				<svg viewBox="0 0 300 300">
				<circle cx="150" cy="150" r="146" stroke-width="2" />
				<polygon points="65,240 225,240 125,270"/>
				<polygon points="65,230 125,220 125,110"/>
				<polygon points="135,220 225,230 135,30"/>
				</svg>
			</span>
            <span class="name">Istio Eng Dashboard</span>
        </a>

        <div id="hamburger">
			<svg class="icon"><use xlink:href="/icons/icons.svg#hamburger"/></svg>
		</div>

        <div id="header-links">
			<button id="authn-button" title="Sign in to GitHub" aria-label="Sign in to GitHub">
				<svg class="icon"><use xlink:href="/icons/icons.svg#github"/></svg>
			</button>

            <div class="menu">
                <button id="gearDropdownButton" class="menu-trigger" title="Options and Settings"
                        aria-label="Options and Settings" aria-controls="gearDropdownContent">
					<svg class="icon"><use xlink:href="/icons/icons.svg#gear"/></svg>
                </button>

                <div id="gearDropdownContent" class="menu-content" aria-labelledby="gearDropdownButton" role="menu">
                    <a tabindex="-1" role="menuitem" class="active" id="light-theme-item">Light Theme</a>
                    <a tabindex="-1" role="menuitem" id="dark-theme-item">Dark Theme</a>
                </div>
            </div>
        </div>
    </nav>
</header>
{{ end }}
`)

func headerHtmlBytes() ([]byte, error) {
	return _headerHtml, nil
}

func headerHtml() (*asset, error) {
	bytes, err := headerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "header.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sidebarHtml = []byte(`{{ define "sidebar" }}
<nav id="sidebar" aria-label="Section Navigation">
    <div class="directory">
        <div class="card">
            <div id="header0" class="header">
	            Istio Eng Dashboard
        	</div>

			{{ $topics := .Entries }}
            {{ $selectedEntry := .SelectedEntry }}
            <div class="body default" aria-labelledby="header0">
				{{ template "sidebar_level" (dict "topics" $topics "collapse" false "top" true "labelledby" "header0" "selectedEntry" $selectedEntry ) }}
			</div>
        </div>
    </div>
</nav>
{{ end }}
`)

func sidebarHtmlBytes() ([]byte, error) {
	return _sidebarHtml, nil
}

func sidebarHtml() (*asset, error) {
	bytes, err := sidebarHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sidebar.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sidebar_levelHtml = []byte(`{{ define "sidebar_level" }}
{{ $topics := .topics }}
{{ $collapse := .collapse }}
{{ $top := .top }}
{{ $labelledby := .labelledby }}
{{ $selectedEntry := .selectedEntry }}

{{ $leafSection := true }}
{{ range $topics }}
    {{ if gt (len .Entries) 0 }}
        {{ $leafSection = false }}
    {{ end }}
{{ end }}

<ul role="{{ if $top }}tree{{ else }}group{{ end }}" aria-expanded="{{ if $collapse }}false{{ else }}true{{ end }}"{{ if $leafSection }} class="leaf-section"{{ end }} {{ if $labelledby}}aria-labelledby="{{ $labelledby }}"{{ end }}>
    {{ range $topics }}
        {{ if gt (len .Entries) 0 }}
            <li role="treeitem" aria-label="{{ .Title }}">
                {{ $collapse := not (.IsAncestor $selectedEntry) }}

                <button{{ if not $collapse }} class="show"{{ end }} aria-hidden="true"></button><a {{ if .IsSame $selectedEntry }}class="current"{{ end }} title="{{ .Description }}" href="{{ .URL }}">{{ .Title}}</a>

                {{ template "sidebar_level" (dict "topics" .Entries "collapse" $collapse "top" false "labelledby" "" "selectedEntry" $selectedEntry ) }}
            </li>
        {{ else }}
            <li role="none">
                {{ if .IsSame $selectedEntry }}
                    <span role="treeitem" class="current" title="{{ .Description }}">{{ .Title }}</span>
                {{ else }}
                    <a role="treeitem" title="{{ .Description }}" href="{{ .URL }}">{{ .Title }}</a>
                {{ end }}
            </li>
        {{ end }}
    {{ end }}
</ul>
{{ end }}
`)

func sidebar_levelHtmlBytes() ([]byte, error) {
	return _sidebar_levelHtml, nil
}

func sidebar_levelHtml() (*asset, error) {
	bytes, err := sidebar_levelHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sidebar_level.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _timeseriesHtml = []byte(`{{ define "timeseries" }}
<div id={{ .Target }}></div>
<script>
  var data = MG.convert.date(JSON.parse({{ .TimeSeries }}), 'date');
	MG.data_graphic({
	title: {{ .Name }},
	description: "This graphic shows a time-series of downloads.",
	data: data,
	area: false,
  interpolate: d3.curveLinear,
	show_tooltips: false,
	width: 600,
	height: 250,
	target: '#' + {{ .Target }},
	x_accessor: 'date',
	y_accessor: 'value',
  });
</script>
{{ end }}
`)

func timeseriesHtmlBytes() ([]byte, error) {
	return _timeseriesHtml, nil
}

func timeseriesHtml() (*asset, error) {
	bytes, err := timeseriesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "timeseries.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _timeseries_initHtml = []byte(`{{ define "timeseries_init" }}
<link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Open+Sans:400,300,700'>
<link rel='stylesheet' href='https://fonts.googleapis.com/css?family=PT+Serif:400,700,400italic'>
<link rel='stylesheet' href='https://netdna.bootstrapcdn.com/font-awesome/4.2.0/css/font-awesome.css'>
<link rel='stylesheet' href='https://cdnjs.cloudflare.com/ajax/libs/metrics-graphics/2.15.6/metricsgraphics.min.css'>
<script src='https://d3js.org/d3.v4.min.js'></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/metrics-graphics/2.15.6/metricsgraphics.min.js"></script>
{{ end }}
`)

func timeseries_initHtmlBytes() ([]byte, error) {
	return _timeseries_initHtml, nil
}

func timeseries_initHtml() (*asset, error) {
	bytes, err := timeseries_initHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "timeseries_init.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"header.html":          headerHtml,
	"sidebar.html":         sidebarHtml,
	"sidebar_level.html":   sidebar_levelHtml,
	"timeseries.html":      timeseriesHtml,
	"timeseries_init.html": timeseries_initHtml,
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
	"header.html":          &bintree{headerHtml, map[string]*bintree{}},
	"sidebar.html":         &bintree{sidebarHtml, map[string]*bintree{}},
	"sidebar_level.html":   &bintree{sidebar_levelHtml, map[string]*bintree{}},
	"timeseries.html":      &bintree{timeseriesHtml, map[string]*bintree{}},
	"timeseries_init.html": &bintree{timeseries_initHtml, map[string]*bintree{}},
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
