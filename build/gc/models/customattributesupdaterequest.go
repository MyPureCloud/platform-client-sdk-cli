package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomattributesupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomattributesupdaterequestDud struct { 
    


    


    


    


    

}

// Customattributesupdaterequest
type Customattributesupdaterequest struct { 
    // Id - Unique identifier for the Custom Attributes record. IDs are created by users.
    Id string `json:"id"`


    // Divisions - The list of division ids. Use [] if divisions aren't used (Unassigned Division). Omitting or setting to [] clears existing values on update.
    Divisions []string `json:"divisions"`


    // SchemaId - The id of the schema that dictates which attributes can be included. Required for create, cannot be updated.
    SchemaId string `json:"schemaId"`


    // Version - The latest version of the Custom Attributes record. Optional for concurrency check on update.
    Version int `json:"version"`


    // CustomAttributes - The map of attribute values.
    CustomAttributes map[string]interface{} `json:"customAttributes"`

}

// String returns a JSON representation of the model
func (o *Customattributesupdaterequest) String() string {
    
     o.Divisions = []string{""} 
    
    
     o.CustomAttributes = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customattributesupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Customattributesupdaterequest

    if CustomattributesupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    CustomattributesupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Divisions []string `json:"divisions"`
        
        SchemaId string `json:"schemaId"`
        
        Version int `json:"version"`
        
        CustomAttributes map[string]interface{} `json:"customAttributes"`
        *Alias
    }{

        


        
        Divisions: []string{""},
        


        


        


        
        CustomAttributes: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

