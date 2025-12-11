package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChecklistfinalizepayloadMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChecklistfinalizepayloadDud struct { 
    

}

// Checklistfinalizepayload
type Checklistfinalizepayload struct { 
    // ExitReason - Exit reason provided at the time of finalizing the checklist.
    ExitReason string `json:"exitReason"`

}

// String returns a JSON representation of the model
func (o *Checklistfinalizepayload) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Checklistfinalizepayload) MarshalJSON() ([]byte, error) {
    type Alias Checklistfinalizepayload

    if ChecklistfinalizepayloadMarshalled {
        return []byte("{}"), nil
    }
    ChecklistfinalizepayloadMarshalled = true

    return json.Marshal(&struct {
        
        ExitReason string `json:"exitReason"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

