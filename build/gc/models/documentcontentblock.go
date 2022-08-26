package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentcontentblockMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentcontentblockDud struct { 
    


    


    

}

// Documentcontentblock
type Documentcontentblock struct { 
    // VarType - The type of the paragraph block.
    VarType string `json:"type"`


    // Text - Text. It must contain a value if the type of the block is Text.
    Text Documenttext `json:"text"`


    // Image - Image. It must contain a value if the type of the block is Image.
    Image Documentbodyimage `json:"image"`

}

// String returns a JSON representation of the model
func (o *Documentcontentblock) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentcontentblock) MarshalJSON() ([]byte, error) {
    type Alias Documentcontentblock

    if DocumentcontentblockMarshalled {
        return []byte("{}"), nil
    }
    DocumentcontentblockMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text Documenttext `json:"text"`
        
        Image Documentbodyimage `json:"image"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

