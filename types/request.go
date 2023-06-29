package types

type evaluateFeatureRequest struct {
	EntityID        string `json:"entityId"`
	EvaluateContext string `json:"evaluateContext"`
}
