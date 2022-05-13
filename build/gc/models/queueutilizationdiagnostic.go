package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueueutilizationdiagnosticMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueueutilizationdiagnosticDud struct { 
    Queue Domainentityref `json:"queue"`


    UsersInQueue int `json:"usersInQueue"`


    ActiveUsersInQueue int `json:"activeUsersInQueue"`


    UsersOnQueue int `json:"usersOnQueue"`


    UsersNotUtilized int `json:"usersNotUtilized"`


    UsersOnQueueWithStation int `json:"usersOnQueueWithStation"`


    UsersOnACampaignCall int `json:"usersOnACampaignCall"`


    UsersOnDifferentEdgeGroup int `json:"usersOnDifferentEdgeGroup"`


    UsersOnANonCampaignCall int `json:"usersOnANonCampaignCall"`

}

// Queueutilizationdiagnostic
type Queueutilizationdiagnostic struct { 
    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Queueutilizationdiagnostic) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queueutilizationdiagnostic) MarshalJSON() ([]byte, error) {
    type Alias Queueutilizationdiagnostic

    if QueueutilizationdiagnosticMarshalled {
        return []byte("{}"), nil
    }
    QueueutilizationdiagnosticMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

