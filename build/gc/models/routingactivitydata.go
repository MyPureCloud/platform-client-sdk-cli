package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingactivitydataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingactivitydataDud struct { 
    


    


    


    

}

// Routingactivitydata
type Routingactivitydata struct { 
    // Group - A mapping from grouping dimension to value
    Group map[string]string `json:"group"`


    // Data - Data for metrics
    Data []Routingactivitymetricvalue `json:"data"`


    // Truncated - Flag for a truncated list of entities. If truncated, the first half of the list of entities will contain the oldest entities and the second half the newest entities.
    Truncated bool `json:"truncated"`


    // Entities - Details for active entities
    Entities []Routingactivityentitydata `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Routingactivitydata) String() string {
     o.Group = map[string]string{"": ""} 
     o.Data = []Routingactivitymetricvalue{{}} 
    
     o.Entities = []Routingactivityentitydata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingactivitydata) MarshalJSON() ([]byte, error) {
    type Alias Routingactivitydata

    if RoutingactivitydataMarshalled {
        return []byte("{}"), nil
    }
    RoutingactivitydataMarshalled = true

    return json.Marshal(&struct {
        
        Group map[string]string `json:"group"`
        
        Data []Routingactivitymetricvalue `json:"data"`
        
        Truncated bool `json:"truncated"`
        
        Entities []Routingactivityentitydata `json:"entities"`
        *Alias
    }{

        
        Group: map[string]string{"": ""},
        


        
        Data: []Routingactivitymetricvalue{{}},
        


        


        
        Entities: []Routingactivityentitydata{{}},
        

        Alias: (*Alias)(u),
    })
}

