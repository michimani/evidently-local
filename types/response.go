package types

type EvaluateFeatureResponse struct {
	Details   string        `json:"details"`
	Reason    string        `json:"reason"`
	Value     VariableValue `json:"value"`
	Variation string        `json:"variation"`
}

type VariableValue map[VariableValueType]any

type VariableValueType string
