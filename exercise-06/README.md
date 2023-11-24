# Cloud Engineering: Aufgabenblatt 6

## Aufgabe 1: Deployment einer App auf einem (einfachen) Kubernetes-Cluster

Diese Übung dient als Lernhilfe für den ersten Kontakt mit Kubernetes (K8S). 

1. Installieren Sie [kubectl](https://kubernetes.io/docs/tasks/tools/) und
   [minikube](https://minikube.sigs.k8s.io/docs/start/) auf Ihrer
   Entwicklungsmaschine.
2. Erstellen Sie mithilfe von minikube ein lokales Cluster mit einem einzigen
   Node und starten Sie es.
3. Erstellen Sie ein Deployment für einen Ihrer Services mithilfe von `kubectl
   create deployment`.
4. Erstellen Sie einen Service für Ihre Pods aus 3. mithilfe von `kubectl
   expose`.
5. Ändern Sie den Typ des Services auf `NodePort` und `LoadBalancer`. Was sind
   die Unterschiede?
6. Starten Sie das Dashboard von minikube mit `minikube dashboard` und
   validieren Sie Ihr Deployment.
7. Testen Sie die Verfügbarkeit Ihrer Anwendung von außerhalb des
   Kubernetes-Netzwerkes (z.B. über die Loopback-IP).
8. Skalieren Sie Ihre Anwendung mithilfe von `kubectl scale` auf 10 Replicas.
9. Validieren Sie wie vorher die Skalierung über das Dashboard und prüfen Sie
   die Verfügbarkeit.

## Aufgabe 2: Erstellen eines Kubernetes-Clusters

> 💡 **Information:**
> [Hier](SETUP_CLUSTER.md) finden Sie eine Hilfestellung für die Einrichtung
> eines Kubernetes-Clusters unter Ubuntu Server 20.04. Alternativ können Sie
> die [offizielle
> Dokumentation](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/)
> von kubeadm zum Erstellen eines Kubernetes-Clusters als Unterstützung lesen.

Nachdem Sie in Aufgabe 1 den ersten Kontakt mit einem Kubernetes-Cluster über
minikube hatten, sollen Sie nun ein eigenes Cluster mit mehreren Nodes
erzeugen. In der Aufgabe wird das Vorgehen mit virtuellen Maschinen
beschrieben. Wenn Sie Zugriff auf entsprechende Hardware besitzen, können Sie
natürlich auch mit dieser anstelle der virtuellen Maschinen arbeiten.

1. Installieren Sie [VirtualBox](https://www.virtualbox.org/) auf Ihrem
   Host-System und laden Sie sich das [Abbild von Ubuntu Server
   20.04 LTS](https://ubuntu.com/download/server#downloads)
   als ISO-Datei herunter. Achten Sie darauf, dass Sie die Version 20.04
   verwenden, da die Installation von Kubernetes bei neueren Versionen wesentlich
   komplizierter wird.
2. Richten Sie beliebig viele (virtuelle) Maschinen ein und installieren Sie
   Ubuntu 20.04 LTS als Betriebssystem. Eine Maschine wird als Control-Plane
   bzw. Master-Node, die restlichen als Worker-Nodes eingerichtet. Achten Sie bei
   der Verwendung von virtuellen Maschinen auf die zur Verfügung stehenden
   Resourcen des Host-Systems. Der Master-Node sollte mindestens 2 CPU und 2GB
   Arbeitsspeicher besitzen während Worker-Nodes mindestens 1 CPU und 1GB
   Arbeitsspeicher besitzen sollten.

   **Beispiel für ein Cluster mit 5 Maschinen:**

   * 1 Master-Node (Control-Plane)
   * 4 Worker-Nodes

3. Damit Ihr Kubernetes-Cluster Pods (Container) verwalten kann, müssen Sie
   zuerst eine Container-Runtime auf jedem Node installieren. Sie können hierfür
   [containerd](https://github.com/containerd/containerd/blob/main/docs/getting-started.md),
   [CRI-O](https://cri-o.io/) oder [Docker Engine (mit
   cri-dockerd)](https://docs.docker.com/engine/install/ubuntu/)
   verwenden.
4. Installieren Sie [kubectl](https://kubernetes.io/docs/reference/kubectl/kubectl/), [kubelet](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/) und
   [kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/)
   auf jedem der Nodes.
5. Initialisieren Sie einen Ihrer Nodes als Control-Plane mithilfe von kubeadm.
6. Installieren Sie für die Kommunikation zwischen Container ein
   Container-Network-Interface (CNI) wie
   [calico](https://www.tigera.io/project-calico/) auf dem Control-Plane.
7. Folgen Sie den Anweisungen zum Registrieren Ihrer Worker-Nodes.

## Aufgabe 3: Deployment Ihrer Anwendung

> 💡 **Information:**
> [Dieses
> Beispiel](https://kubernetes.io/docs/tasks/access-application-cluster/connecting-frontend-backend/)
> zeigt, wie verschiedene Deployments auf einem Cluster miteinander
> kommunizieren können. Dies ist hilfreich, damit z.B. Frontend-Services mit
> entsprechenden Backend-Services sprechen können.

1. Erstellen Sie Manifests zum Deployen Ihrer Microservices (kind: `Deployment`).
2. Erstellen Sie Manifests für das Veröffentlichen Ihrer Microservices mit Kubernetes-Services (kind: `Service`).
3. Deployen Sie Ihre Microservices auf Ihr in Aufgabe 2 erstelltes Cluster.
4. Testen Sie Ihr Deployment, indem Sie auf Ihre Web-Anwendung über den Browser außerhalb des Cluster-Netzwerkes zugreifen (z.B. vom Host-System).
