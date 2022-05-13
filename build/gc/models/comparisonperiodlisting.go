package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ComparisonperiodlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ComparisonperiodlistingDud struct { 
    


    


    

}

// Comparisonperiodlisting
type Comparisonperiodlisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Comparisonperiod `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Comparisonperiodlisting) String() string {
    
     o.Entities = []Comparisonperiod{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Comparisonperiodlisting) MarshalJSON() ([]byte, error) {
    type Alias Comparisonperiodlisting

    if ComparisonperiodlistingMarshalled {
        return []byte("{}"), nil
    }
    ComparisonperiodlistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Comparisonperiod `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Comparisonperiod{{}},
        


        

        Alias: (*Alias)(u),
    })
}

