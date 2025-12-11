package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagemediasettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagemediasettingsDud struct { 
    


    


    


    


    


    


    


    

}

// Messagemediasettings
type Messagemediasettings struct { 
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
    SubTypeSettings map[string]Messagesubtypesettings `json:"subTypeSettings"`


    // EnableInactivityTimeout - Indicates if inactivity timeout is enabled for all subtypes.
    EnableInactivityTimeout bool `json:"enableInactivityTimeout"`


    // InactivityTimeoutSettings - Inactivity timeout settings for messages
    InactivityTimeoutSettings Inactivitytimeoutsettings `json:"inactivityTimeoutSettings"`

}

// String returns a JSON representation of the model
func (o *Messagemediasettings) String() string {
    
    
    
    
    
     o.SubTypeSettings = map[string]Messagesubtypesettings{"": {}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagemediasettings) MarshalJSON() ([]byte, error) {
    type Alias Messagemediasettings

    if MessagemediasettingsMarshalled {
        return []byte("{}"), nil
    }
    MessagemediasettingsMarshalled = true

    return json.Marshal(&struct {
        
        EnableAutoAnswer bool `json:"enableAutoAnswer"`
        
        AlertingTimeoutSeconds int `json:"alertingTimeoutSeconds"`
        
        ServiceLevel Servicelevel `json:"serviceLevel"`
        
        AutoAnswerAlertToneSeconds float64 `json:"autoAnswerAlertToneSeconds"`
        
        ManualAnswerAlertToneSeconds float64 `json:"manualAnswerAlertToneSeconds"`
        
        SubTypeSettings map[string]Messagesubtypesettings `json:"subTypeSettings"`
        
        EnableInactivityTimeout bool `json:"enableInactivityTimeout"`
        
        InactivityTimeoutSettings Inactivitytimeoutsettings `json:"inactivityTimeoutSettings"`
        *Alias
    }{

        


        


        


        


        


        
        SubTypeSettings: map[string]Messagesubtypesettings{"": {}},
        


        


        

        Alias: (*Alias)(u),
    })
}

