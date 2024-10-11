package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemoncreaterulelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemoncreaterulelistingDud struct { 
    


    


    


    


    

}

// Workitemoncreaterulelisting
type Workitemoncreaterulelisting struct { 
    // Entities
    Entities []Workitemoncreaterule `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // After
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Workitemoncreaterulelisting) String() string {
     o.Entities = []Workitemoncreaterule{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemoncreaterulelisting) MarshalJSON() ([]byte, error) {
    type Alias Workitemoncreaterulelisting

    if WorkitemoncreaterulelistingMarshalled {
        return []byte("{}"), nil
    }
    WorkitemoncreaterulelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitemoncreaterule `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Workitemoncreaterule{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

