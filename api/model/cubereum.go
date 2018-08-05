package model

type FilterDet struct {
	Alias    string
	Operator string
	Value    string
	AliasSlc []string
	Master   int
}
type SQLFilter struct {
	Filter map[string]interface{}
	Offset int
	Limit  int
	SortBy string
}

var ListGroupFields map[string]FilterDet
var ListGroupFieldsSort map[string]FilterDet

var ListGroupFieldsGroupBy map[string]FilterDet

//ListGroupFieldsOption : Seller/Seller User filter options allowed value set
var ListGroupFieldsOption map[string]map[string]struct{}

func init() {

	//
	//Seller/Seller User filter options
	ListGroupFields = make(map[string]FilterDet)
	ListGroupFields = map[string]FilterDet{"schoolname": {Operator: "="},
		"address":  {Operator: "="},
		"pincode":  {Operator: "="},
		"area":     {Operator: "="},
		"landmark": {Operator: "="},
	}
	ListGroupFieldsSort = make(map[string]FilterDet)
	ListGroupFieldsSort = map[string]FilterDet{"name": {Operator: "="},
		"medium_of_inst": {Operator: "="},
		"pincode":        {Operator: "="},
	}
	ListGroupFieldsGroupBy = make(map[string]FilterDet)
	ListGroupFieldsGroupBy = map[string]FilterDet{"category": {Operator: "="},
		"medium_of_inst": {Operator: "="},
		"gender":         {Operator: "="},
	}

}

//response struct for groupby
type Name struct {
	Schoolname string `json:"schoolname" gorm:"column:schoolname"`
}

//response struct
type BlrSchool struct {
	District        string `json:"-" gorm:"column:district"`
	Block           string `json: "block" gorm:"column:block"`
	Cluster         string `json:"cluster" gorm:"column:cluster"`
	Schoolid        int    `json:"schoolid" gorm:"column:schoolid"`
	Schoolname      string `json:"schoolname" gorm:"column:schoolname"`
	Category        string `json:"category" gorm:"column:category"`
	Gender          string `json:"gender" gorm:"column:gender"`
	Medium_of_inst  string `json:"medium_of_inst" gorm:"column:medium_of_inst"`
	Address         string `json:"address" gorm:"column:address"`
	Area            string `json:"area" gorm:"column:area"`
	Pincode         int    `json:"pincode" gorm:"column:pincode"`
	Landmark        string `json:"landmark" gorm:"column:landmark"`
	Identification1 string `json:"-" gorm:"column:identification1"`
	Busroutes       string `json:"busroutes" gorm:"column:busroutes"`
	Identification2 string `json:"-" gorm:"column:identification2"`
	Latlong         string `json:"-" gorm:"column:latlong"`
}

//Prepared statement contains too many placeholders
