# Simple Landing Page
This is my personal site which includes a short biography and my social links.

The site itself is hosted as a Github page (HTML only), but also [on Render as a Go project](https://personal-site-jx8a.onrender.com/).

The Github pages site is relatively simple: it's pure HTML using Tailwind CSS.

The Go version of this site only has one dependency in the Chi router, uses Go
HTML templates to display the pages, and even utilizes the Go embed package for
the layouts. A Dockerfile is used to deploy this project on Render. 

Check out the site and feel free to let me know what you think!
