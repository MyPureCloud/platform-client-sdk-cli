package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsprovisioningstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsprovisioningstatusDud struct { 
    


    


    


    

}

// Smsprovisioningstatus
type Smsprovisioningstatus struct { 
    // Action - Provisioning action
    Action string `json:"action"`


    // State - Provisioning state
    State string `json:"state"`


    // VarError - Any error associated with a Failed state
    VarError Errorbody `json:"error"`


    // Version - The phone number version associated with the provisioning action
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Smsprovisioningstatus) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsprovisioningstatus) MarshalJSON() ([]byte, error) {
    type Alias Smsprovisioningstatus

    if SmsprovisioningstatusMarshalled {
        return []byte("{}"), nil
    }
    SmsprovisioningstatusMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        
        State string `json:"state"`
        
        VarError Errorbody `json:"error"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

