package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternaleventexternalidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternaleventexternalidDud struct { 
    


    

}

// Externaleventexternalid
type Externaleventexternalid struct { 
    // Value - The value of the identifier.
    Value string `json:"value"`


    // ExternalSourceId - The id of the external source.
    ExternalSourceId string `json:"externalSourceId"`

}

// String returns a JSON representation of the model
func (o *Externaleventexternalid) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externaleventexternalid) MarshalJSON() ([]byte, error) {
    type Alias Externaleventexternalid

    if ExternaleventexternalidMarshalled {
        return []byte("{}"), nil
    }
    ExternaleventexternalidMarshalled = true

    return json.Marshal(&struct {
        
        Value string `json:"value"`
        
        ExternalSourceId string `json:"externalSourceId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

