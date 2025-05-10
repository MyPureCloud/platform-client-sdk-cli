package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableversionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    DatePublished time.Time `json:"datePublished"`


    


    


    SelfUri string `json:"selfUri"`

}

// Decisiontableversion
type Decisiontableversion struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Version - The decision table version.
    Version int `json:"version"`


    // Status - Current status of this decision table version
    Status string `json:"status"`


    // Description - The decision table description.
    Description string `json:"description"`


    // RowCount - The number of rows in this decision table version.
    RowCount int `json:"rowCount"`


    // RowsUri - The rows URI for this decision table version.
    RowsUri string `json:"rowsUri"`


    


    


    


    // Columns - The column definitions of this decision table version.
    Columns Decisiontablecolumns `json:"columns"`


    // Contract - The contract information for this decision table version.
    Contract Decisiontablecontract `json:"contract"`


    

}

// String returns a JSON representation of the model
func (o *Decisiontableversion) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableversion) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableversion

    if DecisiontableversionMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableversionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Version int `json:"version"`
        
        Status string `json:"status"`
        
        Description string `json:"description"`
        
        RowCount int `json:"rowCount"`
        
        RowsUri string `json:"rowsUri"`
        
        Columns Decisiontablecolumns `json:"columns"`
        
        Contract Decisiontablecontract `json:"contract"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

