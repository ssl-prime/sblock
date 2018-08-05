required information for run these api:

1. install golang  version 1.10.3
you have to clone this repo where by default path is set.
or you can path as per your convenice 
if required package or import location  not  found  then
run these command  go get or go get <import location>
i have table in mysal whose name is school_info

CREATE TABLE `school_info` (
  `school_info_id` int(11) NOT NULL AUTO_INCREMENT,
  `district` varchar(255) DEFAULT NULL,
  `block` varchar(255) DEFAULT NULL,
  `cluster` varchar(255) DEFAULT NULL,
  `schoolid` int(11) NOT NULL,
  `schoolname` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `medium_of_inst` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `area` varchar(255) NOT NULL,
  `pincode` int(6) NOT NULL,
  `landmark` varchar(255) DEFAULT NULL,
  `identification1` varchar(255) DEFAULT NULL,
  `busroutes` varchar(255) DEFAULT NULL,
  `identification2` varchar(255) DEFAULT NULL,
  `latlong` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`school_info_id`),
  UNIQUE KEY `schoolid` (`schoolid`)
) ;


for making connection you need to change 
user_name ,password and database
in sblock->api->util-> dbConn(){
db, err = gorm.Open("mysql", "root:spatico@/grabbd")
}
change file location as i have attached csv with this repo.
now go to lauch folder and use this command :   go run main.go
this will run on 8090 by default
api :

for uploading the csv file into database
(if u need the csv to be converted into json then pls let me know i will give you in separate file)

localhost:8090/saurabh/sblock/v1/convert/csv/tojson/

search info based on certain parameter

localhost:8090/saurabh/sblock/v1/search/school/?pincode=560062&schoolname= "GUHPS ALAHALLI"

sort the data on click on certain column,
localhost:8090/saurabh/sblock/v1/sort/by/?pincode

separate into particular group based on certain column
localhost:8090/saurabh/sblock/v1/group/by/?category
i have only group the name of school based on certain passed parameter.
if something is missing let me know
email : skcse03@gmail.com
mob : 8825385286

