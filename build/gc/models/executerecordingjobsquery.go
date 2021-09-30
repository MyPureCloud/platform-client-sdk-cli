package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExecuterecordingjobsqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExecuterecordingjobsqueryDud struct { 
    

}

// Executerecordingjobsquery
type Executerecordingjobsquery struct { 
    // State - The desired state for the job to be set to.
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Executerecordingjobsquery) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Executerecordingjobsquery) MarshalJSON() ([]byte, error) {
    type Alias Executerecordingjobsquery

    if ExecuterecordingjobsqueryMarshalled {
        return []byte("{}"), nil
    }
    ExecuterecordingjobsqueryMarshalled = true

    return json.Marshal(&struct { 
        State string `json:"state"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

