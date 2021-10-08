package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EndpointMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EndpointDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Endpoint
type Endpoint struct { 
    


    // Name - Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    // DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The ID of the user that last modified the resource.
    ModifiedBy string `json:"modifiedBy"`


    // CreatedBy - The ID of the user that created the resource.
    CreatedBy string `json:"createdBy"`


    


    // ModifiedByApp - The application that last modified the resource.
    ModifiedByApp string `json:"modifiedByApp"`


    // CreatedByApp - The application that created the resource.
    CreatedByApp string `json:"createdByApp"`


    // Count
    Count int `json:"count"`


    // Properties
    Properties map[string]interface{} `json:"properties"`


    // Schema - Schema
    Schema Domainentityref `json:"schema"`


    // Enabled
    Enabled bool `json:"enabled"`


    // Site
    Site Domainentityref `json:"site"`


    // Dids
    Dids []string `json:"dids"`


    

}

// String returns a JSON representation of the model
func (o *Endpoint) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Properties = map[string]interface{}{"": Interface{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Dids = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Endpoint) MarshalJSON() ([]byte, error) {
    type Alias Endpoint

    if EndpointMarshalled {
        return []byte("{}"), nil
    }
    EndpointMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        
        
        
        ModifiedByApp string `json:"modifiedByApp"`
        
        CreatedByApp string `json:"createdByApp"`
        
        Count int `json:"count"`
        
        Properties map[string]interface{} `json:"properties"`
        
        Schema Domainentityref `json:"schema"`
        
        Enabled bool `json:"enabled"`
        
        Site Domainentityref `json:"site"`
        
        Dids []string `json:"dids"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Properties: map[string]interface{}{"": Interface{}},
        

        

        

        

        

        

        

        

        
        Dids: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

