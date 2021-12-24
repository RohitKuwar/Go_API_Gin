package models

type Goal struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	Status          string `json:"status"`
	AssignedTo      []string `json:"assignedTo"`
	AssignedBy      string `json:"assignedBy"`
	AssignedOn      string `json:"assignedOn"`
	CompletionAward string `json:"completionAward"`
}
