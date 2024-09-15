package model

import "fmt"

type PackageMetadata struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func (p *PackageMetadata) ToString() string {
	return fmt.Sprintf("%s %s", p.Name, p.Link)
}

type FileMetadata struct {
	Name    string          `json:"name"`
	Package PackageMetadata `json:"package"`
}

func (f *FileMetadata) ToString() string {
	return fmt.Sprintf("%s %s", f.Name, f.Package.ToString())
}

type FunctionMetadata struct {
	Name      string       `json:"name"`
	Signature string       `json:"signature"`
	Comment   string       `json:"comment"`
	File      FileMetadata `json:"file"`
}

func (f *FunctionMetadata) ToString() string {
	return fmt.Sprintf("%s %s %s %s", f.Name, f.Signature, f.Comment, f.File.ToString())
}
