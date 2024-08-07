package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonebaseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonebaseDud struct { 
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

// Phonebase
type Phonebase struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // PhoneMetaBase - A phone metabase is essentially a database for storing phone configuration settings, which simplifies the configuration process.
    PhoneMetaBase Domainentityref `json:"phoneMetaBase"`


    // Lines - The list of linebases associated with the phone base.
    Lines []Linebase `json:"lines"`


    // Properties
    Properties map[string]interface{} `json:"properties"`


    // Capabilities
    Capabilities Phonecapabilities `json:"capabilities"`


    

}

// String returns a JSON representation of the model
func (o *Phonebase) String() string {
    
    
    
    
    
     o.Lines = []Linebase{{}} 
     o.Properties = map[string]interface{}{"": Interface{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonebase) MarshalJSON() ([]byte, error) {
    type Alias Phonebase

    if PhonebaseMarshalled {
        return []byte("{}"), nil
    }
    PhonebaseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        PhoneMetaBase Domainentityref `json:"phoneMetaBase"`
        
        Lines []Linebase `json:"lines"`
        
        Properties map[string]interface{} `json:"properties"`
        
        Capabilities Phonecapabilities `json:"capabilities"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Lines: []Linebase{{}},
        


        
        Properties: map[string]interface{}{"": Interface{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

