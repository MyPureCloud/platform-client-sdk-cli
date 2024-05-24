package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PausecriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PausecriteriaDud struct { 
    


    

}

// Pausecriteria
type Pausecriteria struct { 
    // UrlFragment
    UrlFragment string `json:"urlFragment"`


    // Condition
    Condition string `json:"condition"`

}

// String returns a JSON representation of the model
func (o *Pausecriteria) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pausecriteria) MarshalJSON() ([]byte, error) {
    type Alias Pausecriteria

    if PausecriteriaMarshalled {
        return []byte("{}"), nil
    }
    PausecriteriaMarshalled = true

    return json.Marshal(&struct {
        
        UrlFragment string `json:"urlFragment"`
        
        Condition string `json:"condition"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

