package config

import (
	"bytes"
	"github.com/coschain/contentos-go/node"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

var configTemplate *template.Template

const DefaultConfigTemplate = `# This is a TOML config file. 
# For more information, see https://github.com/toml-lang/toml

[node]

DataDir = ""
HTTPHost = "{{ .HTTPHost }}"
HTTPPort = {{ .HTTPPort }}

[p2p]

MaxPeers = {{ .P2P.MaxPeers }}
StaticNodes = []
TrustedNodes = []
ListenAddr = "{{ .P2P.ListenAddr }}"
`

func WriteNodeConfigFile(configDirPath string, configName string, config node.Config, mode os.FileMode) error {
	var buffer bytes.Buffer
	var err error

	if configTemplate, err = template.New("configFileTemplate").Parse(DefaultConfigTemplate); err != nil {
		return err
	}

	if err = configTemplate.Execute(&buffer, config); err != nil {
		return err
	}
	configPath := filepath.Join(configDirPath, configName)
	return ioutil.WriteFile(configPath, buffer.Bytes(), mode)
}