package models

import (
	"github.com/michimani/evidentlylocal/types"
)

type Feature struct {
	DefaultVariation string                 `json:"defaultVariation"`
	EntityOverrides  EntityOverride         `json:"entityOverrides"`
	Name             string                 `json:"name"`
	Project          string                 `json:"project"`
	Status           string                 `json:"status"`
	ValueType        types.FeatureValueType `json:"valueType"`
	Variations       []Variation            `json:"variations"`
}

type EntityOverride map[string]string

type Variation struct {
	Name  string                          `json:"name"`
	Value map[types.VariableValueType]any `json:"value"`
}

func (f *Feature) GetValue(variation string) any {
	for _, v := range f.Variations {
		if v.Name == variation {
			return v.Value[f.VariableValueType()]
		}
	}

	return nil
}

func (f *Feature) GetDefaultValue() any {
	return f.GetValue(f.DefaultVariation)
}

func (f *Feature) VariableValueType() types.VariableValueType {
	switch f.ValueType {
	case types.FeatureValueTypeString:
		return types.VariableValueTypeString
	case types.FeatureValueTypeBoolean:
		return types.VariableValueTypeBool
	case types.FeatureValueTypeLong:
		return types.VariableValueTypeLong
	case types.FeatureValueTypeDouble:
		return types.VariableValueTypeDouble
	default:
		// noop
	}

	return ""
}
