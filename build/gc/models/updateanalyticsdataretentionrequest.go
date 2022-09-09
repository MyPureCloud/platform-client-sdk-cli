package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateanalyticsdataretentionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateanalyticsdataretentionrequestDud struct { 
    

}

// Updateanalyticsdataretentionrequest
type Updateanalyticsdataretentionrequest struct { 
    // RetentionDays - Analytics data retention period in days to set for the organization.
    RetentionDays int `json:"retentionDays"`

}

// String returns a JSON representation of the model
func (o *Updateanalyticsdataretentionrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateanalyticsdataretentionrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateanalyticsdataretentionrequest

    if UpdateanalyticsdataretentionrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateanalyticsdataretentionrequestMarshalled = true

    return json.Marshal(&struct {
        
        RetentionDays int `json:"retentionDays"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

