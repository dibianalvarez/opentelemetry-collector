// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/consumer/pdata"
)

func TestTracesUnmarshal_EncodingError(t *testing.T) {
	translator := &mockTranslator{}
	encoder := &mockEncoder{}

	tu := NewTracesUnmarshaler(encoder, translator)
	expectedBytes := []byte{1, 2, 3}
	expectedModel := struct{}{}

	encoder.On("DecodeTraces", expectedBytes).Return(expectedModel, errors.New("decode failed"))

	_, err := tu.Unmarshal(expectedBytes)
	assert.Error(t, err)
	assert.EqualError(t, err, "unmarshal failed: decode failed")
}

func TestTracesUnmarshal_TranslationError(t *testing.T) {
	translator := &mockTranslator{}
	encoder := &mockEncoder{}

	tu := NewTracesUnmarshaler(encoder, translator)
	expectedBytes := []byte{1, 2, 3}
	expectedModel := struct{}{}

	encoder.On("DecodeTraces", expectedBytes).Return(expectedModel, nil)
	translator.On("ToTraces", expectedModel).Return(pdata.NewTraces(), errors.New("translation failed"))

	_, err := tu.Unmarshal(expectedBytes)
	assert.Error(t, err)
	assert.EqualError(t, err, "converting model to pdata failed: translation failed")
}

func TestTracesUnmarshal_Decode(t *testing.T) {
	translator := &mockTranslator{}
	encoder := &mockEncoder{}

	tu := NewTracesUnmarshaler(encoder, translator)
	expectedTraces := pdata.NewTraces()
	expectedBytes := []byte{1, 2, 3}
	expectedModel := struct{}{}

	encoder.On("DecodeTraces", expectedBytes).Return(expectedModel, nil)
	translator.On("ToTraces", expectedModel).Return(expectedTraces, nil)

	actualTraces, err := tu.Unmarshal(expectedBytes)
	assert.NoError(t, err)
	assert.Equal(t, expectedTraces, actualTraces)
}

func TestMetricsUnmarshal_EncodingError(t *testing.T) {
	translator := &mockTranslator{}
	encoder := &mockEncoder{}

	mu := NewMetricsUnmarshaler(encoder, translator)
	expectedBytes := []byte{1, 2, 3}
	expectedModel := struct{}{}

	encoder.On("DecodeMetrics", expectedBytes).Return(expectedModel, errors.New("decode failed"))

	_, err := mu.Unmarshal(expectedBytes)
	assert.Error(t, err)
	assert.EqualError(t, err, "unmarshal failed: decode failed")
}

func TestMetricsUnmarshal_TranslationError(t *testing.T) {
	translator := &mockTranslator{}
	encoder := &mockEncoder{}

	mu := NewMetricsUnmarshaler(encoder, translator)
	expectedBytes := []byte{1, 2, 3}
	expectedModel := struct{}{}

	encoder.On("DecodeMetrics", expectedBytes).Return(expectedModel, nil)
	translator.On("ToMetrics", expectedModel).Return(pdata.NewMetrics(), errors.New("translation failed"))

	_, err := mu.Unmarshal(expectedBytes)
	assert.Error(t, err)
	assert.EqualError(t, err, "converting model to pdata failed: translation failed")
}

func TestMetricsUnmarshal_Decode(t *testing.T) {
	translator := &mockTranslator{}
	encoder := &mockEncoder{}

	mu := NewMetricsUnmarshaler(encoder, translator)
	expectedMetrics := pdata.NewMetrics()
	expectedBytes := []byte{1, 2, 3}
	expectedModel := struct{}{}

	encoder.On("DecodeMetrics", expectedBytes).Return(expectedModel, nil)
	translator.On("ToMetrics", expectedModel).Return(expectedMetrics, nil)

	actualMetrics, err := mu.Unmarshal(expectedBytes)
	assert.NoError(t, err)
	assert.Equal(t, expectedMetrics, actualMetrics)
}

func TestLogsUnmarshal_EncodingError(t *testing.T) {
	translator := &mockTranslator{}
	encoder := &mockEncoder{}

	lu := NewLogsUnmarshaler(encoder, translator)
	expectedBytes := []byte{1, 2, 3}
	expectedModel := struct{}{}

	encoder.On("DecodeLogs", expectedBytes).Return(expectedModel, errors.New("decode failed"))

	_, err := lu.Unmarshal(expectedBytes)
	assert.Error(t, err)
	assert.EqualError(t, err, "unmarshal failed: decode failed")
}

func TestLogsUnmarshal_TranslationError(t *testing.T) {
	translator := &mockTranslator{}
	encoder := &mockEncoder{}

	lu := NewLogsUnmarshaler(encoder, translator)
	expectedBytes := []byte{1, 2, 3}
	expectedModel := struct{}{}

	encoder.On("DecodeLogs", expectedBytes).Return(expectedModel, nil)
	translator.On("ToLogs", expectedModel).Return(pdata.NewLogs(), errors.New("translation failed"))

	_, err := lu.Unmarshal(expectedBytes)
	assert.Error(t, err)
	assert.EqualError(t, err, "converting model to pdata failed: translation failed")
}

func TestLogsUnmarshal_Decode(t *testing.T) {
	translator := &mockTranslator{}
	encoder := &mockEncoder{}

	lu := NewLogsUnmarshaler(encoder, translator)
	expectedLogs := pdata.NewLogs()
	expectedBytes := []byte{1, 2, 3}
	expectedModel := struct{}{}

	encoder.On("DecodeLogs", expectedBytes).Return(expectedModel, nil)
	translator.On("ToLogs", expectedModel).Return(expectedLogs, nil)

	actualLogs, err := lu.Unmarshal(expectedBytes)
	assert.NoError(t, err)
	assert.Equal(t, expectedLogs, actualLogs)
}
