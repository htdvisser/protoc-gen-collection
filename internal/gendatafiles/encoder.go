// Copyright 2021 Hylke Visser
// SPDX-License-Identifier: Apache-2.0

package gendatafiles

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type JSONEncoder struct{}

func (JSONEncoder) FileExtension() string { return "json" }
func (JSONEncoder) EncodeData(v interface{}) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type YAMLEncoder struct{}

func (YAMLEncoder) FileExtension() string { return "yml" }
func (YAMLEncoder) EncodeData(v interface{}) (string, error) {
	b, err := yaml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
