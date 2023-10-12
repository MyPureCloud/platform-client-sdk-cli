package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetflowexecutiondatajobresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetflowexecutiondatajobresultDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Getflowexecutiondatajobresult - This is a list of executionData links that can be used to download the complete executionData
type Getflowexecutiondatajobresult struct { 
    


    // Name
    Name string `json:"name"`


    // Entities - On jobState = Success this field will be populated with the list of results of files for download.
    Entities []Executiondataentity `json:"entities"`


    // JobState - The state of the backend process to prep the files for download.
    JobState string `json:"jobState"`


    

}

// String returns a JSON representation of the model
func (o *Getflowexecutiondatajobresult) String() string {
    
     o.Entities = []Executiondataentity{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getflowexecutiondatajobresult) MarshalJSON() ([]byte, error) {
    type Alias Getflowexecutiondatajobresult

    if GetflowexecutiondatajobresultMarshalled {
        return []byte("{}"), nil
    }
    GetflowexecutiondatajobresultMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Entities []Executiondataentity `json:"entities"`
        
        JobState string `json:"jobState"`
        *Alias
    }{

        


        


        
        Entities: []Executiondataentity{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

