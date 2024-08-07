package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivitydataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivitydataDud struct { 
    


    


    


    

}

// Useractivitydata
type Useractivitydata struct { 
    // Group - A mapping from grouping dimension to value
    Group map[string]string `json:"group"`


    // Data - Data for metrics
    Data []Useractivitymetricvalue `json:"data"`


    // Truncated - Flag for a truncated list of entities. If truncated, the first half of the list of entities will contain the oldest entities and the second half the newest entities.
    Truncated bool `json:"truncated"`


    // Entities - Details for active entities
    Entities []Useractivityentitydata `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Useractivitydata) String() string {
     o.Group = map[string]string{"": ""} 
     o.Data = []Useractivitymetricvalue{{}} 
    
     o.Entities = []Useractivityentitydata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivitydata) MarshalJSON() ([]byte, error) {
    type Alias Useractivitydata

    if UseractivitydataMarshalled {
        return []byte("{}"), nil
    }
    UseractivitydataMarshalled = true

    return json.Marshal(&struct {
        
        Group map[string]string `json:"group"`
        
        Data []Useractivitymetricvalue `json:"data"`
        
        Truncated bool `json:"truncated"`
        
        Entities []Useractivityentitydata `json:"entities"`
        *Alias
    }{

        
        Group: map[string]string{"": ""},
        


        
        Data: []Useractivitymetricvalue{{}},
        


        


        
        Entities: []Useractivityentitydata{{}},
        

        Alias: (*Alias)(u),
    })
}

