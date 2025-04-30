package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentlistpickersectionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentlistpickersectionDud struct { 
    


    


    

}

// Conversationcontentlistpickersection - List Picker object for presenting a section of selectable items.
type Conversationcontentlistpickersection struct { 
    // Title - Required title for the section.
    Title string `json:"title"`


    // MultipleSelection - Whether multiple items can be selected in this section.
    MultipleSelection bool `json:"multipleSelection"`


    // Items - List of items to choice from in the section
    Items []Conversationcontentlistpickeritem `json:"items"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentlistpickersection) String() string {
    
    
     o.Items = []Conversationcontentlistpickeritem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentlistpickersection) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentlistpickersection

    if ConversationcontentlistpickersectionMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentlistpickersectionMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        MultipleSelection bool `json:"multipleSelection"`
        
        Items []Conversationcontentlistpickeritem `json:"items"`
        *Alias
    }{

        


        


        
        Items: []Conversationcontentlistpickeritem{{}},
        

        Alias: (*Alias)(u),
    })
}

