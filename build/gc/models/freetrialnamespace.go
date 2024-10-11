package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FreetrialnamespaceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FreetrialnamespaceDud struct { 
    


    


    

}

// Freetrialnamespace
type Freetrialnamespace struct { 
    // Name
    Name string `json:"name"`


    // FriendlyName
    FriendlyName string `json:"friendlyName"`


    // Limits
    Limits []Freetriallimit `json:"limits"`

}

// String returns a JSON representation of the model
func (o *Freetrialnamespace) String() string {
    
    
     o.Limits = []Freetriallimit{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Freetrialnamespace) MarshalJSON() ([]byte, error) {
    type Alias Freetrialnamespace

    if FreetrialnamespaceMarshalled {
        return []byte("{}"), nil
    }
    FreetrialnamespaceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        FriendlyName string `json:"friendlyName"`
        
        Limits []Freetriallimit `json:"limits"`
        *Alias
    }{

        


        


        
        Limits: []Freetriallimit{{}},
        

        Alias: (*Alias)(u),
    })
}

