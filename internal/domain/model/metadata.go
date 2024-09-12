package metadata

import "fmt"

type PackageMetadata struct {
	Name string
	Link string
}

func (p *PackageMetadata) ToString() string {
	return fmt.Sprintf("%s %s", p.Name, p.Link)
}

type FileMetadata struct {
	Name    string
	Package PackageMetadata
}

func (f *FileMetadata) ToString() string {
	return fmt.Sprintf("%s %s", f.Name, f.Package.ToString())
}

type FunctionMetadata struct {
	Name      string
	Signature string
	Comment   string
	File      FileMetadata
}

func (f *FunctionMetadata) ToString() string {
	return fmt.Sprintf("%s %s %s %s", f.Name, f.Signature, f.Comment, f.File.ToString())
}
