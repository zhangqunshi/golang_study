package main

import "fmt"

type Book struct {
	name string
	price int
}

func main() {
	var book1 Book;
	var book2 Book;
	
	book1.name = "书名1"
	book1.price = 100
	
	book2 = Book {"书名2", 200}
	
	//fmt.Printf( "Book 1 name : %s\n", book1.name)
	//fmt.Printf( "Book 1 price : %d\n", book1.price)
	printBook(book1);
		
	fmt.Printf( "Book 2 : %s\n", book2)
	

}

func printBook( book Book ) {
   fmt.Printf( "Book name : %s\n", book.name);
   fmt.Printf( "Book price : %d\n", book.price)
}