package book

import "gorm.io/gorm"

type BookRepository interface {
	FindAll() []Book //This method is expected to fetch all the books in storage and return them as a slice.
	FindByID(id int64) Book
	Save(book Book) (*Book, error)
	Update(book Book) (*Book, error)
	Delete(book Book) (*Book, error) //This method supposedly deletes the given book and returns a pointer to the deleted book along with error information if an error occurs.
}

type BookRepositoryImplement struct {
	db *gorm.DB
}

func NewBookRepositoryImplement(db *gorm.DB) *BookRepositoryImplement {
	return &BookRepositoryImplement{db}
}

func (r *BookRepositoryImplement) FindAll() []Book {
	// Logic to fetch all books from storage and return them as a slice
	var books []Book
	_ = r.db.Find(&books)
	return books
}

func (r *BookRepositoryImplement) FindByID(id int64) Book {
	// Logic to fetch a book by ID from storage and return it
	var book Book
	_ = r.db.First(&book, id)
	return book
}

func (r *BookRepositoryImplement) Save(book Book) (*Book, error) {
	// Logic to save a book to storage
	// Return a pointer to the saved book along with error information

	bookResult := r.db.Save(&book)

	if bookResult.Error != nil {
		return nil, bookResult.Error
	}
	return &book, nil // Return the pointer to the saved book
}

func (r *BookRepositoryImplement) Update(book Book) (*Book, error) {
	// Logic to update a book in storage
	// Return a pointer to the updated book along with error information

	bookResult := r.db.Save(&book)

	if bookResult.Error != nil {
		return nil, bookResult.Error
	}
	return &book, nil
}

func (r *BookRepositoryImplement) Delete(book Book) (*Book, error) {

	bookResult := r.db.Delete(&book) //NOTE: use GORM to remove the book from storage using the Delete method of r.db (the object representing the database connection).

	if bookResult.Error != nil {
		return nil, bookResult.Error //If there is an error deleting the book, return nil for the book pointer and return an error object obtained from the bookResult.
	}
	return &book, nil //If the deletion is successful (without errors), returns a pointer to the deleted book.
}
