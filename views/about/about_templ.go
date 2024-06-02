// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package about

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AboutWithoutLayoutComponent() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"h-full flex-grow px-4\"><div class=\"group flex h-8 items-center active:border-primary-400  active:border-l-2 active:pl-[22px] active:text-gray-11 active:bg-gray-6 focus:border-primary-400 focus:border-l-2 focus:pl-[22px] focus:text-gray-11 focus:bg-gray-6 focus:outline-none has-[:focus]:border-orange has-[:focus]:border-l-2 has-[:focus]:pl-[22px] has-[:focus]:text-gray-11 has-[:focus]:bg-gray-6 has-[:focus]:outline-none\"><span class=\"text-xl font bold pb-8\">About</span></div><pre class=\"text-white flex flex-col gap-4\"><div tabindex=\"0\" number=\"1\" class=\"group px-4 flex h-8 items-center text-gray-500 hover:bg-terminal-hover hover:border-primary-400 hover:border-l-2 active:border-primary-400  active:border-l-2 active:pl-[22px] active:text-white active:bg-gray-6 focus:border-primary-400 focus:border-l-2 focus:pl-[22px] focus:text-white focus:bg-gray-6 focus:outline-none has-[:focus]:border-orange has-[:focus]:border-l-2 has-[:focus]:pl-[22px] has-[:focus]:text-gray-11 has-[:focus]:bg-gray-6 has-[:focus]:outline-none\">1. hi my name harish gokul. i am a junior software engineer</div><div tabindex=\"0\" number=\"2\" class=\"group px-4 flex h-8 items-center text-gray-500 hover:bg-terminal-hover hover:border-primary-400 hover:border-l-2 active:border-primary-400  active:border-l-2 active:pl-[22px] active:text-white active:bg-gray-6 focus:border-primary-400 focus:border-l-2 focus:pl-[22px] focus:text-white focus:bg-gray-6 focus:outline-none has-[:focus]:border-orange has-[:focus]:border-l-2 has-[:focus]:pl-[22px] has-[:focus]:text-gray-11 has-[:focus]:bg-gray-6 has-[:focus]:outline-none\">2. i love building CLI applications, reading up on database internals, building minimalistic LSP's and exploring cool data structures. </div><div tabindex=\"0\" number=\"3\" class=\"group px-4 flex h-8 items-center text-gray-500 hover:bg-terminal-hover hover:border-primary-400 hover:border-l-2 active:border-primary-400  active:border-l-2 active:pl-[22px] active:text-white active:bg-gray-6 focus:border-primary-400 focus:border-l-2 focus:pl-[22px] focus:text-white focus:bg-gray-6 focus:outline-none has-[:focus]:border-orange has-[:focus]:border-l-2 has-[:focus]:pl-[22px] has-[:focus]:text-gray-11 has-[:focus]:bg-gray-6 has-[:focus]:outline-none\">3. i am an upcoming masters grad at UC Davis.</div><div tabindex=\"0\" number=\"4\" class=\"group text-wrap px-4 py-2 flex h-auto items-center text-gray-500 hover:bg-terminal-hover hover:border-primary-400 hover:border-l-2 active:border-primary-400  active:border-l-2 active:pl-[22px] active:text-white active:bg-gray-6 focus:border-primary-400 focus:border-l-2 focus:pl-[22px] focus:text-white focus:bg-gray-6 focus:outline-none has-[:focus]:border-orange has-[:focus]:border-l-2 has-[:focus]:pl-[22px] has-[:focus]:text-gray-11 has-[:focus]:bg-gray-6 has-[:focus]:outline-none\">4. at my work, i started off as a JS/TS dev developing frontend features and backend REST API's to a wider tech stack where i develop full stack apps, write ETL pipelines,perform data migration and use golang to do build out my side projects.</div><div tabindex=\"0\" number=\"4\" class=\"group text-wrap px-4 flex h-auto items-center text-gray-500 hover:bg-terminal-hover hover:border-primary-400 hover:border-l-2 active:border-primary-400  active:border-l-2 active:pl-[22px] active:text-white active:bg-gray-6 focus:border-primary-400 focus:border-l-2 focus:pl-[22px] focus:text-white focus:bg-gray-6 focus:outline-none has-[:focus]:border-orange has-[:focus]:border-l-2 has-[:focus]:pl-[22px] has-[:focus]:text-gray-11 has-[:focus]:bg-gray-6 has-[:focus]:outline-none\">5. outside of nerding out on tech stuff, i love basketball and american football. </div></pre></div><span id=\"currentText\" hx-swap-oob=\"true\" class=\"bg-secondary-500 text-white px-2 text-sm\">readme.txt</span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func About() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = AboutWithoutLayoutComponent().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
