// Copyright 2019 Istio Authors
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

//go:generate ../../../scripts/gen_topic.sh

package webanalytics

import (
	"net/http"

	"istio.io/bots/policybot/dashboard/types"
	"istio.io/bots/policybot/pkg/storage"
	"istio.io/bots/policybot/pkg/storage/cache"
)

// WebAnalytics lets users understand istio.io site traffic and content quality.
type WebAnalytics struct {
	store storage.Store
	cache *cache.Cache
	page  string
}

// New creates a new WebAnalytics instance.
func New(store storage.Store, cache *cache.Cache) *WebAnalytics {
	return &WebAnalytics{
		store: store,
		cache: cache,
		page:  string(MustAsset("page.html")),
	}
}

// Renders the HTML for this topic.
func (pr *WebAnalytics) Render(req *http.Request) (types.RenderInfo, error) {
	return types.RenderInfo{
		Content: pr.page,
	}, nil
}
