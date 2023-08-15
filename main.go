package main

import "fmt"

func main(){
  var publisher, writer, artist, title, genre string
  var year, pageNumber uint
  var grade float32
  title = "Mr. GoToSleep"
  writer = "Tracey hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  genre = "Mystery"
  year = 1997
  pageNumber = 14
  grade = 6.5
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "year:", year, "genre:", genre, "pages:", pageNumber, "grade:", grade)

  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc."
  genre = "Action"
  year = 2013
  pageNumber = 160
  grade = 9.0

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "year:", year, "genre:", genre, "pages:", pageNumber, "grade:", grade)

  
}