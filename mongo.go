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

var c *mgo.Collection
var dbo *mgo.Session

func MongoDial(d,t string) *mgo.Session {

	if dbo == nil {
		db,e := mgo.Dial("localhost")
		dbo = db
		err(e)
		//defer db.Close()
	}

	dbo.SetMode(mgo.Monotonic,true)

	c = dbo.DB(d).C(t)
	i := mgo.Index{
		Key: []string{"_id"},
		Unique: true,
		DropDups: true,
		Background: true,
	}

	e := c.EnsureIndex(i)
	err(e)
	return dbo
}

func Insert(in bson.M){
	delete(in, "ID");
	e := c.Insert(in)
	err(e)
}

func FindOne(find bson.M,sel bson.M) bson.M {
	var r bson.M
	e := c.Find(find).Select(sel).One(&r)
	err(e)
	return r
}

func FindLast(find bson.M, sel bson.M, sort string) bson.M {
	var r bson.M
	e := c.Find(find).Sort(sort).Select(sel).One(&r)
	err(e)
	return r
}

func FindAll(find bson.M,sort string) []bson.M {
	var r []bson.M
	e := c.Find(find).Sort(sort).All(&r)
	err(e)
	return r
}

func Update(find bson.M,change bson.M,multi bool)  {
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
