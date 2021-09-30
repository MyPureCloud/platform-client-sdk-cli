package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuancemixdlgsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuancemixdlgsettingsDud struct { 
    


    

}

// Nuancemixdlgsettings
type Nuancemixdlgsettings struct { 
    // ChannelId - The Nuance channel ID to use when launching the Nuance bot, which must one of the code names of the bot's registered input channels.
    ChannelId string `json:"channelId"`


    // InputParameters - Name/value pairs of input variables to be sent to the Nuance bot. The values must be in the appropriate format for the variable's type (see https://docs.mix.nuance.com/dialog-grpc/v1/#simple-variable-types for help)
    InputParameters map[string]interface{} `json:"inputParameters"`

}

// String returns a JSON representation of the model
func (o *Nuancemixdlgsettings) String() string {
    
    
    
    
    
    
     o.InputParameters = map[string]interface{}{"": Interface{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuancemixdlgsettings) MarshalJSON() ([]byte, error) {
    type Alias Nuancemixdlgsettings

    if NuancemixdlgsettingsMarshalled {
        return []byte("{}"), nil
    }
    NuancemixdlgsettingsMarshalled = true

    return json.Marshal(&struct { 
        ChannelId string `json:"channelId"`
        
        InputParameters map[string]interface{} `json:"inputParameters"`
        
        *Alias
    }{
        

        

        

        
        InputParameters: map[string]interface{}{"": Interface{}},
        

        
        Alias: (*Alias)(u),
    })
}

