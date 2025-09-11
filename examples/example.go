package examples

import (
	"context"
	"fmt"
	"main/database"
)

func Example(ctx context.Context) {
	// err := book_operations.AddBookToMongoDB(ctx, "Sherlock Holmes", "Sir Arthut Conan Doyle")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("Book Created")

	// id, err := book_operations.GetBookByName("Sherlock Holmes", ctx)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Book ID:", id.Hex())
	// }
	// fmt.Println("Get Book ID")

	// err = book_operations.DeleteBookFromMongoDB(ctx, id)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("Book Deleted")

	// err = database.GetBooksCollection().Drop(ctx)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("Books Collection Dropped")

	// err := user_operations.AddUserToMongoDB(ctx, "Ertugrul", "Test123")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("User Created")

	// id, err := user_operations.GetUserByName("Ertugrul", ctx)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Book ID:", id.Hex())
	// }
	// fmt.Println("Get User ID")
	// return id, err

	// err = user_operations.DeleteUserFromMongoDB(ctx, id)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("User Deleted")

	err := database.GetUsersCollection().Drop(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Users Collection Dropped")
}
