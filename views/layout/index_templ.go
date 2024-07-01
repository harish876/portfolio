// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" data-theme=\"dark\" class=\"dark\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"Harish Portfolio\"><meta name=\"google\" content=\"notranslate\"><link rel=\"stylesheet\" href=\"/css/styles.css\"><link rel=\"stylesheet\" href=\"/css/static.css\"><link rel=\"icon\" type=\"image/x-icon\" href=\"/img/favicon.ico\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layout/index.templ`, Line: 17, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><script src=\"https://unpkg.com/htmx.org@1.9.9\" integrity=\"sha384-QFjmbokDn2DjBjq+fM+8LUIVrAgqcNW2s0PjAxHETgRn9l4fvX31ZxDxvwQnyMOX\" crossorigin=\"anonymous\"></script><script src=\"https://cdn.jsdelivr.net/npm/sweetalert2@11\"></script><script src=\"https://unpkg.com/hyperscript.org@0.9.12\"></script><script defer src=\"https://unpkg.com/@alpinejs/focus@3.x.x/dist/cdn.min.js\"></script><script defer src=\"https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js\"></script><script>\n                /* Script to toggle light and  mode */\n            </script><script type=\"module\" src=\"/js/launch.js\"></script></head><body hx-boost=\"true\" class=\"dark:bg-terminal-base font-mono bg-terminal-base text-white leading-normal tracking-tighter\"><layout id=\"layout\" class=\"w-full h-screen flex flex-col\"><header class=\"bg-primary flex justify-between items-center py-1 px-4\"><div class=\"flex space-x-4\"><button hx-get=\"/\" hx-target=\"body\" hx-swap=\"innerHTML transition:true\" class=\"text-white font-bold tracking-tight btn btn-ghost btn-sm  hover:bg-primary-400 \">home</button> <button hx-get=\"/about\" hx-target=\"#commandDisplay\" hx-swap=\"transition:true\" class=\"text-white btn btn-ghost btn-sm  hover:bg-primary-400\">readme</button> <button hx-get=\"/project\" hx-indicator=\"#spinner\" hx-target=\"#commandDisplay\" hx-swap=\"transition:true\" class=\"text-white btn btn-ghost btn-sm  hover:bg-primary-400\">projects</button> <button hx-get=\"/blog\" hx-indicator=\"#spinner\" hx-target=\"#commandDisplay\" hx-swap=\"transition:true\" class=\"text-white disabled\">blog <span class=\"text-xs\">(coming soon)</span></button></div><div class=\"flex items-center space-x-2\"><a href=\"https://www.linkedin.com/in/harish-gokul-39432718b/\" class=\"text-white btn btn-ghost btn-xs\">LinkedIn</a> <a href=\"https://github.com/harish876\" class=\"text-white btn btn-ghost btn-xs\">Github</a></div></header><div id=\"commandDisplay\" class=\"flex flex-col h-screen flex-grow px-4 py-8 text-xl\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"py-8 group flex h-8 items-center hover:bg-terminal-hover\"><div class=\"bg-primary-400 w-[2px] h-[30px] shrink-0 animate-blink\"></div><form hx-post=\"/commands\" hx-indicator=\"#spinner\" hx-target=\"#commandDisplay\" hx-swap=\"innerHTML transition:true\" hx-on::after-request=\"this.reset()\"><input name=\"command\" id=\"commandBox\" type=\"text\" autocomplete=\"off\" placeholder=\"$ enter command here\" class=\"px-4 border-none focus:ring-transparent text-white leading-tight bg-transparent\"></form></div><footer class=\"bg-primary flex flex-row justify-between items-center\"><div class=\"flex flex-row space-x-4\"><span id=\"currentText\" class=\"bg-secondary-500 text-white px-2 text-sm\">home.txt</span> <span id=\"spinner\" class=\"loading loading-infinity loading-sm htmx-indicator\"></span></div><div class=\"flex flex-row space-x-4\"><span id=\"tabSize\" class=\"text-white px-2 text-sm hover:bg-primary-400 cursor-pointer\">Ln 81, Col 122</span> <span class=\"text-white px-2 text-sm hover:bg-primary-400 cursor-pointer\">htmx</span> <span class=\"text-white px-2 text-sm hover:bg-primary-400 cursor-pointer\">templ</span> <span id=\"tabSize\" class=\"text-white px-2 text-sm hover:bg-primary-400 cursor-pointer\">Tab Size: 4</span> <span id=\"errorText\" class=\"bg-error-base text-white px-2 text-sm\">Errors</span></div></footer></layout><script lang=\"js\">\n\t\t\t     document.body.addEventListener('htmx:beforeSwap', function(evt) {\n\t\t\t    if(evt.detail.xhr.status === 404){\n\t\t\t        // alert the user when a 404 occurs (maybe use a nicer mechanism than alert())\n\t\t\t        alert(\"Error: Could Not Find Resource\");\n\t\t\t    } else if(evt.detail.xhr.status === 422){\n\t\t\t        // allow 422 responses to swap as we are using this as a signal that\n\t\t\t        // a form was submitted with bad data and want to rerender with the\n\t\t\t        // errors\n\t\t\t        //\n\t\t\t        // set isError to false to avoid error logging in console\n                    alert(\"Bad Request Debug\")\n\t\t\t        evt.detail.shouldSwap = true;\n\t\t\t        evt.detail.isError = false;\n\t\t\t    } else if(evt.detail.xhr.status === 418){\n\t\t\t        // if the response code 418 (I'm a teapot) is returned, retarget the\n\t\t\t        // content of the response to the element with the id `teapot`s\n\t\t\t        evt.detail.shouldSwap = true;\n\t\t\t        evt.detail.target = htmx.find(\"#teapot\");\n\t\t\t    }\n\t\t\t});\n\t\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
