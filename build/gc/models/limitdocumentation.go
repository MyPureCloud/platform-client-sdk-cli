package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LimitdocumentationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LimitdocumentationDud struct { 
    


    

}

// Limitdocumentation
type Limitdocumentation struct { 
    // Url
    Url string `json:"url"`


    // Namespaces
    Namespaces []Namespacedocs `json:"namespaces"`

}

// String returns a JSON representation of the model
func (o *Limitdocumentation) String() string {
    
     o.Namespaces = []Namespacedocs{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Limitdocumentation) MarshalJSON() ([]byte, error) {
    type Alias Limitdocumentation

    if LimitdocumentationMarshalled {
        return []byte("{}"), nil
    }
    LimitdocumentationMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        Namespaces []Namespacedocs `json:"namespaces"`
        *Alias
    }{

        


        
        Namespaces: []Namespacedocs{{}},
        

        Alias: (*Alias)(u),
    })
}

