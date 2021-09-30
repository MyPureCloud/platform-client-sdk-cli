package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SectionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SectionDud struct { 
    


    


    


    

}

// Section
type Section struct { 
    // FieldList
    FieldList []Fieldlist `json:"fieldList"`


    // InstructionText
    InstructionText string `json:"instructionText"`


    // Key
    Key string `json:"key"`


    // State
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Section) String() string {
    
    
     o.FieldList = []Fieldlist{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Section) MarshalJSON() ([]byte, error) {
    type Alias Section

    if SectionMarshalled {
        return []byte("{}"), nil
    }
    SectionMarshalled = true

    return json.Marshal(&struct { 
        FieldList []Fieldlist `json:"fieldList"`
        
        InstructionText string `json:"instructionText"`
        
        Key string `json:"key"`
        
        State string `json:"state"`
        
        *Alias
    }{
        

        
        FieldList: []Fieldlist{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

