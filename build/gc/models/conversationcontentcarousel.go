package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentcarouselMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentcarouselDud struct { 
    

}

// Conversationcontentcarousel - Carousel content object.
type Conversationcontentcarousel struct { 
    // Cards - An array of card objects.
    Cards []Conversationcontentcard `json:"cards"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentcarousel) String() string {
     o.Cards = []Conversationcontentcard{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentcarousel) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentcarousel

    if ConversationcontentcarouselMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentcarouselMarshalled = true

    return json.Marshal(&struct {
        
        Cards []Conversationcontentcard `json:"cards"`
        *Alias
    }{

        
        Cards: []Conversationcontentcard{{}},
        

        Alias: (*Alias)(u),
    })
}

