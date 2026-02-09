package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsercustomattributesupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsercustomattributesupdaterequestDud struct { 
    


    


    

}

// Usercustomattributesupdaterequest
type Usercustomattributesupdaterequest struct { 
    // SchemaId - The id of the schema that dictates which attributes can be included.
    SchemaId string `json:"schemaId"`


    // SchemaVersion - The version of the schema.
    SchemaVersion int `json:"schemaVersion"`


    // Attributes - The map of attribute values.
    Attributes map[string]interface{} `json:"attributes"`

}

// String returns a JSON representation of the model
func (o *Usercustomattributesupdaterequest) String() string {
    
    
     o.Attributes = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usercustomattributesupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Usercustomattributesupdaterequest

    if UsercustomattributesupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    UsercustomattributesupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        SchemaId string `json:"schemaId"`
        
        SchemaVersion int `json:"schemaVersion"`
        
        Attributes map[string]interface{} `json:"attributes"`
        *Alias
    }{

        


        


        
        Attributes: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

