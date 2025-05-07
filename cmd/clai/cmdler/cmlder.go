package cmdler

import (
	"fmt"
	"os"
	"strings"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/process"
	"github.com/spf13/cobra"
	"github.com/v4nderstruck/clai/internal"
)

var Cmdler = &cobra.Command{
	Use:   "cmdler [user prompt]",
	Short: "Generate a Cli cmd from prompt query",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		if args[0] == "" {
			return fmt.Errorf("Empty user prompt")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		getModel, err := cmd.Flags().GetString("model")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to access model flag: %v\n", err)
			os.Exit(1)
		}

		claiTool, err := internal.NewClaiTool(getModel)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Wrong model flag: %v\n", err)
			os.Exit(1)
		}

		systemPrompt, err := generateSystemPrompt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not generate a system prompt : %v\n", err)
			os.Exit(1)
		}

		userPrompt := args[0]

		fmt.Printf("systemPrompt: %s\n\n", systemPrompt)
		fmt.Printf("userPrompt: %s\n\n", userPrompt)
		fmt.Printf("ModelHelp: %s\n\n", claiTool.AiModel.ModelHelp())

		cmdler, err := claiTool.AiModel.OneShotPrompt(0, systemPrompt, userPrompt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not generate a system prompt : %v\n", err)
			os.Exit(1)
		}

		fmt.Println(cmdler)
	},
}

var (
	thinkingLevel int32
	model         string
	debug         bool
	streamed      bool
)

const systemPrompt = `You are an Linux Expert. The System is running a %s ` +
	`operating system with a %s shell. You will receive a query by the ` +
	`administrator of the machine to generate a CLI command for the given shell. ` +
	`Assume that all software is installed if not further stated in the admin query. ` +
	`Only generate the CLI command in the response without additional description or formatting in markdown or json. ` +
	`Your commands MUST be shell one-liners.`

func generateSystemPrompt() (string, error) {
	var err error

	// Getting platform version
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("could not get platform information %v", err)
	}
	platformName := fmt.Sprintf("%s %s", info.Platform, info.PlatformVersion)

	// Getting shell
	ppid := os.Getppid()
	parent, err := process.NewProcess(int32(ppid))
	if err != nil {
		return "", fmt.Errorf("could not get shell information %v", err)
	}
	shellName, err := parent.Name()
	if err != nil {
		return "", fmt.Errorf("could not get shell information %v", err)
	}
	shellName = strings.ToLower(shellName)
	shellName = strings.TrimSuffix(shellName, ".exe")

	return fmt.Sprintf(systemPrompt, platformName, shellName), nil
}

func Init() {
	Cmdler.PersistentFlags().Int32VarP(&thinkingLevel, "think", "t", 0, "Set value [0-2] for increasing thinking levels")
	Cmdler.PersistentFlags().StringVarP(&model, "model", "m", "gemini", "Set Name of model family")
	Cmdler.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Print additional debug outputm")
	Cmdler.PersistentFlags().BoolVarP(&streamed, "stream", "s", false, "Print output as streamed text")
}
