package types

type VariableValueType string
type FeatureValueType string
type EvaluationReason string

const (
	VariableValueTypeString VariableValueType = "stringValue"
	VariableValueTypeBool   VariableValueType = "boolValue"
	VariableValueTypeLong   VariableValueType = "longValue"
	VariableValueTypeDouble VariableValueType = "doubleValue"

	FeatureValueTypeBoolean FeatureValueType = "BOOLEAN"
	FeatureValueTypeString  FeatureValueType = "STRING"
	FeatureValueTypeLong    FeatureValueType = "LONG"
	FeatureValueTypeDouble  FeatureValueType = "DOUBLE"

	EvaluationReasonDefault         EvaluationReason = "DEFAULT"
	EvaluationReasonOverride        EvaluationReason = "OVERRIDE_RULE"
	EvaluationReasonLaunchRuleMatch EvaluationReason = "LAUNCH_RULE_MATCH"
)
