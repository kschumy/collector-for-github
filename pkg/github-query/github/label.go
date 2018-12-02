package github

import "time"

const LabelsAsString = "labels(first:100){totalCount,nodes{id,name,updatedAt}}"

//
type LabelsQueryResults struct {
	TotalCount int      `json:"totalCount"`
	Labels     []*Label `json:"nodes"`
}

//
type Label struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetNames returns a list of label names for provided LabelsQueryResults.
func (labels *LabelsQueryResults) GetNames() []string {
	labelNames := []string{}
	for _, label := range labels.Labels {
		labelNames = append(labelNames, label.Name)
	}
	return labelNames
}
