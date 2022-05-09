package grades

import (
	"fmt"
	"sync"
)

//学生信息
type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

//分数
type Grade struct {
	Title string
	Type  GradeType
	Score float32
}
type GradeType string

const ( //考试类型
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

//学生的平均成绩
func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}

	return result / float32(len(s.Grades))
}

//寻找学生 by ID
type Students []Student

func (ss Students) GetByID(id int) (*Student, error) {
	for i := range ss {
		if ss[i].ID == id {
			return &ss[i], nil
		}
	}

	return nil, fmt.Errorf("学生的ID：  %d  未找到", id)
}

//用于外部的访问
var (
	students      Students
	studentsMutex sync.Mutex
)
