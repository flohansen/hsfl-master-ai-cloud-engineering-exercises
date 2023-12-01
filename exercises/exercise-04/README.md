# Cloud Engineering: Aufgabenblatt 4

Sie haben bisher mithilfe eines selbst implementierten HTTP-Routers
entsprechende (Micro-) Services für Ihr Produkt umgesetzt. Außerdem haben Sie
automatisierte Tests und CI-Pipelines entwickelt, um eine gewisse Qualität Ihrer
Dienste zu gewährleisten. In dieser Übung werden Sie Ihre Pipelines zu
CI/CD-Pipelines ausbauen, ein Frontend-Service für Ihre Applikation entwickeln
und Ihre Dienste mithilfe eines API-Gateways bzw. Reverse-Proxy anbinden.
Abschließend containerisieren Sie Ihre Dienste und entwerfen eine Konfiguration,
um diese zu starten.

## Aufgabe 1: Webserver & SSR

1. Entwickeln Sie einen Webserver in Go und verwenden Sie Go-Templates
   (html/template) für Server-Side- Rendering (SSR). Implementieren Sie
   entsprechende Routen für das Rendering von HTML-Seiten.

2. Stellen Sie statische Dateien (z.B. JavaScript, CSS, ...) über den Pfad
   `/static` bereit. Binden Sie entsprechende UI-Logik in Ihrem Template ein.

   _💡 Für JavaScript verwenden Sie script-Tags und für Styling link- Tags,
   die Resourcen über `/static` clientseitig laden._

3. Entwickeln Sie eine einfache Benutzeroberfläche für Ihre Anwendung und
   stellen Sie die wichtigsten Daten Ihrer Businesslogik dar. Beispiel: Angenommen
   es handelt sich um einen Online-Shop. Dann soll hier z.B. eine Produktliste
   gerendert werden.

4. Entwickeln Sie interaktive Elemente mithilfe von TypeScript und binden Sie
   diese mithilfe Ihres statischen File- Servers aus 1.2 ein. Hinweis: Um
   JavaScript clientseitig zu laden, verwenden Sie entsprechende script-Tags im
   Template.

5. Verzichten Sie auf eine Authentifizierung des Benutzers. Gehen Sie davon aus,
   dass Ihr Produkt (vorerst) durch anonyme Nutzer bedient werden kann.

6. Erstellen Sie einen Testdatensatz und verwenden Sie diesen, um das Rendering
   zu testen.
   
   _💡 Sie kön- nen vorerst manuelle Tests fahren. Ergänzen aber
   auf jeden Fall im weiteren Verlauf dieser Veranstaltung entsprechende
   Integrationstests_

7. Anforderungen an Design und User-Experience werden ausdrücklich nicht
   gestellt.

8. Aktualisieren Sie Ihr Architekturbild.

## Aufgabe 2: API-Gateway & Reverse Proxy

1. Entwickeln Sie ein Einstiegspunkt für Ihre Cloudanwendung. Üblicherweise
   werden hierfür sogenannte API- Gateways oder Reverse-Proxies verwendet.
   Erstellen Sie einen neuen Dienst und implementieren Sie einen Reverse-Proxy für
   Ihre Architektur.

2. Der Proxy soll eine Konfiguration verwenden, um eine Verbindung zwischen
   eingehender Requests und Ihren Microservices zu schaffen. Beispiele:

   * `/` → `http://webserver:port`
   * `/static` → `http://assetserver:port/static`
   * `/api/v1/auth` → `http://authserver:port/api/v1/auth`

3. Der Proxy soll nur die Aufgabe haben, Anfragen an entsprechende Dienste
   weiterzuleiten.

4. Definieren Sie eine Struktur für den Reverse-Proxy.

5. Implementieren Sie das http.Handler-Interface für den Reverse-Proxy, damit
   dieser als Parameter in http.ListenAndServe übergeben werden kann.

6. Lesen Sie mithilfe des eingehenden Requests die Quelladresse bzw. den
   Quellpfad aus und finden Sie die entsprechende Zieladresse anhand der
   Reverse-Proxy-Konfiguration aus 2.2.

7. Ändern Sie den Request, damit dieser nun auf die Zieladresse zeigt.
   Überschreiben Sie entsprechende Felder wie Host, URL, Scheme und RequestURI.

8. Senden Sie den HTTP-Request zum Zieldienst. Verwenden Sie hierfür ein
   Interface für HTTP-Clients, damit dieses in Tests durch entsprechende Mocks
   ersetzt und damit das Testen erleichtert werden kann.

9. Nachdem Sie eine Antwort von Ihren Zieldienst bekommen haben, sollten Sie
   alle Response-Header in die Antwort des Reverse-Proxy an den anfragenden Client
   kopiert werden.

10. Kopieren Sie nun den Datenstrom von der Antwort Ihres Zieles in die Antwort
   Ihres Reverse-Proxy. Sie können hierfür die Funktion `io.Copy` verwenden.

11. (Optional): Anstelle von 2.7, 2.8, 2.9 und 2.10 können Sie auch das Modul
   `net/http/httputil` aus der Standardbibliothek verwenden. Achten Sie hierbei auf
   die Testbarkeit Ihrer Implementation und verwenden Sie entsprechende Interfaces.

12. Entwickeln Sie Tests für Ihren Reverse-Proxy.

13. Aktualisieren Sie Ihr Architekturbild.

## Aufgabe 3: Containerisierung & Continuous-Delivery

1. Containerisieren Sie nun Ihre (Micro-) Services. Erstellen Sie hierfür
   entsprechende Dockerfiles in den Ordnern Ihrer Services.

2. Erweitern Sie Ihre Pipelines um einen Build-Schritt, sodass diese nun auch
   Continuous-Delivery erfüllen.

3. Erstellen Sie für Ihre Services eigene Repositories auf [Dockerhub](https://hub.docker.com) und
   pushen Sie die gebauten Docker- Images über Ihre Pipeline in diese.

4. Erstellen Sie eine Docker-Compose-Konfiguration, um Ihre Dienste miteinander
   verknüpft lokal starten und tes- ten zu können.