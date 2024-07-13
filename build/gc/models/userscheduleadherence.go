package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserscheduleadherenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserscheduleadherenceDud struct { 
    Id string `json:"id"`


    


    User Userreference `json:"user"`


    BusinessUnit Businessunitreference `json:"businessUnit"`


    ManagementUnit Managementunitreference `json:"managementUnit"`


    Team Teamreference `json:"team"`


    ScheduledActivityCategory string `json:"scheduledActivityCategory"`


    ScheduledActivityCode Activitycodesummary `json:"scheduledActivityCode"`


    SystemPresence string `json:"systemPresence"`


    OrganizationSecondaryPresenceId string `json:"organizationSecondaryPresenceId"`


    RoutingStatus string `json:"routingStatus"`


    ActualActivityCategory string `json:"actualActivityCategory"`


    IsOutOfOffice bool `json:"isOutOfOffice"`


    AdherenceState string `json:"adherenceState"`


    Impact string `json:"impact"`


    AdherenceExplanation Realtimeadherenceexplanation `json:"adherenceExplanation"`


    TimeOfAdherenceChange time.Time `json:"timeOfAdherenceChange"`


    PresenceUpdateTime time.Time `json:"presenceUpdateTime"`


    ActiveQueues []Queuereference `json:"activeQueues"`


    ActiveQueuesModifiedTime time.Time `json:"activeQueuesModifiedTime"`


    RemovedFromManagementUnit bool `json:"removedFromManagementUnit"`


    SelfUri string `json:"selfUri"`

}

// Userscheduleadherence
type Userscheduleadherence struct { 
    


    // Name
    Name string `json:"name"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Userscheduleadherence) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userscheduleadherence) MarshalJSON() ([]byte, error) {
    type Alias Userscheduleadherence

    if UserscheduleadherenceMarshalled {
        return []byte("{}"), nil
    }
    UserscheduleadherenceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

