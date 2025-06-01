package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationformpageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationformpageDud struct { 
    


    


    

}

// Conversationformpage - Form page object.
type Conversationformpage struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the subtitle.
    Subtitle string `json:"subtitle"`


    // PageComponents - Page components in this form page.
    PageComponents []Conversationformpagecomponent `json:"pageComponents"`

}

// String returns a JSON representation of the model
func (o *Conversationformpage) String() string {
    
    
     o.PageComponents = []Conversationformpagecomponent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationformpage) MarshalJSON() ([]byte, error) {
    type Alias Conversationformpage

    if ConversationformpageMarshalled {
        return []byte("{}"), nil
    }
    ConversationformpageMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        PageComponents []Conversationformpagecomponent `json:"pageComponents"`
        *Alias
    }{

        


        


        
        PageComponents: []Conversationformpagecomponent{{}},
        

        Alias: (*Alias)(u),
    })
}

