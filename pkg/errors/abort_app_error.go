package errors

import "fmt"

type StopExecution struct {
	Message string
}

func (u StopExecution) Error() string {
	return fmt.Sprintf("AbortAppError: Aborting process due to unrecoverable error : %s", u.Message)
}
