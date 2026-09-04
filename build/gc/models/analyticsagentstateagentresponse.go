package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsagentstateagentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsagentstateagentresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Analyticsagentstateagentresponse
type Analyticsagentstateagentresponse struct { 
    // UserId - User Id - only returned if division is covered by agentStateNames permission
    UserId string `json:"userId"`


    // DivisionId - Division Id
    DivisionId string `json:"divisionId"`


    // UserName - User name - only returned if division is covered by agentStateNames permission
    UserName string `json:"userName"`


    // ManagerId - The user that this user reports to
    ManagerId string `json:"managerId"`


    // SessionCount - The count of sessions
    SessionCount int `json:"sessionCount"`


    // Sessions - List of sessions
    Sessions []Analyticsagentstateagentsessionresult `json:"sessions"`


    // SystemPresence - The user's system presence
    SystemPresence string `json:"systemPresence"`


    // OrganizationPresenceId - The identifier for the user's organization presence
    OrganizationPresenceId string `json:"organizationPresenceId"`


    // PresenceDate - The timestamp for when the user's presence began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    PresenceDate time.Time `json:"presenceDate"`


    // RoutingStatus - The user's routing status
    RoutingStatus string `json:"routingStatus"`


    // RoutingStatusDate - The timestamp for when the user's routing status began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RoutingStatusDate time.Time `json:"routingStatusDate"`


    // IsOutOfOffice - Whether the user is out of office
    IsOutOfOffice bool `json:"isOutOfOffice"`


    // ManagementUnitId - The id of the user's management unit
    ManagementUnitId string `json:"managementUnitId"`


    // BusinessUnitId - The id of the user's business unit
    BusinessUnitId string `json:"businessUnitId"`


    // AdherenceState - The user's adherence state
    AdherenceState string `json:"adherenceState"`


    // AdherenceImpact - The user's adherence impact
    AdherenceImpact string `json:"adherenceImpact"`


    // AdherenceDate - The timestamp for when the user's adherence state began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    AdherenceDate time.Time `json:"adherenceDate"`


    // ScheduledActivityCodeId - The id of the user's scheduled activity code
    ScheduledActivityCodeId string `json:"scheduledActivityCodeId"`


    // ScheduledActivityCategory - The user's scheduled activity category
    ScheduledActivityCategory string `json:"scheduledActivityCategory"`


    // ActualActivityCategory - The user's actual activity category
    ActualActivityCategory string `json:"actualActivityCategory"`

}

// String returns a JSON representation of the model
func (o *Analyticsagentstateagentresponse) String() string {
    
    
    
    
    
     o.Sessions = []Analyticsagentstateagentsessionresult{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsagentstateagentresponse) MarshalJSON() ([]byte, error) {
    type Alias Analyticsagentstateagentresponse

    if AnalyticsagentstateagentresponseMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsagentstateagentresponseMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        DivisionId string `json:"divisionId"`
        
        UserName string `json:"userName"`
        
        ManagerId string `json:"managerId"`
        
        SessionCount int `json:"sessionCount"`
        
        Sessions []Analyticsagentstateagentsessionresult `json:"sessions"`
        
        SystemPresence string `json:"systemPresence"`
        
        OrganizationPresenceId string `json:"organizationPresenceId"`
        
        PresenceDate time.Time `json:"presenceDate"`
        
        RoutingStatus string `json:"routingStatus"`
        
        RoutingStatusDate time.Time `json:"routingStatusDate"`
        
        IsOutOfOffice bool `json:"isOutOfOffice"`
        
        ManagementUnitId string `json:"managementUnitId"`
        
        BusinessUnitId string `json:"businessUnitId"`
        
        AdherenceState string `json:"adherenceState"`
        
        AdherenceImpact string `json:"adherenceImpact"`
        
        AdherenceDate time.Time `json:"adherenceDate"`
        
        ScheduledActivityCodeId string `json:"scheduledActivityCodeId"`
        
        ScheduledActivityCategory string `json:"scheduledActivityCategory"`
        
        ActualActivityCategory string `json:"actualActivityCategory"`
        *Alias
    }{

        


        


        


        


        


        
        Sessions: []Analyticsagentstateagentsessionresult{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

