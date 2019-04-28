{{ define "content" }}
  <table class="table">
    <thead>
      <tr>
        <th scope="col">Service</th>
        <th scope="col">Type</th>
        <th scope="col">Ports</th>
        <th scope="col">Status</th>
        <th scope="col"></th>
      </tr>
    </thead>
    <tbody>
      {{range $val := .services }}
        <tr>
          <td>{{check $val}}</td>
          <td>{{$val.Spec.Type}}</td>
          <td>
            {{ range $port := $val.Spec.Ports }}
              <p>{{$port.Port}}</p>
            {{ end }}
          </td>

            <a href="/namespace/{{$val.Name}}">{{$val.Name}}</a>
          </td>
          <td>
            <a href="/namespace/{{$val.Name}}">{{$val.Name}}</a>
          </td>
        </tr>
      {{end}}
    </tbody>
  </table>

{{ end }}
