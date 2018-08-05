package service

import (
	//"fmt"
	"github.com/rightjoin/aqua"
	"sblock/api/util"
)

type Sblock struct {
	aqua.RestService `prefix:"saurabh/sblock" root:"/" version:"1"`
	uploadCSV        aqua.POST `url:"/convert/csv/tojson/"`
	searchSchool     aqua.GET  `url:"/search/school/"`
	sortBy           aqua.GET  `url:"/sort/by/"`
	groupBy          aqua.GET  `url:"/group/by/"`
}

//Sort uploadCSV
func (mvz *Sblock) UploadCSV(j aqua.Aide) (
	response interface{}, err error) {
	response, err = util.UploadCSV(j)
	return
}

//SortBy
func (mvz *Sblock) SearchSchool(j aqua.Aide) (
	response interface{}, err error) {

	response, err = util.SearchSchool(j)
	return

}

//sort BY
func (mvz *Sblock) SortBy(j aqua.Aide) (
	response interface{}, err error) {

	response, err = util.SortBy(j)
	return

}

//group by
func (mvz *Sblock) GroupBy(j aqua.Aide) (
	response interface{}, err error) {

	response, err = util.GroupBy(j)
	return

}
