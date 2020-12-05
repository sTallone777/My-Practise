package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	// mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	// log.Fatal(http.ListenAndServe("localhost:8000", mux))
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/remove", db.remove)
	http.ListenAndServe("localhost:8000", nil)
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	tmpStr := `
	<h1>item list</h1>
	<table>
	<tr style='text-align: left'>
	  <th>item</th>
	  <th>price</th>
	</tr>
	{{range $key, $value := .db}}
	<tr>
	  <td>{{$key}}</td>
	  <td>{{$value}}</a></td>
	</tr>
	{{end}}
	</table>`
	itmeList := template.Must(template.New("itmelist").Parse(tmpStr))
	//itmeList.Execute(os.Stdout, db)
	itmeList.Execute(w, db)
	// for item, price := range db {
	// 	fmt.Fprintf(w, "%s: %s\n", item, price)
	// }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	newPrice := req.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	f, err := strconv.ParseFloat(newPrice, 32)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	db[item] = dollars(f)
	fmt.Fprintf(w, "%s\n", "update "+item+" successed.")
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	newPrice := req.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		f, err := strconv.ParseFloat(newPrice, 32)
		if err != nil {
			fmt.Fprintf(w, "%s\n", err.Error())
			return
		}

		db[item] = dollars(f)
		fmt.Fprintf(w, "%s\n", "add "+item+" successed.")
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "existed item: %q\n", item)
		return
	}

}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "%s\n", "remove "+item+" successed.")
}
