package store

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type VectorDB struct {
	DB *sql.DB
}

func NewVectorDB(db *sql.DB) *VectorDB {
	return &VectorDB{
		DB: db,
	}
}

func (v *VectorDB) AddCourse(courseName string) (uuid.UUID, error) {
	query := `INSERT INTO courses (course_name) VALUES($1) RETURNING id`

	var courseId uuid.UUID
	if err := v.DB.QueryRow(query, courseName).Scan(&courseId); err != nil {
		return uuid.Nil, err
	}
	return courseId, nil
}

func (v *VectorDB) FindCourseByName(courseName string) (uuid.UUID, error) {
	query := `SELECT id FROM courses WHERE course_name=$1`

	var courseId uuid.UUID
	if err := v.DB.QueryRow(query, courseName).Scan(&courseId); err != nil {
		return uuid.Nil, err
	}
	return courseId, nil
}

func (v *VectorDB) AddExamPaper(courseId uuid.UUID, examYear int, imageUrl string) (uuid.UUID, error) {
	query := `INSERT INTO exam_papers (course_id, exam_year, image_url) VALUES ($1, $2, $3) RETURNING id`

	var examId uuid.UUID
	if err := v.DB.QueryRow(query, courseId, examYear, imageUrl).Scan(&examId); err != nil {
		return uuid.Nil, err
	}
	return examId, nil
}

func (v *VectorDB) FindExamPaper(courseId uuid.UUID, examYear string) (uuid.UUID, error) {
	query := `SELECT id FROM exam_papers WHERE course_id=$1 AND exam_year=$2`

	var examId uuid.UUID
	if err := v.DB.QueryRow(query, courseId, examYear).Scan(&examId); err != nil {
		return uuid.Nil, err
	}

	return examId, nil
}

func (v *VectorDB) AddQuestion(courseId uuid.UUID, examId uuid.UUID, question string, marks int, vector string) error {
	query := `INSERT INTO questions (course_id, exam_paper_id, question_text, marks, embedding)
	VALUES($1, $2, $3, $4, $5)`

	_, err := v.DB.Query(query, courseId, examId, question, marks, vector)
	if err != nil {
		return fmt.Errorf("Enable to insert embedding %w", err)
	}

	return nil
}

func ToPGVector(vec []float32) string {
	var sb strings.Builder
	sb.WriteString("[")

	for i, v := range vec {
		sb.WriteString(fmt.Sprintf("%f", v))
		if i != len(vec)-1 {
			sb.WriteString(",")
		}
	}

	sb.WriteString("]")
	return sb.String()
}
