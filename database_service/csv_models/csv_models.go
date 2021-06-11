package csv_models

import (
	"os"

	"github.com/gocarina/gocsv"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

type schoolEntry struct {
	SchoolName string `csv:"schoolname"`
}

type facultyEntry struct {
	FacultyName string `csv:"facultyname"`
	SchoolID    uint   `csv:"schoolid"`
}

type courseEntry struct {
	CourseCode string `csv:"coursecode"`
	CourseName string `csv:"coursename"`
	FacultyID  uint   `csv:"facultyid"`
}

const (
	schoolpath  string = "../database_service/csv/school.csv"
	facultypath string = "../database_service/csv/faculty.csv"
	coursepath  string = "../database_service/csv/course.csv"
)

func ImportSchool() (schools []databaseModel.School, err error) {
	file, err := os.Open(schoolpath)
	if err != nil {
		return schools, err
	}
	defer file.Close()
	var schoolentries []schoolEntry
	err = gocsv.Unmarshal(file, &schoolentries)
	if err != nil {
		return schools, err
	}
	schools = make([]databaseModel.School, len(schoolentries))
	for i, v := range schoolentries {
		schools[i] = convertschentrytoschool(v)
	}
	return
}

func ImportFaculty() (faculties []databaseModel.Faculty, err error) {
	file, err := os.Open(facultypath)
	if err != nil {
		return faculties, err
	}
	defer file.Close()
	var facentries []facultyEntry
	err = gocsv.Unmarshal(file, &facentries)
	if err != nil {
		return faculties, err
	}
	faculties = make([]databaseModel.Faculty, len(facentries))
	for i, v := range facentries {
		faculties[i] = convertfacultyentrytofaculty(v)
	}
	return
}

func ImportCourse() (courses []databaseModel.Course, err error) {
	file, err := os.Open(coursepath)
	if err != nil {
		return courses, err
	}
	defer file.Close()
	var courseentries []courseEntry
	err = gocsv.Unmarshal(file, &courseentries)
	if err != nil {
		return courses, err
	}
	courses = make([]databaseModel.Course, len(courseentries))
	for i, v := range courseentries {
		courses[i] = convertcourseentrytocourse(v)
	}
	return
}

func convertschentrytoschool(entry schoolEntry) (school databaseModel.School) {
	school.SchoolName = entry.SchoolName
	return
}

func convertfacultyentrytofaculty(entry facultyEntry) (faculty databaseModel.Faculty) {
	faculty.FacultyName = entry.FacultyName
	faculty.SchoolID = entry.SchoolID
	return
}

func convertcourseentrytocourse(entry courseEntry) (course databaseModel.Course) {
	course.CourseCode = entry.CourseCode
	course.CourseName = entry.CourseName
	course.FacultyID = entry.FacultyID
	return
}
