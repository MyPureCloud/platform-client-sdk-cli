package models
import (
	"time"
)
// String Used to prevent swagger schemas defined as {"type": "Date-Time"} from throwing errors
// when rendering the api's request body in api.mustache
type Time struct {

}

func (t *Time) String() string {
	return time.Time{}.String()
}