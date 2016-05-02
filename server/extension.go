// Copyright (C) 2015 NTT Innovation Institute, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !v8
// Copyright (C) 2015 NTT Innovation Institute, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/cloudwan/gohan/extension"
	"github.com/cloudwan/gohan/extension/gohanscript"
	"github.com/cloudwan/gohan/extension/golang"
	"github.com/cloudwan/gohan/extension/otto"
	"time"
)

func (server *Server) newEnvironment() extension.Environment {
	envs := []extension.Environment{}
	timelimit := time.Duration(server.timelimit) * time.Second
	for _, extension := range server.extensions {
		switch extension {
		case "javascript":
			envs = append(envs, otto.NewEnvironment(server.db, server.identityService, timelimit))
		case "gohanscript":
			envs = append(envs, gohanscript.NewEnvironment(timelimit))
		case "go":
			envs = append(envs, golang.NewEnvironment())
		}
	}
	return extension.NewEnvironment(envs)
}
