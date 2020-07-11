package main

import (
	"expvar"
	"fmt"
	"log"
	"os"

	"github.com/ardanlabs/conf"
	"github.com/coreyvan/job-tracker/app/tracker-admin/commands"
	"github.com/coreyvan/job-tracker/business/data"
	"github.com/coreyvan/job-tracker/business/data/schema"
	"github.com/pkg/errors"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {
	log := log.New(os.Stdout, "ADMIN : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(log); err != nil {
		if errors.Cause(err) != commands.ErrHelp {
			log.Printf("error: %s", err)
		}
		os.Exit(1)
	}
}

func run(log *log.Logger) error {

	var cfg struct {
		conf.Version
		Args   conf.Args
		DGraph struct {
			URL string `conf:"default:http://localhost:8080"`
		}
	}
	cfg.Version.SVN = build
	cfg.Version.Desc = "job tracker admin functions"

	const prefix = "TRACKER"
	if err := conf.Parse(os.Args[1:], prefix, &cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, err := conf.Usage(prefix, &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		case conf.ErrVersionWanted:
			version, err := conf.VersionString(prefix, &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config version")
			}
			fmt.Println(version)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	// =========================================================================
	// App Starting

	// Print the build version for our logs. Also expose it under /debug/vars.
	expvar.NewString("build").Set(build)
	log.Printf("main : Started : Application initializing : version %q", build)
	defer log.Println("main: Completed")

	out, err := conf.String(&cfg)
	if err != nil {
		return errors.Wrap(err, "generating config for output")
	}
	log.Printf("main: Config:\n%v\n", out)

	// =========================================================================
	// Commands

	gqlConfig := data.GraphQLConfig{
		URL: cfg.DGraph.URL,
	}

	switch cfg.Args.Num(0) {
	case "import":
		if err := commands.Import(gqlConfig, log, cfg.Args.Num(1)); err != nil {
			return errors.Wrap(err, "importing data")
		}
	case "schema":
		schemaConfig := schema.Config{}
		if err := commands.Schema(gqlConfig, schemaConfig, log); err != nil {
			return errors.Wrap(err, "updating schema")
		}
	}
	return nil
}
