package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Contentquickreplyv2Marshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Contentquickreplyv2Dud struct { 
    


    

}

// Contentquickreplyv2 - Quick reply object V2.
type Contentquickreplyv2 struct { 
    // Title - Text to show as the title of the quick reply.
    Title string `json:"title"`


    // Actions - An array of quick reply objects.
    Actions []Contentquickreply `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Contentquickreplyv2) String() string {
    
     o.Actions = []Contentquickreply{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentquickreplyv2) MarshalJSON() ([]byte, error) {
    type Alias Contentquickreplyv2

    if Contentquickreplyv2Marshalled {
        return []byte("{}"), nil
    }
    Contentquickreplyv2Marshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Actions []Contentquickreply `json:"actions"`
        *Alias
    }{

        


        
        Actions: []Contentquickreply{{}},
        

        Alias: (*Alias)(u),
    })
}

