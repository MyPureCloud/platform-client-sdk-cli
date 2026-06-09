package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateorganizationlinkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateorganizationlinkDud struct { 
    


    

}

// Createorganizationlink
type Createorganizationlink struct { 
    // TargetOrganizationId - Id for the linking organization.
    TargetOrganizationId string `json:"targetOrganizationId"`


    // TargetRegion - Region where target organization is hosted.
    TargetRegion string `json:"targetRegion"`

}

// String returns a JSON representation of the model
func (o *Createorganizationlink) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createorganizationlink) MarshalJSON() ([]byte, error) {
    type Alias Createorganizationlink

    if CreateorganizationlinkMarshalled {
        return []byte("{}"), nil
    }
    CreateorganizationlinkMarshalled = true

    return json.Marshal(&struct {
        
        TargetOrganizationId string `json:"targetOrganizationId"`
        
        TargetRegion string `json:"targetRegion"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

