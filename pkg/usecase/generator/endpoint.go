package generator

import (
	"context"
	stdtypes "go/types"
	stdstrings "strings"

	"github.com/swipe-io/swipe/pkg/types"

	"github.com/iancoleman/strcase"
	"github.com/swipe-io/swipe/pkg/domain/model"
	"github.com/swipe-io/swipe/pkg/importer"
	"github.com/swipe-io/swipe/pkg/strings"

	"github.com/swipe-io/swipe/pkg/writer"
)

type EndpointOption struct {
}

type endpoint struct {
	*writer.GoLangWriter

	info model.GenerateInfo
	o    model.ServiceOption
	i    *importer.Importer
}

func (g *endpoint) Imports() []string {
	return g.i.SortedImports()
}

func (g *endpoint) Process(ctx context.Context) error {
	kitEndpointPkg := g.i.Import("endpoint", "github.com/go-kit/kit/endpoint")
	contextPkg := g.i.Import("context", "context")
	typeStr := stdtypes.TypeString(g.o.Type, g.i.QualifyPkg)

	g.W("type EndpointSet struct {\n")
	for _, m := range g.o.Methods {
		g.W("%sEndpoint %s.Endpoint\n", m.Name, kitEndpointPkg)
	}
	g.W("}\n")

	g.W("func MakeEndpointSet(s %s) EndpointSet {\n", typeStr)
	g.W("return EndpointSet{\n")
	for _, m := range g.o.Methods {
		g.W("%[1]sEndpoint: make%[1]sEndpoint(s),\n", m.Name)
	}
	g.W("}\n")
	g.W("}\n")

	for _, m := range g.o.Methods {
		if len(m.Params) > 0 {
			g.W("type %sRequest%s struct {\n", m.LcName, g.o.ID)
			for _, p := range m.Params {
				g.W("%s %s `json:\"%s\"`\n", strings.UcFirst(p.Name()), stdtypes.TypeString(p.Type(), g.i.QualifyPkg), strcase.ToLowerCamel(p.Name()))
			}
			g.W("}\n")
		}

		if m.ResultsNamed {
			g.W("type %sResponse%s struct {\n", m.LcName, g.o.ID)
			for _, p := range m.Results {
				name := p.Name()
				g.W("%s %s `json:\"%s\"`\n", strings.UcFirst(name), stdtypes.TypeString(p.Type(), g.i.QualifyPkg), strcase.ToLowerCamel(name))
			}
			g.W("}\n")
		}

		g.W("func make%sEndpoint(s %s", m.Name, typeStr)
		g.W(") %s.Endpoint {\n", kitEndpointPkg)
		g.W("w := func(ctx %s.Context, request interface{}) (interface{}, error) {\n", contextPkg)

		var callParams []string

		if m.ParamCtx != nil {
			callParams = append(callParams, "ctx")
		}

		callParams = append(callParams, types.Params(m.Params, func(p *stdtypes.Var) []string {
			name := p.Name()
			name = stdstrings.ToUpper(name[:1]) + name[1:]
			return []string{"req." + name}
		}, nil)...)

		if len(m.Params) > 0 {
			g.W("req := request.(%sRequest%s)\n", m.LcName, g.o.ID)
		}

		if len(m.Results) > 0 {
			if len(m.Results) > 1 && m.ResultsNamed {
				for i, p := range m.Results {
					if i > 0 {
						g.W(", ")
					}
					g.W(p.Name())
				}
			} else {
				g.W("result")
			}
			g.W(", ")
		}
		if m.ReturnErr != nil {
			g.W("err")
		}
		g.W(" := ")

		g.WriteFuncCall("s", m.Name, callParams)
		if m.ReturnErr != nil {
			g.WriteCheckErr(func() {
				g.W("return nil, err\n")
			})
		}
		g.W("return ")
		if len(m.Results) > 0 {
			if len(m.Results) > 1 && m.ResultsNamed {
				g.W("%sResponse%s", m.LcName, g.o.ID)
				g.WriteStructAssign(structKeyValue(m.Results, nil))
			} else {
				g.W("result")
			}
		} else {
			g.W("nil")
		}
		g.W(" ,nil\n")
		g.W("}\n")
		g.W("return w\n")
		g.W("}\n\n")
	}
	return nil
}

func (g *endpoint) Filename() string {
	return "endpoint_gen.go"
}

func (g *endpoint) OutputDir() string {
	return ""
}

func (g *endpoint) PkgName() string {
	return ""
}

func NewEndpoint(info model.GenerateInfo, o model.ServiceOption, i *importer.Importer) Generator {
	return &endpoint{
		GoLangWriter: writer.NewGoLangWriter(i),
		info:         info,
		i:            i,
		o:            o,
	}
}