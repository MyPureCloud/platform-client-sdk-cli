package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CursoradherenceadjustmentslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CursoradherenceadjustmentslistingDud struct { 
    


    


    


    

}

// Cursoradherenceadjustmentslisting
type Cursoradherenceadjustmentslisting struct { 
    // Entities
    Entities []Adherenceadjustment `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Cursoradherenceadjustmentslisting) String() string {
     o.Entities = []Adherenceadjustment{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cursoradherenceadjustmentslisting) MarshalJSON() ([]byte, error) {
    type Alias Cursoradherenceadjustmentslisting

    if CursoradherenceadjustmentslistingMarshalled {
        return []byte("{}"), nil
    }
    CursoradherenceadjustmentslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Adherenceadjustment `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Adherenceadjustment{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

