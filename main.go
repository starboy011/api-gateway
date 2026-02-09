


package main
import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/starboy011/api-gateway/server"
)

func main() {

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("api-gateway"),
		newrelic.ConfigLicense("eu01xxc12dc4fe3729b337cc4130261eFFFFNRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	server.SetupBarberShopsServiceRoutes(router)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		txn := app.StartTransaction(req.URL.Path)
		defer txn.End()

		router.ServeHTTP(w, req)
	})

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
