package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentwheelpickeritemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentwheelpickeritemDud struct { 
    


    


    

}

// Conversationcontentwheelpickeritem - Represents a selectable item in a Wheel Picker
type Conversationcontentwheelpickeritem struct { 
    // Id - Unique identifier for the wheel picker item
    Id string `json:"id"`


    // Title - The main text displayed for the item.
    Title string `json:"title"`


    // Value - The value of the item.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentwheelpickeritem) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentwheelpickeritem) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentwheelpickeritem

    if ConversationcontentwheelpickeritemMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentwheelpickeritemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

