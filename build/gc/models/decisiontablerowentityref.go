package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontablerowentityrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontablerowentityrefDud struct { 
    


    

}

// Decisiontablerowentityref
type Decisiontablerowentityref struct { 
    // Id
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Decisiontablerowentityref) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontablerowentityref) MarshalJSON() ([]byte, error) {
    type Alias Decisiontablerowentityref

    if DecisiontablerowentityrefMarshalled {
        return []byte("{}"), nil
    }
    DecisiontablerowentityrefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

