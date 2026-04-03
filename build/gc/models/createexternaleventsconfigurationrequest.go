package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateexternaleventsconfigurationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateexternaleventsconfigurationrequestDud struct { 
    


    


    


    


    

}

// Createexternaleventsconfigurationrequest
type Createexternaleventsconfigurationrequest struct { 
    // Name - The name of the external event configuration.
    Name string `json:"name"`


    // Description - A description of the external event configuration.
    Description string `json:"description"`


    // DivisionId - The division ID associated with this configuration.
    DivisionId string `json:"divisionId"`


    // SchemaId - The dynamic schema ID that defines the structure of external events.
    SchemaId string `json:"schemaId"`


    // Source - The source of the external events e.g. Adobe, Salesforce. This cannot be changed after creation.
    Source string `json:"source"`

}

// String returns a JSON representation of the model
func (o *Createexternaleventsconfigurationrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createexternaleventsconfigurationrequest) MarshalJSON() ([]byte, error) {
    type Alias Createexternaleventsconfigurationrequest

    if CreateexternaleventsconfigurationrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateexternaleventsconfigurationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DivisionId string `json:"divisionId"`
        
        SchemaId string `json:"schemaId"`
        
        Source string `json:"source"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

