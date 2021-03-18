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
 *
 */

package middle

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/actions"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/actions/utils/metrics"

	gorestful "github.com/emicklei/go-restful"
)

const (
	dynamicK8sPath   = actions.PathV1 + "/k8s/dynamic"
	dynamicMesosPath = actions.PathV1 + "/mesos/dynamic"

	clusterIDTag    = "clusterId"
	resourceTypeTag = "resourceType"

	// MetricsPrefix for metrics namespace
	MetricsPrefix = "bkbcs_storage"
)

// MetricsMiddleHandler returns a go-restful metrics report middleware.
func MetricsMiddleHandler(m Middleware) gorestful.FilterFunction {
	return func(req *gorestful.Request, resp *gorestful.Response, chain *gorestful.FilterChain) {
		r := &reporter{
			req:  req,
			resp: resp,
		}
		m.Measure(r, func() {
			chain.ProcessFilter(req, resp)
		})
	}
}

// Options is the configuration for the middleware factory.
type Options struct {
	// Recorder is the way the metrics will be recorder in the different backends.
	Recorder metrics.Recorder

	// GroupedStatus will group the status label in the form of `\dxx`. default will be false.
	GroupedStatus bool
	// DisableMeasureSize will disable the recording metrics about the response size.
	DisableMeasureSize bool
	// DisableMeasureInflight will disable the recording metrics about the inflight requests number.
	DisableMeasureInflight bool
}

// Middleware is a service that knows how to measure an HTTP handler by wrapping another handler.
type Middleware struct {
	cfg Options
}

// New returns the a Middleware service.
func New(cfg Options) Middleware {
	m := Middleware{cfg: cfg}

	return m
}

// Measure abstracts the HTTP handler implementation by only requesting a reporter,
// this reporter will return the required data to be measured.
// it accepts a next function that will be called as the wrapped logic before and after
// measurement actions.
// TODO: url routePath trans to handlerName label
func (m Middleware) Measure(reporter Reporter, next func()) {
	ctx := reporter.Context()
	hid := reporter.RoutePath()

	// Measure inflights if required.
	if !m.cfg.DisableMeasureInflight {
		props := metrics.HTTPProperties{
			Method:  reporter.Method(),
			Handler: hid,
		}
		m.cfg.Recorder.AddInflightRequests(ctx, props, 1)
		defer m.cfg.Recorder.AddInflightRequests(ctx, props, -1)
	}

	// Start the timer and when finishing measure the duration.
	start := time.Now()
	defer func() {
		duration := time.Since(start)

		// code aggregation
		var code string
		if m.cfg.GroupedStatus {
			code = fmt.Sprintf("%dxx", reporter.StatusCode()/100)
		} else {
			code = strconv.Itoa(reporter.StatusCode())
		}

		clusterID, resourceType := parseCLusterIDAndResourceType(reporter)
		// cluster_id & resource_type focus mongo/dynamic
		props := metrics.HTTPReqProperties{
			Handler:      hid,
			Method:       reporter.Method(),
			Code:         code,
			ClusterID:    clusterID,
			ResourceType: resourceType,
		}
		m.cfg.Recorder.ObserveHTTPRequestCounterDuration(ctx, props, duration)

		// Measure size of response if required.
		if !m.cfg.DisableMeasureSize {
			m.cfg.Recorder.ObserveHTTPResponseSize(ctx, props, reporter.BytesWritten())
		}
	}()

	// Call the wrapped logic.
	next()
}

// Reporter knows how to report the data to the Middleware so it can measure the different framework/libraries.
type Reporter interface {
	Method() string
	Context() context.Context
	URLPath() string
	StatusCode() int
	BytesWritten() int64
	GetReq() *gorestful.Request
	RoutePath() string
}

type reporter struct {
	req  *gorestful.Request
	resp *gorestful.Response
}

func (r *reporter) Method() string { return r.req.Request.Method }

func (r *reporter) Context() context.Context { return r.req.Request.Context() }

func (r *reporter) URLPath() string { return r.req.Request.URL.Path }

func (r *reporter) RoutePath() string { return r.req.SelectedRoutePath() }

func (r *reporter) StatusCode() int { return r.resp.StatusCode() }

func (r *reporter) BytesWritten() int64 { return int64(r.resp.ContentLength()) }

func (r *reporter) GetReq() *gorestful.Request { return r.req }

func parseCLusterIDAndResourceType(report Reporter) (string, string) {
	var (
		clusterID    = ""
		resourceType = ""
	)
	if report == nil || len(report.URLPath()) == 0 {
		return clusterID, resourceType
	}

	if strings.HasPrefix(report.URLPath(), dynamicK8sPath) || strings.HasPrefix(report.URLPath(), dynamicMesosPath) {
		clusterID = report.GetReq().PathParameter(clusterIDTag)
		resourceType = report.GetReq().PathParameter(resourceTypeTag)
	}

	return clusterID, resourceType
}
