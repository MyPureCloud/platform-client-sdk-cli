package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivationtriggerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivationtriggerDud struct { 
    


    


    


    

}

// Activationtrigger
type Activationtrigger struct { 
    // TriggerType - Trigger type that activated this checklist.
    TriggerType string `json:"triggerType"`


    // TriggerDate - Date when the checklist was triggered. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    TriggerDate time.Time `json:"triggerDate"`


    // IntentId - The intent ID if checklist was triggered by an intent.
    IntentId string `json:"intentId"`


    // IntentName - The intent name if checklist was triggered by an intent.
    IntentName string `json:"intentName"`

}

// String returns a JSON representation of the model
func (o *Activationtrigger) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activationtrigger) MarshalJSON() ([]byte, error) {
    type Alias Activationtrigger

    if ActivationtriggerMarshalled {
        return []byte("{}"), nil
    }
    ActivationtriggerMarshalled = true

    return json.Marshal(&struct {
        
        TriggerType string `json:"triggerType"`
        
        TriggerDate time.Time `json:"triggerDate"`
        
        IntentId string `json:"intentId"`
        
        IntentName string `json:"intentName"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

