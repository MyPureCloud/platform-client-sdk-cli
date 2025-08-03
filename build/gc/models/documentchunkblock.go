package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentchunkblockMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentchunkblockDud struct { 
    


    


    


    

}

// Documentchunkblock
type Documentchunkblock struct { 
    // Id - The globally unique identifier for the chunk.
    Id string `json:"id"`


    // Text - The chunk text associated with the variation.
    Text string `json:"text"`


    // Confidence - The confidence associated with a chunk with respect to a search query.
    Confidence float32 `json:"confidence"`


    // Document - Reference to document associated with a chunk
    Document Documentchunkreference `json:"document"`

}

// String returns a JSON representation of the model
func (o *Documentchunkblock) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentchunkblock) MarshalJSON() ([]byte, error) {
    type Alias Documentchunkblock

    if DocumentchunkblockMarshalled {
        return []byte("{}"), nil
    }
    DocumentchunkblockMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Text string `json:"text"`
        
        Confidence float32 `json:"confidence"`
        
        Document Documentchunkreference `json:"document"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

