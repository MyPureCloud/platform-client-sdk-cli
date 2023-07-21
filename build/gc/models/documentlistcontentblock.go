package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentlistcontentblockMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentlistcontentblockDud struct { 
    


    


    


    


    

}

// Documentlistcontentblock
type Documentlistcontentblock struct { 
    // VarType - The type of the list block.
    VarType string `json:"type"`


    // Text - Text. It must contain a value if the type of the block is Text.
    Text Documenttext `json:"text"`


    // Image - Image. It must contain a value if the type of the block is Image.
    Image Documentbodyimage `json:"image"`


    // List - List. It must contain a value if the type of the block is UnorderedList or OrderedList.
    List Documentbodylist `json:"list"`


    // Video - Video. It must contain a value if the type of the block is Video.
    Video Documentbodyvideo `json:"video"`

}

// String returns a JSON representation of the model
func (o *Documentlistcontentblock) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentlistcontentblock) MarshalJSON() ([]byte, error) {
    type Alias Documentlistcontentblock

    if DocumentlistcontentblockMarshalled {
        return []byte("{}"), nil
    }
    DocumentlistcontentblockMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text Documenttext `json:"text"`
        
        Image Documentbodyimage `json:"image"`
        
        List Documentbodylist `json:"list"`
        
        Video Documentbodyvideo `json:"video"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

