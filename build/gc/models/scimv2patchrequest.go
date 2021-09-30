package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Scimv2patchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Scimv2patchrequestDud struct { 
    


    

}

// Scimv2patchrequest - Defines a SCIM PATCH request. See section 3.5.2 \"Modifying with PATCH\" in RFC 7644 for details.
type Scimv2patchrequest struct { 
    // Schemas - The list of schemas used in the PATCH request.
    Schemas []string `json:"schemas"`


    // Operations - The list of operations to perform for the PATCH request.
    Operations []Scimv2patchoperation `json:"Operations"`

}

// String returns a JSON representation of the model
func (o *Scimv2patchrequest) String() string {
    
    
     o.Schemas = []string{""} 
    
    
    
     o.Operations = []Scimv2patchoperation{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimv2patchrequest) MarshalJSON() ([]byte, error) {
    type Alias Scimv2patchrequest

    if Scimv2patchrequestMarshalled {
        return []byte("{}"), nil
    }
    Scimv2patchrequestMarshalled = true

    return json.Marshal(&struct { 
        Schemas []string `json:"schemas"`
        
        Operations []Scimv2patchoperation `json:"Operations"`
        
        *Alias
    }{
        

        
        Schemas: []string{""},
        

        

        
        Operations: []Scimv2patchoperation{{}},
        

        
        Alias: (*Alias)(u),
    })
}

