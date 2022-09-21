package model

import (
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	SentenceId int
	Pos        string
	Word       string
}

type Sentence struct {
	gorm.Model
	FileID       int
	Sentence     string
	Pos_sentence string
	Pos          string
	Word []Word
}

type File struct {
	gorm.Model
	CategoryID int
	FileName   string
	Content    string
}


type Category struct {
	gorm.Model
	CategoryName string
}