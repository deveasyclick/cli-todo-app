package models

type Status string

const (
	Pending    Status = "pending"
	InProgress Status = "inProgress"
	Completed  Status = "completed"
	Failed     Status = "failed"
)

func (s Status) IsValid() bool {
	switch s {
	case Pending, InProgress, Completed, Failed:
		return true
	default:
		return false
	}
}

func (s Status) String() string {
	return string(s)
}
