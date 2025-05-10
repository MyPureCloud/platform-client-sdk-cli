package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableDud struct { 
    Id string `json:"id"`


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    DatePublished time.Time `json:"datePublished"`


    Published Decisiontableversionentity `json:"published"`


    Latest Decisiontableversionentity `json:"latest"`


    


    


    SelfUri string `json:"selfUri"`

}

// Decisiontable
type Decisiontable struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The decision table description.
    Description string `json:"description"`


    


    


    


    


    


    // Columns - The column definitions of this decision table.
    Columns Decisiontablecolumns `json:"columns"`


    // PublishedContract - The published contract information for this decision table.
    PublishedContract Decisiontablecontract `json:"publishedContract"`


    

}

// String returns a JSON representation of the model
func (o *Decisiontable) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontable) MarshalJSON() ([]byte, error) {
    type Alias Decisiontable

    if DecisiontableMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Columns Decisiontablecolumns `json:"columns"`
        
        PublishedContract Decisiontablecontract `json:"publishedContract"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

