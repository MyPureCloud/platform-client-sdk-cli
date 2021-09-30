package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActioninputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActioninputDud struct { 
    


    


    

}

// Actioninput - Input requirements of Action.
type Actioninput struct { 
    // InputSchema - JSON Schema that defines the body of the request that the client (edge/architect/postman) is sending to the service, on the /execute path. If the 'flatten' query parameter is omitted or false, this field will be returned. Either inputSchema or inputSchemaFlattened will be returned, not both.
    InputSchema Jsonschemadocument `json:"inputSchema"`


    // InputSchemaFlattened - JSON Schema that defines the body of the request that the client (edge/architect/postman) is sending to the service, on the /execute path. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either inputSchema or inputSchemaFlattened will be returned, not both.
    InputSchemaFlattened Jsonschemadocument `json:"inputSchemaFlattened"`


    // InputSchemaUri - The URI of the input schema
    InputSchemaUri string `json:"inputSchemaUri"`

}

// String returns a JSON representation of the model
func (o *Actioninput) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actioninput) MarshalJSON() ([]byte, error) {
    type Alias Actioninput

    if ActioninputMarshalled {
        return []byte("{}"), nil
    }
    ActioninputMarshalled = true

    return json.Marshal(&struct { 
        InputSchema Jsonschemadocument `json:"inputSchema"`
        
        InputSchemaFlattened Jsonschemadocument `json:"inputSchemaFlattened"`
        
        InputSchemaUri string `json:"inputSchemaUri"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

