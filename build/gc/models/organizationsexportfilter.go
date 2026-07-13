package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationsexportfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationsexportfilterDud struct { 
    


    


    


    


    


    


    

}

// Organizationsexportfilter
type Organizationsexportfilter struct { 
    // Eq - Filtered field should have the same value
    Eq Organizationsexportfieldfilter `json:"eq"`


    // In - Filtered field should match one of the listed values
    In Organizationsexportfieldlistfilter `json:"in"`


    // Lte - Filtered field should be less than or equal to the value
    Lte Organizationsexportcomparisonfieldfilter `json:"lte"`


    // Gte - Filtered field should be greater than or equal to the value
    Gte Organizationsexportcomparisonfieldfilter `json:"gte"`


    // And - Boolean AND combination of filters
    And []Organizationsexportfilter `json:"and"`


    // Or - Boolean OR combination of filters
    Or []Organizationsexportfilter `json:"or"`


    // Not - Boolean negation of filters
    Not *Organizationsexportfilter `json:"not"`

}

// String returns a JSON representation of the model
func (o *Organizationsexportfilter) String() string {
    
    
    
    
     o.And = []Organizationsexportfilter{{}} 
     o.Or = []Organizationsexportfilter{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationsexportfilter) MarshalJSON() ([]byte, error) {
    type Alias Organizationsexportfilter

    if OrganizationsexportfilterMarshalled {
        return []byte("{}"), nil
    }
    OrganizationsexportfilterMarshalled = true

    return json.Marshal(&struct {
        
        Eq Organizationsexportfieldfilter `json:"eq"`
        
        In Organizationsexportfieldlistfilter `json:"in"`
        
        Lte Organizationsexportcomparisonfieldfilter `json:"lte"`
        
        Gte Organizationsexportcomparisonfieldfilter `json:"gte"`
        
        And []Organizationsexportfilter `json:"and"`
        
        Or []Organizationsexportfilter `json:"or"`
        
        Not *Organizationsexportfilter `json:"not"`
        *Alias
    }{

        


        


        


        


        
        And: []Organizationsexportfilter{{}},
        


        
        Or: []Organizationsexportfilter{{}},
        


        

        Alias: (*Alias)(u),
    })
}

