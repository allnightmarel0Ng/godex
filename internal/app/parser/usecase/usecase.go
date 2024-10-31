package usecase

import (
	"bytes"
	"errors"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"net/url"
	"strings"
	"unicode"
	"net/http"
	"fmt"
	"io"

	"github.com/allnightmarel0Ng/godex/internal/domain/model"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
)

type ParserUseCase interface {
	ProduceMessage(toSend model.FunctionMetadata) error
	ExtractFunctions(code []byte, fileName, packageName, link string) ([]model.FunctionMetadata, error)
	ParseUrl(rawUrl string) (string, string, string, error)
	FetchFile(url string) ([]byte, error)
}

type parserUseCase struct {
	producer  *kafka.Producer
	whiteList map[string]bool
}

func NewParserUseCase(producer *kafka.Producer, whiteList map[string]bool) ParserUseCase {
	return &parserUseCase{
		producer:  producer,
		whiteList: whiteList,
	}
}

func (p *parserUseCase) ProduceMessage(toSend model.FunctionMetadata) error {
	log.Printf("trying to produce: %s", toSend.ToString())
	return p.producer.Produce("functions", []byte(toSend.ToString()))
}

func (p *parserUseCase) FetchFile(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", response.Status)
	}

	return io.ReadAll(response.Body)
}

func (p *parserUseCase) ParseUrl(rawUrl string) (string, string, string, error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", "", "", err
	}

	if !strings.HasSuffix(parsedUrl.Path, ".go") {
		return "", "", "", errors.New("not a .go file")
	}

	if !p.whiteList[parsedUrl.Hostname()] {
		return "", "", "", errors.New("not in whitelist")
	}

	tokens := strings.Split(rawUrl, "/")
	tokensLength := len(tokens)

	return tokens[tokensLength-1], tokens[tokensLength-2], rawUrl, nil
}

func (p *parserUseCase) ExtractFunctions(code []byte, fileName, packageName, link string) ([]model.FunctionMetadata, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var functions []model.FunctionMetadata
	ast.Inspect(file, func(n ast.Node) bool {
		function, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		if !unicode.IsUpper(rune(function.Name.Name[0])) {
			return true
		}

		var signature bytes.Buffer
		if function.Type.Params != nil {
			signature.WriteString("(")
			for i, params := range function.Type.Params.List {
				if i > 0 {
					signature.WriteString(",")
				}
				printer.Fprint(&signature, fset, params.Type)
			}
			signature.WriteString(")")
		}
		if function.Type.Results != nil {
			if len(function.Type.Results.List) > 1 {
				signature.WriteString("(")
			}
			for i, results := range function.Type.Results.List {
				if i > 0 {
					signature.WriteString(",")
				}
				printer.Fprint(&signature, fset, results.Type)
			}
			if len(function.Type.Results.List) > 1 {
				signature.WriteString(")")
			}
		}
		comment := "NoComment"
		if function.Doc != nil {
			comment = function.Doc.Text()
		}

		functions = append(functions, model.FunctionMetadata{
			Name:      function.Name.Name,
			Signature: strings.Replace(signature.String(), " ", "", -1),
			Comment:   comment,
			File: model.FileMetadata{
				Name: fileName,
				Package: model.PackageMetadata{
					Name: packageName,
					Link: link,
				},
			},
		})
		return true
	})

	return functions, nil
}
