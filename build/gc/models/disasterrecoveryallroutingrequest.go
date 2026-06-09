package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DisasterrecoveryallroutingrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DisasterrecoveryallroutingrequestDud struct { 
    


    


    

}

// Disasterrecoveryallroutingrequest - Disaster Recovery all numbers routing request body
type Disasterrecoveryallroutingrequest struct { 
    // SourceOrganizationId - Value for login Organization Id
    SourceOrganizationId string `json:"sourceOrganizationId"`


    // SwitchOrganizationId - Organization Id that will receive the routing
    SwitchOrganizationId string `json:"switchOrganizationId"`


    // TargetRegion - Region for rerouting
    TargetRegion string `json:"targetRegion"`

}

// String returns a JSON representation of the model
func (o *Disasterrecoveryallroutingrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Disasterrecoveryallroutingrequest) MarshalJSON() ([]byte, error) {
    type Alias Disasterrecoveryallroutingrequest

    if DisasterrecoveryallroutingrequestMarshalled {
        return []byte("{}"), nil
    }
    DisasterrecoveryallroutingrequestMarshalled = true

    return json.Marshal(&struct {
        
        SourceOrganizationId string `json:"sourceOrganizationId"`
        
        SwitchOrganizationId string `json:"switchOrganizationId"`
        
        TargetRegion string `json:"targetRegion"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

