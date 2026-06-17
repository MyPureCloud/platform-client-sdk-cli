package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenteffectivebidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenteffectivebidDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agenteffectivebid
type Agenteffectivebid struct { 
    // Id - The ID of the schedule bid
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // EffectiveDate - The effective date of the bid relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EffectiveDate time.Time `json:"effectiveDate"`


    // EndDate - The end date of the bid, relative to the business unit time zone in yyyy-MM-dd format. Null denotes an active schedule bid. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndDate time.Time `json:"endDate"`


    // DownloadUrl - The download URL to fetch the list of schedule sets and the agents assigned to them
    DownloadUrl string `json:"downloadUrl"`


    // DownloadTemplate - This field will always be null. Effective schedule sets are returned through the download URL. The schema is included here for documentation purposes
    DownloadTemplate Agentassignedschedulesetlist `json:"downloadTemplate"`


    

}

// String returns a JSON representation of the model
func (o *Agenteffectivebid) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenteffectivebid) MarshalJSON() ([]byte, error) {
    type Alias Agenteffectivebid

    if AgenteffectivebidMarshalled {
        return []byte("{}"), nil
    }
    AgenteffectivebidMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        EffectiveDate time.Time `json:"effectiveDate"`
        
        EndDate time.Time `json:"endDate"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadTemplate Agentassignedschedulesetlist `json:"downloadTemplate"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

