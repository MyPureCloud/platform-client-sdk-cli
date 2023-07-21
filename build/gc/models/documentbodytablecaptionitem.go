package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytablecaptionitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytablecaptionitemDud struct { 
    


    


    


    


    


    

}

// Documentbodytablecaptionitem
type Documentbodytablecaptionitem struct { 
    // VarType - The type of the caption item.
    VarType string `json:"type"`


    // Text - Text. It must contain a value if the type of the block is Text.
    Text Documenttext `json:"text"`


    // Paragraph - Paragraph. It must contain a value if the type of the block is Paragraph.
    Paragraph Documentbodyparagraph `json:"paragraph"`


    // Image - Image. It must contain a value if the type of the block is Image.
    Image Documentbodyimage `json:"image"`


    // Video - Video. It must contain a value if the type of the block is Video.
    Video Documentbodyvideo `json:"video"`


    // List - List. It must contain a value if the type of the block is UnorderedList or OrderedList.
    List Documentbodylist `json:"list"`

}

// String returns a JSON representation of the model
func (o *Documentbodytablecaptionitem) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytablecaptionitem) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytablecaptionitem

    if DocumentbodytablecaptionitemMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytablecaptionitemMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text Documenttext `json:"text"`
        
        Paragraph Documentbodyparagraph `json:"paragraph"`
        
        Image Documentbodyimage `json:"image"`
        
        Video Documentbodyvideo `json:"video"`
        
        List Documentbodylist `json:"list"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

