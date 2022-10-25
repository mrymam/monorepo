package setting

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v2"
)

var (
	//go:embed files/setting.*.yaml
	files       embed.FS
	settingFile = map[string]string{
		"dev": "files/setting.dev.yaml",
	}
	settingDefault = "dev"
)

var setting Setting

func init() {
	buf, err := getSettingFile(os.Getenv("ENV"))
	if err != nil {
		panic(err)
	}

	buf = []byte(os.ExpandEnv(string(buf)))

	if err := yaml.Unmarshal(buf, &setting); err != nil {
		panic(err)
	}
}

func Get() Setting {
	return setting
}

func getSettingFile(env string) ([]byte, error) {
	filename, ok := settingFile[env]
	if !ok {
		filename, ok = settingFile[settingDefault]
		if !ok {
			return nil, fmt.Errorf("setting file is not found.")
		}
	}

	tmpl, err := files.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	t, err := template.New(filepath.Base(filename)).Funcs(
		sprig.TxtFuncMap(),
	).Parse(string(tmpl))

	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err := t.Execute(&b, nil); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
