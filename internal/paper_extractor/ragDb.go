package paperextractor

import (
	"fmt"

	"github.com/ajay-patidar-0/rag/internal/model"
	"github.com/ajay-patidar-0/rag/internal/rag"
	"github.com/ajay-patidar-0/rag/internal/store"
	"github.com/google/uuid"
)

func StorePaper(paper *model.ExamPaper, v *store.VectorDB) error {
	courseId, err := v.FindCourseByName(paper.CourseName)
	if err != nil && courseId == uuid.Nil {
		courseId, err = v.AddCourse(paper.CourseName)
		if err != nil {
			return fmt.Errorf("error in store paper in course id sectin %v", err)
		}
	}

	examId, err := v.FindExamPaper(courseId, paper.ExamYear)
	if err != nil && examId == uuid.Nil {
		examId, err = v.AddExamPaper(courseId, paper.ExamYear, "imageUrl") // update image url
		if err != nil {
			return fmt.Errorf("error in store paper in exam id sectin %v", err)
		}
	}

	for _, question := range paper.Questions {
		text := question.Text
		if len(question.VisualElements) > 0 {
			for _, v := range question.VisualElements {
				text += " " + v.Description
			}
		}

		emb, err := rag.GetEmbedding(text)
		if err != nil {
			return fmt.Errorf("error occur in store paper in get embedding %v ", err)
		}
		vector := store.ToPGVector(emb)
		err = v.AddQuestion(courseId, examId, question.Text, int(question.Marks), vector)
		if err != nil {
			return fmt.Errorf("error in adding questions %v", err)
		}
	}
	return nil
}
