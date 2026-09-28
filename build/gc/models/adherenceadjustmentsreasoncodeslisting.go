package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceadjustmentsreasoncodeslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceadjustmentsreasoncodeslistingDud struct { 
    

}

// Adherenceadjustmentsreasoncodeslisting
type Adherenceadjustmentsreasoncodeslisting struct { 
    // Entities
    Entities []Adherenceadjustmentsreasoncode `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Adherenceadjustmentsreasoncodeslisting) String() string {
     o.Entities = []Adherenceadjustmentsreasoncode{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceadjustmentsreasoncodeslisting) MarshalJSON() ([]byte, error) {
    type Alias Adherenceadjustmentsreasoncodeslisting

    if AdherenceadjustmentsreasoncodeslistingMarshalled {
        return []byte("{}"), nil
    }
    AdherenceadjustmentsreasoncodeslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Adherenceadjustmentsreasoncode `json:"entities"`
        *Alias
    }{

        
        Entities: []Adherenceadjustmentsreasoncode{{}},
        

        Alias: (*Alias)(u),
    })
}

