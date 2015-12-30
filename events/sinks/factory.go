// Copyright 2015 Google Inc. All Rights Reserved.
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

package sinks

import (
	"fmt"

	"github.com/golang/glog"
	"k8s.io/heapster/events/core"
	"k8s.io/heapster/events/sinks/log"
	"k8s.io/heapster/util/flags"
)

type SinkFactory struct {
}

func (this *SinkFactory) Build(uri flags.Uri) (core.EventSink, error) {
	switch uri.Key {
	case "log":
		return logsink.NewLogSink(), nil
	default:
		return nil, fmt.Errorf("Sink not recognized: %s", uri.Key)
	}
}

func (this *SinkFactory) BuildAll(uris flags.Uris) []core.EventSink {
	result := make([]core.EventSink, 0, len(uris))
	for _, uri := range uris {
		sink, err := this.Build(uri)
		if err != nil {
			glog.Errorf("Failed to create sink: %v", err)
			continue
		}
		result = append(result, sink)
	}
	return result
}

func NewSinkFactory() *SinkFactory {
	return &SinkFactory{}
}