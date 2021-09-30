package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsuserdetailsasyncqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsuserdetailsasyncqueryresponseDud struct { 
    


    


    

}

// Analyticsuserdetailsasyncqueryresponse
type Analyticsuserdetailsasyncqueryresponse struct { 
    // UserDetails
    UserDetails []Analyticsuserdetail `json:"userDetails"`


    // Cursor - Optional cursor to indicate where to resume the results
    Cursor string `json:"cursor"`


    // DataAvailabilityDate - Data available up to at least this datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DataAvailabilityDate time.Time `json:"dataAvailabilityDate"`

}

// String returns a JSON representation of the model
func (o *Analyticsuserdetailsasyncqueryresponse) String() string {
    
    
     o.UserDetails = []Analyticsuserdetail{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsuserdetailsasyncqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Analyticsuserdetailsasyncqueryresponse

    if AnalyticsuserdetailsasyncqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsuserdetailsasyncqueryresponseMarshalled = true

    return json.Marshal(&struct { 
        UserDetails []Analyticsuserdetail `json:"userDetails"`
        
        Cursor string `json:"cursor"`
        
        DataAvailabilityDate time.Time `json:"dataAvailabilityDate"`
        
        *Alias
    }{
        

        
        UserDetails: []Analyticsuserdetail{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

