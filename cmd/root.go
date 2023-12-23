/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/seba-ban/manifests-tree/pkg/filter"
	"github.com/seba-ban/manifests-tree/pkg/printer"
	"github.com/seba-ban/manifests-tree/pkg/reader"
	"github.com/seba-ban/manifests-tree/pkg/store"
	"github.com/spf13/cobra"
)

var version = "dev"

var genericFilter filter.GenericFilter
var printerOpts printer.PrinterOpts
var outputFormat string
var dirOpts reader.DirWalkerOpts

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "manifests-tree [flags] [inputs...]",
	Version: version,
	Short:   "CLI for displaying Kubernetes resources included in yaml files.",
	Long: `CLI for displaying Kubernetes resources included in yaml files.
One or more inputs are expected. Acceptable inputs are:

- paths to yaml files;
- paths to directories containing yaml files;
- http(s) urls pointing to yaml;
- '-' to read from stdin.
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Please provide yamls to read")
			os.Exit(1)
		}

		docs, err := reader.ReadYamls(&dirOpts, args...)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		store := store.NewDocumentsStore()
		store.AddMany(docs)

		filteredData := genericFilter.RunFilter(store.TreeData())

		printer, err := printer.GetPrinter(outputFormat, printerOpts)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		printer.Print(filteredData)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringSliceVar(&genericFilter.IncludeKinds, "include-kinds", []string{}, "Include resources with these kinds.")
	rootCmd.Flags().BoolVar(&genericFilter.IncludeKindsStrict, "include-kinds-strict", false, "Include resources with these kinds.")
	rootCmd.Flags().StringSliceVar(&genericFilter.ExcludeKinds, "exclude-kinds", []string{}, "Exclude resources with these kinds.")
	rootCmd.Flags().BoolVar(&genericFilter.ExcludeKindsStrict, "exclude-kinds-strict", false, "Exclude resources with these kinds.")

	rootCmd.Flags().StringSliceVar(&genericFilter.IncludeApiVersions, "include-api-versions", []string{}, "Include resources with these api-versions.")
	rootCmd.Flags().BoolVar(&genericFilter.IncludeApiVersionsStrict, "include-api-versions-strict", false, "Include resources with these api-versions.")
	rootCmd.Flags().StringSliceVar(&genericFilter.ExcludeApiVersions, "exclude-api-versions", []string{}, "Exclude resources with these api-versions.")
	rootCmd.Flags().BoolVar(&genericFilter.ExcludeApiVersionsStrict, "exclude-api-versions-strict", false, "Exclude resources with these api-versions.")

	rootCmd.Flags().StringSliceVar(&genericFilter.IncludeNames, "include-names", []string{}, "Include resources with these names.")
	rootCmd.Flags().BoolVar(&genericFilter.IncludeNamesStrict, "include-names-strict", false, "Include resources with these names.")
	rootCmd.Flags().StringSliceVar(&genericFilter.ExcludeNames, "exclude-names", []string{}, "Exclude resources with these names.")
	rootCmd.Flags().BoolVar(&genericFilter.ExcludeNamesStrict, "exclude-names-strict", false, "Exclude resources with these names.")

	rootCmd.Flags().BoolVar(&genericFilter.IncludeUnrecognized, "include-unrecognized", false, "Include unrecognized resources.")

	rootCmd.Flags().BoolVar(&printerOpts.WithPaths, "with-paths", true, "Display resource paths with line numbers.")
	rootCmd.Flags().BoolVar(&printerOpts.OnlyKinds, "only-kinds", false, "Don't display resource names, only kinds.")

	rootCmd.Flags().BoolVarP(&dirOpts.Recursive, "recursive", "r", true, "When reading directory, read yaml files recursively.")

	rootCmd.Flags().StringVarP(&outputFormat, "output", "o", "tree", "Output format. One of: yaml, json, tree.")
}
