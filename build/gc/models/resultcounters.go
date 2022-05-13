package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResultcountersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResultcountersDud struct { 
    


    

}

// Resultcounters
type Resultcounters struct { 
    // Success
    Success int `json:"success"`


    // Failure
    Failure int `json:"failure"`

}

// String returns a JSON representation of the model
func (o *Resultcounters) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Resultcounters) MarshalJSON() ([]byte, error) {
    type Alias Resultcounters

    if ResultcountersMarshalled {
        return []byte("{}"), nil
    }
    ResultcountersMarshalled = true

    return json.Marshal(&struct {
        
        Success int `json:"success"`
        
        Failure int `json:"failure"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

