package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LabelentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LabelentityDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Labelentity
type Labelentity struct { 
    // Id - The Id of the label.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Labelentity) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Labelentity) MarshalJSON() ([]byte, error) {
    type Alias Labelentity

    if LabelentityMarshalled {
        return []byte("{}"), nil
    }
    LabelentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

