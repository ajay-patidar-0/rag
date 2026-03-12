package model

type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
	Usage      Usage       `json:"usageMetadata"`
}

type Candidate struct {
	Content struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"content"`
}

type Usage struct {
	PromptTokenCount     int `json:"promptTokenCount"`
	CandidatesTokenCount int `json:"candidatesTokenCount"`
	TotalTokenCount      int `json:"totalTokenCount"`
}

// ExamPaper is your custom data structure
type ExamPaper struct {
	CourseName string     `json:"course_name"`
	ExamYear   string     `json:"exam_year"`
	Questions  []Question `json:"questions"`
}

type Question struct {
	QuestionNumber string          `json:"question_number"`
	Text           string          `json:"text"` // Contains LaTeX
	Marks          float64         `json:"marks"`
	Coordinates    []float64       `json:"coordinates"` // [ymin, xmin, ymax, xmax]
	VisualElements []VisualElement `json:"visual_elements"`
}

type VisualElement struct {
	ElementType string    `json:"element_type"`
	Description string    `json:"description"`
	Box2D       []float64 `json:"box_2d"`
}
