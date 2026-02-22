package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationdataschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationdataschemaDud struct { 
    


    


    


    DateCreated time.Time `json:"dateCreated"`


    CreatedBy Domainentityref `json:"createdBy"`


    


    SelfUri string `json:"selfUri"`

}

// Conversationdataschema
type Conversationdataschema struct { 
    // Id - The globally unique identifier for the schema.
    Id string `json:"id"`


    // Version - The schema's version, a positive integer.
    Version int `json:"version"`


    // Enabled - The schema's enabled/disabled status. A disabled schema cannot be assigned to any other entities, but the data on those entities from the schema still exists.
    Enabled bool `json:"enabled"`


    


    


    // JsonSchema - A JSON schema defining the extension to the built-in entity type.
    JsonSchema Conversationjsonschemadocument `json:"jsonSchema"`


    

}

// String returns a JSON representation of the model
func (o *Conversationdataschema) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationdataschema) MarshalJSON() ([]byte, error) {
    type Alias Conversationdataschema

    if ConversationdataschemaMarshalled {
        return []byte("{}"), nil
    }
    ConversationdataschemaMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Version int `json:"version"`
        
        Enabled bool `json:"enabled"`
        
        JsonSchema Conversationjsonschemadocument `json:"jsonSchema"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

