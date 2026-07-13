package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RelationshipsexportfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RelationshipsexportfilterDud struct { 
    


    


    


    


    


    


    

}

// Relationshipsexportfilter
type Relationshipsexportfilter struct { 
    // Eq - Filtered field should have the same value
    Eq Relationshipsexportfieldfilter `json:"eq"`


    // In - Filtered field should match one of the listed values
    In Relationshipsexportfieldlistfilter `json:"in"`


    // Lte - Filtered field should be less than or equal to the value
    Lte Relationshipsexportcomparisonfieldfilter `json:"lte"`


    // Gte - Filtered field should be greater than or equal to the value
    Gte Relationshipsexportcomparisonfieldfilter `json:"gte"`


    // And - Boolean AND combination of filters
    And []Relationshipsexportfilter `json:"and"`


    // Or - Boolean OR combination of filters
    Or []Relationshipsexportfilter `json:"or"`


    // Not - Boolean negation of filters
    Not *Relationshipsexportfilter `json:"not"`

}

// String returns a JSON representation of the model
func (o *Relationshipsexportfilter) String() string {
    
    
    
    
     o.And = []Relationshipsexportfilter{{}} 
     o.Or = []Relationshipsexportfilter{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Relationshipsexportfilter) MarshalJSON() ([]byte, error) {
    type Alias Relationshipsexportfilter

    if RelationshipsexportfilterMarshalled {
        return []byte("{}"), nil
    }
    RelationshipsexportfilterMarshalled = true

    return json.Marshal(&struct {
        
        Eq Relationshipsexportfieldfilter `json:"eq"`
        
        In Relationshipsexportfieldlistfilter `json:"in"`
        
        Lte Relationshipsexportcomparisonfieldfilter `json:"lte"`
        
        Gte Relationshipsexportcomparisonfieldfilter `json:"gte"`
        
        And []Relationshipsexportfilter `json:"and"`
        
        Or []Relationshipsexportfilter `json:"or"`
        
        Not *Relationshipsexportfilter `json:"not"`
        *Alias
    }{

        


        


        


        


        
        And: []Relationshipsexportfilter{{}},
        


        
        Or: []Relationshipsexportfilter{{}},
        


        

        Alias: (*Alias)(u),
    })
}

