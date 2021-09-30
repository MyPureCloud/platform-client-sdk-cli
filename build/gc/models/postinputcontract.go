package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PostinputcontractMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PostinputcontractDud struct { 
    

}

// Postinputcontract - The schemas defining all of the expected requests/inputs.
type Postinputcontract struct { 
    // InputSchema - JSON Schema that defines the body of the request that the client (edge/architect/postman) is sending to the service, on the /execute path.
    InputSchema Jsonschemadocument `json:"inputSchema"`

}

// String returns a JSON representation of the model
func (o *Postinputcontract) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Postinputcontract) MarshalJSON() ([]byte, error) {
    type Alias Postinputcontract

    if PostinputcontractMarshalled {
        return []byte("{}"), nil
    }
    PostinputcontractMarshalled = true

    return json.Marshal(&struct { 
        InputSchema Jsonschemadocument `json:"inputSchema"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

