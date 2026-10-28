// Copyright 2015 Matthew Holt and The Caddy Authors
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


package caddytls

import (
	"sync"
	"testing"
)

func TestECHConfigsToPublishMatchesNamesCaseInsensitively(t *testing.T) {
	ech := &ECH{
		configsMu: new(sync.RWMutex),
		configs: map[string][]echConfig{
			"public.example": {{ConfigID: 1}},
			"other.example":  {{ConfigID: 2}, {ConfigID: 3}},
		},
	}

	for i, tc := range []struct {
		names   []string
		wantIDs int
	}{
		{names: nil, wantIDs: 3},
		{names: []string{"public.example"}, wantIDs: 1},
		{names: []string{"Public.Example"}, wantIDs: 1},
		{names: []string{" PUBLIC.example "}, wantIDs: 1},
		{names: []string{"public.example", "OTHER.example"}, wantIDs: 3},
		{names: []string{"missing.example"}, wantIDs: 0},
		{names: []string{}, wantIDs: 0},
	} {
		list, ids := ech.configsToPublish(tc.names)
		if len(ids) != tc.wantIDs || len(list) != tc.wantIDs {
			t.Errorf("Test %d: names %q: got %d configs / %d IDs, want %d", i, tc.names, len(list), len(ids), tc.wantIDs)
		}
	}
}
