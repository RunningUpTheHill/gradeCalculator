# ESEP Grade Calculator

This library is a simple utility to calculate your grades from three categories:
- assignments
- exams
- essays

Letter Scale:
- A: >=90
- B: >=80
- C: >=70
- D: >=60
- F: < 60

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade-calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade-calculator/
go test -coverprofile=coverage.txt ./...
```