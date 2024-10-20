package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumenttablecontentblockwithhighlightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumenttablecontentblockwithhighlightDud struct { 
    


    


    


    


    


    


    


    

}

// Documenttablecontentblockwithhighlight
type Documenttablecontentblockwithhighlight struct { 
    // VarType - The type of the block for the table cell. This determines which body block object (paragraph, list, video, image or table) would have a value.
    VarType string `json:"type"`


    // Text - Text. It must contain a value if the type of the block is Text.
    Text Documenttext `json:"text"`


    // Image - Image. It must contain a value if the type of the block is Image.
    Image Documentbodyimage `json:"image"`


    // Video - Video. It must contain a value if the type of the block is Video.
    Video Documentbodyvideo `json:"video"`


    // Paragraph - Paragraph. It must contain a value if the type of the block is Paragraph.
    Paragraph Documentbodyparagraphwithhighlight `json:"paragraph"`


    // List - List. It must contain a value if the type of the block is UnorderedList or OrderedList.
    List Documentbodylistwithhighlight `json:"list"`


    // Table - Table. It must contain a value if the type of the block is Table.
    Table Documentbodytablewithhighlight `json:"table"`


    // AnswerHighlight - The block highlight data.
    AnswerHighlight Documentcontenthighlightindex `json:"answerHighlight"`

}

// String returns a JSON representation of the model
func (o *Documenttablecontentblockwithhighlight) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documenttablecontentblockwithhighlight) MarshalJSON() ([]byte, error) {
    type Alias Documenttablecontentblockwithhighlight

    if DocumenttablecontentblockwithhighlightMarshalled {
        return []byte("{}"), nil
    }
    DocumenttablecontentblockwithhighlightMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text Documenttext `json:"text"`
        
        Image Documentbodyimage `json:"image"`
        
        Video Documentbodyvideo `json:"video"`
        
        Paragraph Documentbodyparagraphwithhighlight `json:"paragraph"`
        
        List Documentbodylistwithhighlight `json:"list"`
        
        Table Documentbodytablewithhighlight `json:"table"`
        
        AnswerHighlight Documentcontenthighlightindex `json:"answerHighlight"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

