package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GoogledialogflowcustomsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GoogledialogflowcustomsettingsDud struct { 
    


    


    


    

}

// Googledialogflowcustomsettings
type Googledialogflowcustomsettings struct { 
    // Environment - If set this environment will be used to initiate the dialogflow bot, otherwise the default configuration will be used.  See https://cloud.google.com/dialogflow/docs/agents-versions
    Environment string `json:"environment"`


    // EventName - If set this eventName will be used to initiate the dialogflow bot rather than language processing on the input text.  See https://cloud.google.com/dialogflow/es/docs/events-overview
    EventName string `json:"eventName"`


    // WebhookQueryParameters - Parameters passed to the fulfillment webhook of the bot (if any).
    WebhookQueryParameters map[string]string `json:"webhookQueryParameters"`


    // EventInputParameters - Parameters passed to the event input of the bot.
    EventInputParameters map[string]string `json:"eventInputParameters"`

}

// String returns a JSON representation of the model
func (o *Googledialogflowcustomsettings) String() string {
    
    
    
    
    
    
    
    
    
    
     o.WebhookQueryParameters = map[string]string{"": ""} 
    
    
    
     o.EventInputParameters = map[string]string{"": ""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googledialogflowcustomsettings) MarshalJSON() ([]byte, error) {
    type Alias Googledialogflowcustomsettings

    if GoogledialogflowcustomsettingsMarshalled {
        return []byte("{}"), nil
    }
    GoogledialogflowcustomsettingsMarshalled = true

    return json.Marshal(&struct { 
        Environment string `json:"environment"`
        
        EventName string `json:"eventName"`
        
        WebhookQueryParameters map[string]string `json:"webhookQueryParameters"`
        
        EventInputParameters map[string]string `json:"eventInputParameters"`
        
        *Alias
    }{
        

        

        

        

        

        
        WebhookQueryParameters: map[string]string{"": ""},
        

        

        
        EventInputParameters: map[string]string{"": ""},
        

        
        Alias: (*Alias)(u),
    })
}

