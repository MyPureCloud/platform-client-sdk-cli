package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FieldmappingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FieldmappingDud struct { 
    


    

}

// Fieldmapping
type Fieldmapping struct { 
    // DataActionValueName - The data action output field name that maps to the list item values.
    DataActionValueName string `json:"dataActionValueName"`


    // DataActionSynonymName - The data action output field name that maps to the list item synonyms. Optional if synonyms are not provided by the data action.
    DataActionSynonymName string `json:"dataActionSynonymName"`

}

// String returns a JSON representation of the model
func (o *Fieldmapping) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Fieldmapping) MarshalJSON() ([]byte, error) {
    type Alias Fieldmapping

    if FieldmappingMarshalled {
        return []byte("{}"), nil
    }
    FieldmappingMarshalled = true

    return json.Marshal(&struct {
        
        DataActionValueName string `json:"dataActionValueName"`
        
        DataActionSynonymName string `json:"dataActionSynonymName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

