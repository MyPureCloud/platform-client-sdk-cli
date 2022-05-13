package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NluconfusionmatrixrowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NluconfusionmatrixrowDud struct { 
    


    

}

// Nluconfusionmatrixrow
type Nluconfusionmatrixrow struct { 
    // Name - The name of the intent for the row.
    Name string `json:"name"`


    // Columns - The columns of confusion matrix for the intent
    Columns []Nluconfusionmatrixcolumn `json:"columns"`

}

// String returns a JSON representation of the model
func (o *Nluconfusionmatrixrow) String() string {
    
     o.Columns = []Nluconfusionmatrixcolumn{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nluconfusionmatrixrow) MarshalJSON() ([]byte, error) {
    type Alias Nluconfusionmatrixrow

    if NluconfusionmatrixrowMarshalled {
        return []byte("{}"), nil
    }
    NluconfusionmatrixrowMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Columns []Nluconfusionmatrixcolumn `json:"columns"`
        *Alias
    }{

        


        
        Columns: []Nluconfusionmatrixcolumn{{}},
        

        Alias: (*Alias)(u),
    })
}

