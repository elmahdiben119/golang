package cmd

import "gihub.com/elmahdiben119/copyfiles/src/fs"

func init() {
	RootCommand.PersistentFlags().BoolVarP(&fs.DEBUG, "debug", "d", fs.DEBUG, "enable debug mode for display tracing")
	RootCommand.PersistentFlags().BoolVarP(&fs.MULTI_THREADS, "multithread", "m", fs.MULTI_THREADS, "use all cpus (warning in some case its useless)")
	RootCommand.PersistentFlags().BoolVarP(&fs.PROGRESS, "progress", "p", fs.PROGRESS, "display progress bar")
	// ssh flags
	RootCommand.PersistentFlags().StringVarP(&fs.SSH_USER, "ssh-user", "u", fs.SSH_USER, "ssh user")
	RootCommand.PersistentFlags().StringVarP(&fs.SSH_PASS, "ssh-pass", "p", fs.SSH_PASS, "ssh password")
	RootCommand.PersistentFlags().StringVarP(&fs.SSH_HOST, "ssh-host", "h", fs.SSH_HOST, "ssh host")
	RootCommand.PersistentFlags().StringVarP(&fs.SSH_PORT, "ssh-port", "p", fs.SSH_PORT, "ssh port")

	// ftp flags
	RootCommand.PersistentFlags().StringVarP(&fs.FTP_USER, "ftp-user", "u", fs.FTP_USER, "ftp user")
	RootCommand.PersistentFlags().StringVarP(&fs.FTP_PASS, "ftp-pass", "p", fs.FTP_PASS, "ftp password")
	RootCommand.PersistentFlags().StringVarP(&fs.FTP_HOST, "ftp-host", "h", fs.FTP_HOST, "ftp host")
	RootCommand.PersistentFlags().StringVarP(&fs.FTP_PORT, "ftp-port", "p", fs.FTP_PORT, "ftp port")
}
