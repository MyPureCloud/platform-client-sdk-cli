package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlertsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlertsummaryDud struct { 
    


    


    


    

}

// Alertsummary
type Alertsummary struct { 
    // Entities - The entities who violated the rule condition over the duration of the alert.
    Entities []Alertsummaryentity `json:"entities"`


    // Conversation - The id of the conversation that triggered the alert.  Only used for alerts based on instance-based conversation metrics.
    Conversation Addressableentityref `json:"conversation"`


    // MetricType - The metric type that is monitored.
    MetricType string `json:"metricType"`


    // EntitiesAreTeamMembers - Flag that indicated whether or not the alert is for a rule with a condition for all members of a team.
    EntitiesAreTeamMembers bool `json:"entitiesAreTeamMembers"`

}

// String returns a JSON representation of the model
func (o *Alertsummary) String() string {
     o.Entities = []Alertsummaryentity{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alertsummary) MarshalJSON() ([]byte, error) {
    type Alias Alertsummary

    if AlertsummaryMarshalled {
        return []byte("{}"), nil
    }
    AlertsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Alertsummaryentity `json:"entities"`
        
        Conversation Addressableentityref `json:"conversation"`
        
        MetricType string `json:"metricType"`
        
        EntitiesAreTeamMembers bool `json:"entitiesAreTeamMembers"`
        *Alias
    }{

        
        Entities: []Alertsummaryentity{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

