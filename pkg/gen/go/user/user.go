package user

// Code generated by avro/gen. DO NOT EDIT.

// User is a generated struct.
type User struct {
	ID         string `avro:"id" json:"id" yaml:"Id"`
	Name       string `avro:"name" json:"name" yaml:"Name"`
	Email      string `avro:"email" json:"email" yaml:"Email"`
	Age        *int   `avro:"age" json:"age" yaml:"Age"`
	Registered bool   `avro:"registered" json:"registered" yaml:"Registered"`
}
