// Copyright 2018 JDCLOUD.COM
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
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type PanelPositionForUpdate struct {

    /* 该panel所在列  */
    Col int64 `json:"col"`

    /* 该panel高度  */
    Height int64 `json:"height"`

    /* 该panel在dashboard中的顺序  */
    Order int64 `json:"order"`

    /* 更新panel的uuid  */
    PanelUid string `json:"panelUid"`

    /* 该panel所在行  */
    Row int64 `json:"row"`

    /* 该panel宽度  */
    Width int64 `json:"width"`
}