package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentgenericMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentgenericDud struct { 
    


    


    


    


    


    

}

// Conversationcontentgeneric - Generic content object.
type Conversationcontentgeneric struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Description - Text to show in the description.
    Description string `json:"description"`


    // Image - URL of an image.
    Image string `json:"image"`


    // Video - URL of a video.
    Video string `json:"video"`


    // Actions - Actions to be taken.
    Actions Conversationcontentactions `json:"actions"`


    // Components - An array of component objects.
    Components []Conversationbuttoncomponent `json:"components"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentgeneric) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Components = []Conversationbuttoncomponent{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentgeneric) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentgeneric

    if ConversationcontentgenericMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentgenericMarshalled = true

    return json.Marshal(&struct { 
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        Image string `json:"image"`
        
        Video string `json:"video"`
        
        Actions Conversationcontentactions `json:"actions"`
        
        Components []Conversationbuttoncomponent `json:"components"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        Components: []Conversationbuttoncomponent{{}},
        

        
        Alias: (*Alias)(u),
    })
}

