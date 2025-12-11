package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationnotificationtemplatecardMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationnotificationtemplatecardDud struct { 
    


    


    

}

// Conversationnotificationtemplatecard - Template card object.
type Conversationnotificationtemplatecard struct { 
    // Header - The template header.
    Header Conversationnotificationtemplateheader `json:"header"`


    // Body - The template body.
    Body Conversationnotificationtemplatebody `json:"body"`


    // Buttons - Template buttons.
    Buttons []Conversationnotificationtemplatebutton `json:"buttons"`

}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplatecard) String() string {
    
    
     o.Buttons = []Conversationnotificationtemplatebutton{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationnotificationtemplatecard) MarshalJSON() ([]byte, error) {
    type Alias Conversationnotificationtemplatecard

    if ConversationnotificationtemplatecardMarshalled {
        return []byte("{}"), nil
    }
    ConversationnotificationtemplatecardMarshalled = true

    return json.Marshal(&struct {
        
        Header Conversationnotificationtemplateheader `json:"header"`
        
        Body Conversationnotificationtemplatebody `json:"body"`
        
        Buttons []Conversationnotificationtemplatebutton `json:"buttons"`
        *Alias
    }{

        


        


        
        Buttons: []Conversationnotificationtemplatebutton{{}},
        

        Alias: (*Alias)(u),
    })
}

