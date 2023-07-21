package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetalertqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetalertqueryDud struct { 
    


    


    


    


    


    


    


    

}

// Getalertquery
type Getalertquery struct { 
    // RuleType - The rule type of the alerts the query will return
    RuleType string `json:"ruleType"`


    // QueryType - The type of query being performed.
    QueryType string `json:"queryType"`


    // ActiveStatus - The status of the alerts the query will return.
    ActiveStatus string `json:"activeStatus"`


    // ViewedStatus - The view status of the alerts the query will return.
    ViewedStatus string `json:"viewedStatus"`


    // PageNumber - The page number of the queried response
    PageNumber int `json:"pageNumber"`


    // PageSize - The number of entities to return of the queried response.  The max is 25
    PageSize int `json:"pageSize"`


    // SortBy - The field to sort responses by.  The accepted choices are Name and DateStart
    SortBy string `json:"sortBy"`


    // SortOrder - The order in which response will be sorted.  The accepted choices are Asc and Desc
    SortOrder string `json:"sortOrder"`

}

// String returns a JSON representation of the model
func (o *Getalertquery) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getalertquery) MarshalJSON() ([]byte, error) {
    type Alias Getalertquery

    if GetalertqueryMarshalled {
        return []byte("{}"), nil
    }
    GetalertqueryMarshalled = true

    return json.Marshal(&struct {
        
        RuleType string `json:"ruleType"`
        
        QueryType string `json:"queryType"`
        
        ActiveStatus string `json:"activeStatus"`
        
        ViewedStatus string `json:"viewedStatus"`
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        
        SortBy string `json:"sortBy"`
        
        SortOrder string `json:"sortOrder"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

