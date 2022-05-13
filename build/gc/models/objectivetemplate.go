package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ObjectivetemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ObjectivetemplateDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Objectivetemplate
type Objectivetemplate struct { 
    


    // Name
    Name string `json:"name"`


    // Zones
    Zones []Objectivezone `json:"zones"`


    

}

// String returns a JSON representation of the model
func (o *Objectivetemplate) String() string {
    
     o.Zones = []Objectivezone{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Objectivetemplate) MarshalJSON() ([]byte, error) {
    type Alias Objectivetemplate

    if ObjectivetemplateMarshalled {
        return []byte("{}"), nil
    }
    ObjectivetemplateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Zones []Objectivezone `json:"zones"`
        *Alias
    }{

        


        


        
        Zones: []Objectivezone{{}},
        


        

        Alias: (*Alias)(u),
    })
}

