package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DraftvalidationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DraftvalidationresultDud struct { 
    


    

}

// Draftvalidationresult - Validation results
type Draftvalidationresult struct { 
    // Valid - Indicates if configuration is valid
    Valid bool `json:"valid"`


    // Errors - List of errors causing validation failure
    Errors []Errorbody `json:"errors"`

}

// String returns a JSON representation of the model
func (o *Draftvalidationresult) String() string {
    
    
    
    
    
    
     o.Errors = []Errorbody{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Draftvalidationresult) MarshalJSON() ([]byte, error) {
    type Alias Draftvalidationresult

    if DraftvalidationresultMarshalled {
        return []byte("{}"), nil
    }
    DraftvalidationresultMarshalled = true

    return json.Marshal(&struct { 
        Valid bool `json:"valid"`
        
        Errors []Errorbody `json:"errors"`
        
        *Alias
    }{
        

        

        

        
        Errors: []Errorbody{{}},
        

        
        Alias: (*Alias)(u),
    })
}

