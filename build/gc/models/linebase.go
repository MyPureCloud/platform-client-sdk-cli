package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LinebaseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LinebaseDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    SelfUri string `json:"selfUri"`

}

// Linebase
type Linebase struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // LineMetaBase
    LineMetaBase Domainentityref `json:"lineMetaBase"`


    // Properties
    Properties map[string]interface{} `json:"properties"`


    

}

// String returns a JSON representation of the model
func (o *Linebase) String() string {
    
    
    
    
    
     o.Properties = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Linebase) MarshalJSON() ([]byte, error) {
    type Alias Linebase

    if LinebaseMarshalled {
        return []byte("{}"), nil
    }
    LinebaseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        LineMetaBase Domainentityref `json:"lineMetaBase"`
        
        Properties map[string]interface{} `json:"properties"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Properties: map[string]interface{}{"": Interface{}},
        


        

        Alias: (*Alias)(u),
    })
}

