// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base(layouts []SeatLayout) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" data-theme=\"dark\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Seats Organizer</title><link rel=\"stylesheet\" href=\"/style.css\"><script src=\"https://unpkg.com/htmx.org@1.9.10\" integrity=\"sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC\" crossorigin=\"anonymous\"></script></head><body hx-boost=\"true\"><div class=\"organize-seats\"><form id=\"organize\" hx-encoding=\"multipart/form-data\" hx-post=\"/organize\" hx-target=\"#organize-result\"><div class=\"form-line\">Choose seat layout:\r <select name=\"layout\" hx-get=\"/preview\" hx-target=\"#seat-layout\" hx-trigger=\"load, change\"><option value=\"\" selected disabled hidden></option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, l := range layouts {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(l.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layout/main.templ`, Line: 27, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(l.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layout/main.templ`, Line: 27, Col: 61}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></div><div class=\"form-line\">List of passengers:\r <input type=\"file\" name=\"passengers\"></div><div class=\"form-line\"><button type=\"submit\">Submit</button></div></form><div>Selected seat layout preview:\r<div id=\"seat-layout\"></div></div></div><div id=\"organize-result\"></div><script type=\"text/javascript\">\r\n            function highlightSeat(element) {\r\n                const seatID = element.getAttribute(\"seat-id\")\r\n                const seats = document.getElementsByClassName(seatID)\r\n                for (let seat of seats) {\r\n                    seat.classList.add(\"highlight\")\r\n                }\r\n            }\r\n            function unhighlightSeat(element) {\r\n                const seatID = element.getAttribute(\"seat-id\")\r\n                const seats = document.getElementsByClassName(seatID)\r\n                for (let seat of seats) {\r\n                    seat.classList.remove(\"highlight\")\r\n                }\r\n            }\r\n        </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
