package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamactivityentitydataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamactivityentitydataDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Teamactivityentitydata
type Teamactivityentitydata struct { 
    // ActivityDate - The time at which the activity was observed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ActivityDate time.Time `json:"activityDate"`


    // OrganizationPresenceId - Organization presence identifier
    OrganizationPresenceId string `json:"organizationPresenceId"`


    // PresenceDate - Date of the latest presence change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    PresenceDate time.Time `json:"presenceDate"`


    // QueueId - Queue identifier
    QueueId string `json:"queueId"`


    // QueueMembershipStatus - Queue membership status (e.g. active or inactive)
    QueueMembershipStatus string `json:"queueMembershipStatus"`


    // RoutingStatus - Agent routing status
    RoutingStatus string `json:"routingStatus"`


    // RoutingStatusDate - Date of the latest routing status change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RoutingStatusDate time.Time `json:"routingStatusDate"`


    // SystemPresence - System presence
    SystemPresence string `json:"systemPresence"`


    // TeamId - The team ID the user is a member of
    TeamId string `json:"teamId"`


    // UserId - Unique identifier for the user
    UserId string `json:"userId"`

}

// String returns a JSON representation of the model
func (o *Teamactivityentitydata) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamactivityentitydata) MarshalJSON() ([]byte, error) {
    type Alias Teamactivityentitydata

    if TeamactivityentitydataMarshalled {
        return []byte("{}"), nil
    }
    TeamactivityentitydataMarshalled = true

    return json.Marshal(&struct {
        
        ActivityDate time.Time `json:"activityDate"`
        
        OrganizationPresenceId string `json:"organizationPresenceId"`
        
        PresenceDate time.Time `json:"presenceDate"`
        
        QueueId string `json:"queueId"`
        
        QueueMembershipStatus string `json:"queueMembershipStatus"`
        
        RoutingStatus string `json:"routingStatus"`
        
        RoutingStatusDate time.Time `json:"routingStatusDate"`
        
        SystemPresence string `json:"systemPresence"`
        
        TeamId string `json:"teamId"`
        
        UserId string `json:"userId"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

