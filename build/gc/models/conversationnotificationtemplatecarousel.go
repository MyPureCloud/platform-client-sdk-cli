package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationnotificationtemplatecarouselMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationnotificationtemplatecarouselDud struct { 
    

}

// Conversationnotificationtemplatecarousel - Template carousel object.
type Conversationnotificationtemplatecarousel struct { 
    // Cards - An array of template card objects.
    Cards []Conversationnotificationtemplatecard `json:"cards"`

}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplatecarousel) String() string {
     o.Cards = []Conversationnotificationtemplatecard{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationnotificationtemplatecarousel) MarshalJSON() ([]byte, error) {
    type Alias Conversationnotificationtemplatecarousel

    if ConversationnotificationtemplatecarouselMarshalled {
        return []byte("{}"), nil
    }
    ConversationnotificationtemplatecarouselMarshalled = true

    return json.Marshal(&struct {
        
        Cards []Conversationnotificationtemplatecard `json:"cards"`
        *Alias
    }{

        
        Cards: []Conversationnotificationtemplatecard{{}},
        

        Alias: (*Alias)(u),
    })
}

