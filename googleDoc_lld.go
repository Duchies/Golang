// https://github.com/adityatandon15/LLD/blob/main/Lecture%2007/standardUml.png

package main

import (
	"fmt"
	"os"
	"strings"
)

// DocumentElement interface defines the render method
type DocumentElement interface {
	render() string
}

// TextElement implements DocumentElement for text content
type TextElement struct {
	text string
}

func (t TextElement) render() string {
	return t.text
}

// ImageElement implements DocumentElement for images
type ImageElement struct {
	imagePath string
}

func (i ImageElement) render() string {
	return "[Image: " + i.imagePath + "]"
}

// NewLineElement represents a line break
type NewLineElement struct{}

func (n NewLineElement) render() string {
	return "\n"
}

// TabSpaceElement represents a tab space
type TabSpaceElement struct{}

func (t TabSpaceElement) render() string {
	return "\t"
}

// Document holds a collection of elements
type Document struct {
	elements []DocumentElement
}

func (d *Document) addElement(element DocumentElement) {
	d.elements = append(d.elements, element)
}

func (d *Document) render() string {
	var result strings.Builder
	for _, element := range d.elements {
		result.WriteString(element.render())
	}
	return result.String()
}

// Persistence interface defines the save method
type Persistence interface {
	save(data string) error
}

// FileStorage implements Persistence for file system storage
type FileStorage struct{}

func (f FileStorage) save(data string) error {
	err := os.WriteFile("document.txt", []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("error: unable to open file for writing")
	}
	fmt.Println("Document saved to document.txt")
	return nil
}

// DBStorage is a placeholder for database storage
type DBStorage struct{}

func (d DBStorage) save(data string) error {
	// Save to DB implementation would go here
	return nil
}

// DocumentEditor manages client interactions
type DocumentEditor struct {
	document        *Document
	storage         Persistence
	renderedDocument string
}

func NewDocumentEditor(document *Document, storage Persistence) *DocumentEditor {
	return &DocumentEditor{
		document: document,
		storage:  storage,
	}
}

func (e *DocumentEditor) addText(text string) {
	e.document.addElement(TextElement{text: text})
}

func (e *DocumentEditor) addImage(imagePath string) {
	e.document.addElement(ImageElement{imagePath: imagePath})
}

func (e *DocumentEditor) addNewLine() {
	e.document.addElement(NewLineElement{})
}

func (e *DocumentEditor) addTabSpace() {
	e.document.addElement(TabSpaceElement{})
}

func (e *DocumentEditor) renderDocument() string {
	if e.renderedDocument == "" {
		e.renderedDocument = e.document.render()
	}
	return e.renderedDocument
}

func (e *DocumentEditor) saveDocument() error {
	return e.storage.save(e.renderDocument())
}

func main() {
	document := &Document{}
	persistence := FileStorage{}

	editor := NewDocumentEditor(document, persistence)

	// Simulate a client using the editor with common text formatting features
	editor.addText("Hello, world!")
	editor.addNewLine()
	editor.addText("This is a real-world document editor example.")
	editor.addNewLine()
	editor.addTabSpace()
	editor.addText("Indented text after a tab space.")
	editor.addNewLine()
	editor.addImage("picture.jpg")

	// Render and display the final document
	fmt.Println(editor.renderDocument())

	err := editor.saveDocument()
	if err != nil {
		fmt.Println(err)
	}
}
