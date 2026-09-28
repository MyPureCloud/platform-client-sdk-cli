package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateadherenceadjustmentsreasoncodesbulkrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateadherenceadjustmentsreasoncodesbulkrequestDud struct { 
    

}

// Updateadherenceadjustmentsreasoncodesbulkrequest
type Updateadherenceadjustmentsreasoncodesbulkrequest struct { 
    // ReasonCodes - The reason codes to update
    ReasonCodes []Updateadherenceadjustmentsreasoncodesbulkitem `json:"reasonCodes"`

}

// String returns a JSON representation of the model
func (o *Updateadherenceadjustmentsreasoncodesbulkrequest) String() string {
     o.ReasonCodes = []Updateadherenceadjustmentsreasoncodesbulkitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateadherenceadjustmentsreasoncodesbulkrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateadherenceadjustmentsreasoncodesbulkrequest

    if UpdateadherenceadjustmentsreasoncodesbulkrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateadherenceadjustmentsreasoncodesbulkrequestMarshalled = true

    return json.Marshal(&struct {
        
        ReasonCodes []Updateadherenceadjustmentsreasoncodesbulkitem `json:"reasonCodes"`
        *Alias
    }{

        
        ReasonCodes: []Updateadherenceadjustmentsreasoncodesbulkitem{{}},
        

        Alias: (*Alias)(u),
    })
}

