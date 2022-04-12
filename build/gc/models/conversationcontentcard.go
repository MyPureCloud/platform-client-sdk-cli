package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentcardMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentcardDud struct { 
    


    


    


    


    


    

}

// Conversationcontentcard
type Conversationcontentcard struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Description - Text to show in the description.
    Description string `json:"description"`


    // DefaultAction - Default action to be taken.
    DefaultAction Conversationcardaction `json:"defaultAction"`


    // Actions - A List of action objects.
    Actions []Conversationcardaction `json:"actions"`


    // Image
    Image string `json:"image"`


    // Video
    Video string `json:"video"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentcard) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Actions = []Conversationcardaction{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentcard) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentcard

    if ConversationcontentcardMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentcardMarshalled = true

    return json.Marshal(&struct { 
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        DefaultAction Conversationcardaction `json:"defaultAction"`
        
        Actions []Conversationcardaction `json:"actions"`
        
        Image string `json:"image"`
        
        Video string `json:"video"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Actions: []Conversationcardaction{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

