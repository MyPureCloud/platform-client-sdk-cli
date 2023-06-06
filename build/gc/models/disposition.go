package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DispositionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DispositionDud struct { 
    


    


    

}

// Disposition
type Disposition struct { 
    // Name - Name of the disposition. Either a platform predefined value, or the name of the disposition in the disposition table..
    Name string `json:"name"`


    // Analyzer - The final media analyzer result that triggered the disposition result, if any.
    Analyzer string `json:"analyzer"`


    // DispositionParameters - Contains various parameters related to call analysis.
    DispositionParameters Dispositionparameters `json:"dispositionParameters"`

}

// String returns a JSON representation of the model
func (o *Disposition) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Disposition) MarshalJSON() ([]byte, error) {
    type Alias Disposition

    if DispositionMarshalled {
        return []byte("{}"), nil
    }
    DispositionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Analyzer string `json:"analyzer"`
        
        DispositionParameters Dispositionparameters `json:"dispositionParameters"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

