package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalorganizationidentifierclaimrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalorganizationidentifierclaimrequestDud struct { 
    


    

}

// Externalorganizationidentifierclaimrequest
type Externalorganizationidentifierclaimrequest struct { 
    // Operation - The operation to perform claim/release
    Operation string `json:"operation"`


    // Identifier - The identifier that should be claimed/released from an external org
    Identifier Externalorganizationidentifier `json:"identifier"`

}

// String returns a JSON representation of the model
func (o *Externalorganizationidentifierclaimrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalorganizationidentifierclaimrequest) MarshalJSON() ([]byte, error) {
    type Alias Externalorganizationidentifierclaimrequest

    if ExternalorganizationidentifierclaimrequestMarshalled {
        return []byte("{}"), nil
    }
    ExternalorganizationidentifierclaimrequestMarshalled = true

    return json.Marshal(&struct {
        
        Operation string `json:"operation"`
        
        Identifier Externalorganizationidentifier `json:"identifier"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

