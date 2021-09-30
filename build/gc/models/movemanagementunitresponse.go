package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MovemanagementunitresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MovemanagementunitresponseDud struct { 
    


    

}

// Movemanagementunitresponse
type Movemanagementunitresponse struct { 
    // BusinessUnit - The new business unit
    BusinessUnit Businessunitreference `json:"businessUnit"`


    // Status - The status of the move.  Will always be 'Processing' unless the Management Unit is already in the requested Business Unit in which case it will be 'Complete'
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Movemanagementunitresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Movemanagementunitresponse) MarshalJSON() ([]byte, error) {
    type Alias Movemanagementunitresponse

    if MovemanagementunitresponseMarshalled {
        return []byte("{}"), nil
    }
    MovemanagementunitresponseMarshalled = true

    return json.Marshal(&struct { 
        BusinessUnit Businessunitreference `json:"businessUnit"`
        
        Status string `json:"status"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

