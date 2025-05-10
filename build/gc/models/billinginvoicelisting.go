package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillinginvoicelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillinginvoicelistingDud struct { 
    


    


    


    

}

// Billinginvoicelisting
type Billinginvoicelisting struct { 
    // Entities
    Entities []Billinginvoice `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Billinginvoicelisting) String() string {
     o.Entities = []Billinginvoice{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billinginvoicelisting) MarshalJSON() ([]byte, error) {
    type Alias Billinginvoicelisting

    if BillinginvoicelistingMarshalled {
        return []byte("{}"), nil
    }
    BillinginvoicelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Billinginvoice `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Billinginvoice{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

