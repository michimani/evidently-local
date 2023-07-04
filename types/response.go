package types

type EvaluateFeatureResponse struct {
	Details   string           `json:"details"`
	Reason    EvaluationReason `json:"reason"`
	Value     VariableValue    `json:"value"`
	Variation string           `json:"variation"`
}

type VariableValue map[VariableValueType]any

type BatchEvaluateFeatureResponse struct {
	Results []EvaluationResult `json:"results"`
}

type EvaluationResult struct {
	Details   string           `json:"details"`
	EntityID  string           `json:"entityId"`
	Feature   string           `json:"feature"`
	Project   string           `json:"project"`
	Reason    EvaluationReason `json:"reason"`
	Variation string           `json:"variation"`
	Value     VariableValue    `json:"value"`
}
