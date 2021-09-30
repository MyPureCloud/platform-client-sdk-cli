package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VmpairinginfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VmpairinginfoDud struct { 
    


    


    


    

}

// Vmpairinginfo
type Vmpairinginfo struct { 
    // MetaData - This is to be used to complete the setup process of a locally deployed virtual edge device.
    MetaData Metadata `json:"meta-data"`


    // EdgeId
    EdgeId string `json:"edge-id"`


    // AuthToken
    AuthToken string `json:"auth-token"`


    // OrgId
    OrgId string `json:"org-id"`

}

// String returns a JSON representation of the model
func (o *Vmpairinginfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Vmpairinginfo) MarshalJSON() ([]byte, error) {
    type Alias Vmpairinginfo

    if VmpairinginfoMarshalled {
        return []byte("{}"), nil
    }
    VmpairinginfoMarshalled = true

    return json.Marshal(&struct { 
        MetaData Metadata `json:"meta-data"`
        
        EdgeId string `json:"edge-id"`
        
        AuthToken string `json:"auth-token"`
        
        OrgId string `json:"org-id"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

