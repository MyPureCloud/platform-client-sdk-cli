package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionasyncaggregatequeryresponseDud struct { 
    


    

}

// Actionasyncaggregatequeryresponse
type Actionasyncaggregatequeryresponse struct { 
    // Results
    Results []Actionaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Actionasyncaggregatequeryresponse) String() string {
     o.Results = []Actionaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Actionasyncaggregatequeryresponse

    if ActionasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    ActionasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Actionaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Actionaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

