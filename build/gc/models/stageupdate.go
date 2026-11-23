package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StageupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StageupdateDud struct { 
    

}

// Stageupdate
type Stageupdate struct { 
    // DateDue - The due date of the Stage. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateDue time.Time `json:"dateDue"`

}

// String returns a JSON representation of the model
func (o *Stageupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stageupdate) MarshalJSON() ([]byte, error) {
    type Alias Stageupdate

    if StageupdateMarshalled {
        return []byte("{}"), nil
    }
    StageupdateMarshalled = true

    return json.Marshal(&struct {
        
        DateDue time.Time `json:"dateDue"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

