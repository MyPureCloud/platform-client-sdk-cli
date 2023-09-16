package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlertsummaryentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlertsummaryentityDud struct { 
    


    


    


    


    


    

}

// Alertsummaryentity
type Alertsummaryentity struct { 
    // EntityType - Specifies the type of entity being evaluated
    EntityType string `json:"entityType"`


    // User - User id of the entity being monitored
    User Addressableentityref `json:"user"`


    // Group - Group id of the entity being monitored
    Group Addressableentityref `json:"group"`


    // Queue - Queue id of the entity being monitored
    Queue Addressableentityref `json:"queue"`


    // Team - Team id of the entity being monitored
    Team Addressableentityref `json:"team"`


    // Alerting - Flag that indicated if the entity is current causing the alert to be triggered
    Alerting bool `json:"alerting"`

}

// String returns a JSON representation of the model
func (o *Alertsummaryentity) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alertsummaryentity) MarshalJSON() ([]byte, error) {
    type Alias Alertsummaryentity

    if AlertsummaryentityMarshalled {
        return []byte("{}"), nil
    }
    AlertsummaryentityMarshalled = true

    return json.Marshal(&struct {
        
        EntityType string `json:"entityType"`
        
        User Addressableentityref `json:"user"`
        
        Group Addressableentityref `json:"group"`
        
        Queue Addressableentityref `json:"queue"`
        
        Team Addressableentityref `json:"team"`
        
        Alerting bool `json:"alerting"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

