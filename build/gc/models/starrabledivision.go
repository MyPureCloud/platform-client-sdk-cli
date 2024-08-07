package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StarrabledivisionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StarrabledivisionDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Starrabledivision
type Starrabledivision struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Starrabledivision) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Starrabledivision) MarshalJSON() ([]byte, error) {
    type Alias Starrabledivision

    if StarrabledivisionMarshalled {
        return []byte("{}"), nil
    }
    StarrabledivisionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

