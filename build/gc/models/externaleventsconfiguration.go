package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternaleventsconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternaleventsconfigurationDud struct { 
    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Externaleventsconfiguration
type Externaleventsconfiguration struct { 
    // Id - The unique identifier for the external event configuration.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Description - A description of the external event configuration.
    Description string `json:"description"`


    // DivisionId - The division ID (UUID) associated with this configuration.
    DivisionId string `json:"divisionId"`


    // DivisionIdActive - Indicates whether the divisionId field is valid.
    DivisionIdActive bool `json:"divisionIdActive"`


    // SchemaId - The dynamic schema ID (UUID) that defines the structure of external events.
    SchemaId string `json:"schemaId"`


    // SchemaActive - Indicates whether the schema is active or inactive.
    SchemaActive bool `json:"schemaActive"`


    // Source - The source of the external events e.g. Adobe, Salesforce.
    Source string `json:"source"`


    // DateLastModified - The timestamp when the configuration was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateLastModified time.Time `json:"dateLastModified"`


    

}

// String returns a JSON representation of the model
func (o *Externaleventsconfiguration) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externaleventsconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Externaleventsconfiguration

    if ExternaleventsconfigurationMarshalled {
        return []byte("{}"), nil
    }
    ExternaleventsconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DivisionId string `json:"divisionId"`
        
        DivisionIdActive bool `json:"divisionIdActive"`
        
        SchemaId string `json:"schemaId"`
        
        SchemaActive bool `json:"schemaActive"`
        
        Source string `json:"source"`
        
        DateLastModified time.Time `json:"dateLastModified"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

