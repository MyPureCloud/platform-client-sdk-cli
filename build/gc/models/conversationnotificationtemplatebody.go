package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationnotificationtemplatebodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationnotificationtemplatebodyDud struct { 
    


    

}

// Conversationnotificationtemplatebody - Template body object.
type Conversationnotificationtemplatebody struct { 
    // Text - Body text. For WhatsApp, ignored.
    Text string `json:"text"`


    // Parameters - Template parameters for placeholders in template.
    Parameters []Conversationnotificationtemplateparameter `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplatebody) String() string {
    
     o.Parameters = []Conversationnotificationtemplateparameter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationnotificationtemplatebody) MarshalJSON() ([]byte, error) {
    type Alias Conversationnotificationtemplatebody

    if ConversationnotificationtemplatebodyMarshalled {
        return []byte("{}"), nil
    }
    ConversationnotificationtemplatebodyMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Parameters []Conversationnotificationtemplateparameter `json:"parameters"`
        *Alias
    }{

        


        
        Parameters: []Conversationnotificationtemplateparameter{{}},
        

        Alias: (*Alias)(u),
    })
}

