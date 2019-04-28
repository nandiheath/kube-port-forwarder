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
            <td>{{$val.Name}}</td>
            <td>{{$val.Type}}</td>
            <td>
              {{ range $port := $val.Ports }}
                <form method="POST">
                  <p>{{$port.Port}} -> {{$port.TargetPort}} {{$port.Forwarded}}</p>
                  <input type="hidden" name="service" value={{$val.Name}}>
                  <input type="hidden" name="from_port" value={{$port.Port}}>
                  <input type="text" name="to_port">
                  <input type="submit" value="submit">
                </form>
              {{ end }}
            </td>
        </tr>
      {{end}}
    </tbody>
  </table>

{{ end }}
