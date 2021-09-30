package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LicenseassignmentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LicenseassignmentrequestDud struct { 
    


    


    

}

// Licenseassignmentrequest
type Licenseassignmentrequest struct { 
    // LicenseId - The id of the license to assign/unassign.
    LicenseId string `json:"licenseId"`


    // UserIdsAdd - The ids of users to assign this license to.
    UserIdsAdd []string `json:"userIdsAdd"`


    // UserIdsRemove - The ids of users to unassign this license from.
    UserIdsRemove []string `json:"userIdsRemove"`

}

// String returns a JSON representation of the model
func (o *Licenseassignmentrequest) String() string {
    
    
    
    
    
    
     o.UserIdsAdd = []string{""} 
    
    
    
     o.UserIdsRemove = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Licenseassignmentrequest) MarshalJSON() ([]byte, error) {
    type Alias Licenseassignmentrequest

    if LicenseassignmentrequestMarshalled {
        return []byte("{}"), nil
    }
    LicenseassignmentrequestMarshalled = true

    return json.Marshal(&struct { 
        LicenseId string `json:"licenseId"`
        
        UserIdsAdd []string `json:"userIdsAdd"`
        
        UserIdsRemove []string `json:"userIdsRemove"`
        
        *Alias
    }{
        

        

        

        
        UserIdsAdd: []string{""},
        

        

        
        UserIdsRemove: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

