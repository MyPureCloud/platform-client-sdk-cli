package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasepriorityupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasepriorityupdateDud struct { 
    

}

// Casepriorityupdate
type Casepriorityupdate struct { 
    // Priority - The priority of the Case.
    Priority string `json:"priority"`

}

// String returns a JSON representation of the model
func (o *Casepriorityupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casepriorityupdate) MarshalJSON() ([]byte, error) {
    type Alias Casepriorityupdate

    if CasepriorityupdateMarshalled {
        return []byte("{}"), nil
    }
    CasepriorityupdateMarshalled = true

    return json.Marshal(&struct {
        
        Priority string `json:"priority"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

