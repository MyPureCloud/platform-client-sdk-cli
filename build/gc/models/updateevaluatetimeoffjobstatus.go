package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateevaluatetimeoffjobstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateevaluatetimeoffjobstatusDud struct { 
    

}

// Updateevaluatetimeoffjobstatus
type Updateevaluatetimeoffjobstatus struct { 
    // Status - The status of the job
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Updateevaluatetimeoffjobstatus) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateevaluatetimeoffjobstatus) MarshalJSON() ([]byte, error) {
    type Alias Updateevaluatetimeoffjobstatus

    if UpdateevaluatetimeoffjobstatusMarshalled {
        return []byte("{}"), nil
    }
    UpdateevaluatetimeoffjobstatusMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

