package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimerrorDud struct { 
    Schemas []string `json:"schemas"`


    Status string `json:"status"`


    ScimType string `json:"scimType"`


    Detail string `json:"detail"`

}

// Scimerror - Defines a SCIM error.
type Scimerror struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimerror) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimerror) MarshalJSON() ([]byte, error) {
    type Alias Scimerror

    if ScimerrorMarshalled {
        return []byte("{}"), nil
    }
    ScimerrorMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

