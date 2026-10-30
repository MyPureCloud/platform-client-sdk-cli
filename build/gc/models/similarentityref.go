package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SimilarentityrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SimilarentityrefDud struct { 
    


    


    


    DateCreated time.Time `json:"dateCreated"`


    SelfUri string `json:"selfUri"`

}

// Similarentityref
type Similarentityref struct { 
    // Id - The unique identifier of the entity
    Id string `json:"id"`


    // Name - The name of the entity
    Name string `json:"name"`


    // Description - The description of the entity
    Description string `json:"description"`


    


    

}

// String returns a JSON representation of the model
func (o *Similarentityref) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Similarentityref) MarshalJSON() ([]byte, error) {
    type Alias Similarentityref

    if SimilarentityrefMarshalled {
        return []byte("{}"), nil
    }
    SimilarentityrefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

