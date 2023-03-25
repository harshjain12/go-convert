// Copyright 2022 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yaml

type (
	Pipeline struct {
		Concurrency *Concurrency      `yaml:"concurrency,omitempty"`
		Defaults    *Defaults         `yaml:"defaults,omitempty"`
		Env         map[string]string `yaml:"env,omitempty"`
		Jobs        map[string]*Job   `yaml:"jobs,omitempty"`
		Name        string            `yaml:"name,omitempty"`
		On          *On               `yaml:"on,omitempty"`
		Permissions *Permissions      `yaml:"permissions,omitempty"`
		RunName     string            `yaml:"run-name,omitempty"`
	}

	Defaults struct {
		Run *Run `yaml:"run,omitempty"`
	}

	Event struct {
		Types []string `yaml:"types,omitempty"`
	}

	Input struct {
		Default     interface{} `yaml:"default,omitempty"`
		Description string      `yaml:"description,omitempty"`
		Options     interface{} `yaml:"options,omitempty"`
		Required    bool        `yaml:"required,omitempty"`
		Type        string      `yaml:"type,omitempty"`
	}

	Job struct {
		Concurrency   *Concurrency        `yaml:"concurrency,omitempty"`
		Container     string              `yaml:"container,omitempty"`         // TODO struct or string
		ContinueOnErr bool                `yaml:"continue-on-error,omitempty"` // TODO string instead of bool? `continue-on-error: ${{ matrix.experimental }}`
		Defaults      *Defaults           `yaml:"defaults,omitempty"`
		Env           map[string]string   `yaml:"env,omitempty"`
		Environment   interface{}         `yaml:"environment,omitempty"` // TODO
		If            string              `yaml:"if,omitempty"`
		Name          string              `yaml:"name,omitempty"`
		Needs         Stringorslice       `yaml:"needs,omitempty"`
		Outputs       map[string]string   `yaml:"outputs,omitempty"`
		Permissions   *Permissions        `yaml:"permissions,omitempty"`
		RunsOn        string              `yaml:"runs-on,omitempty"`
		Services      map[string]*Service `yaml:"services,omitempty"`
		Steps         []*Step             `yaml:"steps,omitempty"`
		Strategy      *Strategy           `yaml:"strategy,omitempty"`
		TimeoutInMin  int                 `yaml:"timeout-in-minutes,omitempty"`
		Uses          string              `yaml:"uses,omitempty"`
	}

	Matrix struct {
		Exclude []map[string]interface{} `yaml:"exclude,omitempty"`
		Include []map[string]interface{} `yaml:"include,omitempty"`
		Matrix  map[string][]string      `yaml:",inline"`
	}

	PullRequest struct {
		Branches        []string `yaml:"branches,omitempty"`
		BranchesIgnore  []string `yaml:"branches-ignore,omitempty"`
		Paths           []string `yaml:"paths,omitempty"`
		PathsIgnore     []string `yaml:"paths-ignore,omitempty"`
		Tags            []string `yaml:"tags,omitempty"`
		TagsIgnore      []string `yaml:"tags-ignore,omitempty"`
		Types           []string `yaml:"types,omitempty"`
		ReviewApproved  bool     `yaml:"review-approved,omitempty"`
		ReviewDismissed bool     `yaml:"review-dismissed,omitempty"`
	}

	PullRequestTarget struct {
		Branches       []string `yaml:"branches,omitempty"`
		BranchesIgnore []string `yaml:"branches-ignore,omitempty"`
		Types          []string `yaml:"types,omitempty"`
	}

	Push struct {
		Branches       []string `yaml:"branches,omitempty"`
		BranchesIgnore []string `yaml:"branches-ignore,omitempty"`
		Paths          []string `yaml:"paths,omitempty"`
		PathsIgnore    []string `yaml:"paths-ignore,omitempty"`
		Tags           []string `yaml:"tags,omitempty"`
		TagsIgnore     []string `yaml:"tags-ignore,omitempty"`
	}

	Run struct {
		Shell      string `yaml:"shell,omitempty"`
		WorkingDir string `yaml:"working-directory,omitempty"`
	}

	Service struct {
		Env      map[string]string `yaml:"env,omitempty"`
		Image    string            `yaml:"image,omitempty"`
		Networks []string          `yaml:"networks,omitempty"`
		Options  []string          `yaml:"options,omitempty"`
		Ports    []string          `yaml:"ports,omitempty"`
		Volumes  []string          `yaml:"volumes,omitempty"`
	}

	Step struct {
		Name string                 `yaml:"name,omitempty"`
		Env  map[string]string      `yaml:"env,omitempty"`
		If   string                 `yaml:"if,omitempty"`
		Run  string                 `yaml:"run,omitempty"`
		With map[string]interface{} `yaml:"with,omitempty"`
		Uses string                 `yaml:"uses,omitempty"`
	}

	Strategy struct {
		Matrix *Matrix `yaml:"matrix,omitempty"`
	}

	WorkflowCall struct {
		Inputs    map[string]interface{}         `yaml:"inputs,omitempty"`
		Outputs   map[string]interface{}         `yaml:"outputs,omitempty"`
		Secrets   map[string]*WorkflowCallSecret `yaml:"secrets,omitempty"`
		Workflows []string                       `yaml:"workflows,omitempty"`
	}

	WorkflowCallSecret struct {
		Description string `yaml:"description,omitempty"`
		Required    bool   `yaml:"required,omitempty"`
	}

	WorkflowDispatch struct {
		Inputs map[string]*Input `yaml:"inputs,omitempty"`
	}

	WorkflowRun struct {
		Branches       []string `yaml:"branches,omitempty"`
		BranchesIgnore []string `yaml:"branches-ignore,omitempty"`
		Types          []string `yaml:"types,omitempty"`
		Workflows      []string `yaml:"workflows,omitempty"`
	}
)
