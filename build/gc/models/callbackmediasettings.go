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


    // AlertingTimeoutSeconds
    AlertingTimeoutSeconds int `json:"alertingTimeoutSeconds"`


    // ServiceLevel
    ServiceLevel Servicelevel `json:"serviceLevel"`


    // SubTypeSettings - Map of media subtype to media subtype specific settings.
    SubTypeSettings map[string]Basemediasettings `json:"subTypeSettings"`

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
        
        SubTypeSettings map[string]Basemediasettings `json:"subTypeSettings"`
        *Alias
    }{

        


        


        


        
        SubTypeSettings: map[string]Basemediasettings{"": {}},
        

        Alias: (*Alias)(u),
    })
}
