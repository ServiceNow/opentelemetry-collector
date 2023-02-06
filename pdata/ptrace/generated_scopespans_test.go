// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestScopeSpans_MoveTo(t *testing.T) {
	ms := generateTestScopeSpans()
	dest := NewScopeSpans()
	ms.MoveTo(dest)
	assert.Equal(t, NewScopeSpans(), ms)
	assert.Equal(t, generateTestScopeSpans(), dest)
}

func TestScopeSpans_CopyTo(t *testing.T) {
	ms := NewScopeSpans()
	orig := NewScopeSpans()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestScopeSpans()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestScopeSpans_Scope(t *testing.T) {
	ms := NewScopeSpans()
	internal.FillTestInstrumentationScope(internal.InstrumentationScope(ms.Scope()))
	assert.Equal(t, pcommon.InstrumentationScope(internal.GenerateTestInstrumentationScope()), ms.Scope())
}

func TestScopeSpans_SchemaUrl(t *testing.T) {
	ms := NewScopeSpans()
	assert.Equal(t, "", ms.SchemaUrl())
	ms.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	assert.Equal(t, "https://opentelemetry.io/schemas/1.5.0", ms.SchemaUrl())
}

func TestScopeSpans_Spans(t *testing.T) {
	ms := NewScopeSpans()
	assert.Equal(t, NewSpanSlice(), ms.Spans())
	fillTestSpanSlice(ms.Spans())
	assert.Equal(t, generateTestSpanSlice(), ms.Spans())
}

func generateTestScopeSpans() ScopeSpans {
	tv := NewScopeSpans()
	fillTestScopeSpans(tv)
	return tv
}

func fillTestScopeSpans(tv ScopeSpans) {
	internal.FillTestInstrumentationScope(internal.NewInstrumentationScope(&tv.orig.Scope))
	tv.orig.SchemaUrl = "https://opentelemetry.io/schemas/1.5.0"
	fillTestSpanSlice(newSpanSlice(&tv.orig.Spans))
}