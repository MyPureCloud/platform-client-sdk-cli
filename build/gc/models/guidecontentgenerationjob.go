package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidecontentgenerationjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidecontentgenerationjobDud struct { 
    Id string `json:"id"`


    Status string `json:"status"`


    Errors []Errorbody `json:"errors"`


    


    


    SelfUri string `json:"selfUri"`

}

// Guidecontentgenerationjob
type Guidecontentgenerationjob struct { 
    


    


    


    // Guide
    Guide Addressableentityref `json:"guide"`


    // GuideContent
    GuideContent Generatedguidecontent `json:"guideContent"`


    

}

// String returns a JSON representation of the model
func (o *Guidecontentgenerationjob) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidecontentgenerationjob) MarshalJSON() ([]byte, error) {
    type Alias Guidecontentgenerationjob

    if GuidecontentgenerationjobMarshalled {
        return []byte("{}"), nil
    }
    GuidecontentgenerationjobMarshalled = true

    return json.Marshal(&struct {
        
        Guide Addressableentityref `json:"guide"`
        
        GuideContent Generatedguidecontent `json:"guideContent"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

