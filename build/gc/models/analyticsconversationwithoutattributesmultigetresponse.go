package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsconversationwithoutattributesmultigetresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsconversationwithoutattributesmultigetresponseDud struct { 
    

}

// Analyticsconversationwithoutattributesmultigetresponse
type Analyticsconversationwithoutattributesmultigetresponse struct { 
    // Conversations
    Conversations []Analyticsconversationwithoutattributes `json:"conversations"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationwithoutattributesmultigetresponse) String() string {
    
    
     o.Conversations = []Analyticsconversationwithoutattributes{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsconversationwithoutattributesmultigetresponse) MarshalJSON() ([]byte, error) {
    type Alias Analyticsconversationwithoutattributesmultigetresponse

    if AnalyticsconversationwithoutattributesmultigetresponseMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsconversationwithoutattributesmultigetresponseMarshalled = true

    return json.Marshal(&struct { 
        Conversations []Analyticsconversationwithoutattributes `json:"conversations"`
        
        *Alias
    }{
        

        
        Conversations: []Analyticsconversationwithoutattributes{{}},
        

        
        Alias: (*Alias)(u),
    })
}

