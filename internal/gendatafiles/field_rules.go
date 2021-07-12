// Copyright 2021 Hylke Visser
// SPDX-License-Identifier: Apache-2.0

package gendatafiles

import (
	"time"

	"github.com/envoyproxy/protoc-gen-validate/validate"
)

type FieldRules struct {
	// Repeated
	MinItems uint64 `json:"min_items,omitempty" yaml:"min_items,omitempty"`
	MaxItems uint64 `json:"max_items,omitempty" yaml:"max_items,omitempty"`
	Unique   bool   `json:"unique,omitempty" yaml:"unique,omitempty"`
	// Map
	MinPairs uint64 `json:"min_pairs,omitempty" yaml:"min_pairs,omitempty"`
	MaxPairs uint64 `json:"max_pairs,omitempty" yaml:"max_pairs,omitempty"`
	NoSparse bool   `json:"no_sparse,omitempty" yaml:"no_sparse,omitempty"`
	// Enum
	DefinedOnly bool `json:"defined_only,omitempty" yaml:"defined_only,omitempty"`
	// Message
	Skip     bool `json:"skip,omitempty" yaml:"skip,omitempty"`
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
	// String
	Len      uint64 `json:"len,omitempty" yaml:"len,omitempty"`
	MinLen   uint64 `json:"min_len,omitempty" yaml:"min_len,omitempty"`
	MaxLen   uint64 `json:"max_len,omitempty" yaml:"max_len,omitempty"`
	LenBytes uint64 `json:"len_bytes,omitempty" yaml:"len_bytes,omitempty"`
	MinBytes uint64 `json:"min_bytes,omitempty" yaml:"min_bytes,omitempty"`
	MaxBytes uint64 `json:"max_bytes,omitempty" yaml:"max_bytes,omitempty"`
	Pattern  string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	Email    bool   `json:"email,omitempty" yaml:"email,omitempty"`
	Hostname bool   `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	URI      bool   `json:"uri,omitempty" yaml:"uri,omitempty"`
	URIRef   bool   `json:"uri_ref,omitempty" yaml:"uri_ref,omitempty"`
	Address  bool   `json:"address,omitempty" yaml:"address,omitempty"`
	UUID     bool   `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	// String and Bytes
	Prefix      interface{} `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Suffix      interface{} `json:"suffix,omitempty" yaml:"suffix,omitempty"`
	Contains    interface{} `json:"contains,omitempty" yaml:"contains,omitempty"`
	NotContains interface{} `json:"not_contains,omitempty" yaml:"not_contains,omitempty"`
	IP          bool        `json:"ip,omitempty" yaml:"ip,omitempty"`
	IPv4        bool        `json:"ipv4,omitempty" yaml:"ipv4,omitempty"`
	IPv6        bool        `json:"ipv6,omitempty" yaml:"ipv6,omitempty"`
	// Timestamp
	LtNow  bool          `json:"lt_now,omitempty" yaml:"lt_now,omitempty"`
	GtNow  bool          `json:"gt_now,omitempty" yaml:"gt_now,omitempty"`
	Within time.Duration `json:"within,omitempty" yaml:"within,omitempty"`
	// Most types
	Const interface{} `json:"const,omitempty" yaml:"const,omitempty"`
	Lt    interface{} `json:"lt,omitempty" yaml:"lt,omitempty"`
	Lte   interface{} `json:"lte,omitempty" yaml:"lte,omitempty"`
	Gt    interface{} `json:"gt,omitempty" yaml:"gt,omitempty"`
	Gte   interface{} `json:"gte,omitempty" yaml:"gte,omitempty"`
	In    interface{} `json:"in,omitempty" yaml:"in,omitempty"`
	NotIn interface{} `json:"not_in,omitempty" yaml:"not_in,omitempty"`
	// Ignore empty (skip validations on zero value)
	IgnoreEmpty bool `json:"ignore_empty,omitempty" yaml:"ignore_empty,omitempty"`
}

func (f *Field) AddFieldRules(src *validate.FieldRules) {
	if src == nil {
		return
	}
	fieldType := f.src.Type()
	if rules := src.GetRepeated(); rules != nil && f.Repeated != nil {
		f.Rules.MinItems = rules.GetMinItems()
		f.Rules.MaxItems = rules.GetMaxItems()
		f.Rules.Unique = rules.GetUnique()
		f.Repeated.Rules.AddFieldRules(rules.GetItems(), fieldType.Element())
		f.Rules.IgnoreEmpty = rules.GetIgnoreEmpty()
	}
	if rules := src.GetMap(); rules != nil && f.MapKey != nil && f.MapValue != nil {
		f.Rules.MinPairs = rules.GetMinPairs()
		f.Rules.MaxPairs = rules.GetMaxPairs()
		f.Rules.NoSparse = rules.GetNoSparse()
		f.MapKey.Rules.AddFieldRules(rules.GetKeys(), fieldType.Key())
		f.MapValue.Rules.AddFieldRules(rules.GetValues(), fieldType.Element())
		f.Rules.IgnoreEmpty = rules.GetIgnoreEmpty()
	}
	f.Rules.AddFieldRules(src, fieldType)
}

func (f *FieldRules) addFloatRules(src *validate.FloatRules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addDoubleRules(src *validate.DoubleRules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addInt32Rules(src *validate.Int32Rules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addInt64Rules(src *validate.Int64Rules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addUint32Rules(src *validate.UInt32Rules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addUint64Rules(src *validate.UInt64Rules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addSint32Rules(src *validate.SInt32Rules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addSint64Rules(src *validate.SInt64Rules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addFixed32Rules(src *validate.Fixed32Rules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addFixed64Rules(src *validate.Fixed64Rules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addSfixed32Rules(src *validate.SFixed32Rules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addSfixed64Rules(src *validate.SFixed64Rules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	if src.Lt != nil {
		f.Lt = src.Lt
	}
	if src.Lte != nil {
		f.Lte = src.Lte
	}
	if src.Gt != nil {
		f.Gt = src.Gt
	}
	if src.Gte != nil {
		f.Gte = src.Gte
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addBoolRules(src *validate.BoolRules) {
	if src.Const != nil {
		f.Const = *src.Const
	}
}

func (f *FieldRules) addStringRules(src *validate.StringRules) {
	if src.Const != nil {
		f.Const = *src.Const
	}
	f.Len = src.GetLen()
	f.MinLen = src.GetMinLen()
	f.MaxLen = src.GetMaxLen()
	f.LenBytes = src.GetLenBytes()
	f.MinBytes = src.GetMinBytes()
	f.MaxBytes = src.GetMaxBytes()
	f.Pattern = src.GetPattern()
	if src.Prefix != nil {
		f.Prefix = src.Prefix
	}
	if src.Suffix != nil {
		f.Suffix = src.Suffix
	}
	if src.Contains != nil {
		f.Contains = src.Contains
	}
	if src.NotContains != nil {
		f.NotContains = src.NotContains
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IP = src.GetIp()
	f.IPv4 = src.GetIpv4()
	f.IPv6 = src.GetIpv6()
	f.Email = src.GetEmail()
	f.Hostname = src.GetHostname()
	f.URI = src.GetUri()
	f.URIRef = src.GetUriRef()
	f.Address = src.GetAddress()
	f.UUID = src.GetUuid()
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addBytesRules(src *validate.BytesRules) {
	if src.Const != nil {
		f.Const = src.Const
	}
	f.Len = src.GetLen()
	f.MinLen = src.GetMinLen()
	f.MaxLen = src.GetMaxLen()
	f.Pattern = src.GetPattern()
	if src.Prefix != nil {
		f.Prefix = src.Prefix
	}
	if src.Suffix != nil {
		f.Suffix = src.Suffix
	}
	if src.Contains != nil {
		f.Contains = src.Contains
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
	f.IP = src.GetIp()
	f.IPv4 = src.GetIpv4()
	f.IPv6 = src.GetIpv6()
	f.IgnoreEmpty = src.GetIgnoreEmpty()
}

func (f *FieldRules) addEnumRules(src *validate.EnumRules) {
	f.DefinedOnly = src.GetDefinedOnly()
	if src.Const != nil {
		f.Const = *src.Const
	}
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
}

func (f *FieldRules) addMessageRules(src *validate.MessageRules) {
	f.Skip = src.GetSkip()
	f.Required = src.GetRequired()
}

func (f *FieldRules) addAnyRules(src *validate.AnyRules) {
	f.Required = src.GetRequired()
	if src.In != nil {
		f.In = src.In
	}
	if src.NotIn != nil {
		f.NotIn = src.NotIn
	}
}

func (f *FieldRules) addDurationRules(src *validate.DurationRules) {
	if src.Const != nil {
		f.Const = src.Const.AsDuration()
	}
	if src.Lt != nil {
		f.Lt = src.Lt.AsDuration()
	}
	if src.Lte != nil {
		f.Lte = src.Lte.AsDuration()
	}
	if src.Gt != nil {
		f.Gt = src.Gt.AsDuration()
	}
	if src.Gte != nil {
		f.Gte = src.Gte.AsDuration()
	}
	if src.In != nil {
		in := make([]time.Duration, len(src.In))
		for i, p := range src.In {
			in[i] = p.AsDuration()
		}
		f.In = in
	}
	if src.NotIn != nil {
		notIn := make([]time.Duration, len(src.NotIn))
		for i, p := range src.NotIn {
			notIn[i] = p.AsDuration()
		}
		f.NotIn = notIn
	}
}

func (f *FieldRules) addTimestampRules(src *validate.TimestampRules) {
	f.Required = src.GetRequired()
	if src.Const != nil {
		f.Const = src.Const.AsTime()
	}
	if src.Lt != nil {
		f.Lt = src.Lt.AsTime()
	}
	if src.Lte != nil {
		f.Lte = src.Lte.AsTime()
	}
	if src.Gt != nil {
		f.Gt = src.Gt.AsTime()
	}
	if src.Gte != nil {
		f.Gte = src.Gte.AsTime()
	}
	f.LtNow = src.GetLtNow()
	f.GtNow = src.GetGtNow()
	if src.Within != nil {
		f.Within = src.Within.AsDuration()
	}
}

func (f *FieldRules) AddFieldRules(src *validate.FieldRules, t PGSFieldType) {
	if src == nil {
		return
	}
	if rules := src.GetFloat(); rules != nil {
		f.addFloatRules(rules)
	}
	if rules := src.GetDouble(); rules != nil {
		f.addDoubleRules(rules)
	}
	if rules := src.GetInt32(); rules != nil {
		f.addInt32Rules(rules)
	}
	if rules := src.GetInt64(); rules != nil {
		f.addInt64Rules(rules)
	}
	if rules := src.GetUint32(); rules != nil {
		f.addUint32Rules(rules)
	}
	if rules := src.GetUint64(); rules != nil {
		f.addUint64Rules(rules)
	}
	if rules := src.GetSint32(); rules != nil {
		f.addSint32Rules(rules)
	}
	if rules := src.GetSint64(); rules != nil {
		f.addSint64Rules(rules)
	}
	if rules := src.GetFixed32(); rules != nil {
		f.addFixed32Rules(rules)
	}
	if rules := src.GetFixed64(); rules != nil {
		f.addFixed64Rules(rules)
	}
	if rules := src.GetSfixed32(); rules != nil {
		f.addSfixed32Rules(rules)
	}
	if rules := src.GetSfixed64(); rules != nil {
		f.addSfixed64Rules(rules)
	}
	if rules := src.GetBool(); rules != nil {
		f.addBoolRules(rules)
	}
	if rules := src.GetString_(); rules != nil {
		f.addStringRules(rules)
	}
	if rules := src.GetBytes(); rules != nil {
		f.addBytesRules(rules)
	}
	if rules := src.GetEnum(); rules != nil && t.IsEnum() {
		f.addEnumRules(rules)
	}
	if rules := src.GetMessage(); rules != nil && t.IsEmbed() {
		f.addMessageRules(rules)
	}
	if rules := src.GetAny(); rules != nil {
		f.addAnyRules(rules)
	}
	if rules := src.GetDuration(); rules != nil {
		f.addDurationRules(rules)
	}
	if rules := src.GetTimestamp(); rules != nil {
		f.addTimestampRules(rules)
	}
}
