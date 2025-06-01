package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentwheelpickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentwheelpickerDud struct { 
    


    

}

// Conversationcontentwheelpicker - Wheel Picker object for representing single selection of selectable items.
type Conversationcontentwheelpicker struct { 
    // Id - Optional unique identifier to help map component replies to form messages where multiple Wheel Pickers can be present.
    Id string `json:"id"`


    // Items - The main text displayed for the item.
    Items []Conversationcontentwheelpickeritem `json:"items"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentwheelpicker) String() string {
    
     o.Items = []Conversationcontentwheelpickeritem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentwheelpicker) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentwheelpicker

    if ConversationcontentwheelpickerMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentwheelpickerMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Items []Conversationcontentwheelpickeritem `json:"items"`
        *Alias
    }{

        


        
        Items: []Conversationcontentwheelpickeritem{{}},
        

        Alias: (*Alias)(u),
    })
}

