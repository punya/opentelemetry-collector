// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zpagesextension

import (
	"go.opentelemetry.io/collector/config"
)

// Config has the configuration for the extension enabling the zPages extension.
type Config struct {
	config.ExtensionSettings `mapstructure:",squash"`

	// Endpoint is the address and port in which the zPages will be listening to.
	// Use localhost:<port> to make it available only locally, or ":<port>" to
	// make it available on all network interfaces.
	Endpoint string `mapstructure:"endpoint"`
}
