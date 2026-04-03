package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternaleventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternaleventDud struct { 
    


    


    


    

}

// Externalevent
type Externalevent struct { 
    // Id - The ID of the event.
    Id string `json:"id"`


    // DateCreated - Timestamp indicating when the event actually took place. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // Identifiers - The identifiers of the contact the event is for.
    Identifiers Externaleventidentifiers `json:"identifiers"`


    // Attributes - The event attributes.
    Attributes map[string]interface{} `json:"attributes"`

}

// String returns a JSON representation of the model
func (o *Externalevent) String() string {
    
    
    
     o.Attributes = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalevent) MarshalJSON() ([]byte, error) {
    type Alias Externalevent

    if ExternaleventMarshalled {
        return []byte("{}"), nil
    }
    ExternaleventMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        Identifiers Externaleventidentifiers `json:"identifiers"`
        
        Attributes map[string]interface{} `json:"attributes"`
        *Alias
    }{

        


        


        


        
        Attributes: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

