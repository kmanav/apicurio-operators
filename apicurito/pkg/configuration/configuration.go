/*
 * Copyright (C) 2020 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/imdario/mergo"

	"github.com/apicurio/apicurio-operators/apicurito/pkg/apis/apicur/v1alpha1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

var ConfigFile string

type Config struct {
	Image string
}

type ConfigLoader interface {
	Config(apicurito *v1alpha1.Apicurito) error
}

// Returns all processed configuration for Apicurito
func (c *Config) Config(apicurito *v1alpha1.Apicurito) error {
	if err := c.loadFromFile(ConfigFile); err != nil {
		return err
	}

	if err := c.setPropertiesFromEnv(); err != nil {
		return err
	}

	if err := c.setPropertiesFromApi(apicurito); err != nil {
		return err
	}

	return nil
}

// Load configuration from config file. Config file is expected to be a yaml
// The returned configuration is parsed to JSON and returned as Config
func (c *Config) loadFromFile(config string) (err error) {
	data, err := ioutil.ReadFile(config)
	if err != nil {
		return err
	}

	if strings.HasSuffix(config, ".yaml") || strings.HasSuffix(config, ".yml") {
		data, err = yaml.ToJSON(data)
		if err != nil {
			return err
		}
	}

	if err := json.Unmarshal(data, c); err != nil {
		return err
	}

	return
}

// From apicurito CR, Unmarshal it into a property object
func (c *Config) setPropertiesFromApi(apicurito *v1alpha1.Apicurito) (err error) {
	cApi := &Config{}
	jsonProperties, err := json.Marshal(apicurito.Spec)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonProperties, cApi)
	if err != nil {
		return err
	}

	err = mergo.Merge(c, cApi, mergo.WithOverride)
	return
}

// Set fields in the configuration from environment variables if they
// are defined
func (c *Config) setPropertiesFromEnv() (err error) {
	cEnv := Config{Image: os.Getenv("APICURITO_IMAGE")}

	err = mergo.Merge(c, cEnv, mergo.WithOverride)
	return
}
