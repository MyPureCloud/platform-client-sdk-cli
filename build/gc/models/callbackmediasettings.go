package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallbackmediasettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallbackmediasettingsDud struct { 
    


    


    


    


    


    


    


    


    

}

// Callbackmediasettings
type Callbackmediasettings struct { 
    // EnableAutoAnswer - Indicates if auto-answer is enabled for the given media type or subtype (default is false).  Subtype settings take precedence over media type settings.
    EnableAutoAnswer bool `json:"enableAutoAnswer"`


    // AlertingTimeoutSeconds - The alerting timeout for the media type, in seconds
    AlertingTimeoutSeconds int `json:"alertingTimeoutSeconds"`


    // ServiceLevel - The targeted service level for the media type
    ServiceLevel Servicelevel `json:"serviceLevel"`


    // AutoAnswerAlertToneSeconds - How long to play the alerting tone for an auto-answer interaction
    AutoAnswerAlertToneSeconds float64 `json:"autoAnswerAlertToneSeconds"`


    // ManualAnswerAlertToneSeconds - How long to play the alerting tone for a manual-answer interaction
    ManualAnswerAlertToneSeconds float64 `json:"manualAnswerAlertToneSeconds"`


    // SubTypeSettings - Map of media subtype to media subtype specific settings.
    SubTypeSettings map[string]Basemediasettings `json:"subTypeSettings"`


    // EnableAutoDialAndEnd - Flag to enable Auto-Dial and Auto-End automation for callbacks on this queue.
    EnableAutoDialAndEnd bool `json:"enableAutoDialAndEnd"`


    // AutoDialDelaySeconds - Time in seconds after agent connects to callback before outgoing call is auto-dialed. Allowable values in range 0 - 1200 seconds. Defaults to 300 seconds.
    AutoDialDelaySeconds int `json:"autoDialDelaySeconds"`


    // AutoEndDelaySeconds - Time in seconds after agent disconnects from the outgoing call before the encasing callback is auto-ended. Allowable values in range 0 - 1200 seconds. Defaults to 300 seconds.
    AutoEndDelaySeconds int `json:"autoEndDelaySeconds"`

}

// String returns a JSON representation of the model
func (o *Callbackmediasettings) String() string {
    
    
    
    
    
     o.SubTypeSettings = map[string]Basemediasettings{"": {}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callbackmediasettings) MarshalJSON() ([]byte, error) {
    type Alias Callbackmediasettings

    if CallbackmediasettingsMarshalled {
        return []byte("{}"), nil
    }
    CallbackmediasettingsMarshalled = true

    return json.Marshal(&struct {
        
        EnableAutoAnswer bool `json:"enableAutoAnswer"`
        
        AlertingTimeoutSeconds int `json:"alertingTimeoutSeconds"`
        
        ServiceLevel Servicelevel `json:"serviceLevel"`
        
        AutoAnswerAlertToneSeconds float64 `json:"autoAnswerAlertToneSeconds"`
        
        ManualAnswerAlertToneSeconds float64 `json:"manualAnswerAlertToneSeconds"`
        
        SubTypeSettings map[string]Basemediasettings `json:"subTypeSettings"`
        
        EnableAutoDialAndEnd bool `json:"enableAutoDialAndEnd"`
        
        AutoDialDelaySeconds int `json:"autoDialDelaySeconds"`
        
        AutoEndDelaySeconds int `json:"autoEndDelaySeconds"`
        *Alias
    }{

        


        


        


        


        


        
        SubTypeSettings: map[string]Basemediasettings{"": {}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

