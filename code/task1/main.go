package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Student is a struct to record name ,id
type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//StudentSet is a map [name]Student
type StudentSet struct {
	M map[string]*Student `json:"data"`
}

func main() {
	studentSet := NewStudentSets()
	studentSet.Add(1, "bingbing")
	studentSet.Add(2, "wuming")
	fmt.Println(studentSet)
	Dump("test.txt", studentSet)
}

// 初始化学生Map
func NewStudentSets() *StudentSet {
	return &StudentSet{M: make(map[string]*Student, 0)}
}
func (s *StudentSet) Add(id int, name string) (err error) {
	_, ok := s.M[name]
	if ok {
		err = fmt.Errorf("student %s already exists.", name)
		return
	}
	s.M[name] = &Student{id, name}
	return
}

func (s *StudentSet) list() string {
	var str string = "Id\t\tName\n"
	for k, v := range s.M {
		str += fmt.Sprintf("%d\t\t%s", v.Id, k)
	}
	return str
}

// Remove student that specified
func (s *StudentSet) remove(name string) error {
	_, ok := s.M[name]
	if ok {
		delete(s.M, name)
		return nil
	}
	return fmt.Errorf("not found")
}

// Clear all students info
func (s *StudentSet) clear() {
	s.M = make(map[string]*Student, 0)
}

// String implements String method
func (s *StudentSet) String() string {
	return s.list()
}

// Dump students info to the file that specified
func Dump(fileName string, stu *StudentSet) (err error) {
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fd.Close()

	bs, err := json.Marshal(stu)
	if err != nil {
		return
	}

	_, err = fd.Write(bs)
	if err != nil {
		return
	}

	return
}

// Load students info to the file that specified
func Load(fileName string) (stu *StudentSet, err error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = json.Unmarshal(bs, &stu)
	if err != nil {
		return
	}

	return
}
