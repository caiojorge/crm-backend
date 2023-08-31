package model

/*
The Active Record pattern is an architectural pattern found in software engineering
that stores in-memory object data in relational databases. In this pattern, an object
instance is tied to a single row in the database table. The object's member variables
correspond to the columns of the table and the behavior of the object maps to CRUD operations
in the database.
*/

// ActiveRecord is an interface to implement the Active Record pattern
type ActiveRecord interface {
	Create() bool
	Update() bool
	Delete() bool
}
