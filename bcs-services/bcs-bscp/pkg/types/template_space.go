/*
Tencent is pleased to support the open source community by making Basic Service Configuration Platform available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	"bscp.io/pkg/criteria/errf"
	"bscp.io/pkg/dal/table"
	"bscp.io/pkg/runtime/filter"
)

// ListTemplateSpacesOption defines options to list group.
type ListTemplateSpacesOption struct {
	BizID  uint32             `json:"biz_id"`
	Filter *filter.Expression `json:"filter"`
	Page   *BasePage          `json:"page"`
}

// Validate the list group options
func (opt *ListTemplateSpacesOption) Validate(po *PageOption) error {
	if opt.BizID <= 0 {
		return errf.New(errf.InvalidParameter, "invalid biz id, should >= 1")
	}

	if opt.Filter == nil {
		return errf.New(errf.InvalidParameter, "filter is nil")
	}

	exprOpt := &filter.ExprOption{
		// remove biz_id because it's a required field in the option.
		RuleFields: table.TemplateSpaceColumns.WithoutColumn("biz_id"),
	}
	if err := opt.Filter.Validate(exprOpt); err != nil {
		return err
	}

	if opt.Page == nil {
		return errf.New(errf.InvalidParameter, "page is null")
	}

	if err := opt.Page.Validate(po); err != nil {
		return err
	}

	return nil
}

// ListTemplateSpaceDetails defines the response details of requested ListTemplateSpacesOption.
type ListTemplateSpaceDetails struct {
	Count   uint32                 `json:"count"`
	Details []*table.TemplateSpace `json:"details"`
}
