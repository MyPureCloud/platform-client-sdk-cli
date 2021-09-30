package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsconversationasyncqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsconversationasyncqueryresponseDud struct { 
    


    


    

}

// Analyticsconversationasyncqueryresponse
type Analyticsconversationasyncqueryresponse struct { 
    // Cursor - Optional cursor to indicate where to resume the results
    Cursor string `json:"cursor"`


    // DataAvailabilityDate - Data available up to at least this datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DataAvailabilityDate time.Time `json:"dataAvailabilityDate"`


    // Conversations
    Conversations []Analyticsconversation `json:"conversations"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationasyncqueryresponse) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Conversations = []Analyticsconversation{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsconversationasyncqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Analyticsconversationasyncqueryresponse

    if AnalyticsconversationasyncqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsconversationasyncqueryresponseMarshalled = true

    return json.Marshal(&struct { 
        Cursor string `json:"cursor"`
        
        DataAvailabilityDate time.Time `json:"dataAvailabilityDate"`
        
        Conversations []Analyticsconversation `json:"conversations"`
        
        *Alias
    }{
        

        

        

        

        

        
        Conversations: []Analyticsconversation{{}},
        

        
        Alias: (*Alias)(u),
    })
}

