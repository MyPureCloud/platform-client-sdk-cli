package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VendorconnectionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VendorconnectionrequestDud struct { 
    


    


    

}

// Vendorconnectionrequest
type Vendorconnectionrequest struct { 
    // Publisher - Publisher of the integration or connector who registered the new connection. Typically, inin.
    Publisher string `json:"publisher"`


    // VarType - Integration or connector type that registered the new connection. Example, wfm-rta-integration
    VarType string `json:"type"`


    // Name - Name of the integration or connector instance that registered the new connection. Example, my-wfm
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Vendorconnectionrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Vendorconnectionrequest) MarshalJSON() ([]byte, error) {
    type Alias Vendorconnectionrequest

    if VendorconnectionrequestMarshalled {
        return []byte("{}"), nil
    }
    VendorconnectionrequestMarshalled = true

    return json.Marshal(&struct {
        
        Publisher string `json:"publisher"`
        
        VarType string `json:"type"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

