package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsdataretentionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsdataretentionresponseDud struct { 
    


    


    

}

// Analyticsdataretentionresponse
type Analyticsdataretentionresponse struct { 
    // RetentionDays - Analytics data retention period in days for the organization.
    RetentionDays int `json:"retentionDays"`


    // DateCreated - Date and time when the analytics data retention was set. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date and time when the analytics data retention was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`

}

// String returns a JSON representation of the model
func (o *Analyticsdataretentionresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsdataretentionresponse) MarshalJSON() ([]byte, error) {
    type Alias Analyticsdataretentionresponse

    if AnalyticsdataretentionresponseMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsdataretentionresponseMarshalled = true

    return json.Marshal(&struct {
        
        RetentionDays int `json:"retentionDays"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

