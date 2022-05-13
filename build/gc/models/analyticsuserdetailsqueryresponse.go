package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsuserdetailsqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsuserdetailsqueryresponseDud struct { 
    


    


    

}

// Analyticsuserdetailsqueryresponse
type Analyticsuserdetailsqueryresponse struct { 
    // UserDetails
    UserDetails []Analyticsuserdetail `json:"userDetails"`


    // Aggregations
    Aggregations []Aggregationresult `json:"aggregations"`


    // TotalHits
    TotalHits int `json:"totalHits"`

}

// String returns a JSON representation of the model
func (o *Analyticsuserdetailsqueryresponse) String() string {
     o.UserDetails = []Analyticsuserdetail{{}} 
     o.Aggregations = []Aggregationresult{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsuserdetailsqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Analyticsuserdetailsqueryresponse

    if AnalyticsuserdetailsqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsuserdetailsqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        UserDetails []Analyticsuserdetail `json:"userDetails"`
        
        Aggregations []Aggregationresult `json:"aggregations"`
        
        TotalHits int `json:"totalHits"`
        *Alias
    }{

        
        UserDetails: []Analyticsuserdetail{{}},
        


        
        Aggregations: []Aggregationresult{{}},
        


        

        Alias: (*Alias)(u),
    })
}

