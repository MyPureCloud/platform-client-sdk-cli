package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyexternaleventsschemalistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyexternaleventsschemalistingDud struct { 
    


    


    

}

// Journeyexternaleventsschemalisting
type Journeyexternaleventsschemalisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Journeyexternaleventsschema `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Journeyexternaleventsschemalisting) String() string {
    
     o.Entities = []Journeyexternaleventsschema{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyexternaleventsschemalisting) MarshalJSON() ([]byte, error) {
    type Alias Journeyexternaleventsschemalisting

    if JourneyexternaleventsschemalistingMarshalled {
        return []byte("{}"), nil
    }
    JourneyexternaleventsschemalistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Journeyexternaleventsschema `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Journeyexternaleventsschema{{}},
        


        

        Alias: (*Alias)(u),
    })
}

