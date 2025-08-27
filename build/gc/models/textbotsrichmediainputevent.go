package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotsrichmediainputeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotsrichmediainputeventDud struct { 
    


    

}

// Textbotsrichmediainputevent - RichMedia Input to the bot.
type Textbotsrichmediainputevent struct { 
    // Inputs - The Rich Media content inputs.
    Inputs []Conversationmessagecontent `json:"inputs"`


    // Messages - The Rich Media message events.
    Messages []Conversationmessageevent `json:"messages"`

}

// String returns a JSON representation of the model
func (o *Textbotsrichmediainputevent) String() string {
     o.Inputs = []Conversationmessagecontent{{}} 
     o.Messages = []Conversationmessageevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotsrichmediainputevent) MarshalJSON() ([]byte, error) {
    type Alias Textbotsrichmediainputevent

    if TextbotsrichmediainputeventMarshalled {
        return []byte("{}"), nil
    }
    TextbotsrichmediainputeventMarshalled = true

    return json.Marshal(&struct {
        
        Inputs []Conversationmessagecontent `json:"inputs"`
        
        Messages []Conversationmessageevent `json:"messages"`
        *Alias
    }{

        
        Inputs: []Conversationmessagecontent{{}},
        


        
        Messages: []Conversationmessageevent{{}},
        

        Alias: (*Alias)(u),
    })
}

