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

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package pmetricotlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExportPartialSuccess_MoveTo(t *testing.T) {
	ms := generateTestExportPartialSuccess()
	dest := NewExportPartialSuccess()
	ms.MoveTo(dest)
	assert.Equal(t, NewExportPartialSuccess(), ms)
	assert.Equal(t, generateTestExportPartialSuccess(), dest)
}

func TestExportPartialSuccess_CopyTo(t *testing.T) {
	ms := NewExportPartialSuccess()
	orig := NewExportPartialSuccess()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestExportPartialSuccess()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestExportPartialSuccess_RejectedDataPoints(t *testing.T) {
	ms := NewExportPartialSuccess()
	assert.Equal(t, int64(0), ms.RejectedDataPoints())
	ms.SetRejectedDataPoints(int64(13))
	assert.Equal(t, int64(13), ms.RejectedDataPoints())
}

func TestExportPartialSuccess_ErrorMessage(t *testing.T) {
	ms := NewExportPartialSuccess()
	assert.Equal(t, "", ms.ErrorMessage())
	ms.SetErrorMessage("error message")
	assert.Equal(t, "error message", ms.ErrorMessage())
}

func generateTestExportPartialSuccess() ExportPartialSuccess {
	tv := NewExportPartialSuccess()
	fillTestExportPartialSuccess(tv)
	return tv
}

func fillTestExportPartialSuccess(tv ExportPartialSuccess) {
	tv.orig.RejectedDataPoints = int64(13)
	tv.orig.ErrorMessage = "error message"
}
