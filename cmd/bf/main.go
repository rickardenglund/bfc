package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"bfc/gen"
	"bfc/lex"
	"bfc/parse"
	"bfc/runner"
)

func main() {
	root := cobra.Command{
		Use:   "bf",
		Short: "brainfuck compiler",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	root.AddCommand(createRunCommand())
	root.AddCommand(createBuildCommand())

	err := root.Execute()
	if err != nil {
		panic(err)
	}
}

func createBuildCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:  "build",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			shouldPrint, err := cmd.Flags().GetBool("print")
			if err != nil {
				return fmt.Errorf("failed to read code: %w", err)
			}

			outFile, err := cmd.Flags().GetString("out")

			prog, err := read(args[0])
			if err != nil {
				return fmt.Errorf("failed to parse code: %w", err)
			}

			assemblyBytes := gen.Generate(prog)
			if shouldPrint {
				fmt.Printf("%s", assemblyBytes)
				return nil
			}

			err = asm(assemblyBytes)
			if err != nil {
				return fmt.Errorf("failed to assemble: %w", err)
			}

			err = link(outFile)

			return nil
		},
	}

	cmd.Flags().BoolP("print", "p", false, "print state after execution")
	cmd.Flags().StringP("out", "o", "a.out", "where to place the binary")

	return &cmd
}

func asm(assemblyBytes []byte) error {
	_, err := exec.LookPath("as")
	if err != nil {
		return fmt.Errorf("did not find as command: %w", err)
	}

	as := exec.Command("as", "-o", "a.o")
	as.Stdin = bytes.NewBuffer(assemblyBytes)
	out := bytes.NewBuffer(nil)
	as.Stdout = out
	err = as.Run()
	if err != nil {
		fmt.Printf("%s\n", out.String())
		return fmt.Errorf("failed to assemble: %w", err)
	}

	return nil

}

func link(outfile string) error {
	_, err := exec.LookPath("ld")
	if err != nil {
		return fmt.Errorf("did not find ld command: %w", err)
	}

	as := exec.Command("ld", "-e", "_start", "-o", outfile, "a.o")
	out := bytes.NewBuffer(nil)
	as.Stdout = out
	err = as.Run()
	if err != nil {
		fmt.Printf("%s\n", out.String())
		return fmt.Errorf("failed to link: %w", err)
	}

	return nil

}
func createRunCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:  "run",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			shouldPrintState, err := cmd.Flags().GetBool("print")
			if err != nil {
				return err
			}

			prog, err := read(args[0])
			if err != nil {
				return err
			}

			s := runner.Run(prog)
			if shouldPrintState {
				fmt.Printf("\nState after running\n%s\n", s.String())
			}

			return nil

		},
	}

	cmd.Flags().BoolP("print", "p", false, "print state after execution")

	return &cmd
}

func read(path string) ([]parse.Instruction, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tokens, err := lex.Lex([]rune(string(bs)))
	if err != nil {
		return nil, err
	}

	program, err := parse.Parse(tokens)
	if err != nil {
		return nil, err
	}

	return program, nil
}
