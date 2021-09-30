package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClientappconfigurationinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClientappconfigurationinfoDud struct { 
    Current Integrationconfiguration `json:"current"`


    Effective Effectiveconfiguration `json:"effective"`

}

// Clientappconfigurationinfo - Configuration information for the integration
type Clientappconfigurationinfo struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Clientappconfigurationinfo) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clientappconfigurationinfo) MarshalJSON() ([]byte, error) {
    type Alias Clientappconfigurationinfo

    if ClientappconfigurationinfoMarshalled {
        return []byte("{}"), nil
    }
    ClientappconfigurationinfoMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

