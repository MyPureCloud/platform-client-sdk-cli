package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TriggerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TriggerDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Trigger - Defines a process automation trigger.
type Trigger struct { 
    


    // Name - The name of the trigger
    Name string `json:"name"`


    // TopicName - The topic that will cause the trigger to be invoked
    TopicName string `json:"topicName"`


    // Target - The target to invoke when a matching event is received
    Target Triggertarget `json:"target"`


    // Version - Version of this trigger
    Version int `json:"version"`


    // Enabled - Whether or not the trigger is enabled
    Enabled bool `json:"enabled"`


    // MatchCriteria - The configuration for when a trigger is considered to be a match for an event
    MatchCriteria []Matchcriteria `json:"matchCriteria"`


    // EventTTLSeconds - Optional length of time that events are meaningful after origination. Events older than this threshold may be dropped if the platform is delayed in processing events. Unset means events are valid indefinitely, otherwise must be set to at least 10 seconds. Only one of eventTTLSeconds or delayBySeconds can be set.
    EventTTLSeconds int `json:"eventTTLSeconds"`


    // DelayBySeconds - Optional delay invoking target after trigger fires. Must be in the range of 60 to 900 seconds. Only one of eventTTLSeconds or delayBySeconds can be set.
    DelayBySeconds int `json:"delayBySeconds"`


    // Description - Description of the trigger. Can be up to 512 characters in length.
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Trigger) String() string {
    
    
    
    
    
     o.MatchCriteria = []Matchcriteria{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trigger) MarshalJSON() ([]byte, error) {
    type Alias Trigger

    if TriggerMarshalled {
        return []byte("{}"), nil
    }
    TriggerMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        TopicName string `json:"topicName"`
        
        Target Triggertarget `json:"target"`
        
        Version int `json:"version"`
        
        Enabled bool `json:"enabled"`
        
        MatchCriteria []Matchcriteria `json:"matchCriteria"`
        
        EventTTLSeconds int `json:"eventTTLSeconds"`
        
        DelayBySeconds int `json:"delayBySeconds"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        


        
        MatchCriteria: []Matchcriteria{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

