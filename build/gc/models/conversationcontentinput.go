package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentinputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentinputDud struct { 
    


    


    


    


    


    

}

// Conversationcontentinput - Single or multiline text input component.
type Conversationcontentinput struct { 
    // Id - Unique identifier for the input
    Id string `json:"id"`


    // Title - The main text displayed for the input field(s).
    Title string `json:"title"`


    // Subtitle - Additional text providing more details about the input field(s).
    Subtitle string `json:"subtitle"`


    // PlaceholderText - Placeholder text for input field(s).
    PlaceholderText string `json:"placeholderText"`


    // IsRequired - Whether the input field(s) are required to be filled in.
    IsRequired bool `json:"isRequired"`


    // IsMultipleLine - Whether the input should allow multiple lines of text.
    IsMultipleLine bool `json:"isMultipleLine"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentinput) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentinput) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentinput

    if ConversationcontentinputMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentinputMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        PlaceholderText string `json:"placeholderText"`
        
        IsRequired bool `json:"isRequired"`
        
        IsMultipleLine bool `json:"isMultipleLine"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

