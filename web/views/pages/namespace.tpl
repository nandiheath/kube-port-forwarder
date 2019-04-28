{{ define "content" }}
  <table class="table">
    <thead>
      <tr>
        <th scope="col">Service</th>
        <th scope="col">Type</th>
        <th scope="col">Status</th>
        <th scope="col">Ports</th>
        <th scope="col"></th>
      </tr>
    </thead>
    <tbody>
      {{range $val := .services }}
        <tr>
            <td>{{$val.Name}}</td>
            <td>{{$val.Type}}</td>
            <td>
              {{ if $val.Forwarded }}
                <h3><span class="badge badge-{{ if eq $val.Status "Failed" }}danger{{ else }}success{{end}}">{{ $val.Status }}</span></h3>
              {{ else }}
                <h3><span class="badge badge-warning">Idle</span></h3>
              {{ end }}
            </td>
            <td>
              {{ if and $val.Forwarded  (ne $val.Status "Failed") }}
                <span class="badge badge-success">
                  {{ $val.ForwardedFrom }}
                </span>
                ->
                <span class="badge badge-info">
                  {{ $val.ForwardedTo }}
                </span>
              {{ else }}
                {{ range $port := $val.Ports }}
                  <form method="POST">
                    <div class="from-group row" style="margin-bottom: 5px;">
                      <input type="hidden" name="service" value={{$val.Name}}>
                      <input type="hidden" name="from_port" value={{$port.Port}}>
                      <label class="col-sm-2 col-form-label">{{$port.Port}}</label>
                      <div class="col-sm-4">
                        <input class="form-control " type="text" name="to_port">
                      </div>
                      <div class="col-sm-4">
                        <input class="btn btn-info" type="submit" value="forward">
                      </div>
                    </div>
                  </form>
                {{ end }}
              {{ end }}
            </td>

        </tr>
      {{end}}
    </tbody>
  </table>

{{ end }}
