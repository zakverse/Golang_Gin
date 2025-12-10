package models

type Type struct {
    ID          uint      `json:"id"`
    CodeType    int       `json:"code_type"`
    Description string    `json:"description"`
}
/*INSERT INTO `type` (code_type, description)
VALUES
(1, 'Televisi'),
(2, 'Komputer'),
(3, 'Peralatan Game'),
(4, 'Camera'),
(5, 'Kipas Angin'),
(6, 'Setrika'),
(7, 'Rice Cooker');*/

var DataType = []Type{
	{ID: 1, CodeType: 1, Description: "Televisi"},
	{ID: 2, CodeType: 2, Description: "Komputer"},
	{ID: 3, CodeType: 3, Description: "Peralatan Game"},
	{ID: 4, CodeType: 4, Description: "Camera"},
	{ID: 5, CodeType: 5, Description: "Kipas Angin"},
	{ID: 6, CodeType: 6, Description: "Setrika"},
	{ID: 7, CodeType: 7, Description: "Rice Cooker"},
}