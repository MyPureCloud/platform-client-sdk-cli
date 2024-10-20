package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentlistcontentblockwithhighlightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentlistcontentblockwithhighlightDud struct { 
    


    


    


    


    


    

}

// Documentlistcontentblockwithhighlight
type Documentlistcontentblockwithhighlight struct { 
    // VarType - The type of the list block.
    VarType string `json:"type"`


    // Text - Text. It must contain a value if the type of the block is Text.
    Text Documenttext `json:"text"`


    // Image - Image. It must contain a value if the type of the block is Image.
    Image Documentbodyimage `json:"image"`


    // Video - Video. It must contain a value if the type of the block is Video.
    Video Documentbodyvideo `json:"video"`


    // List - List. It must contain a value if the type of the block is UnorderedList or OrderedList.
    List Documentbodylistwithhighlight `json:"list"`


    // AnswerHighlight - The block highlight data.
    AnswerHighlight Documentcontenthighlightindex `json:"answerHighlight"`

}

// String returns a JSON representation of the model
func (o *Documentlistcontentblockwithhighlight) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentlistcontentblockwithhighlight) MarshalJSON() ([]byte, error) {
    type Alias Documentlistcontentblockwithhighlight

    if DocumentlistcontentblockwithhighlightMarshalled {
        return []byte("{}"), nil
    }
    DocumentlistcontentblockwithhighlightMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text Documenttext `json:"text"`
        
        Image Documentbodyimage `json:"image"`
        
        Video Documentbodyvideo `json:"video"`
        
        List Documentbodylistwithhighlight `json:"list"`
        
        AnswerHighlight Documentcontenthighlightindex `json:"answerHighlight"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

