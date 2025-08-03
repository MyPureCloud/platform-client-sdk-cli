package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentvariationquerychunkblockMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentvariationquerychunkblockDud struct { 
    


    

}

// Documentvariationquerychunkblock
type Documentvariationquerychunkblock struct { 
    // Id - The globally unique identifier for the chunk.
    Id string `json:"id"`


    // Text - The chunk text associated with the variation.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Documentvariationquerychunkblock) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentvariationquerychunkblock) MarshalJSON() ([]byte, error) {
    type Alias Documentvariationquerychunkblock

    if DocumentvariationquerychunkblockMarshalled {
        return []byte("{}"), nil
    }
    DocumentvariationquerychunkblockMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Text string `json:"text"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

