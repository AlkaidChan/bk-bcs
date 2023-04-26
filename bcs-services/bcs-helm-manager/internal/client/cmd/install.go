/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-helm-manager/internal/common"
	helmmanager "github.com/Tencent/bk-bcs/bcs-services/bcs-helm-manager/proto/bcs-helm-manager"

	"github.com/spf13/cobra"
)

var (
	flagValueFile []string
	flagArgs      string
	installCMD    = &cobra.Command{
		Use:   "install",
		Short: "install",
		Long:  "install chart release",
		Run:   Install,
	}
)

func init() {
	installCMD.PersistentFlags().StringVarP(
		&flagProject, "project", "p", "", "project id for operation")
	installCMD.PersistentFlags().StringVarP(
		&flagRepository, "repository", "r", "", "repository name for operation")
	installCMD.PersistentFlags().StringVarP(
		&flagNamespace, "namespace", "n", "", "release namespace for operation")
	installCMD.PersistentFlags().StringVarP(
		&flagCluster, "cluster", "", "", "release cluster id for operation")
	installCMD.PersistentFlags().StringSliceVarP(
		&flagValueFile, "file", "f", nil, "value file for installation")
	installCMD.PersistentFlags().StringVarP(
		&flagArgs, "args", "", "", "args to append to helm command")
}

// Install provide the actions to do installCMD
func Install(cmd *cobra.Command, args []string) {
	req := &helmmanager.InstallReleaseV1Req{}

	if len(args) < 3 {
		fmt.Printf("install args need at least 3, install [name] [chart] [version]\n")
		os.Exit(1)
	}
	values, err := getValues()
	if err != nil {
		fmt.Printf("read values file failed, %s\n", err.Error())
		os.Exit(1)
	}

	req.Name = common.GetStringP(args[0])
	req.Namespace = &flagNamespace
	req.ClusterID = &flagCluster
	req.ProjectCode = &flagProject
	req.Repository = &flagRepository
	req.Chart = common.GetStringP(args[1])
	req.Version = common.GetStringP(args[2])
	req.Values = values
	req.ValueFile = &flagValueFile[0]
	if flagArgs != "" {
		req.Args = strings.Split(flagArgs, " ")
	}

	c := newClientWithConfiguration()
	err = c.Release().Install(cmd.Context(), req)
	if err != nil {
		fmt.Printf("install release failed, %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("success to install release %s in version %s namespace %s cluster %s "+
		"with appVersion %s revision %d\n",
		req.GetName(), req.GetVersion(), req.GetNamespace(), req.GetClusterID())
}

func getValues() ([]string, error) {
	values := make([]string, 0, 10)
	for _, vf := range flagValueFile {
		content, err := os.ReadFile(vf)
		if err != nil {
			return nil, err
		}
		values = append(values, string(content))
	}

	return values, nil
}
