package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AttributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AttributeDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Attribute
type Attribute struct { 
    


    // Name - The attribute name.
    Name string `json:"name"`


    // Version
    Version int `json:"version"`


    // Description
    Description string `json:"description"`


    // CreatedBy
    CreatedBy Domainentityref `json:"createdBy"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // ModifiedBy
    ModifiedBy Domainentityref `json:"modifiedBy"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Attribute) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Attribute) MarshalJSON() ([]byte, error) {
    type Alias Attribute

    if AttributeMarshalled {
        return []byte("{}"), nil
    }
    AttributeMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Description string `json:"description"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        ModifiedBy Domainentityref `json:"modifiedBy"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

