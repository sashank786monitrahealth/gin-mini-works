Based on work: https://www.youtube.com/watch?v=LOn1GUsjOF4

Today we will be making a webapi using Go and a popular web framework Gin. Lets get started:
First thing you want to do is make a folder and we are going to call this folder newsfeeder. 
As the name suggests, it is going to be an app that posts news items and lets people read news items. 

* mkdir newsfeeder
* cd newsfeeder
* touch main.go
* go mod init newsfeeder


This creates a file - go.mod. It contains a inventory of all your dependencies. It is similar to:
package.json file in node js.

Also, it allows your go compiler to know where your project is indexed. Where the root folder of your project is. In our case, it is here - where our main.go resides.

* go run main.go

It is working.

* touch makefile

add a makefile. 