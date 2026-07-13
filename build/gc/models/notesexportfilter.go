package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NotesexportfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NotesexportfilterDud struct { 
    


    


    


    


    


    


    

}

// Notesexportfilter
type Notesexportfilter struct { 
    // Eq - Filtered field should have the same value
    Eq Notesexportfieldfilter `json:"eq"`


    // In - Filtered field should match one of the listed values
    In Notesexportfieldlistfilter `json:"in"`


    // Lte - Filtered field should be less than or equal to the value
    Lte Notesexportcomparisonfieldfilter `json:"lte"`


    // Gte - Filtered field should be greater than or equal to the value
    Gte Notesexportcomparisonfieldfilter `json:"gte"`


    // And - Boolean AND combination of filters
    And []Notesexportfilter `json:"and"`


    // Or - Boolean OR combination of filters
    Or []Notesexportfilter `json:"or"`


    // Not - Boolean negation of filters
    Not *Notesexportfilter `json:"not"`

}

// String returns a JSON representation of the model
func (o *Notesexportfilter) String() string {
    
    
    
    
     o.And = []Notesexportfilter{{}} 
     o.Or = []Notesexportfilter{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Notesexportfilter) MarshalJSON() ([]byte, error) {
    type Alias Notesexportfilter

    if NotesexportfilterMarshalled {
        return []byte("{}"), nil
    }
    NotesexportfilterMarshalled = true

    return json.Marshal(&struct {
        
        Eq Notesexportfieldfilter `json:"eq"`
        
        In Notesexportfieldlistfilter `json:"in"`
        
        Lte Notesexportcomparisonfieldfilter `json:"lte"`
        
        Gte Notesexportcomparisonfieldfilter `json:"gte"`
        
        And []Notesexportfilter `json:"and"`
        
        Or []Notesexportfilter `json:"or"`
        
        Not *Notesexportfilter `json:"not"`
        *Alias
    }{

        


        


        


        


        
        And: []Notesexportfilter{{}},
        


        
        Or: []Notesexportfilter{{}},
        


        

        Alias: (*Alias)(u),
    })
}

