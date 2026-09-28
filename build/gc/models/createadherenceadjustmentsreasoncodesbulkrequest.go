package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateadherenceadjustmentsreasoncodesbulkrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateadherenceadjustmentsreasoncodesbulkrequestDud struct { 
    

}

// Createadherenceadjustmentsreasoncodesbulkrequest
type Createadherenceadjustmentsreasoncodesbulkrequest struct { 
    // ReasonCodes - The reason codes to create
    ReasonCodes []Createadherenceadjustmentsreasoncoderequest `json:"reasonCodes"`

}

// String returns a JSON representation of the model
func (o *Createadherenceadjustmentsreasoncodesbulkrequest) String() string {
     o.ReasonCodes = []Createadherenceadjustmentsreasoncoderequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createadherenceadjustmentsreasoncodesbulkrequest) MarshalJSON() ([]byte, error) {
    type Alias Createadherenceadjustmentsreasoncodesbulkrequest

    if CreateadherenceadjustmentsreasoncodesbulkrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateadherenceadjustmentsreasoncodesbulkrequestMarshalled = true

    return json.Marshal(&struct {
        
        ReasonCodes []Createadherenceadjustmentsreasoncoderequest `json:"reasonCodes"`
        *Alias
    }{

        
        ReasonCodes: []Createadherenceadjustmentsreasoncoderequest{{}},
        

        Alias: (*Alias)(u),
    })
}

