package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SourceentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SourceentityDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Sourceentity
type Sourceentity struct { 
    


    // VarType - The type of the source entity
    VarType string `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Sourceentity) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sourceentity) MarshalJSON() ([]byte, error) {
    type Alias Sourceentity

    if SourceentityMarshalled {
        return []byte("{}"), nil
    }
    SourceentityMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

