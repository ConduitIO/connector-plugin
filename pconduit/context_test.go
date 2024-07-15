// Copyright © 2024 Meroxa, Inc.
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

package pconduit

import (
	"context"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	"github.com/matryer/is"
	"testing"
)

func TestContextUtils_ConnectorToken(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	want := "test-token"
	ctx = ContextWithConnectorToken(ctx, want)
	got := ConnectorTokenFromContext(ctx)

	is.Equal(want, got)
}

func TestContextUtils_ConnectorID(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	want := "test-connector-id"
	ctx = ContextWithConnectorID(ctx, want)
	got := ConnectorIDFromContext(ctx)

	is.Equal(want, got)
}

func TestContextUtils_LogLevel(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	want := pconnector.LogLevelDebug
	ctx = ContextWithLogLevel(ctx, want)
	got := LogLevelFromContext(ctx)

	is.Equal(want, got)
}

func TestContextUtils_Enrich(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	cfg := pconnector.PluginConfig{
		Token:       "test-token",
		ConnectorID: "test-connector-id",
		LogLevel:    pconnector.LogLevelDebug,
	}

	got := Enrich(ctx, cfg)
	is.Equal(cfg.Token, ConnectorTokenFromContext(got))
	is.Equal(cfg.ConnectorID, ConnectorIDFromContext(got))
	is.Equal(cfg.LogLevel, LogLevelFromContext(got))
}
