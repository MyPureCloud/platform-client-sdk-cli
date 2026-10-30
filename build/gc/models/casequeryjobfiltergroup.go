package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasequeryjobfiltergroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasequeryjobfiltergroupDud struct { 
    


    

}

// Casequeryjobfiltergroup
type Casequeryjobfiltergroup struct { 
    // Operator - Boolean operator used to combine the filters within this group.
    Operator string `json:"operator"`


    // Filters - List of filter objects combined using the group operator.
    Filters []Casequeryjobfilter `json:"filters"`

}

// String returns a JSON representation of the model
func (o *Casequeryjobfiltergroup) String() string {
    
     o.Filters = []Casequeryjobfilter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casequeryjobfiltergroup) MarshalJSON() ([]byte, error) {
    type Alias Casequeryjobfiltergroup

    if CasequeryjobfiltergroupMarshalled {
        return []byte("{}"), nil
    }
    CasequeryjobfiltergroupMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Filters []Casequeryjobfilter `json:"filters"`
        *Alias
    }{

        


        
        Filters: []Casequeryjobfilter{{}},
        

        Alias: (*Alias)(u),
    })
}

