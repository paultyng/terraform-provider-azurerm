package main

import (
	"html/template"
	"io"
)

type Resource struct {
	Provider         string // azurerm
	Name             string // azurerm_image
	NameSuffix       string // image
	ShortDescription string // Get information about an Image
	Description      string // Use this data source to access information about an Image.
	// TODO: +Example usage, etc?
	// TODO: resource category
	Attributes []Attribute
}

type Attribute struct {
	Name        string
	Description string
	Optional    bool
	Required    bool
	Computed    bool

	Attributes []Attribute
	Min        int
	Max        int
}

func GeneratePage(r *Resource, w io.Writer) error {
	tmpl, err := template.ParseFiles("websitedoc/resource.tmpl")
	if err != nil {
		return err
	}
	return tmpl.Execute(w, r)
}
