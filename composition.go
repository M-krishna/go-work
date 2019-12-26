package main

import "fmt"

/* 
	composition instead of inheritance
 	composition can be achieved by embedding one struct type into another
	blog post is a perfect example
*/


// blog post contains author
type author struct{
	firstName string
	lastName string
	bio string
}

// each post contains title, content and author
type post struct{
	title string
	content string
	author
}

/* Embedding Slice of struct */
type Website struct {
	posts []post
}

func (w Website) Contents() {
	fmt.Println("Contents of the Website")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}


// function to get the fullname of the author
func (a author) fullName() string{
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

// function to get the details about the blog post
func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.author.fullName())
	fmt.Println("Bio: ", p.author.bio)
}

func constructAuthor(firstname, lastname, Bio string) author{
	return author{
		firstName: firstname,
		lastName: lastname,
		bio: Bio,
	}
}

func constructBlogPost(Title string, Content string, a author) post{
	return post{
		title: Title,
		content: Content,
		author: a,
	}
}

func main(){
	krishna := constructAuthor("Krishna", "Murugan", "I'm a Programmer!")
	blog_post_1 := constructBlogPost("First Blog Post", "My First Line", krishna)
	blog_post_2 := constructBlogPost("2nd Blog Post", "My First Line", krishna)
	blog_post_3 := constructBlogPost("3rd Blog Post", "My First Line", krishna)
	blog_post_4 := constructBlogPost("4th Blog Post", "My First Line", krishna)
	w := Website{
		posts: []post{blog_post_1, blog_post_2, blog_post_3, blog_post_4},
	}
	fmt.Println(krishna)
	fmt.Println(blog_post_1)
	fmt.Println(krishna.fullName())
	blog_post_1.details()
	fmt.Println()

	/*
		Embedding Slice of structs.
	*/
	w.Contents()
}
