package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentlistpickeritemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentlistpickeritemDud struct { 
    


    


    


    

}

// Conversationcontentlistpickeritem - Represents a selectable item in a list picker.
type Conversationcontentlistpickeritem struct { 
    // Id - Unique identifier for the list picker item
    Id string `json:"id"`


    // Title - The main text displayed for the item.
    Title string `json:"title"`


    // Subtitle - Additional text providing more details about the item.
    Subtitle string `json:"subtitle"`


    // ImageUrl - URL of an image to be displayed alongside the item.
    ImageUrl string `json:"imageUrl"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentlistpickeritem) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentlistpickeritem) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentlistpickeritem

    if ConversationcontentlistpickeritemMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentlistpickeritemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        ImageUrl string `json:"imageUrl"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

