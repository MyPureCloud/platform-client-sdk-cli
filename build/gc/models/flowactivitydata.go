package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowactivitydataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowactivitydataDud struct { 
    


    


    


    

}

// Flowactivitydata
type Flowactivitydata struct { 
    // Group - A mapping from grouping dimension to value
    Group map[string]string `json:"group"`


    // Data - Data for metrics
    Data []Flowactivitymetricvalue `json:"data"`


    // Truncated - Flag for a truncated list of entities. If truncated, the first half of the list of entities will contain the oldest entities and the second half the newest entities.
    Truncated bool `json:"truncated"`


    // Entities - Details for active entities
    Entities []Flowactivityentitydata `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Flowactivitydata) String() string {
     o.Group = map[string]string{"": ""} 
     o.Data = []Flowactivitymetricvalue{{}} 
    
     o.Entities = []Flowactivityentitydata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowactivitydata) MarshalJSON() ([]byte, error) {
    type Alias Flowactivitydata

    if FlowactivitydataMarshalled {
        return []byte("{}"), nil
    }
    FlowactivitydataMarshalled = true

    return json.Marshal(&struct {
        
        Group map[string]string `json:"group"`
        
        Data []Flowactivitymetricvalue `json:"data"`
        
        Truncated bool `json:"truncated"`
        
        Entities []Flowactivityentitydata `json:"entities"`
        *Alias
    }{

        
        Group: map[string]string{"": ""},
        


        
        Data: []Flowactivitymetricvalue{{}},
        


        


        
        Entities: []Flowactivityentitydata{{}},
        

        Alias: (*Alias)(u),
    })
}

