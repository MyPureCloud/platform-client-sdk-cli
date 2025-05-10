package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableversionentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableversionentityDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Decisiontableversionentity
type Decisiontableversionentity struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Version - The decision table version.
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Decisiontableversionentity) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableversionentity) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableversionentity

    if DecisiontableversionentityMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableversionentityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

