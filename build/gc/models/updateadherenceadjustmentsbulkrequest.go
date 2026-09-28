package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateadherenceadjustmentsbulkrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateadherenceadjustmentsbulkrequestDud struct { 
    

}

// Updateadherenceadjustmentsbulkrequest
type Updateadherenceadjustmentsbulkrequest struct { 
    // Adjustments - The adherence adjustments to update
    Adjustments []Updateadherenceadjustmentsbulkitem `json:"adjustments"`

}

// String returns a JSON representation of the model
func (o *Updateadherenceadjustmentsbulkrequest) String() string {
     o.Adjustments = []Updateadherenceadjustmentsbulkitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateadherenceadjustmentsbulkrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateadherenceadjustmentsbulkrequest

    if UpdateadherenceadjustmentsbulkrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateadherenceadjustmentsbulkrequestMarshalled = true

    return json.Marshal(&struct {
        
        Adjustments []Updateadherenceadjustmentsbulkitem `json:"adjustments"`
        *Alias
    }{

        
        Adjustments: []Updateadherenceadjustmentsbulkitem{{}},
        

        Alias: (*Alias)(u),
    })
}

