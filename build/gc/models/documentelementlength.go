package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentelementlengthMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentelementlengthDud struct { 
    


    

}

// Documentelementlength
type Documentelementlength struct { 
    // Value - The length value of the element in the selected unit.
    Value float32 `json:"value"`


    // Unit - The unit of length.
    Unit string `json:"unit"`

}

// String returns a JSON representation of the model
func (o *Documentelementlength) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentelementlength) MarshalJSON() ([]byte, error) {
    type Alias Documentelementlength

    if DocumentelementlengthMarshalled {
        return []byte("{}"), nil
    }
    DocumentelementlengthMarshalled = true

    return json.Marshal(&struct {
        
        Value float32 `json:"value"`
        
        Unit string `json:"unit"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

