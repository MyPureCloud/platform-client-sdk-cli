package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IdentifierclaimrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IdentifierclaimrequestDud struct { 
    


    

}

// Identifierclaimrequest
type Identifierclaimrequest struct { 
    // Operation - The operation to perform claim/release
    Operation string `json:"operation"`


    // Identifier - The identifier that should be claimed/released from a contact
    Identifier Contactidentifier `json:"identifier"`

}

// String returns a JSON representation of the model
func (o *Identifierclaimrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Identifierclaimrequest) MarshalJSON() ([]byte, error) {
    type Alias Identifierclaimrequest

    if IdentifierclaimrequestMarshalled {
        return []byte("{}"), nil
    }
    IdentifierclaimrequestMarshalled = true

    return json.Marshal(&struct {
        
        Operation string `json:"operation"`
        
        Identifier Contactidentifier `json:"identifier"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

