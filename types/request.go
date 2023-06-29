package types

type EvaluateFeatureRequest struct {
	EntityID        string `json:"entityId"`
	EvaluateContext string `json:"evaluateContext"`
}
