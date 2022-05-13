package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesearchresponseDud struct { 
    SearchId string `json:"searchId"`


    Total int `json:"total"`


    PageCount int `json:"pageCount"`


    PageSize int `json:"pageSize"`


    PageNumber int `json:"pageNumber"`


    Results []Knowledgesearchdocument `json:"results"`

}

// Knowledgesearchresponse
type Knowledgesearchresponse struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgesearchresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesearchresponse

    if KnowledgesearchresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesearchresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

