package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationlinkresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationlinkresponseDud struct { 
    


    


    


    


    


    

}

// Organizationlinkresponse
type Organizationlinkresponse struct { 
    // SourceOrganizationId - Organization Id for the login organization.
    SourceOrganizationId string `json:"sourceOrganizationId"`


    // TargetOrganizationId - Organization Id for the linking organization.
    TargetOrganizationId string `json:"targetOrganizationId"`


    // SourceRegion - Region where context organization is hosted, ie. us-east-1
    SourceRegion string `json:"sourceRegion"`


    // TargetRegion - Region where linking organization is hosted, ie. us-east-2
    TargetRegion string `json:"targetRegion"`


    // TargetName - Name for the linking organization.
    TargetName string `json:"targetName"`


    // Status - Status of the linking.
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Organizationlinkresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationlinkresponse) MarshalJSON() ([]byte, error) {
    type Alias Organizationlinkresponse

    if OrganizationlinkresponseMarshalled {
        return []byte("{}"), nil
    }
    OrganizationlinkresponseMarshalled = true

    return json.Marshal(&struct {
        
        SourceOrganizationId string `json:"sourceOrganizationId"`
        
        TargetOrganizationId string `json:"targetOrganizationId"`
        
        SourceRegion string `json:"sourceRegion"`
        
        TargetRegion string `json:"targetRegion"`
        
        TargetName string `json:"targetName"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

