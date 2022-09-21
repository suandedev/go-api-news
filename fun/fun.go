package fun

import (
	"go-news/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	// connect sqlite db
	db, err := gorm.Open(sqlite.Open("words.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate the schema
	db.AutoMigrate(&model.Word{}, &model.Sentence{}, &model.File{}, &model.Category{})

	return db
}

func CreateWord(sentenceId int, pos string, word string) *model.Word {
	// connect db
	db := ConnectDb()

	// append word
	res := model.Word{SentenceId: sentenceId, Pos: pos, Word: word}

	// save word
	db.Create(&res)

	return &res
}

func GetWords() *[]model.Word {
	// connect db
	db := ConnectDb()

	// get all words
	var words []model.Word
	db.Find(&words)

	return &words
}

func GetWord(id int) *model.Word {
	// connect db
	db := ConnectDb()

	// get word
	var word model.Word
	db.First(&word, id)

	return &word
}

func GetWordBySentenceId(sentenceId int) *[]model.Word {
	// connect db
	db := ConnectDb()

	// get word by sentence id
	var words []model.Word
	db.Where("sentence_id = ?", sentenceId).Find(&words)

	return &words
}

func UpdateWord(id int, sentenceId int, pos string, word string) *model.Word {
	// connect db
	db := ConnectDb()

	// update word by id
	var res model.Word
	db.Model(&res).Where("id = ?", id).Updates(model.Word{SentenceId: sentenceId, Pos: pos, Word: word})

	return &res
}

func DeleteWord(id int) *string {
	// connect db
	db := ConnectDb()

	// delete word by id
	var word model.Word
	db.Delete(&word, id)

	return &word.Word
}

func CreateSentence(fileID int, sentence string, pos_sentence string, pos string) *model.Sentence {
	// connect db
	db := ConnectDb()

	// append word
	res := model.Sentence{FileID: fileID, Sentence: sentence, Pos_sentence: pos_sentence, Pos: pos}

	// save word
	db.Create(&res)

	return &res
}

func GetSentences() *gorm.DB {
	// connect db
	db := ConnectDb()

	// get all words
	var sentences []model.Sentence
	// join sentence and word where sentence.id = word.sentence_id
	db.Select("sentences.id, sentences.file_id, sentences.sentence, sentences.pos_sentence, sentences.pos, words.id, words.sentence_id, words.pos, words.word").Joins("left join words on sentences.id = words.sentence_id").Find(&sentences)

	return db
}

func GetSentence(id int) *model.Sentence {
	// connect db
	db := ConnectDb()

	// get word
	var sentence model.Sentence
	db.First(&sentence, id)

	return &sentence
}

func GetSentenceByFileId(fileId int) *[]model.Sentence {
	// connect db
	db := ConnectDb()

	// get word by sentence id
	var sentences []model.Sentence
	db.Where("file_id = ?", fileId).Find(&sentences)

	return &sentences
}

func UpdateSentence(id int, fileID int, sentence string, pos_sentence string, pos string) *model.Sentence {
	// connect db
	db := ConnectDb()

	// update word by id
	var res model.Sentence
	db.Model(&res).Where("id = ?", id).Updates(model.Sentence{FileID: fileID, Sentence: sentence, Pos_sentence: pos_sentence, Pos: pos})

	return &res
}

func DeleteSentence(id int) *string {
	// connect db
	db := ConnectDb()

	// delete word by id
	var sentence model.Sentence
	db.Delete(&sentence, id)

	return &sentence.Sentence
}

func CreateFile(categoriId int, fileName string, content string) *model.File {
	// connect db
	db := ConnectDb()

	// append word
	res := model.File{CategoryID: categoriId, FileName: fileName, Content: content}

	// save word
	db.Create(&res)

	return &res
}

func GetFiles() *[]model.File {
	// connect db
	db := ConnectDb()

	// get all words
	var files []model.File
	db.Find(&files)

	return &files
}

func GetFileById(id int) *model.File {
	// connect db
	db := ConnectDb()

	// get word
	var file model.File
	db.First(&file, id)

	return &file
}

func GetFileByCategoriId(categoryId int) *[]model.File {
	// connect db
	db := ConnectDb()

	// get word by sentence id
	var files []model.File
	db.Where("category_id = ?", categoryId).Find(&files)

	return &files
}

func UpdateFile(id int, categoriId int, fileName string, content string) *model.File {
	// connect db
	db := ConnectDb()

	// update word by id
	var res model.File
	db.Model(&res).Where("id = ?", id).Updates(model.File{CategoryID: categoriId, FileName: fileName, Content: content})

	return &res
}

func DeleteFile(id int) *string {
	// connect db
	db := ConnectDb()

	// delete word by id
	var file model.File
	db.Delete(&file, id)

	return &file.FileName
}

func CreateCategory(categoryName string) *model.Category {
	// connect db
	db := ConnectDb()

	// append word
	res := model.Category{CategoryName: categoryName}

	// save word
	db.Create(&res)

	return &res
}

func GetCategories() *[]model.Category {
	// connect db
	db := ConnectDb()

	// get all words
	var categories []model.Category
	db.Find(&categories)

	return &categories
}

func GetCategoryById(id int) *model.Category {
	// connect db
	db := ConnectDb()

	// get word
	var category model.Category
	db.First(&category, id)

	return &category
}

func UpdateCategory(id int, categoryName string) *model.Category {
	// connect db
	db := ConnectDb()

	// update word by id
	var res model.Category
	db.Model(&res).Where("id = ?", id).Updates(model.Category{CategoryName: categoryName})

	return &res
}

func DeleteCategory(id int) *string {
	// connect db
	db := ConnectDb()

	// delete word by id
	var category model.Category
	db.Delete(&category, id)

	return &category.CategoryName
}