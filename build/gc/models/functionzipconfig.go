package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FunctionzipconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FunctionzipconfigDud struct { 
    


    


    


    

}

// Functionzipconfig - Action function current zip file upload settings and state.
type Functionzipconfig struct { 
    // Status - Status of zip upload.
    Status string `json:"status"`


    // Id - Zip file Identifier
    Id string `json:"id"`


    // Name - Zip file name
    Name string `json:"name"`


    // DateCreated - Date and time zip record was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`

}

// String returns a JSON representation of the model
func (o *Functionzipconfig) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Functionzipconfig) MarshalJSON() ([]byte, error) {
    type Alias Functionzipconfig

    if FunctionzipconfigMarshalled {
        return []byte("{}"), nil
    }
    FunctionzipconfigMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        DateCreated time.Time `json:"dateCreated"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

