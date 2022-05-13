package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssociatedvaluefieldMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssociatedvaluefieldDud struct { 
    


    

}

// Associatedvaluefield
type Associatedvaluefield struct { 
    // DataType - The data type of the value field.
    DataType string `json:"dataType"`


    // Name - The field name for extracting value from event.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Associatedvaluefield) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Associatedvaluefield) MarshalJSON() ([]byte, error) {
    type Alias Associatedvaluefield

    if AssociatedvaluefieldMarshalled {
        return []byte("{}"), nil
    }
    AssociatedvaluefieldMarshalled = true

    return json.Marshal(&struct {
        
        DataType string `json:"dataType"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

