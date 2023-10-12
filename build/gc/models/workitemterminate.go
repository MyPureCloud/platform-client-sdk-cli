package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemterminateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemterminateDud struct { 
    

}

// Workitemterminate
type Workitemterminate struct { 
    // StatusId - The ID of the status the workitem should be updated to when terminating. The status must be a 'Closed' category status.
    StatusId string `json:"statusId"`

}

// String returns a JSON representation of the model
func (o *Workitemterminate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemterminate) MarshalJSON() ([]byte, error) {
    type Alias Workitemterminate

    if WorkitemterminateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemterminateMarshalled = true

    return json.Marshal(&struct {
        
        StatusId string `json:"statusId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

