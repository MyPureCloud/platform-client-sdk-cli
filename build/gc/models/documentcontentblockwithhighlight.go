package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentcontentblockwithhighlightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentcontentblockwithhighlightDud struct { 
    


    


    


    


    

}

// Documentcontentblockwithhighlight
type Documentcontentblockwithhighlight struct { 
    // VarType - The type of the paragraph block.
    VarType string `json:"type"`


    // Text - Text. It must contain a value if the type of the block is Text.
    Text Documenttext `json:"text"`


    // Image - Image. It must contain a value if the type of the block is Image.
    Image Documentbodyimage `json:"image"`


    // Video - Video. It must contain a value if the type of the block is Video.
    Video Documentbodyvideo `json:"video"`


    // AnswerHighlight - The block highlight data.
    AnswerHighlight Documentcontenthighlightindex `json:"answerHighlight"`

}

// String returns a JSON representation of the model
func (o *Documentcontentblockwithhighlight) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentcontentblockwithhighlight) MarshalJSON() ([]byte, error) {
    type Alias Documentcontentblockwithhighlight

    if DocumentcontentblockwithhighlightMarshalled {
        return []byte("{}"), nil
    }
    DocumentcontentblockwithhighlightMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text Documenttext `json:"text"`
        
        Image Documentbodyimage `json:"image"`
        
        Video Documentbodyvideo `json:"video"`
        
        AnswerHighlight Documentcontenthighlightindex `json:"answerHighlight"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

