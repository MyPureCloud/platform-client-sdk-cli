package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NamespacedocsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NamespacedocsDud struct { 
    


    

}

// Namespacedocs
type Namespacedocs struct { 
    // FriendlyName
    FriendlyName string `json:"friendlyName"`


    // Limits
    Limits []Limitdocs `json:"limits"`

}

// String returns a JSON representation of the model
func (o *Namespacedocs) String() string {
    
     o.Limits = []Limitdocs{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Namespacedocs) MarshalJSON() ([]byte, error) {
    type Alias Namespacedocs

    if NamespacedocsMarshalled {
        return []byte("{}"), nil
    }
    NamespacedocsMarshalled = true

    return json.Marshal(&struct {
        
        FriendlyName string `json:"friendlyName"`
        
        Limits []Limitdocs `json:"limits"`
        *Alias
    }{

        


        
        Limits: []Limitdocs{{}},
        

        Alias: (*Alias)(u),
    })
}

