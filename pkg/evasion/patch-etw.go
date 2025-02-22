package evasion

import (
	"embed"
	"strings"

	"github.com/guervild/uru/pkg/common"
	"github.com/guervild/uru/pkg/models"
)

type PatchEtwEvasion struct {
	Name         string
	Description  string
	Debug        bool
	UseBanana 	 string
	SubName 	 string
}

func NewPatchEtwEvasion() models.ObjectModel {

	return &PatchEtwEvasion{
		Name:         "patchetw",
		Description:  `Path etw. Can use BananaPhone if set. (credits: method taken from Merlin, @Ne0nd0g)
  Argument(s):
    UseBanana : use BananaPhone to perfom syscall`,		
		UseBanana: "false",
		SubName: common.RandomStringOnlyChar(5),
	}
}

func (e *PatchEtwEvasion) GetImports() []string {

	imports := []string{
		`"fmt"`,
		`"syscall"`,
		`"unsafe"`,
		`"golang.org/x/sys/windows"`,
	}

	if strings.ToLower(e.UseBanana) == "true" {
		imports = append(imports, `bananaphone "github.com/C-Sto/BananaPhone/pkg/BananaPhone"`)
	}

	return imports
}

func (e *PatchEtwEvasion) RenderInstanciationCode(data embed.FS) (string, error) {
	return common.CommonRendering(data, "templates/evasions/patchetw/instanciation.go.tmpl", e)
}

func (e *PatchEtwEvasion) RenderFunctionCode(data embed.FS) (string, error) {
	return common.CommonRendering(data, "templates/evasions/commons/patch_functions.go.tmpl", e)
}
