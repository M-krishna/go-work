package main

import "fmt"

// composition instead of inheritance
// composition can be achieved by embedding one struct type into another
// blog post is a perfect example

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
	blog_post := constructBlogPost("First Blog Post", "My First Line", krishna)
	fmt.Println(krishna)
	fmt.Println(blog_post)
	fmt.Println(krishna.fullName())
	blog_post.details()
}
