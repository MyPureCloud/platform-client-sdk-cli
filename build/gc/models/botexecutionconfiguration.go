package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotexecutionconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotexecutionconfigurationDud struct { 
    


    


    

}

// Botexecutionconfiguration - Model for setting the launch configuration for a Nuance bot available to Genesys Cloud
type Botexecutionconfiguration struct { 
    // BotId - The Nuance bot ID
    BotId string `json:"botId"`


    // ExecutionHost - The hostname to use when contacting Nuance to execute this bot
    ExecutionHost string `json:"executionHost"`


    // BotCredentials - The bot's launch credentials
    BotCredentials Nuancebotcredentials `json:"botCredentials"`

}

// String returns a JSON representation of the model
func (o *Botexecutionconfiguration) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botexecutionconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Botexecutionconfiguration

    if BotexecutionconfigurationMarshalled {
        return []byte("{}"), nil
    }
    BotexecutionconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        BotId string `json:"botId"`
        
        ExecutionHost string `json:"executionHost"`
        
        BotCredentials Nuancebotcredentials `json:"botCredentials"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

