package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodyblockMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodyblockDud struct { 
    


    


    


    


    


    

}

// Documentbodyblock
type Documentbodyblock struct { 
    // VarType - The type of the block for the body. This determines which body block object (paragraph, list, video, image or table) would have a value.
    VarType string `json:"type"`


    // Image - Image. It must contain a value if the type of the block is Image.
    Image Documentbodyimage `json:"image"`


    // Video - Video. It must contain a value if the type of the block is Video.
    Video Documentbodyvideo `json:"video"`


    // Paragraph - Paragraph. It must contain a value if the type of the block is Paragraph.
    Paragraph Documentbodyparagraph `json:"paragraph"`


    // List - List. It must contain a value if the type of the block is UnorderedList or OrderedList.
    List Documentbodylist `json:"list"`


    // Table - Table. It must contain a value if type of the block is Table.
    Table Documentbodytable `json:"table"`

}

// String returns a JSON representation of the model
func (o *Documentbodyblock) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodyblock) MarshalJSON() ([]byte, error) {
    type Alias Documentbodyblock

    if DocumentbodyblockMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodyblockMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Image Documentbodyimage `json:"image"`
        
        Video Documentbodyvideo `json:"video"`
        
        Paragraph Documentbodyparagraph `json:"paragraph"`
        
        List Documentbodylist `json:"list"`
        
        Table Documentbodytable `json:"table"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

