/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package dbm dbm授权
package dbm

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubebkbcs/apis/bkbcs/v1"
	"github.com/parnurzeal/gorequest"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-webhook-server/options"
)

// NewDBMClient create DBM client
func NewDBMClient(os *options.Env) (*AuthDbmClient, error) {
	if len(os.ExternalSysConfig) == 0 {
		return nil, fmt.Errorf("create DBM client fialed, empty configuration")
	}
	client := AuthDbmClient{}
	err := json.Unmarshal([]byte(os.ExternalSysConfig), &client)
	if err != nil {
		return nil, err
	}

	client.AppCode = os.AppCode
	client.AppSecret = os.AppSecret
	client.Operator = os.AppOperator

	return &client, nil
}

// DoPri implement ExternalPrivilege interface
func (dc *AuthDbmClient) DoPri(os *options.Env, dbPrivConfig *v1.DbPrivConfigStatus) error {
	var (
		reqURL   = fmt.Sprintf("%s/%s/plugin/mysql/authorize/authorize_apply", dc.Host, dc.Environment)
		respData = &AuthorizeResponse{}
	)

	auth := fmt.Sprintf("{\"bk_app_code\": \"%s\", \"bk_app_secret\": \"%s\", \"bk_username\": \"%s\"}",
		dc.AppCode, dc.AppSecret, dc.Operator)

	req := AuthorizeRequest{
		App:            dbPrivConfig.AppName,
		User:           dbPrivConfig.CallUser,
		AccessDB:       dbPrivConfig.DbName,
		SourceIPs:      dbPrivConfig.NodeIp,
		TargetInstance: dbPrivConfig.TargetDb,
		// 暂时不生效，后续等dbm完善接口再修改为实际用户
		Operator: dbPrivConfig.Operator,
		Type:     os.CallType,
	}

	_, _, errs := gorequest.New().
		SetLogger(newLogger()).
		Timeout(defaultTimeOut).
		Post(reqURL).
		Set("Content-Type", "application/json").
		Set("X-Bkapi-Authorization", auth).
		SetDebug(dc.Debug).
		Send(&req).
		EndStruct(respData)
	if len(errs) > 0 {
		blog.Errorf("call DoPri failed: %v,req: %+v", errs[0], req)
		return errs[0]
	}

	if respData == nil || respData.Task == nil {
		return fmt.Errorf("call DoPri failed, empty response")
	}
	if respData.Code != 0 {
		return fmt.Errorf("call DoPri failed, %s", respData.Message)
	}

	dc.Task.TaskID = respData.Task.TaskID
	dc.Task.Platform = respData.Task.Platform

	return nil
}

// CheckFinalStatus implement ExternalPrivilege interface
func (dc *AuthDbmClient) CheckFinalStatus() error {
	if len(dc.Task.TaskID) == 0 || len(dc.Task.Platform) == 0 {
		return fmt.Errorf("taskid or platform is empty when call CheckFinalStatus")
	}

	taskid, _ := strconv.Atoi(dc.Task.TaskID)
	var (
		reqURL = fmt.Sprintf("%s/%s/plugin/mysql/authorize/query_authorize_apply_result?task_id=%d&platform=%s",
			dc.Host, dc.Environment, taskid, dc.Task.Platform)
		respData = &QueryResponse{}
	)

	auth := fmt.Sprintf("{\"bk_app_code\": \"%s\", \"bk_app_secret\": \"%s\", \"bk_username\": \"%s\"}",
		dc.AppCode, dc.AppSecret, dc.Operator)

	_, _, errs := gorequest.New().
		SetLogger(newLogger()).
		Timeout(defaultTimeOut).
		Get(reqURL).
		Set("Content-Type", "application/json").
		Set("X-Bkapi-Authorization", auth).
		SetDebug(dc.Debug).
		EndStruct(respData)
	if len(errs) > 0 {
		blog.Errorf("call CheckFinalStatus failed: %v", errs[0])
		return errs[0]
	}
	if respData == nil || respData.Status == nil {
		return fmt.Errorf("call CheckFinalStatus failed, empty response")
	}
	if respData.Code != 0 {
		return fmt.Errorf("call CheckFinalStatus failed, %s", respData.Message)
	}

	switch respData.Status.Status {
	case taskStatusRunning, taskStatusPending:
		return fmt.Errorf("task[%d] is %s, %s", taskid, respData.Status.Status, respData.Status.Msg)
	case taskStatusSucceeded:
		blog.Infof("task[%d] is %s, %s", taskid, respData.Status.Status, respData.Status.Msg)
		return nil
	default:
		return fmt.Errorf("task[%d] is %s, %s", taskid, respData.Status.Status, respData.Status.Msg)
	}
}
