package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EngagementrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EngagementrequestDud struct { 
    


    

}

// Engagementrequest
type Engagementrequest struct { 
    // Visibility - Represents the visibility of summary
    Visibility string `json:"visibility"`


    // Status - Represents the engagements made by the agent in response to the generated summary
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Engagementrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Engagementrequest) MarshalJSON() ([]byte, error) {
    type Alias Engagementrequest

    if EngagementrequestMarshalled {
        return []byte("{}"), nil
    }
    EngagementrequestMarshalled = true

    return json.Marshal(&struct {
        
        Visibility string `json:"visibility"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

