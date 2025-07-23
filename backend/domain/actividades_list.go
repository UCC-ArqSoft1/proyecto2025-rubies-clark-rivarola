package domain

// ActivityListItem es lo que devuelve GET /activities
type ActivityListItem struct {
	ID             uint    `json:"id"`
	Title          string  `json:"title"`
	CategoryName   string  `json:"category_name"`
	InstructorName string  `json:"instructor_name"`
	Day            string  `json:"day"`
	StartTime      string  `json:"start_time"`
	Duration       uint16  `json:"duration"`
	Capacity       uint16  `json:"capacity"`
	ImageURL       *string `json:"image_url,omitempty"`
}
