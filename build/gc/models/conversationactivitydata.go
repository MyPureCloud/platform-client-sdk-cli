package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationactivitydataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationactivitydataDud struct { 
    


    


    


    

}

// Conversationactivitydata
type Conversationactivitydata struct { 
    // Group - A mapping from grouping dimension to value
    Group map[string]string `json:"group"`


    // Data - Data for metrics
    Data []Conversationactivitymetricvalue `json:"data"`


    // Truncated - Flag for a truncated list of entities. If truncated, the first half of the list of entities will contain the oldest entities and the second half the newest entities.
    Truncated bool `json:"truncated"`


    // Entities - Details for active entities
    Entities []Conversationactivityentitydata `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Conversationactivitydata) String() string {
     o.Group = map[string]string{"": ""} 
     o.Data = []Conversationactivitymetricvalue{{}} 
    
     o.Entities = []Conversationactivityentitydata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationactivitydata) MarshalJSON() ([]byte, error) {
    type Alias Conversationactivitydata

    if ConversationactivitydataMarshalled {
        return []byte("{}"), nil
    }
    ConversationactivitydataMarshalled = true

    return json.Marshal(&struct {
        
        Group map[string]string `json:"group"`
        
        Data []Conversationactivitymetricvalue `json:"data"`
        
        Truncated bool `json:"truncated"`
        
        Entities []Conversationactivityentitydata `json:"entities"`
        *Alias
    }{

        
        Group: map[string]string{"": ""},
        


        
        Data: []Conversationactivitymetricvalue{{}},
        


        


        
        Entities: []Conversationactivityentitydata{{}},
        

        Alias: (*Alias)(u),
    })
}

