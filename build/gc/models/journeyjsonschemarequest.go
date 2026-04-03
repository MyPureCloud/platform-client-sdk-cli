package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyjsonschemarequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyjsonschemarequestDud struct { 
    


    


    


    


    

}

// Journeyjsonschemarequest - A JSON Schema for create/update requests.
type Journeyjsonschemarequest struct { 
    // Schema - The JSON Schema specification link. The only value currently supported is \"http://json-schema.org/draft-04/schema#\".
    Schema string `json:"$schema"`


    // Title - The title of the schema. Must be unique across all enabled External Event schemas.
    Title string `json:"title"`


    // Description - The schema description.
    Description string `json:"description"`


    // Required - The list of required schema properties. All fields are optional unless listed. Optional fields can't be changed to required after the schema is saved.
    Required []string `json:"required"`


    // Properties - The map of schema properties and their limits.
    Properties map[string]interface{} `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Journeyjsonschemarequest) String() string {
    
    
    
     o.Required = []string{""} 
     o.Properties = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyjsonschemarequest) MarshalJSON() ([]byte, error) {
    type Alias Journeyjsonschemarequest

    if JourneyjsonschemarequestMarshalled {
        return []byte("{}"), nil
    }
    JourneyjsonschemarequestMarshalled = true

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

