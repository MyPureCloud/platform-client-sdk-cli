package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LabelresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LabelresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Labelresponse
type Labelresponse struct { 
    


    // Name - The name of the label.
    Name string `json:"name"`


    // Color - The color of the label.
    Color string `json:"color"`


    // DateCreated - The creation date and time of the label. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The last modification date and time of the label. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DocumentCount - Number of documents assigned to this label.
    DocumentCount int `json:"documentCount"`


    // ExternalId - The external id associated with the label.
    ExternalId string `json:"externalId"`


    

}

// String returns a JSON representation of the model
func (o *Labelresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Labelresponse) MarshalJSON() ([]byte, error) {
    type Alias Labelresponse

    if LabelresponseMarshalled {
        return []byte("{}"), nil
    }
    LabelresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Color string `json:"color"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        DocumentCount int `json:"documentCount"`
        
        ExternalId string `json:"externalId"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

