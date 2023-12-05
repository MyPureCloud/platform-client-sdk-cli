package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotchannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotchannelDud struct { 
    


    


    


    

}

// Botchannel - Channel information relevant to a bot flow.
type Botchannel struct { 
    // Name - The name of the channel.
    Name string `json:"name"`


    // InputModes - The input modes for the channel.
    InputModes []string `json:"inputModes"`


    // OutputModes - The output modes for the channel.
    OutputModes []string `json:"outputModes"`


    // UserAgent - Information about the end user agent calling the bot flow.
    UserAgent Textbotuseragent `json:"userAgent"`

}

// String returns a JSON representation of the model
func (o *Botchannel) String() string {
    
     o.InputModes = []string{""} 
     o.OutputModes = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botchannel) MarshalJSON() ([]byte, error) {
    type Alias Botchannel

    if BotchannelMarshalled {
        return []byte("{}"), nil
    }
    BotchannelMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        InputModes []string `json:"inputModes"`
        
        OutputModes []string `json:"outputModes"`
        
        UserAgent Textbotuseragent `json:"userAgent"`
        *Alias
    }{

        


        
        InputModes: []string{""},
        


        
        OutputModes: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

