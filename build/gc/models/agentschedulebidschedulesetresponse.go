package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentschedulebidschedulesetresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentschedulebidschedulesetresponseDud struct { 
    


    

}

// Agentschedulebidschedulesetresponse
type Agentschedulebidschedulesetresponse struct { 
    // DownloadUrl - The download URL to fetch the schedule set of the bid group to which the agent belongs
    DownloadUrl string `json:"downloadUrl"`


    // DownloadTemplate - Schedule sets always come through downloadUrl, the schema included here is just for documentation
    DownloadTemplate Bidgroupscheduleset `json:"downloadTemplate"`

}

// String returns a JSON representation of the model
func (o *Agentschedulebidschedulesetresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentschedulebidschedulesetresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentschedulebidschedulesetresponse

    if AgentschedulebidschedulesetresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentschedulebidschedulesetresponseMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadTemplate Bidgroupscheduleset `json:"downloadTemplate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

