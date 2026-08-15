package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManagedeleteprotectionresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManagedeleteprotectionresultDud struct { 
    

}

// Managedeleteprotectionresult
type Managedeleteprotectionresult struct { 
    // FailedUpdates - List of failed delete protection status updates
    FailedUpdates []Faileddeleteprotectionupdate `json:"failedUpdates"`

}

// String returns a JSON representation of the model
func (o *Managedeleteprotectionresult) String() string {
     o.FailedUpdates = []Faileddeleteprotectionupdate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Managedeleteprotectionresult) MarshalJSON() ([]byte, error) {
    type Alias Managedeleteprotectionresult

    if ManagedeleteprotectionresultMarshalled {
        return []byte("{}"), nil
    }
    ManagedeleteprotectionresultMarshalled = true

    return json.Marshal(&struct {
        
        FailedUpdates []Faileddeleteprotectionupdate `json:"failedUpdates"`
        *Alias
    }{

        
        FailedUpdates: []Faileddeleteprotectionupdate{{}},
        

        Alias: (*Alias)(u),
    })
}

