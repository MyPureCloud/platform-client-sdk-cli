package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateroomrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateroomrequestDud struct { 
    


    

}

// Createroomrequest
type Createroomrequest struct { 
    // Description - Room's description
    Description string `json:"description"`


    // Subject - Room's subject
    Subject string `json:"subject"`

}

// String returns a JSON representation of the model
func (o *Createroomrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createroomrequest) MarshalJSON() ([]byte, error) {
    type Alias Createroomrequest

    if CreateroomrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateroomrequestMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        
        Subject string `json:"subject"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

