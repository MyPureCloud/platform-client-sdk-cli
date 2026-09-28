package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CurrentagentcursoradherenceadjustmentslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CurrentagentcursoradherenceadjustmentslistingDud struct { 
    


    


    


    

}

// Currentagentcursoradherenceadjustmentslisting
type Currentagentcursoradherenceadjustmentslisting struct { 
    // Entities
    Entities []Currentagentadherenceadjustment `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Currentagentcursoradherenceadjustmentslisting) String() string {
     o.Entities = []Currentagentadherenceadjustment{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Currentagentcursoradherenceadjustmentslisting) MarshalJSON() ([]byte, error) {
    type Alias Currentagentcursoradherenceadjustmentslisting

    if CurrentagentcursoradherenceadjustmentslistingMarshalled {
        return []byte("{}"), nil
    }
    CurrentagentcursoradherenceadjustmentslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Currentagentadherenceadjustment `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Currentagentadherenceadjustment{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

