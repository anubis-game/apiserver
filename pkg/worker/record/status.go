package record

const (
	Created Status = "created"
	Failure Status = "failure"
	Retried Status = "retried"
	Success Status = "success"
	Unknown Status = "unknown"
	Waiting Status = "waiting"
)

type Status string
