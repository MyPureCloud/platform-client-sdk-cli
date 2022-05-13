package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Scimv2patchoperationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Scimv2patchoperationDud struct { 
    


    


    

}

// Scimv2patchoperation - Defines a SCIM PATCH operation. The path and value follow very specific rules based on operation types. See section 3.5.2 \"Modifying with PATCH\" in RFC 7644 for details.
type Scimv2patchoperation struct { 
    // Op - The PATCH operation to perform.
    Op string `json:"op"`


    // Path - The attribute path that describes the target of the operation. Required for a \"remove\" operation.
    Path string `json:"path"`


    // Value - The value to set in the path.
    Value interface{} `json:"value"`

}

// String returns a JSON representation of the model
func (o *Scimv2patchoperation) String() string {
    
    
     o.Value = Interface{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimv2patchoperation) MarshalJSON() ([]byte, error) {
    type Alias Scimv2patchoperation

    if Scimv2patchoperationMarshalled {
        return []byte("{}"), nil
    }
    Scimv2patchoperationMarshalled = true

    return json.Marshal(&struct {
        
        Op string `json:"op"`
        
        Path string `json:"path"`
        
        Value interface{} `json:"value"`
        *Alias
    }{

        


        


        
        Value: Interface{},
        

        Alias: (*Alias)(u),
    })
}

