package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumenttextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumenttextDud struct { 
    


    


    


    

}

// Documenttext
type Documenttext struct { 
    // Text - Text.
    Text string `json:"text"`


    // Marks - The unique list of marks (whether it is bold and/or underlined etc.) for the text.
    Marks []string `json:"marks"`


    // Hyperlink - The URL of the page OR an email OR the reference to the knowledge article that the hyperlink goes to. Possible URL value types are https://<url link> | mailto:<email> | grn:knowledge:::documentVariation/<knowledgeBaseId>/<documentId>/<variationId> | grn:knowledge:::document/<knowledgeBaseId>/<documentId> | grn:knowledge:::category/<knowledgeBaseId>/<categoryId> | grn:knowledge:::label/<knowledgeBaseId>/<labelId>
    Hyperlink string `json:"hyperlink"`


    // Properties - The properties for the text.
    Properties Documenttextproperties `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Documenttext) String() string {
    
     o.Marks = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documenttext) MarshalJSON() ([]byte, error) {
    type Alias Documenttext

    if DocumenttextMarshalled {
        return []byte("{}"), nil
    }
    DocumenttextMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Marks []string `json:"marks"`
        
        Hyperlink string `json:"hyperlink"`
        
        Properties Documenttextproperties `json:"properties"`
        *Alias
    }{

        


        
        Marks: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

