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

// Package gse provides gse api client.
package gse

import (
	"context"
	"fmt"

	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/cc"
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/components"
	pbfs "github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/protocol/feed-server"
)

// TransferFileData defines transfer file task data
type TransferFileData struct {
	Result TransferFileDataResult `json:"result"`
}

// TransferFileDataResult defines transfer file task result
type TransferFileDataResult struct {
	TaskID string `json:"task_id"`
}

// CreateTransferFileTask create sync transfer file task
func CreateTransferFileTask(ctx context.Context, sourceAgentID, sourceContainerID, sourceFileDir, sourceUser,
	filename string, targetAgentID, targetContainerID, targetFileDir, targetUser string) (string, error) {

	// 1. if sourceContainerID is set, means source is node, else is container
	// 2. if targetContainerID is set, means target is node, else is container

	url := fmt.Sprintf("%s/api/v2/task/extensions/async_transfer_file", cc.FeedServer().GSE.Host)
	authHeader := fmt.Sprintf("{\"bk_app_code\": \"%s\", \"bk_app_secret\": \"%s\"}",
		cc.FeedServer().GSE.AppCode, cc.FeedServer().GSE.AppSecret)
	resp, err := components.GetClient().R().
		SetContext(ctx).
		SetHeader("X-Bkapi-Authorization", authHeader).
		SetBody(map[string]interface{}{
			"timeout_seconds": 600,
			"auto_mkdir":      true,
			"upload_speed":    0,
			"download_speed":  0,
			"tasks": []map[string]interface{}{
				{
					"source": map[string]interface{}{
						"file_name": filename,
						"store_dir": sourceFileDir,
						"agent": map[string]interface{}{
							"user":            sourceUser,
							"bk_agent_id":     sourceAgentID,
							"bk_container_id": sourceContainerID,
						},
					},
					"target": map[string]interface{}{
						"file_name": filename,
						"store_dir": targetFileDir,
						"agents": []map[string]interface{}{
							{
								"user":            targetUser,
								"bk_agent_id":     targetAgentID,
								"bk_container_id": targetContainerID,
							},
						},
					},
				},
			},
		}).
		Post(url)

	if err != nil {
		return "", err
	}

	data := &TransferFileData{}
	if err := components.UnmarshalBKResult(resp, data); err != nil {
		return "", err
	}

	return data.Result.TaskID, nil
}

// TransferFileResultData defines transfer file task result data
type TransferFileResultData struct {
	Version string                         `json:"version"`
	Result  []TransferFileResultDataResult `json:"result"`
}

// TransferFileResultDataResult defines transfer file task result data result
type TransferFileResultDataResult struct {
	Content   TransferFileResultDataResultContent `json:"content"`
	ErrorCode int                                 `json:"error_code"`
	ErrorMsg  string                              `json:"error_msg"`
}

// TransferFileResultDataResultContent defines transfer file task result data result content
type TransferFileResultDataResultContent struct {
	DestAgentID       string `json:"dest_agent_id"`
	DestContainerID   string `json:"dest_container_id"`
	DestFileDir       string `json:"dest_file_dir"`
	DestFileName      string `json:"dest_file_name"`
	Mode              int    `json:"mode"`
	Progress          int    `json:"progress"`
	SourceAgentID     string `json:"source_agent_id"`
	SourceContainerID string `json:"source_container_id"`
	SourceFileDir     string `json:"source_file_dir"`
	SourceFileName    string `json:"source_file_name"`
	Speed             int    `json:"speed"`
	Status            int    `json:"status"`
	StatusInfo        string `json:"status_info"`
	Type              string `json:"type"`
	StartTime         int64  `json:"start_time"`
	EndTime           int64  `json:"end_time"`
	Size              int64  `json:"size"`
}

// TransferFileResult query transfer file task result
func TransferFileResult(ctx context.Context, taskID string) (pbfs.AsyncDownloadStatus, error) {

	url := fmt.Sprintf("%s/api/v2/task/extensions/get_transfer_file_result", cc.FeedServer().GSE.Host)
	authHeader := fmt.Sprintf("{\"bk_app_code\": \"%s\", \"bk_app_secret\": \"%s\"}",
		cc.FeedServer().GSE.AppCode, cc.FeedServer().GSE.AppSecret)
	resp, err := components.GetClient().R().
		SetContext(ctx).
		SetHeader("X-Bkapi-Authorization", authHeader).
		SetBody(map[string]interface{}{
			"task_id": taskID,
		}).
		Post(url)

	if err != nil {
		return pbfs.AsyncDownloadStatus_FAILED, err
	}

	data := &TransferFileResultData{}
	if err := components.UnmarshalBKResult(resp, data); err != nil {
		return pbfs.AsyncDownloadStatus_FAILED, err
	}

	// any task failed, return failed
	// any task downloading, return downloading
	// all task success, return success
	for _, result := range data.Result {
		switch result.ErrorCode {
		case 0:
		case 115:
			return pbfs.AsyncDownloadStatus_DOWNLOADING, nil
		default:
			return pbfs.AsyncDownloadStatus_FAILED, fmt.Errorf(result.ErrorMsg)
		}
	}
	return pbfs.AsyncDownloadStatus_SUCCESS, nil
}
