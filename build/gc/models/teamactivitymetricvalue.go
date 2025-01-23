package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamactivitymetricvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamactivitymetricvalueDud struct { 
    


    


    


    


    

}

// Teamactivitymetricvalue
type Teamactivitymetricvalue struct { 
    // Metric - Metric
    Metric string `json:"metric"`


    // Qualifier - Metric qualifier
    Qualifier string `json:"qualifier"`


    // SecondaryQualifier - Secondary metric qualifier
    SecondaryQualifier string `json:"secondaryQualifier"`


    // EntityIds - Entity ids for matching entities if details were requested
    EntityIds []string `json:"entityIds"`


    // Count - Metric count
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Teamactivitymetricvalue) String() string {
    
    
    
     o.EntityIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamactivitymetricvalue) MarshalJSON() ([]byte, error) {
    type Alias Teamactivitymetricvalue

    if TeamactivitymetricvalueMarshalled {
        return []byte("{}"), nil
    }
    TeamactivitymetricvalueMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        Qualifier string `json:"qualifier"`
        
        SecondaryQualifier string `json:"secondaryQualifier"`
        
        EntityIds []string `json:"entityIds"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        


        


        
        EntityIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

