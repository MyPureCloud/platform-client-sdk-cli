package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReschedulingmanagementunitresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReschedulingmanagementunitresponseDud struct { 
    


    

}

// Reschedulingmanagementunitresponse
type Reschedulingmanagementunitresponse struct { 
    // ManagementUnit - The management unit
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // Applied - Whether the rescheduling run is applied for the given management unit
    Applied bool `json:"applied"`

}

// String returns a JSON representation of the model
func (o *Reschedulingmanagementunitresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reschedulingmanagementunitresponse) MarshalJSON() ([]byte, error) {
    type Alias Reschedulingmanagementunitresponse

    if ReschedulingmanagementunitresponseMarshalled {
        return []byte("{}"), nil
    }
    ReschedulingmanagementunitresponseMarshalled = true

    return json.Marshal(&struct { 
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        Applied bool `json:"applied"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

