package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalidDud struct { 
    


    

}

// Externalid
type Externalid struct { 
    // ExternalSource - The external source of the identifier. Max: 255 characters. Leading and trailing whitespace stripped.
    ExternalSource Externalsource `json:"externalSource"`


    // Value - The string value of the identifier. Max: 255 characters. Leading and trailing whitespace stripped.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Externalid) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalid) MarshalJSON() ([]byte, error) {
    type Alias Externalid

    if ExternalidMarshalled {
        return []byte("{}"), nil
    }
    ExternalidMarshalled = true

    return json.Marshal(&struct {
        
        ExternalSource Externalsource `json:"externalSource"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

