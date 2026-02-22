package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationjsonschemarequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationjsonschemarequestDud struct { 
    


    


    


    


    

}

// Conversationjsonschemarequest - A JSON Schema for create/update requests.
type Conversationjsonschemarequest struct { 
    // Schema - The JSON Schema specification link. The only value currently supported is \"http://json-schema.org/draft-04/schema#\".
    Schema string `json:"$schema"`


    // Title - The title of the schema. Must be unique across all enabled Custom Attributes schemas.
    Title string `json:"title"`


    // Description - The schema description.
    Description string `json:"description"`


    // Required - The list of required schema properties. All fields are optional unless listed. New fields added after initial schema creation must be optional before being able to update to required.
    Required []string `json:"required"`


    // Properties - The map of schema properties and their limits.
    Properties map[string]interface{} `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Conversationjsonschemarequest) String() string {
    
    
    
     o.Required = []string{""} 
     o.Properties = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationjsonschemarequest) MarshalJSON() ([]byte, error) {
    type Alias Conversationjsonschemarequest

    if ConversationjsonschemarequestMarshalled {
        return []byte("{}"), nil
    }
    ConversationjsonschemarequestMarshalled = true

    return json.Marshal(&struct {
        
        Schema string `json:"$schema"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        Required []string `json:"required"`
        
        Properties map[string]interface{} `json:"properties"`
        *Alias
    }{

        


        


        


        
        Required: []string{""},
        


        
        Properties: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

