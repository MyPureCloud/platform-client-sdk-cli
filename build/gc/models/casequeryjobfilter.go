package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasequeryjobfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasequeryjobfilterDud struct { 
    


    


    

}

// Casequeryjobfilter
type Casequeryjobfilter struct { 
    // Name - Attribute name. Valid filter names are: 'caseplanId', 'ownerId', 'status', 'priority', 'dateDue', 'externalContactId', 'customerIntentId', 'dateCreated', 'divisionId', 'reference'.
    Name string `json:"name"`


    // Operator - Filter operator.
    Operator string `json:"operator"`


    // Values - List of values to be used in the filter.
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Casequeryjobfilter) String() string {
    
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casequeryjobfilter) MarshalJSON() ([]byte, error) {
    type Alias Casequeryjobfilter

    if CasequeryjobfilterMarshalled {
        return []byte("{}"), nil
    }
    CasequeryjobfilterMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Operator string `json:"operator"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

