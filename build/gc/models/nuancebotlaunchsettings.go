package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuancebotlaunchsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuancebotlaunchsettingsDud struct { 
    

}

// Nuancebotlaunchsettings - Model for setting the launch configuration for Nuance bots available to Genesys Cloud
type Nuancebotlaunchsettings struct { 
    // BotExecutionConfigurations - The list of Nuance bots that are configured as available to the Genesys Cloud system
    BotExecutionConfigurations []Botexecutionconfiguration `json:"botExecutionConfigurations"`

}

// String returns a JSON representation of the model
func (o *Nuancebotlaunchsettings) String() string {
     o.BotExecutionConfigurations = []Botexecutionconfiguration{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuancebotlaunchsettings) MarshalJSON() ([]byte, error) {
    type Alias Nuancebotlaunchsettings

    if NuancebotlaunchsettingsMarshalled {
        return []byte("{}"), nil
    }
    NuancebotlaunchsettingsMarshalled = true

    return json.Marshal(&struct {
        
        BotExecutionConfigurations []Botexecutionconfiguration `json:"botExecutionConfigurations"`
        *Alias
    }{

        
        BotExecutionConfigurations: []Botexecutionconfiguration{{}},
        

        Alias: (*Alias)(u),
    })
}

