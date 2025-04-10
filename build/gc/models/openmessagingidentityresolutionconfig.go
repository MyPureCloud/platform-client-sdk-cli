package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenmessagingidentityresolutionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenmessagingidentityresolutionconfigDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Openmessagingidentityresolutionconfig
type Openmessagingidentityresolutionconfig struct { 
    


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    // ResolveIdentities - Whether the channel should resolve identities
    ResolveIdentities bool `json:"resolveIdentities"`


    // ExternalSource - The external source used for stitching this channel - used only for Open Messaging.
    ExternalSource Addressableentityref `json:"externalSource"`


    

}

// String returns a JSON representation of the model
func (o *Openmessagingidentityresolutionconfig) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openmessagingidentityresolutionconfig) MarshalJSON() ([]byte, error) {
    type Alias Openmessagingidentityresolutionconfig

    if OpenmessagingidentityresolutionconfigMarshalled {
        return []byte("{}"), nil
    }
    OpenmessagingidentityresolutionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        ResolveIdentities bool `json:"resolveIdentities"`
        
        ExternalSource Addressableentityref `json:"externalSource"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

