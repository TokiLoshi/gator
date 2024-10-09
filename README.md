# Blog Aggregator (aggreGATOR 🐊)

## A Guided Project from Boot.dev

Gator is an RSS feed Aggregator CLI tool built with Go, PostgreSQL and Goose, that allows you to enter in feeds you would like to track, and then you can run the aggregator to get the latest posts from your feed and browse from the comfort of your own terminal.

# Requirements

Gator requires Go(version 1.23+) and Postgres(version 16+) to be installed.

1. Go: If you do not already have Go installed you could use the (webi installer)[https://webinstall.dev/golang/] or browse the (official installation documentation)[https://go.dev/doc/install]. Gator is built suing the go toolchain (version 1.23+), please ensure you are on a similar version for best results.
2. Postgres: Please ensure you already have Postgres installed, ideally using a client like psql. If you don't have postgres installed and you are on a mac you can install postgres with brew
   `brew install postgresql@16`
   Alternatively fro Linux/WSL(Debian) the (official documentation to learn more)[https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql].

Then you can install gator:
`go install github.com/TokiLoshi/gator@latest`

You can activate the psql shell for basic operations either using the following commands:

1. Mac: `psql postgres`
2. Linux: `sudo -u postgres psql`

# How to set up the config and run the program

Go ahead and create your own local database calling it "gator" in the psql shell, and then exit the shell.

After running ` go build .`` run the program you can run  `./gator``` + the command you would like to use, and an argument if necessary.

For example:
`./gator reset` will reset all of the users in your database.
`./gator register` + username will register the username to your database and automatically log them in using the middlewareLoggedIn function. To add a feed you can `./gator addfeed` followed by the url you would like to follow, this will add the feed to your list of feeds for this logged in user. It will also automatically follow this feed. If you lose interest at any time you can use the `./gator unfollow` +feed to unfollow. To scrape for posts you can use the `./gator agg` + request timer (e.g 1s, 10s, 1m, 1hr) to scrape posts. This can take a while and if you need to terminate the process at any time use the ctrl + c command to exit. The agg (scrape) command will save new posts to the database and check for updates, please be careful to not overload sites with scraping traffic. Once you have updated the posts you can run the `./gator browse` + limit command to search for the latest post (ordered) and limit to the number of posts you would like to add.

# Notes:

This project will very likely be getting some updates in the future as well as some additional testing. If you have any questions, comments or feature requests please feel free to add an issue and I'll do my best to address it. I'm still pretty new to Go so I'm sure there is a lot to be improved, I also welcome any feedback anyone is kind enough to take the time to provide. Thank you!
