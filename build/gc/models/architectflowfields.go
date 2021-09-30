package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArchitectflowfieldsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArchitectflowfieldsDud struct { 
    


    

}

// Architectflowfields
type Architectflowfields struct { 
    // ArchitectFlow - The architect flow.
    ArchitectFlow Addressableentityref `json:"architectFlow"`


    // FlowRequestMappings - Collection of Architect Flow Request Mappings to use.
    FlowRequestMappings []Requestmapping `json:"flowRequestMappings"`

}

// String returns a JSON representation of the model
func (o *Architectflowfields) String() string {
    
    
    
    
    
    
     o.FlowRequestMappings = []Requestmapping{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Architectflowfields) MarshalJSON() ([]byte, error) {
    type Alias Architectflowfields

    if ArchitectflowfieldsMarshalled {
        return []byte("{}"), nil
    }
    ArchitectflowfieldsMarshalled = true

    return json.Marshal(&struct { 
        ArchitectFlow Addressableentityref `json:"architectFlow"`
        
        FlowRequestMappings []Requestmapping `json:"flowRequestMappings"`
        
        *Alias
    }{
        

        

        

        
        FlowRequestMappings: []Requestmapping{{}},
        

        
        Alias: (*Alias)(u),
    })
}

