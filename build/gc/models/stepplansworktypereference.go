package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StepplansworktypereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StepplansworktypereferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Stepplansworktypereference
type Stepplansworktypereference struct { 
    // Id - The id of the worktype.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Stepplansworktypereference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stepplansworktypereference) MarshalJSON() ([]byte, error) {
    type Alias Stepplansworktypereference

    if StepplansworktypereferenceMarshalled {
        return []byte("{}"), nil
    }
    StepplansworktypereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

