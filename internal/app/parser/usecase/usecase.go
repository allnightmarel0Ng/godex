package usecase

import (
	"fmt"
	"go/ast"

	metadata "github.com/allnightmarel0Ng/godex/internal/domain/model"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
)

type ParserUseCase interface {
}

type parserUseCase struct {
	producer *kafka.Producer
}

func (p *parserUseCase) ProduceMessage(toSend metadata.FunctionMetadata) error {
	return p.producer.Produce("functions", []byte(toSend.ToString()))
}

func (p *parserUseCase) ExtractFunctions(file *ast.File, fileName, packageName, link string) []metadata.FunctionMetadata {
	var functions []metadata.FunctionMetadata
	ast.Inspect(file, func(n ast.Node) bool {
		function, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		var signature string
		if function.Type.Params != nil {
			signature += "("
			for i, params := range function.Type.Params.List {
				if i > 0 {
					signature += ", "
				}
				for _, name := range params.Names {
					signature += name.Name + " "
				}
				signature += fmt.Sprintf("%s", params.Type)
			}
			signature += ")"
		}
		if function.Type.Results != nil {
			signature += " "
			if len(function.Type.Results.List) > 1 {
				signature += "("
			}
			for i, results := range function.Type.Results.List {
				if i > 0 {
					signature += ", "
				}
				signature += fmt.Sprintf("%s", results.Type)
			}
			if len(function.Type.Results.List) > 1 {
				signature += ")"
			}
		}
		comment := ""
		if function.Doc != nil {
			comment = function.Doc.Text()
		}

		functions = append(functions, metadata.FunctionMetadata{
			Name:      function.Name.Name,
			Signature: signature,
			Comment:   comment,
			File: metadata.FileMetadata{
				Name: fileName,
				Package: metadata.PackageMetadata{
					Name: packageName,
					Link: link,
				},
			},
		})
		return true
	})

	return functions
}
