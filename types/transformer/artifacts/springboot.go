/*
 *  Copyright IBM Corporation 2021
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package artifacts

import (
	transformertypes "github.com/konveyor/move2kube/types/transformer"
)

// SpringBootConfig stores spring boot related configuration information
type SpringBootConfig struct {
	SpringBootVersion  string   `yaml:"springBootVersion,omitempty" json:"springBootVersion,omitempty"`
	SpringBootAppName  string   `yaml:"springBootAppName,omitempty" json:"springBootAppName,omitempty"`
	SpringBootProfiles []string `yaml:"springBootProfiles,omitempty" json:"springBootProfiles,omitempty"`
}

const (
	// SpringBootConfigType stores the springboot config
	SpringBootConfigType transformertypes.ConfigType = "SpringBoot"
)
