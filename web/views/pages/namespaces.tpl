{{ define "content" }}
  {{range $val := .namespaces }}
    <li>
     <a href="/namespace/{{$val.Name}}">{{$val.Name}}</a>
     </li>
  {{end}}
{{ end }}