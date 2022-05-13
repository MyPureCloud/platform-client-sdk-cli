package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsconversationqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsconversationqueryresponseDud struct { 
    


    


    

}

// Analyticsconversationqueryresponse
type Analyticsconversationqueryresponse struct { 
    // Aggregations
    Aggregations []Aggregationresult `json:"aggregations"`


    // Conversations
    Conversations []Analyticsconversationwithoutattributes `json:"conversations"`


    // TotalHits
    TotalHits int `json:"totalHits"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationqueryresponse) String() string {
     o.Aggregations = []Aggregationresult{{}} 
     o.Conversations = []Analyticsconversationwithoutattributes{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsconversationqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Analyticsconversationqueryresponse

    if AnalyticsconversationqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsconversationqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Aggregations []Aggregationresult `json:"aggregations"`
        
        Conversations []Analyticsconversationwithoutattributes `json:"conversations"`
        
        TotalHits int `json:"totalHits"`
        *Alias
    }{

        
        Aggregations: []Aggregationresult{{}},
        


        
        Conversations: []Analyticsconversationwithoutattributes{{}},
        


        

        Alias: (*Alias)(u),
    })
}

