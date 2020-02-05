package consensus

import (
	"fmt"
	"time"
)

type ConsensusError struct {
	Time    time.Time
	Message string
}

func (e *ConsensusError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.Time, e.Message)
}
