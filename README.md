# web-service-gin

Developing a RESTful API with Go and Gin
This tutorial introduces the basics of writing a RESTful web service API with Go and the Gin Web Framework (Gin).

You'll get the most out of this tutorial if you have a basic familiarity with Go and its tooling. If this is your first exposure to Go, please see Tutorial: Get started with Go for a quick introduction.

Gin simplifies many coding tasks associated with building web applications, including web services. In this tutorial, you'll use Gin to route requests, retrieve request details, and marshal JSON for responses.

In this tutorial, you will build a RESTful API server with two endpoints. Your example project will be a repository of data about vintage jazz records.

The tutorial includes the following sections:

- Design API endpoints.
- Create a folder for your code.
- Create the data.
- Write a handler to return all items.
- Write a handler to add a new item.
- Write a handler to return a specific item.
- Design API endpoints
You'll build an API that provides access to a store selling vintage recordings on vinyl. So you'll need to provide endpoints through which a client can get and add albums for users.

When developing an API, you typically begin by designing the endpoints. Your API's users will have more success if the endpoints are easy to understand.

Here are the endpoints you'll create in this tutorial.
```sh
/albums
```
GET – Get a list of all albums, returned as JSON.
POST – Add a new album from request data sent as JSON.

```sh
/albums/:id
```
GET – Get an album by its ID, returning the album data as JSON.
Next, you'll create a project for your code.

https://shell.cloud.google.com/?walkthrough_tutorial_url=https%3A%2F%2Fraw.githubusercontent.com%2Fgolang%2Ftour%2Fmaster%2Ftutorial%2Fweb-service-gin.md&pli=1&show=ide&environment_deployment=ide

# Step 1
## Step 1 of 7
Create a project for your code
To begin, create a project for the code you'll write.

Ensure that the cloudshell_open folder is selected.

Click the File Menu, then click New Folder.

In the New Folder dialog, enter web-service-gin for the folder name, then click OK.

Click the File Menu, then click Open Workspace.

In the Open Workspace dialog, select the cloudshell_open/web-service-gin folder you just created, then click Open.

Click the New Terminal menu command to open a new Cloud Shell terminal.

The terminal prompt should be in the web-service-gin directory.

Create a module in which you can manage dependencies.

Run the go mod init command, giving it the path of the module your code will be in.
```sh
go mod init \
    example.com/web-service-gin
go: creating new go.mod: module \
    example.com/web-service-gin
```
This command creates a go.mod file in which dependencies you might add will be listed for tracking. For more, be sure to see Managing dependencies.

Next, you'll design data structures for handling data.


# Step 2
## Step 2 of 7
Create the data
To keep things simple for the tutorial, you'll store data in memory. A more typical API would interact with a database.

Note that storing data in memory means that the set of albums will be lost each time you stop the server, then recreated when you start it.

Write the code
Click the File Menu, then click New File.

In the New File dialog, enter main.go for the file name, then click OK.

Into main.go, at the top of the file, paste the package declaration below.
```Go
package main
```
A standalone program (as opposed to a library) is always in package main.

Beneath the package declaration, paste the following declaration of an album struct. You'll use this to store album data in memory.

Struct tags such as json:"artist" specify what a field's name should be when the struct's contents are serialized into JSON. Without them, the JSON would use the struct's capitalized field names – a style not as common in JSON.
```Go
// album represents data about a record album.
type album struct {
        ID     string  `json:"id"`
        Title  string  `json:"title"`
        Artist string  `json:"artist"`
        Price  float64 `json:"price"`
}
```
Beneath the struct declaration you just added, paste the following slice of album structs containing data you'll use to start.

```Go
// albums slice to seed record album data.
var albums = []album{
        {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
        {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
        {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
```
Next, you'll write code to implement your first endpoint.