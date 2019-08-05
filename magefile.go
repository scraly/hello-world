// +build mage

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Default    = Build
	goFiles    = getGoFiles()
	goSrcFiles = getGoSrcFiles()
)

var curDir = func() string {
	name, _ := os.Getwd()
	return name
}()

var version = "0.0.1"

// Calculate file paths
var toolsBinDir = normalizePath(path.Join(curDir, "tools", "bin"))

func init() {
	time.Local = time.UTC

	// Add local bin in PATH
	err := os.Setenv("PATH", fmt.Sprintf("%s:%s", toolsBinDir, os.Getenv("PATH")))
	if err != nil {
		panic(err)
	}
}

func Build() {
	banner := figure.NewFigure("Hello World", "", true)
	banner.Print()

	fmt.Println("")
	color.Red("# Build Info ---------------------------------------------------------------")
	fmt.Printf("Go version : %s\n", runtime.Version())
	fmt.Printf("Git revision : %s\n", hash())
	fmt.Printf("Git branch : %s\n", branch())
	// TODO: replace version with tag()
	fmt.Printf("Tag : %s\n", version)

	fmt.Println("")

	color.Red("# Core packages ------------------------------------------------------------")
	mg.SerialDeps(Go.Deps, Go.Generate, Go.Format, Go.Lint, Go.Test)

	fmt.Println("")
	color.Red("# Artifacts ----------------------------------------------------------------")
	mg.Deps(Bin.HelloWorld)
}

// -----------------------------------------------------------------------------

type Gen mg.Namespace

// Generate swagger files
// func (Gen) Swagger() {
// 	color.Blue("### Generate swagger files")

// 	mustGoGenerate("Swagger", "github.com/scraly/hello-world/internal/app/swagger")
// 	mustGoGenerate("Swagger", "github.com/scraly/hello-world/pkg/swagger")
// }

// Generate mocks for tests
func (Gen) Mocks() {
	color.Blue("### Mocks")

	mustGoGenerate("Mocks", "github.com/scraly/hello-world/internal/app")
}

// -----------------------------------------------------------------------------

type Go mg.Namespace

// Generate go code
func (Go) Generate() error {
	color.Cyan("## Generate code")
	// mg.SerialDeps(Gen.Mocks)
	return nil
}

// Test run go test
func (Go) Test() error {
	color.Cyan("## Running unit tests")
	sh.Run("mkdir", "-p", "test-results/junit")
	return sh.RunV("gotestsum", "--junitfile", "test-results/junit/unit-tests.xml", "--", "-short", "-race", "-cover", "-coverprofile", "test-results/cover.out", "./...")
}

// AnalyzeCoverage calculates a single coverage percentage
func (Go) AnalyzeCoverage() error {
	color.Cyan("## Analyze tests coverage")
	return sh.RunV("go-agg-cov", "-coverFile", "test-results/cover.out", "-businessLogicFile", "businesslogic.txt", "-minCoverageThreshold", "50")
}

// Test run go test
func (Go) IntegrationTest() {
	color.Cyan("## Running integration tests")
	sh.Run("mkdir", "-p", "test-results/junit")
}

// Verifying dependencies
func (Go) Verify() error {
	fmt.Println("## Verifying dependencies")
	return sh.RunV("go", "mod", "verify")
}

// Tidy add/remove depenedencies.
func (Go) Tidy() error {
	fmt.Println("## Cleaning go modules")
	return sh.RunV("go", "mod", "tidy", "-v")
}

// Deps install dependency tools.
func (Go) Deps() error {
	color.Cyan("## Vendoring dependencies")
	return sh.RunV("go", "mod", "vendor")
}

// Format runs gofmt on everything
func (Go) Format() error {
	color.Cyan("## Format everything")
	args := []string{"-s", "-w"}
	args = append(args, goFiles...)
	return sh.RunV("gofumpt", args...)
}

// Import runs goimports on everything
func (Go) Import() error {
	color.Cyan("## Process imports")
	args := []string{"-w"}
	args = append(args, goSrcFiles...)
	return sh.RunV("goreturns", args...)
}

// Lint run linter.
func (Go) Lint() error {
	mg.Deps(Go.Format)
	color.Cyan("## Lint go code")
	return sh.RunV("golangci-lint", "run")
}

type Bin mg.Namespace

// Build licman microservice
func (Bin) HelloWorld() error {
	return goBuild("github.com/scraly/hello-world/cli/hello-world", "hello-world")
}

func goBuild(packageName, out string) error {
	fmt.Printf(" > Building %s [%s]\n", out, packageName)

	// TODO: version replace with tag()
	varsSetByLinker := map[string]string{
		"github.com/scraly/hello-world/internal/version.Version":   version,
		"github.com/scraly/hello-world/internal/version.Revision":  hash(),
		"github.com/scraly/hello-world/internal/version.Branch":    branch(),
		"github.com/scraly/hello-world/internal/version.BuildUser": os.Getenv("USER"),
		"github.com/scraly/hello-world/internal/version.BuildDate": time.Now().Format(time.RFC3339),
		"github.com/scraly/hello-world/internal/version.GoVersion": runtime.Version(),
	}
	var linkerArgs string
	for name, value := range varsSetByLinker {
		linkerArgs += fmt.Sprintf(" -X %s=%s", name, value)
	}
	linkerArgs = fmt.Sprintf("-s -w %s", linkerArgs)

	return sh.Run("go", "build", "-mod", "vendor", "-ldflags", linkerArgs, "-o", fmt.Sprintf("bin/%s", out), packageName)
}

// -----------------------------------------------------------------------------

func getGoFiles() []string {
	var goFiles []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "vendor/") {
			return filepath.SkipDir
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		goFiles = append(goFiles, path)
		return nil
	})

	return goFiles
}

func getGoSrcFiles() []string {
	var goSrcFiles []string

	for _, path := range goFiles {
		if !strings.HasSuffix(path, "_test.go") {
			continue
		}

		goSrcFiles = append(goSrcFiles, path)
	}

	return goSrcFiles
}

// tag returns the git tag for the current branch or "" if none.
func tag() string {
	s, _ := sh.Output("git", "describe", "--tags")
	fmt.Println(s)
	return s
}

// hash returns the git hash for the current repo or "" if none.
func hash() string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return hash
}

// branch returns the git branch for current repo
func branch() string {
	hash, _ := sh.Output("git", "rev-parse", "--abbrev-ref", "HEAD")
	return hash
}

func mustStr(r string, err error) string {
	if err != nil {
		panic(err)
	}
	return r
}

func mustGoGenerate(txt, name string) {
	fmt.Printf(" > %s [%s]\n", txt, name)
	err := sh.RunV("go", "generate", name)
	if err != nil {
		panic(err)
	}
}

func runIntegrationTest(txt, name string) {
	fmt.Printf(" > %s [%s]\n", txt, name)
	err := sh.RunV("gotestsum", "--junitfile", fmt.Sprintf("test-results/junit/integration-%s.xml", strings.ToLower(txt)), name, "--", "-tags=integration", "-race")
	if err != nil {
		panic(err)
	}
}

// normalizePath turns a path into an absolute path and removes symlinks
func normalizePath(name string) string {
	absPath := mustStr(filepath.Abs(name))
	return absPath
}