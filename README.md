post_to_tumblr
==============

Script that posts to tumblr via email. When run it does the following:

* Read last post from `posts/` directory
* Parses human readable title from file name, i.e. `1_Hello_world.md` becomes `Hello world`
* Sends email to tumblr email address using Gmail.

Requirements
------------

* A Tumblr account
* A Gmail account
* Working Go installation

Example
-------
`go run main.go -f <from_your_email@gmail.com> -u <your_gmail_password>`
