# Cloud Engineering: Aufgabenblatt 1

## Aufgabe 1: Research

1. Betreiben Sie Recherche über Monolithen und Microservices. Versuchen Sie
   diese beiden Softwarearchitekturen zu verstehen und Vergleichen Sie sie.
2. Das Arbeiten mit Microservices schließt die Verwendung Monolithen nicht
   gänzlich aus. Recherchieren Sie über den Monolith-First-Ansatz und dessen
   Vorteile.

## Aufgabe 2: Architecture Design

1. Überlegen Sie sich eine Architektur für Ihr Projekt und stellen Sie
   entsprechende Komponenten und deren Beziehungen in einem Diagramm dar.
   Arbeiten Sie dabei nach dem Monolith-First-Ansatz.
2. Überführen Sie den Monolithen aus 1. nun in eine
   Microservice-Architektur. Stellen Sie sicher, dass jeder Microservice maximal
   eine Aufgabe innerhalb der Architektur übernimmt. Erstellen Sie auch hier ein
   entsprechendes Architekturdiagramm.
3. Dokumentieren Sie beide Architekturen. Gehen Sie vor allem auf die
   Komponenten (Microservices, Klassen, etc.), deren Aufgaben und Beziehungen
   ein.

## Aufgabe 3: Project initialization

1. Erstellen Sie einen Fork von diesem
   [Repository](https://github.com/flohansen/hsfl-master-ai-cloud-engineering) und
   bearbeiten Sie die `README.md`.
2. Fügen Sie alle Gruppenmitglieder zum GitHub-Projekt hinzu und geben Sie Ihnen
   die Rolle Admin oder `Write`.
3. Erstellen Sie entsprechende Unterprojekte für Ihre Microservices und fügen
   Sie jeweils eine `README.md` hinzu, welche die Komponente erläutert.
   Nutzen Sie dafür Ihre Dokumentation aus 2.3.
4. Synchronisieren Sie `remote` mit `local` mittels eines
   `push` Befehls. Achten Sie auf die Konventionen für Commit-Messages in
   den Hinweisen.

## Hinweise

### Projektbeispiele
Hier sind ein paar Beispielprojekte, an denen Sie sich orientieren können.

1. Multi-Vendor Marketplace: Eine Plattform, auf der verschiedene Verkäufer
   Produkte beliebiger Kategorien zum Verkauf anbieten.
2. Machine-Learning Platform: Eine Plattform, auf der Datensätze abgelegt und
   bearbeitet werden können, um anschließend Machine-Learning Modelle zu
   trainieren.
3. Airline Tracker: Ein System, welches die Positionen von Luftfahrzeugen
   verschiedener Airlines berechnet und darstellt.
4. Security Monitoring System: Ein Tool, welches verschiedene Rechner als
   Knotenpunkte einbindet und effizient auf IT-Sicherheit mittels Angriffspfaden
   prüft.
5. Search Engine: Ein Webcrawler, der systematisch das Internet durchsucht und
   Ressourcen als Suchmaschine verfügbar macht.

### Aufbau eines Git-Repository
Bitte achten Sie darauf, dass Sie nur Dateien und Verzeichnisse zu einem
Git-Projekt hinzufügen, die essentiell sind. Grundsätzlich gilt, dass wenn eine
entsprechende Datei durch den Source-Code erzeugen könnte (z.B. Executables,
Libraries, etc.), dann wird diese nicht in das Projekt committet.
Konfigurationen für Texteditoren bzw. IDEs sind zum Beispiel ebenfalls solche
Dateien, da der entsprechende Editor keine Abhängigkeit von dem Projekt sein
sollte. Auch Umgebungsvariablen oder Dateien, die Secrets wie Passwörter etc.
enthalten, sollten aus Sicherheitsgründen nicht in die Git-Historie aufgenommen
werden. GitHub selbst bietet eine große
[Liste](https://github.com/github/gitignore) für verschiedene
`.gitignore`-Dateien für verschiedene Use-Cases, um auszuhelfen.

### Commit-Nachrichten
Wir nutzen die Konventionen von [Conventional
Commits](https://www.conventionalcommits.org/) für
Commit-Nachrichten in Git. Grundsätzlich sollten die Nachrichten aussagekräftig
sein, um Änderungen schnell identifizieren zu können. Versuchen Sie pro Commit
nur ein Feature zu bearbeiten und relativ hochfrequent Änderungen zu committen.
Hier ein paar Beispiele.

    git commit -m "initial commit"
    git commit -m "docs: add README to project"
    git commit -m "docs(user-service): add README"
    git commit -m "feat(user-service): add authentication handler and implement OAuth2.0"
    git commit -m "refactor(delivery-service): rename interfaces"