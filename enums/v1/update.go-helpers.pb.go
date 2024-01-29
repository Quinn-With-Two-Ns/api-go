// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package enums

import (
	"fmt"
)

var (
	UpdateWorkflowExecutionLifecycleStage_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Admitted":    1,
		"Accepted":    2,
		"Completed":   3,
	}
)

// UpdateWorkflowExecutionLifecycleStageFromString parses a UpdateWorkflowExecutionLifecycleStage value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to UpdateWorkflowExecutionLifecycleStage
func UpdateWorkflowExecutionLifecycleStageFromString(s string) (UpdateWorkflowExecutionLifecycleStage, error) {
	if v, ok := UpdateWorkflowExecutionLifecycleStage_value[s]; ok {
		return UpdateWorkflowExecutionLifecycleStage(v), nil
	} else if v, ok := UpdateWorkflowExecutionLifecycleStage_shorthandValue[s]; ok {
		return UpdateWorkflowExecutionLifecycleStage(v), nil
	}
	return UpdateWorkflowExecutionLifecycleStage(0), fmt.Errorf("%s is not a valid UpdateWorkflowExecutionLifecycleStage", s)
}

var (
	UpdateRequestedEventOrigin_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Reapply":     1,
	}
)

// UpdateRequestedEventOriginFromString parses a UpdateRequestedEventOrigin value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to UpdateRequestedEventOrigin
func UpdateRequestedEventOriginFromString(s string) (UpdateRequestedEventOrigin, error) {
	if v, ok := UpdateRequestedEventOrigin_value[s]; ok {
		return UpdateRequestedEventOrigin(v), nil
	} else if v, ok := UpdateRequestedEventOrigin_shorthandValue[s]; ok {
		return UpdateRequestedEventOrigin(v), nil
	}
	return UpdateRequestedEventOrigin(0), fmt.Errorf("%s is not a valid UpdateRequestedEventOrigin", s)
}
