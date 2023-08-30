/*
Copyright 2021 The Kubernetes Authors.

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

package external

import (
	"fmt"

	"github.com/spf13/pflag"

	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/config/v3"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v3/pkg/model/resource"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
	"sigs.k8s.io/yaml"
)

var _ plugin.CreateWebhookSubcommand = &createWebhookSubcommand{}

type createWebhookSubcommand struct {
	Path   string
	Args   []string
	config config.Config
}

func (p *createWebhookSubcommand) InjectResource(*resource.Resource) error {
	// Do nothing since resource flags are passed to the external plugin directly.
	return nil
}

func (p *createWebhookSubcommand) UpdateMetadata(_ plugin.CLIMetadata, subcmdMeta *plugin.SubcommandMetadata) {
	setExternalPluginMetadata("webhook", p.Path, subcmdMeta)
}

func (p *createWebhookSubcommand) BindFlags(fs *pflag.FlagSet) {
	bindExternalPluginFlags(fs, "webhook", p.Path, p.Args)
}

func (p *createWebhookSubcommand) Scaffold(fs machinery.Filesystem) error {
	configBytes, err := yaml.Marshal(p.config)
	if err != nil {
		return fmt.Errorf("Error marshalling config: %v\n", err)
	}

	var cfg v3.Cfg

	err = yaml.Unmarshal(configBytes, &cfg)
	if err != nil {
		return fmt.Errorf("Error unmarshalling config: %v\n", err)
	}

	req := external.PluginRequest{
		APIVersion: defaultAPIVersion,
		Command:    "create webhook",
		Args:       p.Args,
		Config:     cfg,
	}

	err = handlePluginResponse(fs, req, p.Path, p)
	if err != nil {
		return err
	}

	return nil
}

func (p *createWebhookSubcommand) InjectConfig(c config.Config) error {
	p.config = c
	return nil
}

func (p *createWebhookSubcommand) GetConfig() config.Config {
	return p.config
}
