package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsercustomattributesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsercustomattributesDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Usercustomattributes
type Usercustomattributes struct { 
    


    // Name
    Name string `json:"name"`


    // UserId - The user's Id which the attributes are being assigned to.
    UserId string `json:"userId"`


    // Schema - The schema that dictates which attributes can be included.
    Schema Dataschema `json:"schema"`


    // Attributes - The map of attribute values.
    Attributes map[string]interface{} `json:"attributes"`


    

}

// String returns a JSON representation of the model
func (o *Usercustomattributes) String() string {
    
    
    
     o.Attributes = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usercustomattributes) MarshalJSON() ([]byte, error) {
    type Alias Usercustomattributes

    if UsercustomattributesMarshalled {
        return []byte("{}"), nil
    }
    UsercustomattributesMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        UserId string `json:"userId"`
        
        Schema Dataschema `json:"schema"`
        
        Attributes map[string]interface{} `json:"attributes"`
        *Alias
    }{

        


        


        


        


        
        Attributes: map[string]interface{}{"": Interface{}},
        


        

        Alias: (*Alias)(u),
    })
}

