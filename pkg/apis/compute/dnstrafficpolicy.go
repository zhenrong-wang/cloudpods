// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compute

import (
	"yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/apis"
)

type DnsTrafficPolicyCreateInput struct {
	apis.EnabledStatusInfrasResourceBaseCreateInput

	Provider    string `json:"provider"`
	PolicyType  string `json:"policy_type"`
	PolicyValue string `json:"policy_value"`
	// 额外参数
	Options *jsonutils.JSONDict `json:"options"`
}

type DnsTrafficPolicyDetails struct {
	apis.EnabledStatusInfrasResourceBaseDetails
}

type DnsTrafficPolicyListInput struct {
	apis.EnabledStatusInfrasResourceBaseListInput

	Provider   []string `json:"provider"`
	PolicyType string   `json:"policy_type"`
}
