package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InactivitytimeoutsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InactivitytimeoutsettingsDud struct { 
    


    


    

}

// Inactivitytimeoutsettings
type Inactivitytimeoutsettings struct { 
    // TimeoutSeconds - Timeout in seconds for inactivity on the interaction
    TimeoutSeconds int `json:"timeoutSeconds"`


    // ActionType - Action to take when timeout occurs
    ActionType string `json:"actionType"`


    // FlowId - Flow ID for architect flow action
    FlowId Domainentityref `json:"flowId"`

}

// String returns a JSON representation of the model
func (o *Inactivitytimeoutsettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Inactivitytimeoutsettings) MarshalJSON() ([]byte, error) {
    type Alias Inactivitytimeoutsettings

    if InactivitytimeoutsettingsMarshalled {
        return []byte("{}"), nil
    }
    InactivitytimeoutsettingsMarshalled = true

    return json.Marshal(&struct {
        
        TimeoutSeconds int `json:"timeoutSeconds"`
        
        ActionType string `json:"actionType"`
        
        FlowId Domainentityref `json:"flowId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

