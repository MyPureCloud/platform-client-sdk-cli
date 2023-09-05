package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IvrdivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IvrdivisionviewDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Ivrdivisionview - An ivr.
type Ivrdivisionview struct { 
    // Id - The ivr identifier
    Id string `json:"id"`


    // Name - The ivr name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Ivrdivisionview) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ivrdivisionview) MarshalJSON() ([]byte, error) {
    type Alias Ivrdivisionview

    if IvrdivisionviewMarshalled {
        return []byte("{}"), nil
    }
    IvrdivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

