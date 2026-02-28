package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeretrievedreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeretrievedreferenceDud struct { 
    


    


    


    

}

// Knowledgeretrievedreference
type Knowledgeretrievedreference struct { 
    // Confidence - The confidence associated with retrieved reference respect to a search query.
    Confidence float64 `json:"confidence"`


    // Text - The matching text for search query.
    Text string `json:"text"`


    // FileName - The file name from which reference is retrieved
    FileName string `json:"fileName"`


    // Url - The url of the file.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Knowledgeretrievedreference) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeretrievedreference) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeretrievedreference

    if KnowledgeretrievedreferenceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeretrievedreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Confidence float64 `json:"confidence"`
        
        Text string `json:"text"`
        
        FileName string `json:"fileName"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

