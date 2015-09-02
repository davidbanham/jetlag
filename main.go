package main

import (
  "os"
  "os/exec"
  "log"
  "html/template"
  "net/http"
  "github.com/davidbanham/required_env"
  "fmt"
  "strings"
)

func main() {
  required_env.Ensure(map[string]string{"ROUTER": "", "PORT": "3000", "COUNTRIES": "US,AU"})

  countries := strings.Split(os.Getenv("COUNTRIES"), ",")
  allowed_countries := map[string]bool{"": true}

  for _, key := range countries {
    allowed_countries[key] = true
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    target := r.FormValue("country")

    if !allowed_countries[target] {
      w.WriteHeader(403)
      w.Write([]byte("That country is not permitted!"))
      return
    }

    if target != "" {
      fmt.Println(target)
      err := fly(target, w)
      if err != nil {
        w.Write([]byte(err.Error()))
        return
      }
    }
    t, err := template.New("foo").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Jetlag</title>
  </head>
  <body>
    {{if .LastAction}}
      <h4>Just flew to {{ .LastAction }}</h4>
    {{end}}
    {{range .Countries}}
      <div>
        <form action="/" method="POST">
          <input type="hidden" name="country" value="{{ . }}"></input>
          <button type="submit">{{ . }}</button>
        </form>
      </div>
    {{end}}
  </body>
</html>
`)
    if err != nil {
      w.Write([]byte(err.Error()))
      return
    }
    data := struct {
      Countries []string
      LastAction string
    }{
      Countries: countries,
      LastAction: target,
    }
    err = t.Execute(w, data)
    if err != nil {
      w.Write([]byte(err.Error()))
      return
    }
  })
  log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func fly(target string, w http.ResponseWriter) error {

  router := os.Getenv("ROUTER")
  
  scriptName := target + "-vpn.sh"

  cmd := exec.Command("ssh", router, "killall openvpn")
  cmd.Stdout = w
  cmd.Stderr = w
  err := cmd.Run()
  if err != nil {
    if e := err.Error(); e != "exit status 1" {
      return err
    }
  }
  cmd2 := exec.Command("ssh", router, "/root/"+scriptName)
  cmd2.Stdout = w
  cmd2.Stderr = w
  err = cmd2.Run()
  if err != nil {
    return err
  }
  return nil
}
