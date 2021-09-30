package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NluconfusionmatrixcolumnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NluconfusionmatrixcolumnDud struct { 
    


    

}

// Nluconfusionmatrixcolumn
type Nluconfusionmatrixcolumn struct { 
    // Name - The name of the intent for the column.
    Name string `json:"name"`


    // Value - The confusion value between the intents
    Value float32 `json:"value"`

}

// String returns a JSON representation of the model
func (o *Nluconfusionmatrixcolumn) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nluconfusionmatrixcolumn) MarshalJSON() ([]byte, error) {
    type Alias Nluconfusionmatrixcolumn

    if NluconfusionmatrixcolumnMarshalled {
        return []byte("{}"), nil
    }
    NluconfusionmatrixcolumnMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Value float32 `json:"value"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

