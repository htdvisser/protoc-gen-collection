// Copyright 2021 Hylke Visser
// SPDX-License-Identifier: Apache-2.0

package gendatafiles

import "google.golang.org/genproto/googleapis/api/annotations"

type HTTPRule struct {
	Method        string `json:"method" yaml:"method"`
	Path          string `json:"path" yaml:"path"`
	Input         string `json:"input,omitempty" yaml:"input,omitempty"`
	InputMessage  string `json:"input_message,omitempty" yaml:"input_message,omitempty"`
	Output        string `json:"output,omitempty" yaml:"output,omitempty"`
	OutputMessage string `json:"output_message,omitempty" yaml:"output_message,omitempty"`
}

func (m *Method) AddHTTPRules(src *annotations.HttpRule) {
	if src == nil {
		return
	}
	httpRule := HTTPRule{}
	if body := src.GetBody(); body != "" && body != "*" {
		httpRule.Input = body
		for _, field := range m.src.Input().Fields() {
			if field.Name().String() == body {
				httpRule.InputMessage = field.FullyQualifiedName()
				break
			}
		}
	}
	if body := src.GetResponseBody(); body != "" && body != "*" {
		httpRule.Output = body
		for _, field := range m.src.Output().Fields() {
			if field.Name().String() == body {
				httpRule.OutputMessage = field.FullyQualifiedName()
				break
			}
		}
	}
	switch pattern := src.Pattern.(type) {
	case *annotations.HttpRule_Get:
		httpRule.Method, httpRule.Path = "GET", pattern.Get
	case *annotations.HttpRule_Put:
		httpRule.Method, httpRule.Path = "PUT", pattern.Put
	case *annotations.HttpRule_Post:
		httpRule.Method, httpRule.Path = "POST", pattern.Post
	case *annotations.HttpRule_Delete:
		httpRule.Method, httpRule.Path = "DELETE", pattern.Delete
	case *annotations.HttpRule_Patch:
		httpRule.Method, httpRule.Path = "PATCH", pattern.Patch
	case *annotations.HttpRule_Custom:
		httpRule.Method, httpRule.Path = pattern.Custom.GetKind(), pattern.Custom.GetPath()
	}
	m.HTTP = append(m.HTTP, httpRule)
	for _, additional := range src.AdditionalBindings {
		m.AddHTTPRules(additional)
	}
}
