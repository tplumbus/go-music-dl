package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"
	// Commit is set at build time via ldflags
	Commit = "none"
	// Date is set at build time via ldflags
	Date = "unknown"
)

func main() {
	app := &cli.App{
		Name:    "go-music-dl",
		Usage:   "Download music from various platforms",
		Version: fmt.Sprintf("%s (commit: %s, built: %s)", Version, Commit, Date),
		Authors: []*cli.Author{
			{
				Name: "go-music-dl contributors",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output directory for downloaded files",
				Value:   ".",
				EnvVars: []string{"MUSIC_DL_OUTPUT"},
			},
			&cli.StringFlag{
				Name:    "quality",
				Aliases: []string{"q"},
				Usage:   "Audio quality (low, medium, high, lossless)",
				Value:   "lossless",
				EnvVars: []string{"MUSIC_DL_QUALITY"},
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Enable verbose logging",
				EnvVars: []string{"MUSIC_DL_VERBOSE"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "download",
				Aliases:   []string{"dl"},
				Usage:     "Download a song or playlist by URL or ID",
				ArgsUsage: "<url-or-id>",
				Action:    downloadCmd,
			},
			{
				Name:      "search",
				Aliases:   []string{"s"},
				Usage:     "Search for music by keyword",
				ArgsUsage: "<keyword>",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "limit",
						Usage: "Maximum number of results to return",
						Value: 20,
					},
					&cli.StringFlag{
						Name:  "platform",
						Usage: "Platform to search on (netease, qq, kugou, kuwo, migu)",
						Value: "netease",
					},
				},
				Action: searchCmd,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// downloadCmd handles the download subcommand
func downloadCmd(c *cli.Context) error {
	if c.NArg() == 0 {
		return cli.ShowCommandHelp(c, "download")
	}

	target := c.Args().First()
	outputDir := c.String("output")
	quality := c.String("quality")
	verbose := c.Bool("verbose")

	if verbose {
		fmt.Printf("Downloading: %s\n", target)
		fmt.Printf("Output dir: %s\n", outputDir)
		fmt.Printf("Quality: %s\n", quality)
	}

	// TODO: implement actual download logic
	fmt.Printf("Downloading %s to %s (quality: %s)\n", target, outputDir, quality)
	return nil
}

// searchCmd handles the search subcommand
func searchCmd(c *cli.Context) error {
	if c.NArg() == 0 {
		return cli.ShowCommandHelp(c, "search")
	}

	keyword := c.Args().First()
	limit := c.Int("limit")
	platform := c.String("platform")

	// TODO: implement actual search logic
	fmt.Printf("Searching for \"%s\" on %s (limit: %d)\n", keyword, platform, limit)
	return nil
}
