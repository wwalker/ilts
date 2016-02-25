# NAME

ilts - In Line Time Stamps : for adding timestamps to stdout of a program without log formatting

# SYNOPSIS

    ilts <-P|--printf-format format>
         <-p|--prefix filename_prefix>
           <-n|--no-stdout>
           <-s|--suffix filename_suffix>
           <-t|--time>
           <-T|--time-format time_format>
             <-u|--utc>
           <-a|--append>
         <-S|--start-line>
         <-E|--end-line>

# DESCRIPTION

ilts reads stdin and writes to stdout.  It reads each line, then outputs, to stdout, a timestamp, " - ", then the line.

optionally ilts will also write to a log file, in addition to stdout. This occurs if -p is used.  -n may be used in conjuction with -p to write to a log file only.


# OPTIONS

     -p|--printf-format format
       format to print each line of input.  defaults to "%s - %s\n"
     -p|--prefix filepath
        the use of -p turns on logging to the file prefixed with filepath.
          -n|--no-stdout
            the use of -n turns off writing to stdout and may not be used without -p
          -t|--time
            the use of -t appends the current time to filepath in the format YYYYmmdd-HHMMSS in the current timezone
              -u|--utc
                the use of -u in conjuction with -t (or -T) changes the time to UTC
          -T|--time-format Go Time Format
            the use of -T turns on letting the user specifying the Format
          -s|--suffix fileext
            the use of -s turns on a file extension added to the filepath after the timestamp (if enabled)
          -a|--append
            the use of -a appends to the log file rather than truncating the file
     -S|--start-line
     -E|--end-line
        the use of -S or -E enables a start and end line.  these are useful in case the app generates no output, or delay a long time befor the first line of output or after the last line of output.  The messages are "Execution begins" and "Execution ends"

# BUGS

See https://github.com/wwalker/ilts/issues

May spontaneously spew forth unspeakable evil.

# AUTHOR

Wayne Walker < wwalker@solid-constructs.com >

# SEE ALSO

     tee(1)
