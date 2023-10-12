package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CriteriacategoryinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CriteriacategoryinfoDud struct { 
    


    

}

// Criteriacategoryinfo
type Criteriacategoryinfo struct { 
    // CategoryId
    CategoryId string `json:"categoryId"`


    // DisplayOrder
    DisplayOrder int `json:"displayOrder"`

}

// String returns a JSON representation of the model
func (o *Criteriacategoryinfo) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Criteriacategoryinfo) MarshalJSON() ([]byte, error) {
    type Alias Criteriacategoryinfo

    if CriteriacategoryinfoMarshalled {
        return []byte("{}"), nil
    }
    CriteriacategoryinfoMarshalled = true

    return json.Marshal(&struct {
        
        CategoryId string `json:"categoryId"`
        
        DisplayOrder int `json:"displayOrder"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

