package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserobservationdatacontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserobservationdatacontainerDud struct { 
    


    

}

// Userobservationdatacontainer
type Userobservationdatacontainer struct { 
    // Group - A mapping from dimension to value
    Group map[string]string `json:"group"`


    // Data
    Data []Observationmetricdata `json:"data"`

}

// String returns a JSON representation of the model
func (o *Userobservationdatacontainer) String() string {
     o.Group = map[string]string{"": ""} 
     o.Data = []Observationmetricdata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userobservationdatacontainer) MarshalJSON() ([]byte, error) {
    type Alias Userobservationdatacontainer

    if UserobservationdatacontainerMarshalled {
        return []byte("{}"), nil
    }
    UserobservationdatacontainerMarshalled = true

    return json.Marshal(&struct {
        
        Group map[string]string `json:"group"`
        
        Data []Observationmetricdata `json:"data"`
        *Alias
    }{

        
        Group: map[string]string{"": ""},
        


        
        Data: []Observationmetricdata{{}},
        

        Alias: (*Alias)(u),
    })
}

