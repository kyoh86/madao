package main

import (
	"errors"
	"fmt"

	"github.com/kyoh86/madao/madao"
	"github.com/spf13/cobra"
)

var moveRawFlags = struct {
	ScopeFiles  []string
	FormatPatch string
}{}

type moveFlags struct {
	ScopeFiles  []madao.Glob
	FormatPatch string
}

var moveCommand = &cobra.Command{
	Use:     "move [flags] <source> <destination>",
	Aliases: []string{"mv"},
	Short:   "Move lines between files",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		var patches []madao.Patch
		f, err := parseMoveCommandArgs(cmd, args)
		if err != nil {
			return fmt.Errorf("command validation error: %w", err)
		}

		sourceContent, drawPatch, err := madao.Draw(ctx, f.SourceRange)
		if err != nil {
			return fmt.Errorf("drawing source: %w", err)
		}
		patches = append(patches, drawPatch)

		sourceIDs, err := madao.SelectContentIDs(ctx, sourceContent)
		if err != nil {
			return fmt.Errorf("search IDs from source: %w", err)
		}

		newContent, err := madao.ReplaceLinksInContent(
			ctx,
			sourceContent,
			f.Source,
			f.Destination,
		)
		if err != nil {
			return fmt.Errorf("process links in the source: %w", err)
		}

		replacePatches, err := madao.ReplaceLinksInDocuments(
			ctx,
			f.ScopeFileGlobs,
			f.Source,
			f.Destination,
		)
		if err != nil {
			return fmt.Errorf("replace links in documents in the scope: %w", err)
		}
		patches = append(patches, replacePatches...)

		placePatch, err := madao.Insert(ctx, f.Source, f.Destination, newContent)
		if err != nil {
			return fmt.Errorf("insert contents to destination document: %w", err)
		}

		patches = append(patches, placePatch)

		if f.FormatPatch != "" {
			madao.DumpPatch(patches...)
		} else {
			madao.ApplyPatch(patches...)
		}
		return nil
	},
}

func parseMoveCommandArgs(
	cmd *cobra.Command,
	args []string,
) (*moveFlags, error) {
	switch len(args) {
	case 0, 1:
		return nil, errors.New("shortage of the arguments")
	case 2:
		// noop
	default:
		return nil, errors.New("surplus of the arguments")
	}

	source, err := madao.ParseFileRange(args[0])
	if err != nil {
		return nil, fmt.Errorf("parse <source> argument: %w", err)
	}

	destination, err := madao.ParseFileRange(args[1])
	if err != nil {
		return nil, fmt.Errorf("parse <destination> argument: %w", err)
	}

	if destination.HasEndLine() {
		return nil, errors.New(
			"invalid <destination> argument: it cannot accept a line of end",
		)
	}
	// TODO: parse moveRawFlags.ScopeFiles to build Glob matcher?
	return &moveFlags{
		// TODO:
	}, nil
}

func init() {
	moveCommand.Flags().
		StringSliceVarP(&moveRawFlags.ScopeFiles, "scope-files", "", nil, "Glob patterns for related Markdown documents. default is ./**/*.md")
	moveCommand.Flags().
		StringVarP(&moveRawFlags.FormatPatch, "format-patch", "p", "", "Generate a patch file in stead of processing files directory")
	facadeCommand.AddCommand(moveCommand)
}
