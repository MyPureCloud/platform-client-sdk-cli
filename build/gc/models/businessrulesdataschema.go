package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessrulesdataschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessrulesdataschemaDud struct { 
    


    


    


    AppliesTo []string `json:"appliesTo"`


    


    CreatedBy Domainentityref `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    


    SelfUri string `json:"selfUri"`

}

// Businessrulesdataschema
type Businessrulesdataschema struct { 
    // Id - The globally unique identifier for the schema.  Only required if a schema is used for custom fields during external entity creation or updates.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Version - The schema's version, a positive integer. Required for updates.
    Version int `json:"version"`


    


    // Enabled - The schema's enabled/disabled status. A disabled schema cannot be assigned to any other entities, but the data on those entities from the schema still exists.
    Enabled bool `json:"enabled"`


    


    


    // JsonSchema - A JSON schema defining the extension to the built-in entity type.
    JsonSchema Jsonschemadocument `json:"jsonSchema"`


    

}

// String returns a JSON representation of the model
func (o *Businessrulesdataschema) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessrulesdataschema) MarshalJSON() ([]byte, error) {
    type Alias Businessrulesdataschema

    if BusinessrulesdataschemaMarshalled {
        return []byte("{}"), nil
    }
    BusinessrulesdataschemaMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Enabled bool `json:"enabled"`
        
        JsonSchema Jsonschemadocument `json:"jsonSchema"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

