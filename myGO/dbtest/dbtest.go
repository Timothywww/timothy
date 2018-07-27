package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Person struct {
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
}

func Register() Person {
	result2 := Person{}
	var name string
	var phone string
	fmt.Println("Please enter your name")
	fmt.Scanln(&name)
	fmt.Println("Please enter your Phone")
	fmt.Scanln(&phone)

	result2.Name = name
	result2.Phone = phone
	return result2
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	db := session.DB("test")
	c := db.C("people")

	for {
		fmt.Println("Welcome!")
		fmt.Println("Menu: 1.Register 2.Update Phone 3.Delete 4.List all 5.Quit")

		var input string
		fmt.Scanln(&input)
		//sign up
		if input == "1" {
			result := Register()

			err = c.Find(bson.M{"name": result.Name}).One(&result)
			if err == mgo.ErrNotFound {
				err = c.Insert(result)
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Println("User already existed")
			}
		} else if input == "2" {
			result2 := Person{}
			var name string
			var phone string
			fmt.Println("Please enter your name")
			fmt.Scanln(&name)
			fmt.Println("Please enter new Phone")
			fmt.Scanln(&phone)

			err = c.Find(bson.M{"name": name}).One(&result2)
			if err != nil {
				fmt.Println("User not found")
			}

			Querier := bson.M{"name": name}
			change := bson.M{"$set": bson.M{"phone": phone}}
			err = c.Update(Querier, change)
			if err != nil {
				panic(err)
			}
		} else if input == "3" {
			result3 := Person{}
			var name string
			fmt.Println("Please enter your name")
			fmt.Scanln(&name)

			err = c.Find(bson.M{"name": name}).One(&result3)
			if err != nil {
				fmt.Println("User not found")
			}
			err = c.Remove(bson.M{"name": name})
			if err != nil {
				panic(err)
			}
		} else if input == "4" {
			var result []interface{}
			iter := c.Find(nil).Iter()
			err := iter.All(&result)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%v", result)
		} else if input == "5" {
			break
		} else {
			continue
		}
	}
}

// mongodb  知识点
// 数据类型和Go内置类型的关联对比
// bson 数据结构
// index 索引以及相关属性
