package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceadjustmentsreasoncodereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceadjustmentsreasoncodereferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Adherenceadjustmentsreasoncodereference
type Adherenceadjustmentsreasoncodereference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Adherenceadjustmentsreasoncodereference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceadjustmentsreasoncodereference) MarshalJSON() ([]byte, error) {
    type Alias Adherenceadjustmentsreasoncodereference

    if AdherenceadjustmentsreasoncodereferenceMarshalled {
        return []byte("{}"), nil
    }
    AdherenceadjustmentsreasoncodereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

