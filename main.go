package main

import (
	"fmt"
	"lol/lol"
	"net/http"
	"strings"
)

// indexHandler handles HTTP requests to the root URL ("/").
func indexHandler(w http.ResponseWriter, r *http.Request) {
	bodyContent := "Add some text in the text box and enjoy :)"

	// Check if the request method is POST.
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "400: Bad Request; Unable to parse form", http.StatusBadRequest)
			return
		}
		// Get the input text and style from the form.
		text := r.Form.Get("text")
		style := r.Form.Get("banner")
		// Generate ASCII art based on the input text and style.
		// save, _ := PrintArt(text, style)
		// Ensure the style file has a ".txt" extension.
		if !strings.Contains(style, ".txt") {
			style += ".txt"
		}
		fmt.Println(text)
		save := lol.AsciiArtTransform(text, style)

		store := "<br>"
		for i := 0; i < 8*len(save); i++ {
			store += save[i]
		}
		store = `<pre style="color: #333366; text-align: left;">` + store + "</pre>"
		bodyContent = store
		// bodyContent = save // Use the generated ASCII art as the response content.
	}

	// HTML template for the web page.
	tmpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>ASCII Art Web</title>
		<link rel="stylesheet" href="/static/style.css">
	</head>
	<body>
		<div class="container">
			<form action="/" method="POST"> <!-- Change action to "/" to match the route -->
				<div>
					<h1>ASCII-art-web</h1>
					<h2>Style</h2>
				</div>
				<div>
					<input type="radio" name="banner" value="standard">Standard
					<input type="radio" name="banner" value="shadow">Shadow
					<input type="radio" name="banner" value="thinkertoy">Thinkertoy
				</div>
				<div>
					<h3>Input text</h3>
				</div>
				<div>
					<textarea placeholder="Input text here" id="add-text" name="text" rows="4" cols="50"></textarea>
				</div>
				<div>
					<input type="submit" value="Generate">
				</div>
			</form>
		</div>
		<div class="box">
			<pre>%s</pre> <!-- Display the generated ASCII art -->
		</div>
		
	</body>
	</html>
	`
	// Send the HTML response with the generated ASCII art.
	fmt.Fprintf(w, tmpl, bodyContent)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Register the indexHandler to handle requests at the root URL ("/").
	http.HandleFunc("/", indexHandler)
	// Start an HTTP server on port 8080.
	http.ListenAndServe(":8080", nil)
}
