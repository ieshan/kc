package helper

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

var homeDir, _ = os.UserHomeDir()
var kubeConfigPath = filepath.Join(homeDir, ".kube", "config")

func GetKubeConfig() (*KubectlConfig, error) {
	file, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		return nil, err
	}
	cfg := KubectlConfig{}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}


type KubectlConfig struct {
	Kind           string                    `json:"kind" yaml:"kind"`
	ApiVersion     string                    `json:"apiVersion" yaml:"apiVersion"`
	CurrentContext string                    `json:"current-context" yaml:"current-context"`
	//Clusters       []*KubectlClusterWithName `json:"clusters" yaml:"clusters"`
	Contexts       []*KubectlContextWithName `json:"contexts" yaml:"contexts"`
	//Users          []*KubectlUserWithName    `json:"users" yaml:"users"`
}

type KubectlClusterWithName struct {
	Name    string         `json:"name" yaml:"name"`
	Cluster KubectlCluster `json:"cluster" yaml:"cluster"`
}

type KubectlCluster struct {
	Server                   string `json:"server,omitempty" yaml:"server,omitempty"`
	CertificateAuthorityData []byte `json:"certificate-authority-data,omitempty" yaml:"certificate-authority-data,omitempty"`
}

type KubectlContextWithName struct {
	Name    string         `json:"name" yaml:"name"`
	Context KubectlContext `json:"context" yaml:"context"`
}

type KubectlContext struct {
	Cluster string `json:"cluster" yaml:"cluster"`
	User    string `json:"user" yaml:"user"`
}

type KubectlUserWithName struct {
	Name string      `json:"name" yaml:"name"`
	User KubectlUser `json:"user" yaml:"user"`
}

type KubectlUser struct {
	ClientCertificateData []byte `json:"client-certificate-data,omitempty" yaml:"client-certificate-data,omitempty"`
	ClientKeyData         []byte `json:"client-key-data,omitempty" yaml:"client-key-data,omitempty"`
	Password              string `json:"password,omitempty" yaml:"password,omitempty"`
	Username              string `json:"username,omitempty" yaml:"username,omitempty"`
	Token                 string `json:"token,omitempty" yaml:"token,omitempty"`
}
