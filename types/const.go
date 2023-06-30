package types

type VariableValueType string
type FeatureValueType string

const (
	VariableValueTypeString VariableValueType = "stringValue"
	VariableValueTypeBool   VariableValueType = "boolValue"
	VariableValueTypeLong   VariableValueType = "longValue"
	VariableValueTypeDouble VariableValueType = "doubleValue"

	FeatureValueTypeBoolean FeatureValueType = "BOOLEAN"
	FeatureValueTypeString  FeatureValueType = "STRING"
	FeatureValueTypeLong    FeatureValueType = "LONG"
	FeatureValueTypeDouble  FeatureValueType = "DOUBLE"
)
