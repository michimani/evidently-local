package types

type EvaluateFeatureRequest struct {
	EntityID        string `json:"entityId"`
	EvaluateContext string `json:"evaluateContext"`
}

type BatchEvaluateFeatureRequest struct {
	Requests []EvaluationRequest `json:"requests"`
}

type EvaluationRequest struct {
	EntityID          string `json:"entityId"`
	EvaluationContext string `json:"evaluationContext"`
	Feature           string `json:"feature"`
}
