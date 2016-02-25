package main // Author wwalker
import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/alecthomas/kingpin"
)

type Cfg struct {
	logfile        *os.File
	printfFormat   string
	prefix         string
	suffix         string
	timeFormat     string
	noStdout       bool
	timeInFilename bool
	timeZoneUtc    bool
	appendToLog    bool
	startLine      bool
	endLine        bool
}

func (config *Cfg) parseArgs() {
	app := kingpin.New("ilts - In Line Time Stamper", "ilts prepends messages with a time stamp and writes to stdout")

	app.Flag("prefix", "turns on logging to the file prefixed with filepath").Short('p').StringVar(&config.prefix)
	app.Flag("printf-format", "printf style for how to print the timestamp with the message ; defaults to \"%s - %s\n\"").Short('P').StringVar(&config.printfFormat)
	app.Flag("suffix", "suffix appended to filepath").Short('s').StringVar(&config.suffix)
	app.Flag("nostdout", "disable writing to stdout").Short('n').BoolVar(&config.noStdout)
	app.Flag("time", "add timestamp to filepath after prefix").Short('t').BoolVar(&config.timeInFilename)
	app.Flag("time-format", "add timestamp to filepath after prefix").Short('T').StringVar(&config.timeFormat)
	app.Flag("utc", "set timezone for timestamp to UTC").Short('u').BoolVar(&config.timeZoneUtc)
	app.Flag("append", "append to rather than truncate existin logfile").Short('a').BoolVar(&config.appendToLog)
	app.Flag("start-line", "immediately write a timestamped \"Starting\" upon startup").Short('S').BoolVar(&config.startLine)
	app.Flag("end-line", "write a timestamped \"Ending\" when exiting").Short('E').BoolVar(&config.endLine)

	kingpin.MustParse(app.Parse(os.Args[1:]))
	config.unsupportedFlags()
	if config.timeFormat == "" {
		config.timeFormat = "2006-01-02_15:04:05.000000"
	}
	// TODO
	if config.printfFormat == "" {
		config.printfFormat = "%s - %s\n"
	}
}

func (config *Cfg) unsupportedFlags() {
	fatal := false
	if config.timeZoneUtc {
		fmt.Printf("--utc|-u is not currently supported")
	}
	if fatal {
		os.Exit(1)
	}
}

func die(message string) {
	fmt.Printf(message + "\n")
	os.Exit(1)
}

func main() {
	var cfg *Cfg

	cfg = &Cfg{}
	cfg.parseArgs()
	cfg.openLogFile()

	if cfg.startLine {
		cfg.printMessage("Execution begins")
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		cfg.printMessage(line)
	}
	if cfg.endLine {
		cfg.printMessage("Execution ends")
	}
}

func (cfg *Cfg) openLogFile() {
	var err error
	var filepath string

	now := time.Now()
	if cfg.prefix != "" {
		filepath = cfg.prefix
		if cfg.timeInFilename {
			filepath = filepath + now.Format(cfg.timeFormat)
		}
		if cfg.suffix != "" {
			filepath = filepath + cfg.suffix
		}
		if cfg.appendToLog {
			cfg.logfile, err = os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				die("Failed to open file: <" + filepath + "> for append")
			}
		} else {
			cfg.logfile, err = os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
			if err != nil {
				die("Failed to open new file: <" + filepath + ">")
			}
		}
	}
}

func (cfg *Cfg) printMessage(message string) {
	now := time.Now()
	nowString := now.Format(cfg.timeFormat)

	if cfg.prefix != "" {
		fmt.Fprintf(cfg.logfile, cfg.printfFormat, nowString, message)
	}
	if !cfg.noStdout {
		fmt.Printf(cfg.printfFormat, nowString, message)
	}
}
