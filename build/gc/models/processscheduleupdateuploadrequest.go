package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProcessscheduleupdateuploadrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProcessscheduleupdateuploadrequestDud struct { 
    


    


    

}

// Processscheduleupdateuploadrequest
type Processscheduleupdateuploadrequest struct { 
    // UploadKey - The uploadKey provided by the request to get an upload URL
    UploadKey string `json:"uploadKey"`


    // TeamIds - The list of teams to which the users being modified belong. Only required if the requesting user has conditional permission to wfm:schedule:edit
    TeamIds []string `json:"teamIds"`


    // ManagementUnitIdsForAddedTeamUsers - The set of muIds to which agents belong if agents are being newly added to the schedule, if the requesting user has conditional permission to wfm:schedule:edit
    ManagementUnitIdsForAddedTeamUsers []string `json:"managementUnitIdsForAddedTeamUsers"`

}

// String returns a JSON representation of the model
func (o *Processscheduleupdateuploadrequest) String() string {
    
    
    
    
    
    
     o.TeamIds = []string{""} 
    
    
    
     o.ManagementUnitIdsForAddedTeamUsers = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Processscheduleupdateuploadrequest) MarshalJSON() ([]byte, error) {
    type Alias Processscheduleupdateuploadrequest

    if ProcessscheduleupdateuploadrequestMarshalled {
        return []byte("{}"), nil
    }
    ProcessscheduleupdateuploadrequestMarshalled = true

    return json.Marshal(&struct { 
        UploadKey string `json:"uploadKey"`
        
        TeamIds []string `json:"teamIds"`
        
        ManagementUnitIdsForAddedTeamUsers []string `json:"managementUnitIdsForAddedTeamUsers"`
        
        *Alias
    }{
        

        

        

        
        TeamIds: []string{""},
        

        

        
        ManagementUnitIdsForAddedTeamUsers: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

