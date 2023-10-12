package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamactivitydataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamactivitydataDud struct { 
    


    


    


    

}

// Teamactivitydata
type Teamactivitydata struct { 
    // Group - A mapping from grouping dimension to value
    Group map[string]string `json:"group"`


    // Data - Data for metrics
    Data []Teamactivitymetricvalue `json:"data"`


    // Truncated - Flag for a truncated list of entities. If truncated, the first half of the list of entities will contain the oldest entities and the second half the newest entities.
    Truncated bool `json:"truncated"`


    // Entities - Details for active entities
    Entities []Teamactivityentitydata `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Teamactivitydata) String() string {
     o.Group = map[string]string{"": ""} 
     o.Data = []Teamactivitymetricvalue{{}} 
    
     o.Entities = []Teamactivityentitydata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamactivitydata) MarshalJSON() ([]byte, error) {
    type Alias Teamactivitydata

    if TeamactivitydataMarshalled {
        return []byte("{}"), nil
    }
    TeamactivitydataMarshalled = true

    return json.Marshal(&struct {
        
        Group map[string]string `json:"group"`
        
        Data []Teamactivitymetricvalue `json:"data"`
        
        Truncated bool `json:"truncated"`
        
        Entities []Teamactivityentitydata `json:"entities"`
        *Alias
    }{

        
        Group: map[string]string{"": ""},
        


        
        Data: []Teamactivitymetricvalue{{}},
        


        


        
        Entities: []Teamactivityentitydata{{}},
        

        Alias: (*Alias)(u),
    })
}

