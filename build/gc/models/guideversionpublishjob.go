package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuideversionpublishjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuideversionpublishjobDud struct { 
    Id string `json:"id"`


    Status string `json:"status"`


    Errors []Errorbody `json:"errors"`


    


    


    SelfUri string `json:"selfUri"`

}

// Guideversionpublishjob
type Guideversionpublishjob struct { 
    


    


    


    // Guide
    Guide Addressableentityref `json:"guide"`


    // GuideVersion
    GuideVersion Guideversion `json:"guideVersion"`


    

}

// String returns a JSON representation of the model
func (o *Guideversionpublishjob) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guideversionpublishjob) MarshalJSON() ([]byte, error) {
    type Alias Guideversionpublishjob

    if GuideversionpublishjobMarshalled {
        return []byte("{}"), nil
    }
    GuideversionpublishjobMarshalled = true

    return json.Marshal(&struct {
        
        Guide Addressableentityref `json:"guide"`
        
        GuideVersion Guideversion `json:"guideVersion"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

