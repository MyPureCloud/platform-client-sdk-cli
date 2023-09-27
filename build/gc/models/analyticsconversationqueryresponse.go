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
    // Conversations
    Conversations []Analyticsconversationwithoutattributes `json:"conversations"`


    // Aggregations
    Aggregations []Aggregationresult `json:"aggregations"`


    // TotalHits
    TotalHits int `json:"totalHits"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationqueryresponse) String() string {
     o.Conversations = []Analyticsconversationwithoutattributes{{}} 
     o.Aggregations = []Aggregationresult{{}} 
    

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
        
        Conversations []Analyticsconversationwithoutattributes `json:"conversations"`
        
        Aggregations []Aggregationresult `json:"aggregations"`
        
        TotalHits int `json:"totalHits"`
        *Alias
    }{

        
        Conversations: []Analyticsconversationwithoutattributes{{}},
        


        
        Aggregations: []Aggregationresult{{}},
        


        

        Alias: (*Alias)(u),
    })
}

