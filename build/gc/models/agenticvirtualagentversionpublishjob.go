package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentversionpublishjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentversionpublishjobDud struct { 
    Id string `json:"id"`


    Status string `json:"status"`


    Errors []Errorbody `json:"errors"`


    SelfUri string `json:"selfUri"`

}

// Agenticvirtualagentversionpublishjob
type Agenticvirtualagentversionpublishjob struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentversionpublishjob) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentversionpublishjob) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentversionpublishjob

    if AgenticvirtualagentversionpublishjobMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentversionpublishjobMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

