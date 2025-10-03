package esepunittests

type GradeCalculator struct {
	list []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		list: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.list = append(gc.list, Grade{Name: name, Grade: grade, Type: gradeType})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	var a, e, s []Grade
	for _, g := range gc.list {
		switch g.Type {
			case Assignment:
				a = append(a, g)
			case Exam:
				e = append(e, g)
			case Essay:
				s = append(s, g)
		}}
	assignmentAverage := computeAverage(a)
	examAverage := computeAverage(e)
	essayAverage := computeAverage(s)

	weighted_grade := float64(assignmentAverage)*.5 + float64(examAverage)*.35 + float64(essayAverage)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	sum := 0

	if len(grades) == 0 {
		return 0
	}
	
	for _, grade := range grades {
		sum += grade.Grade
	}

	return sum / len(grades)
}
