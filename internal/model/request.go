package model

type RequestBody struct {
	Contents         []Content        `json:"contents"`
	GenerationConfig GenerationConfig `json:"generationConfig"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text       string      `json:"text,omitempty"`
	InlineData *InlineData `json:"inline_data,omitempty"`
}

type InlineData struct {
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}

type GenerationConfig struct {
	ResponseMimeType string         `json:"response_mime_type"`
	ResponseSchema   map[string]any `json:"response_schema"`
}

var responseSchema = map[string]any{
	"type": "OBJECT",
	"properties": map[string]any{
		"course_name": map[string]any{"type": "STRING"},
		"exam_year":   map[string]any{"type": "STRING"},
		"questions": map[string]any{
			"type": "ARRAY",
			"items": map[string]any{
				"type": "OBJECT",
				"properties": map[string]any{
					"question_number": map[string]any{"type": "STRING"},
					"text": map[string]any{
						"type":        "STRING",
						"description": "Question text with LaTeX formatting",
					},
					"marks": map[string]any{"type": "NUMBER"},
					"coordinates": map[string]any{
						"type":        "ARRAY",
						"description": "[ymin, xmin, ymax, xmax]",
						"items": map[string]any{
							"type": "NUMBER",
						},
					},
					"visual_elements": map[string]any{
						"type": "ARRAY",
						"items": map[string]any{
							"type": "OBJECT",
							"properties": map[string]any{
								"element_type": map[string]any{"type": "STRING"},
								"description":  map[string]any{"type": "STRING"},
								"box_2d": map[string]any{
									"type": "ARRAY",
									"items": map[string]any{
										"type": "NUMBER",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func CreateRequestBody(encodedImage []string) *RequestBody {
	parts := []Part{
		{
			Text: `Act as a specialized OCR and Document Parser. Analyze the provided exam paper and extract all data into the specified JSON format.
Formatting Rules:
    LaTeX: All mathematical equations, chemical formulas, and scientific symbols MUST be wrapped in LaTeX (e.g., E=mc2).
    Coordinates: For every question and every associated image/diagram, provide the bounding box as [ymin, xmin, ymax, xmax] using normalized coordinates (0-1000).
    Structure: Extract the 'course_name', 'exam_year', and an array of 'questions'.
    Questions: Each question must include 'text' (in LaTeX), 'marks', and 'visual_elements'. If a diagram exists, provide its 'type' and 'coordinates'.`,
		},
	}

	for _, enImg := range encodedImage {
		parts = append(parts, Part{
			InlineData: &InlineData{
				MimeType: "image/jpeg",
				Data:     enImg,
			},
		})
	}

	return &RequestBody{
		Contents: []Content{
			{
				Parts: parts,
			},
		},
		GenerationConfig: GenerationConfig{
			ResponseMimeType: "application/json",
			ResponseSchema:   responseSchema,
		},
	}
}
