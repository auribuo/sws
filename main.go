package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var dir string
var port int
var ssl bool
var certFile string
var keyFile string

var rootCmd = &cobra.Command{
	Use:   "sws",
	Short: "sws is a simple web server",
	Long:  `sws is a simple web server that can be used to serve static files from a directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(NewServer(dir, port, ssl, certFile, keyFile).ListenAndServe())
	},
}

func init() {
	rootCmd.Flags().StringVarP(&dir, "dir", "d", ".", "The directory to serve files from.")
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "The port to serve on.")
	rootCmd.Flags().BoolVarP(&ssl, "ssl", "s", false, "Use https.")
	rootCmd.Flags().StringVarP(&certFile, "cert", "c", fmt.Sprintf("%s/.certs/localhost/localhost.crt", os.Getenv("HOME")), "The certificate file.")
	rootCmd.Flags().StringVarP(&keyFile, "key", "k", fmt.Sprintf("%s/.certs/localhost/localhost.key", os.Getenv("HOME")), "The key file.")
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
