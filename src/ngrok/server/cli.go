package server

import (
	"flag"
)

type Options struct {
	domainHost   string
	httpAddr   string
	tunnelAddr string
	domain     string
	logto      string
	loglevel   string
	secretPath   string
}

func parseArgs() *Options {
	domainHost := flag.String("domainHost", "ngrok.com", "http/https tunnel host, default: request host header")
	httpAddr := flag.String("httpAddr", ":80", "Public address for HTTP connections, empty string to disable")
	tunnelAddr := flag.String("tunnelAddr", ":4443", "Public address listening for ngrok client")
	domain := flag.String("domain", "ngrok.com", "Domain where the tunnels are hosted")
	logto := flag.String("log", "stdout", "Write log messages to this file. 'stdout' and 'none' have special meanings")
	loglevel := flag.String("log-level", "DEBUG", "The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")
	secretPath := flag.String("secretPath", "/etc/ngrok-secrets", "Path to authentication information")
	flag.Parse()

	return &Options{
		domainHost:   *domainHost,
		httpAddr:   *httpAddr,
		tunnelAddr: *tunnelAddr,
		domain:     *domain,
		logto:      *logto,
		loglevel:   *loglevel,
		secretPath: *secretPath,
	}
}
