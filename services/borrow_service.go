package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorrowBook(ctx context.Context, userID primitive.ObjectID, bookID primitive.ObjectID) error {
	user, err := GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if len(user.BooksTaken) >= 2 {
		return errors.New("Borrowing limit exceeded")
	}

	book, err := GetBookByID(ctx, bookID)
	if err != nil {
		return err
	}

	if book.Borrowed >= book.Count {
		return errors.New("Book unavailable")
	}

	for _, existingBookID := range user.BooksTaken {
		if existingBookID == bookID {
			return errors.New("Already borrowed")
		}
	}

	err = AddBookToUser(ctx, userID, bookID)
	if err != nil {
		return err
	}

	err = UpdateBookBorrowedCount(ctx, bookID, 1)
	return err
}

func ReturnBook(ctx context.Context, userID primitive.ObjectID, bookID primitive.ObjectID) error {
	user, err := GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	bookFound := false
	for _, existingBookID := range user.BooksTaken {
		if existingBookID == bookID {
			bookFound = true
			break
		}
	}

	if !bookFound {
		return errors.New("Not borrowed")
	}

	err = RemoveBookFromUser(ctx, userID, bookID)
	if err != nil {
		return err
	}

	err = UpdateBookBorrowedCount(ctx, bookID, -1)
	return err
}
