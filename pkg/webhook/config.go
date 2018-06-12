/*
Copyright 2018 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhook

import (
	"crypto/tls"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/golang/glog"
)

// Config contains the server (the webhook) cert and key.
type Config struct {
	CertFile string
	KeyFile  string
}

// Get a clientset with in-cluster config.
func GetClient() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		glog.Fatalf("Cannot setup in-cluster configuration: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatalf("unbale to config client: %v", err)
	}
	return clientset
}

func ConfigTLS(config Config, clientset *kubernetes.Clientset) *tls.Config {

	if len(config.CertFile) == 0 {
		glog.Fatalf("Empty certificate file %q", config.CertFile)
	}

	if _, err := os.Stat(config.CertFile); err != nil {
		glog.Fatalf("Cannot stat file %q: %v", config.CertFile, err)
	}

	if len(config.KeyFile) == 0 {
		glog.Fatalf("Empty private key file %q", config.KeyFile)
	}

	if _, err := os.Stat(config.KeyFile); err != nil {
		glog.Fatalf("Cannot stat file %q: %v", config.KeyFile, err)
	}

	sCert, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
	if err != nil {
		glog.Fatal(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{sCert},
		// TODO: uses mutual tls after we agree on what cert the apiserver should use.
		// ClientAuth:   tls.RequireAndVerifyClientCert,
	}
}
