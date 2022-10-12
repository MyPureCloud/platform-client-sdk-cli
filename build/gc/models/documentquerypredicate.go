package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentquerypredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentquerypredicateDud struct { 
    


    


    

}

// Documentquerypredicate
type Documentquerypredicate struct { 
    // Fields - Specifies the document fields to be matched against.
    Fields []string `json:"fields"`


    // Values - Specifies the values of the fields to be matched against.
    Values []string `json:"values"`


    // VarType - Specifies the matching criteria between the fields and values.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Documentquerypredicate) String() string {
     o.Fields = []string{""} 
     o.Values = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentquerypredicate) MarshalJSON() ([]byte, error) {
    type Alias Documentquerypredicate

    if DocumentquerypredicateMarshalled {
        return []byte("{}"), nil
    }
    DocumentquerypredicateMarshalled = true

    return json.Marshal(&struct {
        
        Fields []string `json:"fields"`
        
        Values []string `json:"values"`
        
        VarType string `json:"type"`
        *Alias
    }{

        
        Fields: []string{""},
        


        
        Values: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

