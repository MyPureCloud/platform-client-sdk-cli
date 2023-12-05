package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatetriggerrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatetriggerrequestDud struct { 
    


    


    


    


    


    


    


    


    

}

// Updatetriggerrequest
type Updatetriggerrequest struct { 
    // Version - Version of this trigger
    Version int `json:"version"`


    // Enabled - Boolean indicating if Trigger is enabled
    Enabled bool `json:"enabled"`


    // Target - The target to invoke when a matching event is received
    Target Triggertarget `json:"target"`


    // MatchCriteria - The configuration for when a trigger is considered to be a match for an event
    MatchCriteria []Matchcriteria `json:"matchCriteria"`


    // Name - The name of the trigger
    Name string `json:"name"`


    // TopicName - The topic that will cause the trigger to be invoked. Must match existing trigger topicName.
    TopicName string `json:"topicName"`


    // EventTTLSeconds - Optional length of time that events are meaningful after origination. Events older than this threshold may be dropped if the platform is delayed in processing events. Unset means events are valid indefinitely, otherwise must be set to at least 10 seconds. Only one of eventTTLSeconds or delayBySeconds can be set.
    EventTTLSeconds int `json:"eventTTLSeconds"`


    // DelayBySeconds - Optional delay invoking target after trigger fires. Must be in the range of 60 to 900 seconds. Only one of eventTTLSeconds or delayBySeconds can be set.
    DelayBySeconds int `json:"delayBySeconds"`


    // Description - Description of the trigger. Can be up to 512 characters in length.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Updatetriggerrequest) String() string {
    
    
    
     o.MatchCriteria = []Matchcriteria{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatetriggerrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatetriggerrequest

    if UpdatetriggerrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatetriggerrequestMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        Enabled bool `json:"enabled"`
        
        Target Triggertarget `json:"target"`
        
        MatchCriteria []Matchcriteria `json:"matchCriteria"`
        
        Name string `json:"name"`
        
        TopicName string `json:"topicName"`
        
        EventTTLSeconds int `json:"eventTTLSeconds"`
        
        DelayBySeconds int `json:"delayBySeconds"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        
        MatchCriteria: []Matchcriteria{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

