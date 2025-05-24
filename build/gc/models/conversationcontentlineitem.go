package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentlineitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentlineitemDud struct { 
    


    


    


    

}

// Conversationcontentlineitem - Item included in an order.
type Conversationcontentlineitem struct { 
    // Name - The display name for the item.
    Name string `json:"name"`


    // Price - The price of the item.
    Price float64 `json:"price"`


    // Description - Additional details about the item (e.g. the length of time to deliver for shipping options).
    Description string `json:"description"`


    // Identifier - A client defined value used to identify the item.
    Identifier string `json:"identifier"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentlineitem) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentlineitem) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentlineitem

    if ConversationcontentlineitemMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentlineitemMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Price float64 `json:"price"`
        
        Description string `json:"description"`
        
        Identifier string `json:"identifier"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

