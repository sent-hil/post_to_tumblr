I wanted to write my posts in markdown and in my favorite text editor: vim, so
I wrote a little [script][1] that does that. The posts are stored in 'posts/'
dir in the same dir where the script is run from. The script first gets the
last file (based on title, all posts are prefixed with a number), parses title
of post from the file name, reads the contents and sends an email to a tumblr
with all that info.

The script itself took only 20 minutes to write, but figuring out Go's smtp
library took way longer. If you plan to use the script for yourself, be sure
to change tumblr's email address to you own. The script accepts -f flag for
from/username, i.e. your email address and -p flag for password. You didn't
think I'd hardcode my email and password, did you? Run it like you'd do any Go
script `go run main.go -f <from> -p <password>`

The script itself is no means complete, but it works. This and the last
post were posted using it.

[1]: https://github.com/sent-hil/post_to_tumblr/blob/master/main.go
