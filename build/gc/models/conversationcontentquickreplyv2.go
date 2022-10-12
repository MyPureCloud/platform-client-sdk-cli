package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Conversationcontentquickreplyv2Marshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Conversationcontentquickreplyv2Dud struct { 
    


    

}

// Conversationcontentquickreplyv2 - Quick reply object V2.
type Conversationcontentquickreplyv2 struct { 
    // Title - Text to show as the title of the quick reply.
    Title string `json:"title"`


    // Actions - An array of quick reply objects.
    Actions []Conversationcontentquickreply `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentquickreplyv2) String() string {
    
     o.Actions = []Conversationcontentquickreply{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentquickreplyv2) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentquickreplyv2

    if Conversationcontentquickreplyv2Marshalled {
        return []byte("{}"), nil
    }
    Conversationcontentquickreplyv2Marshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Actions []Conversationcontentquickreply `json:"actions"`
        *Alias
    }{

        


        
        Actions: []Conversationcontentquickreply{{}},
        

        Alias: (*Alias)(u),
    })
}

