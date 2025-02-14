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

// MavenConfig stores maven related configuration information
type MavenConfig struct {
	MavenAppName  string   `yaml:"mavenAppName,omitempty" json:"mavenAppName,omitempty"`
	ArtifactType  string   `yaml:"artifactType"`
	MavenProfiles []string `yaml:"mavenProfiles,omitempty" json:"mavenProfiles,omitempty"`
}

const (
	// MavenConfigType stores the maven config
	MavenConfigType transformertypes.ConfigType = "Maven"
	// MavenPomPathType stores the Maven POM file Path
	MavenPomPathType transformertypes.PathType = "MavenPom"
)
