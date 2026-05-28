package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryexternalactivityopportunitiesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryexternalactivityopportunitiesresponseDud struct { 
    

}

// Queryexternalactivityopportunitiesresponse
type Queryexternalactivityopportunitiesresponse struct { 
    // ExternalActivities - The external activities and associated opportunity details
    ExternalActivities []Queryexternalactivityopportunityresult `json:"externalActivities"`

}

// String returns a JSON representation of the model
func (o *Queryexternalactivityopportunitiesresponse) String() string {
     o.ExternalActivities = []Queryexternalactivityopportunityresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryexternalactivityopportunitiesresponse) MarshalJSON() ([]byte, error) {
    type Alias Queryexternalactivityopportunitiesresponse

    if QueryexternalactivityopportunitiesresponseMarshalled {
        return []byte("{}"), nil
    }
    QueryexternalactivityopportunitiesresponseMarshalled = true

    return json.Marshal(&struct {
        
        ExternalActivities []Queryexternalactivityopportunityresult `json:"externalActivities"`
        *Alias
    }{

        
        ExternalActivities: []Queryexternalactivityopportunityresult{{}},
        

        Alias: (*Alias)(u),
    })
}

