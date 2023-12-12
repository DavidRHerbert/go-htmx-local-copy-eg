# go-htmx-local-copy-eg
Example of using go and htmx which has no external dependencies.
This simple web page that uses htmx to update the page without reloading.
htmx.min.js is loaded from the root where no Internet access is required.
This repo documents the user problem I ran into, and thus the solution is provided for others. 

# htmx installation by downloading a local copy

https://htmx.org/docs/#installing

The htmx org states on their web page that Htmx is a dependency-free, browser-oriented javascript library.
This was ideal for my use case of an organisations internal web communications app where WWW WAN is unavailable.

I developed with CDN Loading
```html
<script src="https://unpkg.com/htmx.org@1.9.9" integrity="sha384-QFjmbokDn2DjBjq+fM+8LUIVrAgqcNW2s0PjAxHETgRn9l4fvX31ZxDxvwQnyMOX" crossorigin="anonymous"></script>
```

But when I swaped to local down loaded copy 
```html
<script src="/path/to/htmx.min.js"></script>
```

I ran into browser error issues and hung loading.

# The solution Background 

After much a doo I managed to solve this case.
A solution not being available on the internet,I have provided this bairbones code example for your public reference. 
( You will find references to this case in stackoverflow and github issues )

# The 3 step Solution

1) use go to handle the serving of static files 👍 
```go
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

2) Move the htmx.min.js into a root/static folder 👍 
 https://github.com/bigskysoftware/htmx/releases/download/v1.9.8/htmx.min.js

3) serve in html 🥇
```html
<script src="static/htmx.min.js"></script>
```





