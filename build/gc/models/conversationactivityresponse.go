package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationactivityresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationactivityresponseDud struct { 
    


    

}

// Conversationactivityresponse
type Conversationactivityresponse struct { 
    // Results - Query results
    Results []Conversationactivitydata `json:"results"`


    // EntityIdDimension - Dimension that is used as an entityId
    EntityIdDimension string `json:"entityIdDimension"`

}

// String returns a JSON representation of the model
func (o *Conversationactivityresponse) String() string {
     o.Results = []Conversationactivitydata{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationactivityresponse) MarshalJSON() ([]byte, error) {
    type Alias Conversationactivityresponse

    if ConversationactivityresponseMarshalled {
        return []byte("{}"), nil
    }
    ConversationactivityresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Conversationactivitydata `json:"results"`
        
        EntityIdDimension string `json:"entityIdDimension"`
        *Alias
    }{

        
        Results: []Conversationactivitydata{{}},
        


        

        Alias: (*Alias)(u),
    })
}

