package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceadjustmentslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceadjustmentslistingDud struct { 
    

}

// Adherenceadjustmentslisting
type Adherenceadjustmentslisting struct { 
    // Entities
    Entities []Adherenceadjustment `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Adherenceadjustmentslisting) String() string {
     o.Entities = []Adherenceadjustment{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceadjustmentslisting) MarshalJSON() ([]byte, error) {
    type Alias Adherenceadjustmentslisting

    if AdherenceadjustmentslistingMarshalled {
        return []byte("{}"), nil
    }
    AdherenceadjustmentslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Adherenceadjustment `json:"entities"`
        *Alias
    }{

        
        Entities: []Adherenceadjustment{{}},
        

        Alias: (*Alias)(u),
    })
}

