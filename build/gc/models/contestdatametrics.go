package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestdatametricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestdatametricsDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    SelfUri string `json:"selfUri"`

}

// Contestdatametrics
type Contestdatametrics struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Contestdatametrics) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestdatametrics) MarshalJSON() ([]byte, error) {
    type Alias Contestdatametrics

    if ContestdatametricsMarshalled {
        return []byte("{}"), nil
    }
    ContestdatametricsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

