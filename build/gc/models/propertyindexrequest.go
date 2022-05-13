package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PropertyindexrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PropertyindexrequestDud struct { 
    


    


    

}

// Propertyindexrequest
type Propertyindexrequest struct { 
    // SessionId - Attach properties to a segment in the indicated session
    SessionId string `json:"sessionId"`


    // TargetDate - Attach properties to a segment covering a specific point in time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    TargetDate time.Time `json:"targetDate"`


    // Properties - The list of properties to index
    Properties []Analyticsproperty `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Propertyindexrequest) String() string {
    
    
     o.Properties = []Analyticsproperty{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Propertyindexrequest) MarshalJSON() ([]byte, error) {
    type Alias Propertyindexrequest

    if PropertyindexrequestMarshalled {
        return []byte("{}"), nil
    }
    PropertyindexrequestMarshalled = true

    return json.Marshal(&struct {
        
        SessionId string `json:"sessionId"`
        
        TargetDate time.Time `json:"targetDate"`
        
        Properties []Analyticsproperty `json:"properties"`
        *Alias
    }{

        


        


        
        Properties: []Analyticsproperty{{}},
        

        Alias: (*Alias)(u),
    })
}

