/**
Mongo CRUD
reoxey
**/
package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

var db *mgo.Database
var dbo *mgo.Session

func MongoDial(d string) *mgo.Session {

	if dbo == nil {
		dbo,e := mgo.Dial("localhost")
		
		err(e)
		dbo.SetMode(mgo.Monotonic,true)
//		defer dbo.Close()
	}

	db = dbo.DB(d)
	return dbo
}

func CreateCollection(t string) *mgo.Collection {
	c := db.C(t)
	
	i := mgo.Index{
		Key: []string{"_id"},
		Unique: true,
		DropDups: true,
		Background: true,
	}
	
	e := c.EnsureIndex(i)
	err(e)
	
	return c
}

func Insert(c *mgo.Collection, in bson.M){
	fmt.Println(c)
	delete(in, "ID");
	e := c.Insert(in)
	err(e)
}

func FindOne(c *mgo.Collection, find bson.M,sel bson.M) bson.M {
	var r bson.M
	e := c.Find(find).Select(sel).One(&r)
	err(e)
	return r
}

func FindAll(c *mgo.Collection, find bson.M,sort string) []bson.M {
	var r []bson.M
	e := c.Find(find).Sort(sort).All(&r)
	err(e)
	return r
}

func Update(c *mgo.Collection, find bson.M,change bson.M,multi bool)  {
	var e error
	if multi {
		_,e = c.UpdateAll(find, change)
	} else {
		e = c.Update(find, change)
	}
	err(e)
}

func Drop(d string)  {
	e := dbo.DB(d).DropDatabase()
	err(e)
}

func err(e error)  {
	if e != nil{
		fmt.Print(e)
	}
}
