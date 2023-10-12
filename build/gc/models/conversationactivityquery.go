package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationactivityqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationactivityqueryDud struct { 
    


    


    


    

}

// Conversationactivityquery
type Conversationactivityquery struct { 
    // Metrics - List of requested metrics
    Metrics []Conversationactivityquerymetric `json:"metrics"`


    // GroupBy - Dimension(s) to group by
    GroupBy []string `json:"groupBy"`


    // Filter - Filter to return a subset of observations. Expresses boolean logical predicates as well as dimensional filters
    Filter Conversationactivityqueryfilter `json:"filter"`


    // Order - Sort the result set in ascending/descending order. Default is ascending
    Order string `json:"order"`

}

// String returns a JSON representation of the model
func (o *Conversationactivityquery) String() string {
     o.Metrics = []Conversationactivityquerymetric{{}} 
     o.GroupBy = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationactivityquery) MarshalJSON() ([]byte, error) {
    type Alias Conversationactivityquery

    if ConversationactivityqueryMarshalled {
        return []byte("{}"), nil
    }
    ConversationactivityqueryMarshalled = true

    return json.Marshal(&struct {
        
        Metrics []Conversationactivityquerymetric `json:"metrics"`
        
        GroupBy []string `json:"groupBy"`
        
        Filter Conversationactivityqueryfilter `json:"filter"`
        
        Order string `json:"order"`
        *Alias
    }{

        
        Metrics: []Conversationactivityquerymetric{{}},
        


        
        GroupBy: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

