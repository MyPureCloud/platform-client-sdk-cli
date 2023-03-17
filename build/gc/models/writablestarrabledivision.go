package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WritablestarrabledivisionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WritablestarrabledivisionDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Writablestarrabledivision
type Writablestarrabledivision struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Writablestarrabledivision) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Writablestarrabledivision) MarshalJSON() ([]byte, error) {
    type Alias Writablestarrabledivision

    if WritablestarrabledivisionMarshalled {
        return []byte("{}"), nil
    }
    WritablestarrabledivisionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

