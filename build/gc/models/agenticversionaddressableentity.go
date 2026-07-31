package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticversionaddressableentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticversionaddressableentityDud struct { 
    Version string `json:"version"`


    

}

// Agenticversionaddressableentity
type Agenticversionaddressableentity struct { 
    


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Agenticversionaddressableentity) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticversionaddressableentity) MarshalJSON() ([]byte, error) {
    type Alias Agenticversionaddressableentity

    if AgenticversionaddressableentityMarshalled {
        return []byte("{}"), nil
    }
    AgenticversionaddressableentityMarshalled = true

    return json.Marshal(&struct {
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

