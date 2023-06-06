//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var tools = []string{
	"github.com/bufbuild/buf/cmd/buf",
	"github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking",
	"github.com/bufbuild/buf/cmd/protoc-gen-buf-lint",
	"github.com/golangci/golangci-lint/cmd/golangci-lint",
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway",
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2",
	"golang.org/x/tools/cmd/cover",
	"golang.org/x/tools/cmd/goimports",
	"google.golang.org/grpc/cmd/protoc-gen-go-grpc",
	"google.golang.org/protobuf/cmd/protoc-gen-go",
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime",
	"google.golang.org/grpc",
	"google.golang.org/protobuf/types/known/anypb",
	"github.com/golang/protobuf/ptypes/any",
	"google.golang.org/protobuf/types/known/structpb",
	"github.com/golang/protobuf/ptypes/any",
	"github.com/golang/protobuf/jsonpb",
}

var project_map = map[string]string{
	"vectoredge.io/build_system":      "./build_system",
	"vectoredge.io/provision_service": "./provision_service",
	"vectoredge.io/connector_service": "./connector_service",
	"vectoredge.io/metadata_service":  "./metadata_service",
	"vectoredge.io/proto_layer":       "./proto_layer",
	"vectoredge.io/utils":             "./utils",
	"vectoredge.io/communication":     "./communication",
}

var Default = Build

func Bench() error {
	fmt.Println("Running benchmarks...")

	if err := sh.RunV("go", "test", "-run", "XXX", "-bench", ".", "-benchmem", "-short", "./..."); err != nil {
		return err
	}

	fmt.Println("Done.")
	return nil
}

// Bootstrap installs the tools required for development

func Bootstrap() error {
	fmt.Println("Bootstrapping tools...")

	if err := os.MkdirAll("_tools", 0755); err != nil {
		return fmt.Errorf("Error creating dir: %w", err)
	}

	// create module if go.mod doesn't exist

	if _, err := os.Stat("_tools/go.mod"); os.IsNotExist(err) {
		cmd := exec.Command("go", "mod", "init", "tools")
		cmd.Dir = "_tools"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	getPackage := []string{"get", "-u", "-v"}
	getPackage = append(getPackage, tools...)

	cmd := exec.Command("go", getPackage...)
	cmd.Dir = "_tools"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	install := []string{"install", "-v"}
	install = append(install, tools...)

	cmd = exec.Command("go", install...)
	cmd.Dir = "_tools"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	if _, err := os.Stat("../go.work"); os.IsNotExist(err) {

		println("\n[WORKSPACE]: making a go.work file")

		working_dir, _ := os.Getwd()
		working_dir_split := strings.Split(working_dir, "/")
		parent_dir := strings.Join(working_dir_split[:len(working_dir_split)-1], "/")

		cmd := exec.Command("go", "work", "init")
		cmd.Dir = parent_dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}

		println("[WORKSPACE]: editing the go.work file")

		for url, local_path := range project_map {
			if _, err := os.Stat(fmt.Sprintf("../%v/go.mod", local_path)); os.IsNotExist(err) {
				continue
			}

			cmd := exec.Command(
				"go",
				"work",
				"edit",
				"-replace",
				fmt.Sprintf("%v@v0.0.1=%v", url, local_path),
			)
			cmd.Dir = parent_dir
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				return err
			}

			cmd = exec.Command(
				"go",
				"work",
				"use",
				fmt.Sprintf("%v", local_path),
			)
			cmd.Dir = parent_dir
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				return err
			}
		}

		println("[WORKSPACE]: finished the go.work file")
	}
	return nil
}

func Proto() error {
	mg.Deps(Bootstrap)

	fmt.Println("Generating proto files...")
	os.Chdir("../proto_layer")
	if _, err := os.Stat("buf.yaml"); err != nil {
		if err := sh.Run("buf", "mod", "init"); err != nil {
			fmt.Println("error in buf mod init")
			return err
		}
	}

	if err := sh.Run("buf", "mod", "update"); err != nil {
		fmt.Println("error in buf mod update")
		return err
	}

	if err := sh.RunV("buf", "generate"); err != nil {

		os.Chdir("../")
		return err
	}
	os.Chdir("../")
	return nil
}

// Clean cleans up the built files

func Clean() error {
	fmt.Println("Cleaning...")

	if err := sh.RunV("go", "mod", "tidy"); err != nil {
		return fmt.Errorf("tidying go.mod: %w", err)
	}
	working_dir, _ := os.Getwd()
	working_dir_split := strings.Split(working_dir, "/")
	parent_dir := strings.Join(working_dir_split[:len(working_dir_split)-1], "/")
	fmt.Println("the parent dir is ", parent_dir)
	clean := []string{
		"dist/*",
		"pkg/*",
		"bin/*",
		"build/*",
		fmt.Sprintf("%v/proto_layer/gen/*", parent_dir),
		"shared_lib/shared_obj/*",
	}
	for _, dir := range clean {
		fmt.Printf("[CLEANING]: %v\n", dir)
		if err := os.RemoveAll(dir); err != nil {
			return fmt.Errorf("removing dir: %q: %w", dir, err)
		}
	}

	return nil
}

type buildMode uint8

const (
	// buildModeDev builds the project for development, without bundling assets
	buildModeDev buildMode = iota
	// BuildModeProd builds the project similar to a release build
	buildModeProd
)

func build(mode buildMode) error {
	buildDate := time.Now().UTC().Format(time.RFC3339)
	buildArgs := make([]string, 0)

	switch mode {
	case buildModeProd:
		buildArgs = append(buildArgs, "-tags", "assets")
	}

	gitCommit, err := sh.Output("git", "rev-parse", "HEAD")
	if err != nil {
		return fmt.Errorf("getting git commit: %w", err)
	}

	buildArgs = append(
		[]string{
			"build",
			"-trimpath",
			"-ldflags",
			fmt.Sprintf("-X main.commit=%s -X main.date=%s", gitCommit, buildDate),
		},
		buildArgs...)
	buildArgs = append(buildArgs, "-o", "./bin/flipt", "./cmd/flipt/")

	return sh.RunV("go", buildArgs...)
}

// Build builds the project similar to release build
func Build() error {
	mg.Deps(Clean)

	if err := build(buildModeProd); err != nil {
		return err
	}
	return nil
}

// Dev builds the project for development, without bundling assets
func Dev() error {
	mg.Deps(Clean)
	fmt.Println("Building...")

	if err := build(buildModeDev); err != nil {
		return err
	}

	fmt.Println("Done.")
	fmt.Printf("\nRun the following to start Flipt server:\n")
	fmt.Printf("\n%v\n", color.CyanString(`./bin/flipt --config config/local.yml`))
	fmt.Printf("\nIn another shell, run the following to start the UI in dev mode:\n")
	fmt.Printf("\n%v\n", color.CyanString(`cd ui && npm run dev`))
	return nil
}
