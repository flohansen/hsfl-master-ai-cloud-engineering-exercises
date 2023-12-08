# Cloud Engineering: Aufgabenblatt 8

## Aufgabe 1: Anpassen des Lasttests

In der letzten Aufgabe haben Sie einen Lasttest implementiert, welche mit einer
linear ansteigenden Anzahl von Requests einen HTTP-Server testet. In dieser
Aufgabe sollen Sie diesen weiterentwickeln.

1. Räumen Sie die bisherige Implementation auf falls erforderlich.
2. Es soll nun eine Spezifikationsliste für die Rampe eingelesen werden. Einträge
   sollen die Dauer und die Ziel RPS (Requests per Second) festlegen.
3. Passen Sie Ihren Code soweit an, sodass mithilfe der Spezifikationsliste die
   Kurve für einen Lasttest definiert werden kann.
4. Zwischen den Einträgen der Spezifikationsliste soll die RPS linear
   interpoliert werden.
5. Testen Sie Ihre Implementation

## Aufgabe 2: Implementierung von gRPC in einem bestehenden Microservices-System

Sie haben bereits Microservices in Go entwickelt und in einem Kubernetes-Cluster
bereitgestellt. Jetzt sollen Sie gRPC als Kommunikationsprotokoll zwischen den
Microservices implementieren, um die Kommunikation effizienter und skalierbarer
zu gestalten.

1. Wählen Sie einen oder mehrere Ihrer vorhandenen Microservices aus.
2. Implementieren Sie gRPC-Protobuf-Datenstrukturen und -Dienste für die
   Kommunikation zwischen den ausgewählten Microservices.
3. Aktualisieren Sie den ausgewählten Microservice, um gRPC als
   Kommunikationsprotokoll zu verwenden.
4. Stellen Sie sicher, dass die anderen Microservices weiterhin über ihre bisherigen
   Kommunikationsmechanismen erreichbar sind, um eine schrittweise Migration zu
   ermöglichen.
5. Testen Sie die gRPC-Implementierung, um sicherzustellen, dass die Kommunikation
   zwischen den Microservices ordnungsgemäß funktioniert.

## Aufgabe 3: Request-Coalescing implementieren

Um die Netzwerkbelastung zu reduzieren und die Effizienz der Anwendung zu
verbessern, sollen Sie Request-Coalescing in einem oder mehreren Ihrer
Microservices implementieren.

1. Wählen Sie einen Microservice aus, der häufige Anfragen von Clients erhält.
2. Implementieren Sie eine Request-Coalescing-Logik, die mehrere Anfragen zu einer
   zusammenfasst, um sie als eine einzige Anfrage zu verarbeiten.
3. Stellen Sie sicher, dass die Coalescing-Logik sinnvoll arbeitet und keine
   Datenverluste oder Inkonsistenzen verursacht.
4. Implementieren Sie eine Testumgebung, um die Auswirkungen der
   Request-Coalescing-Implementierung zu überprüfen. Nutzen Sie hierfür Ihren
   Lasttest, ein Kubernetes-Cluster und entsprechendes Monitoring.
5. Dokumentieren Sie die implementierte Request-Coalescing-Logik und ihre
   Auswirkungen auf die Leistung und Effizienz der betroffenen Microservices.

Hinweis: Bei der Implementierung von gRPC und Request-Coalescing ist es
wichtig, die Auswirkungen auf die Performance, Skalierbarkeit und
Zuverlässigkeit Ihrer Microservices zu überwachen und zu bewerten.
