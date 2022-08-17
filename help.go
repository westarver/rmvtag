package rmvtag

import (
	"fmt"
	"io"
	"os"
)

//────────────────────┤ getUsage ├────────────────────

func getUsage() string {
	var help = `Usage: rmv-tag [command] [flags...]
	
Commands:     
*+[help] <command>    : Show help and exit, default command
*[remove | rem]       : Remove tagged lines from source.
*[comment | com]      : Do not remove lines, but comment them out.
*[uncomment | unc]    : Restore commented tagged lines.
*[list]	              : List the lines that are tagged.
	
Flags:				
[--mock | -m]                 : Leave input unchanged. send mock changes to stdout.
[--copy | -c] <copies>...    : Filenames to copy the originals to.
[--source | -s] <sources>... : Input files to operate on.
#[--range | -r] <ranges>...   : Restrict the operation to range(s) of lines --range 25 50 100 110.
[--region | -R] <regions>...  : Restrict the operation to named region(s) --region my-reg other-reg.

Long Description:
--source:   Sources are optional.  Missing flag will enable obtaining source file names from stdin.
            Wildcard, filename(s) or blank are all valid.
            Filenames with spaces require quotes.
            Terminate the space separated list with -- if other options follow.

--copy:     Copies are optional.  Missing flag will copy originals to a default file name.
            Terminate the space separated list with -- if other options follow.
			
--range:    Ranges are optional. Use the form -r 10 20 98 108... Numbers will be taken in pairs.
            To use end of file as the range end use a number as large as or larger than the last line number.
            Terminate the space separated list with -- if other options follow.
            Specifying a range or ranges implies a single source file, so any extras will be ignored.
            	
--region:   Regions are optional. Use the form -R my_region ... and separate the regions with a space.
            Terminate the space separated list with -- if other options follow.
            Specifying a region or regions implies a single source file, so any extras will be ignored.
            The use of regions requires that regions are deliniated by region tags.
            See the help for the 'region-tag' tool

More:		
examples: 
rmv-tag [enter]            will use stdin and stdout and use comment as the default command
rmv-tag -c a b -- -s c d   will use a and b for originals and c and d for input  
rmv-tag -s a b             will cause b to be interpreted as a save-to file with no sources given
rmv-tag list               will generate a listing of tagged lines to stdout
rmv-tag com                will comment out all lines included in a tagged range or region
rmv-tag rem -m -s *.go     will send the changed sources to stdout and leave the originals unchanged

The default range is the entire source file
The default region is the entire source file
`
	return help
}

//────────────────────┤ ShowHelp ├────────────────────

func ShowHelp(w io.Writer, command ...string) {

	var help = `Usage: rmv-tag [command] [flags...]
	
    Commands:     
    [help] <command>    : Show help and exit, default command
    [remove | rem]       : Remove tagged lines from source.
    [comment | com]      : Do not remove lines, but comment them out.
    [uncomment | unc]    : Restore commented tagged lines.
    [list]	              : List the lines that are tagged.
        
    Flags:				
    [--mock | -m]                 : Leave input unchanged. send mock changes to stdout.
    [--copy | -c] <copies>...    : Filenames to copy the originals to.
    [--source | -s] <sources>... : Input files to operate on.
    [--range | -r] <ranges>...   : Restrict the operation to range(s) of lines --range 25 50 100 110.
    [--region | -R] <regions>...  : Restrict the operation to named region(s) --region my-reg other-reg.
    
    Long Description:
    --source:   Sources are optional.  Missing flag will enable obtaining source file names from stdin.
                Wildcard, filename(s) or blank are all valid.
                Filenames with spaces require quotes.
                Terminate the space separated list with -- if other options follow.
    
    --copy:     Copies are optional.  Missing flag will copy originals to a default file name.
                Terminate the space separated list with -- if other options follow.
                
    --range:    Ranges are optional. Use the form -r 10 20 98 108... Numbers will be taken in pairs.
                To use end of file as the range end use a number as large as or larger than the last line number.
                Terminate the space separated list with -- if other options follow.
                Specifying a range or ranges implies a single source file, so any extras will be ignored.
                    
    --region:   Regions are optional. Use the form -R my_region ... and separate the regions with a space.
                Terminate the space separated list with -- if other options follow.
                Specifying a region or regions implies a single source file, so any extras will be ignored.
                The use of regions requires that regions are deliniated by region tags.
                See the help for the 'region-tag' tool
    
    More:		
    examples: 
    rmv-tag [enter]            will use stdin and stdout and use comment as the default command
    rmv-tag -c a b -- -s c d   will use a and b for originals and c and d for input  
    rmv-tag -s a b             will cause b to be interpreted as a save-to file with no sources given
    rmv-tag list               will generate a listing of tagged lines to stdout
    rmv-tag com                will comment out all lines included in a tagged range or region
    rmv-tag rem -m -s *.go     will send the changed sources to stdout and leave the originals unchanged
    
    The default range is the entire source file
    The default region is the entire source file
    `
	if len(command) != 0 {
		fmt.Fprintf(w, "%s\n", command[0])
		os.Exit(0)
	} else {
		fmt.Fprintln(w, help)
	}
}
