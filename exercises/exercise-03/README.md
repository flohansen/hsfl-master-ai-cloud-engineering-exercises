# Cloud Engineering: Aufgabenblatt 3

In der letzten Übung haben Sie Ihre ersten Microservices und automatisierte
Tests für deren Funktionalität entwickelt. In dieser Übung sollen Sie sich in
die Konzepte hinter Continuous-Integration und -Delivery/Deployment einarbeiten.
Sie erlernen Fähigkeiten zum Entwickeln von Pipelines, die bestimmte Prozesse
des Lebenszyklus' Ihres Softwareprodukts automatisieren, Ihre implementierten
Tests ausführen und damit eine entsprechende Qualität Ihrer Produkte
gewährleisten.

## Aufgabe 1: Continuous-Integration

1. Arbeiten Sie sich in Pipelines in der Softwareentwicklung ein. Versuchen Sie
   dabei die Prinzipien hinter Continuous-Integration (CI) und
   Continuous-Delivery/Deployment (CD) zu verstehen. Studieren Sie anschließend
   Pipelines und deren Rolle im CI/CD-Kontext.

2. Entwickeln Sie für jeden Ihrer Microservices eine CI/CD-Pipeline. Verwenden
   Sie GitHub-Workflows, um entsprechende Pipelines zu implementieren.

   _💡 Verwenden Sie für jeden Microservice eine eigene
   YAML-Konfigurationsdatei, um Workflows zu definieren._

3. Die Pipelines sollen automatisch ausgeführt werden, wenn sich der Quellcode
   von einem Microservice durch einen Git-Push im Repository ändert. Achten Sie
   darauf, dass lediglich die Pipelines von betroffenen Microservices ausgeführt
   werden.

   _💡 Nutzen Sie Path-Filter, um zwischen Commits in Ihrem Monorepo
   bezüglich der darin enthaltenen Microservices unterscheiden zu können._

4. Konzentrieren Sie sich beim Entwickeln der Pipelines vorerst auf
   Continuous-Integration. Die Pipelines sollen die Funktionalität Ihrer
   Microservices überprüfen, aber keine Builds etc. erstellen.

   _💡 Continuous-Delivery bzw. Deployment wird erst in späteren
   Aufgaben zum Einsatz kommen, wenn Containerization und -orchestration eingeführt
   wird._

## Aufgabe 2: GitHub-Badges

1. Fügen Sie in Ihren READMEs Badges hinzu, die den aktuellen Status der
   Workflows anzeigen.

   _💡 GitHub Docs enthält Beispiele hierfür._

2. (optional) Fügen Sie Badges zu den READMEs hinzu, welche die Testabdeckung
   Ihrer Services anzeigen. 

   _💡 Sie können Ihr Repository mit dem Dienst
   app.codecov.io synchronisieren und Coverage-Reports als Badges rendern lassen._

3. (optional) Fügen Sie Badges zu den READMEs hinzu, welche die Code-Qualität
   Ihrer Projekte anzeigt.

   _💡 Sie können den Dienst von goreportcard.com hierfür verwenden._