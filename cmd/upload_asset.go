package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/itscaro/gitlab-utils/git"
	"github.com/spf13/cobra"
)

var uploadAssetCmdOpts struct {
	configFile string
	projectUrl string
	project    string
	tag        string
	file       string
	url        string
	assetName  string
}

func createUploadAssetCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:           "upload-asset",
		Short:         "",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          runUploadAssetCmd,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			printMemory()
		},
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Could not determine working directory")
	}
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.configFile, "config", "c", filepath.Join(dir, "label.yml"), "Project")
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.project, "project", "p", "", "Project")
	_ = cmd.MarkPersistentFlagRequired("project")
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.projectUrl, "project-url", "u", "", "Project Url")
	_ = cmd.MarkPersistentFlagRequired("project-url")
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.tag, "tag", "t", "", "Tag")
	_ = cmd.MarkPersistentFlagRequired("tag")
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.file, "file", "", "", "File to upload")
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.url, "url", "", "", "File to upload")
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.assetName, "name", "", "", "Name of the asset")

	return cmd
}

func runUploadAssetCmd(cmd *cobra.Command, args []string) error {
	createClient()

	return git.UploadAsset(
		uploadAssetCmdOpts.projectUrl,
		uploadAssetCmdOpts.project,
		uploadAssetCmdOpts.tag,
		uploadAssetCmdOpts.assetName,
		uploadAssetCmdOpts.file,
		uploadAssetCmdOpts.url,
	)
}
