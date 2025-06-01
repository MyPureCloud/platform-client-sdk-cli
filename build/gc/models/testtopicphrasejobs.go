package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TesttopicphrasejobsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TesttopicphrasejobsDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Testtopicphrasejobs
type Testtopicphrasejobs struct { 
    


    // State
    State string `json:"state"`


    // CreatedBy
    CreatedBy Addressableentityref `json:"createdBy"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // Action
    Action string `json:"action"`


    // EntityType
    EntityType string `json:"entityType"`


    

}

// String returns a JSON representation of the model
func (o *Testtopicphrasejobs) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testtopicphrasejobs) MarshalJSON() ([]byte, error) {
    type Alias Testtopicphrasejobs

    if TesttopicphrasejobsMarshalled {
        return []byte("{}"), nil
    }
    TesttopicphrasejobsMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        CreatedBy Addressableentityref `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        Action string `json:"action"`
        
        EntityType string `json:"entityType"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

