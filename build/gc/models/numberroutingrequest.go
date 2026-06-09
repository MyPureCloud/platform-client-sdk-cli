package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NumberroutingrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NumberroutingrequestDud struct { 
    


    


    

}

// Numberroutingrequest
type Numberroutingrequest struct { 
    // OrganizationId - Target organization Id where number will be routed to
    OrganizationId string `json:"organizationId"`


    // NumberId - Number Id to be routed.
    NumberId string `json:"numberId"`


    // TargetRegion - Region where target organization is hosted.
    TargetRegion string `json:"targetRegion"`

}

// String returns a JSON representation of the model
func (o *Numberroutingrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Numberroutingrequest) MarshalJSON() ([]byte, error) {
    type Alias Numberroutingrequest

    if NumberroutingrequestMarshalled {
        return []byte("{}"), nil
    }
    NumberroutingrequestMarshalled = true

    return json.Marshal(&struct {
        
        OrganizationId string `json:"organizationId"`
        
        NumberId string `json:"numberId"`
        
        TargetRegion string `json:"targetRegion"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

