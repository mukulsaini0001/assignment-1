package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//type Address[]struct {
//street string
//landmark string
//city strng
//Pincode int
// State string
//}
var Address map[string]string

type Company struct {
	Name           string
	Age_in_years   float32
	Origin         string
	Head_officee   string
	address        []map[string]string
	Sponsors       []map[string]string
	Revenue        string
	No_Of_employee int
	Str_text       []string
	Int_text       []int
}

var Sponsors []map[string]string

//"address":[
//{
//"street" : "91 Springboard"
//"landmark" : "Axis Bank"
//"city" : "Noida"
//"pincode" : 201301,
//"state" ; "Uttar Pradesh"
//},
//{
// "street":"91 Springboard",
//"landmark":"Axis Bank",
//"city":"Noida",
//"pincode":201301,
//"state":"Uttar Pradesh"
//}
//],

func main() {
	s := `{
		"name" : "Tolexo Online Pvt. Ltd",
		"age_in_years" : 8.5,
		"origin" : "Noida",
		"head_office" : "Noida, Uttar Pradesh",
		"address" : [
		{
		"Street" : "91 Springboard",
		"Landmark" : "Axis Bank",
		"City" : "Noida",
		"Pincode" : 201301,
		"State" : "Uttar Pradesh"
		}
		],
		"sponsers" : {
		"name" : "One"
		},
		"revenue" : "19.8 million$",
		"no_of_employee" : 630,
		"str_text" : ["one","two"],
		"int_text" : [1,3,4]
		}`
	var comp Company
	json.Unmarshal([]byte(s), &comp)
	v := reflect.ValueOf(comp)
	fmt.Printf("value %#v", v)
	valuescomp := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		valuescomp[i] = v.Field(i)
	}
	for i := range valuescomp {
		fmt.Print(valuescomp[i], "\t")
		fmt.Println(reflect.TypeOf(valuescomp[i]))
		switch valuescomp[i].(type) {
		case string:
		case float32:
		case int:
		default:
			a := reflect.ValueOf((valuescomp[i]))
			typ := reflect.TypeOf(a)
			fmt.Println(valuescomp[i], "Hello")
			for j := 0; j < a.NumField(); i++ {
				f := a.Field(j)
				fmt.Printf("fff %v", typ.Field(j))
				fmt.Printf("%d:%s %s = %v\n", j, typ.Field(j).Name, f.Type(), f.Interface())
				fmt.Println("")
			}
		}

	}

}
