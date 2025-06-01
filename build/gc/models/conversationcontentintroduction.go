package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentintroductionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentintroductionDud struct { 
    


    


    


    

}

// Conversationcontentintroduction - Introduction content object.
type Conversationcontentintroduction struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the subtitle.
    Subtitle string `json:"subtitle"`


    // ImageUrl - URL of an image.
    ImageUrl string `json:"imageUrl"`


    // ButtonText - Text to show on the button.
    ButtonText string `json:"buttonText"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentintroduction) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentintroduction) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentintroduction

    if ConversationcontentintroductionMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentintroductionMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        ImageUrl string `json:"imageUrl"`
        
        ButtonText string `json:"buttonText"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

