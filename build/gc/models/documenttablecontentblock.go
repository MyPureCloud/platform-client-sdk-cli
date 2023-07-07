package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumenttablecontentblockMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumenttablecontentblockDud struct { 
    


    


    


    


    


    


    

}

// Documenttablecontentblock
type Documenttablecontentblock struct { 
    // VarType - The type of the block for the table cell. This determines which body block object (paragraph, list, video, image or table) would have a value.
    VarType string `json:"type"`


    // Paragraph - Paragraph. It must contain a value if the type of the block is Paragraph.
    Paragraph Documentbodyparagraph `json:"paragraph"`


    // Text - Text. It must contain a value if the type of the block is Text.
    Text Documenttext `json:"text"`


    // Image - Image. It must contain a value if the type of the block is Image.
    Image Documentbodyimage `json:"image"`


    // Video - Video. It must contain a value if the type of the block is Video.
    Video Documentbodyvideo `json:"video"`


    // List - List. It must contain a value if the type of the block is UnorderedList or OrderedList.
    List Documentbodylist `json:"list"`


    // Table - Table. It must contain a value if the type of the block is Table.
    Table Documentbodytable `json:"table"`

}

// String returns a JSON representation of the model
func (o *Documenttablecontentblock) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documenttablecontentblock) MarshalJSON() ([]byte, error) {
    type Alias Documenttablecontentblock

    if DocumenttablecontentblockMarshalled {
        return []byte("{}"), nil
    }
    DocumenttablecontentblockMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Paragraph Documentbodyparagraph `json:"paragraph"`
        
        Text Documenttext `json:"text"`
        
        Image Documentbodyimage `json:"image"`
        
        Video Documentbodyvideo `json:"video"`
        
        List Documentbodylist `json:"list"`
        
        Table Documentbodytable `json:"table"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

