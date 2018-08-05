package util

import (
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rightjoin/aqua"
	"os"
	"sblock/api/model"
	"strconv"
)

func UploadCSV(j aqua.Aide) (interface{}, error) {

	//prods := []model.BlrSchool{}
	var (
		db         *gorm.DB
		err        error
		insertData []interface{}
	)
	cnt := 0

	f, _ := os.Open(`/home/spatico/skcse03/go/src/sblock/school.csv`)
	reader := csv.NewReader(f)
	reader.Comma = '|'
	schoolInfo, err := reader.ReadAll()
	for i, prod := range schoolInfo {
		//fmt.Println(prod)
		if i < 1000 {
			if i == 0 {
				// skip header line
				continue
			}
			if schoolID, err := strconv.Atoi(prod[3]); err == nil {
				if pinCode, err := strconv.Atoi(prod[10]); err == nil {
					cnt++
					insertData = append(insertData, prod[0], prod[1], prod[2], schoolID, prod[4], prod[5],
						prod[6], prod[7], prod[8], prod[9], pinCode, prod[11], prod[12],
						prod[13], prod[14], prod[15])
				}
			} else {
				fmt.Println(err)
			}
		}
	}
	kt := len(insertData)
	if kt > 1000 {
		kt = 1000
	}
	if db, err = dbConn(); err == nil {
		var qry string
		if cnt > 0 {
			qry = "values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
			for i := 1; i < cnt; i++ {
				qry += ", (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
			}

			insrt := `INSERT INTO school_info (district, block, cluster, schoolid,
		 		 schoolname, category, gender, medium_of_inst, address, area,
		 		 pincode, landmark, identification1, busroutes, identification2,
		 		 latlong ) ` + qry + `ON DUPLICATE KEY UPDATE schoolid= values(schoolid)`

			if err := db.Exec(insrt, insertData...).Error; err != nil {
				fmt.Println("cubereum", err)
			} else {
				fmt.Println("Saurabh")
			}
		}
	}
	return "srt", err

}

func SearchSchool(qParams aqua.Aide) (interface{}, error) {
	qParams.LoadVars()
	//filterOpt := make(map[string]model.FilterDet)
	qstring := ` where `
	f := 0
	fmt.Println(qParams.QueryVar)
	for key, _ := range model.ListGroupFields {
		if value, ok := qParams.QueryVar[key]; ok == true {
			if f == 1 {
				qstring += " and "
				qstring += " " + key + " = " + value
			} else {
				f = 1
				qstring += key + " = " + value
			}

		}
	}
	if f == 0 {
		qstring = ``
	}

	var (
		db       *gorm.DB
		response []model.BlrSchool
		err      error
	)
	if db, err = dbConn(); err == nil {
		schoolQry := `Select * from school_info` + qstring
		fmt.Println(schoolQry)
		err = db.Raw(schoolQry).Find(&response).Error

	}

	return response, err
}

//
func SortBy(qParams aqua.Aide) (interface{}, error) {
	qParams.LoadVars()
	//filterOpt := make(map[string]model.FilterDet)
	qstring := ` order by  `
	f := 0
	fmt.Println(qParams.QueryVar)
	for key, _ := range model.ListGroupFieldsSort {
		if _, ok := qParams.QueryVar[key]; ok == true {
			if f == 1 {
				qstring += " , "
				qstring += " " + key
			} else {
				f = 1
				qstring += key
			}

		}
	}
	if f == 0 {
		qstring = ` order by schoolname `
	}

	var (
		db       *gorm.DB
		response []model.BlrSchool
		err      error
	)
	if db, err = dbConn(); err == nil {
		schoolQry := `Select * from school_info` + qstring
		fmt.Println(schoolQry)
		err = db.Raw(schoolQry).Find(&response).Error

	}

	return response, err
}

func GroupBy(qParams aqua.Aide) (interface{}, error) {
	qParams.LoadVars()
	//filterOpt := make(map[string]model.FilterDet)
	qstring := ` group by  `
	f := 0
	fmt.Println(qParams.QueryVar)
	for key, _ := range model.ListGroupFieldsGroupBy {
		if _, ok := qParams.QueryVar[key]; ok == true {
			if f == 1 {
				qstring += " , "
				qstring += " " + key
			} else {
				f = 1
				qstring += key
			}

		}
	}
	if f == 0 {
		qstring = ``
	}

	var (
		db       *gorm.DB
		response []model.Name
		err      error
	)

	if db, err = dbConn(); err == nil {
		schoolQry := `Select group_concat(schoolname) as schoolname 
		from school_info` + qstring
		fmt.Println(schoolQry)
		err = db.Raw(schoolQry).Find(&response).Error

	}

	return response, err
}
func dbConn() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:spatico@/grabbd")
	return
}
