package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UtilizationlabelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UtilizationlabelDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Utilizationlabel
type Utilizationlabel struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Utilizationlabel) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Utilizationlabel) MarshalJSON() ([]byte, error) {
    type Alias Utilizationlabel

    if UtilizationlabelMarshalled {
        return []byte("{}"), nil
    }
    UtilizationlabelMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

