package github

const AssigneesAsString = "assignees(first:100){totalCount,nodes{id,login,name,company}}"

type Assignees struct {
	TotalCount   int        `json:"totalCount"`
	AssigneeList []Assignee `json:"nodes"`
}

type Assignee struct {
	ID      string `json:"id"`
	Login   string `json:"login"`
	Name    string `json:"name"`
	Company string `json:"company"`
}

//
//func (assignees Assignees) GetLogins() []string {
//	fmt.Printf("%+v", assignees)
//	assigneeNames := []string{}
//	for _, assignee := range assignees.AssigneeList {
//		assigneeNames = append(assigneeNames, assignee.Name)
//	}
//	return assigneeNames
//
//}
