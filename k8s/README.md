   1 - Install metrics-server https://github.com/kubernetes-sigs/metrics-server
   
    kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/download/v0.3.7/components.yaml    

   2 - Update  metrics-server deployment
   
     containers:
        - name: metrics-server
          image: 'k8s.gcr.io/metrics-server/metrics-server:v0.3.7'
          args:
            - '--cert-dir=/tmp'
            - '--secure-port=4443'
            - '--metric-resolution=30s'
            - '--kubelet-preferred-address-types=InternalIP'
            - '--kubelet-insecure-tls'
   
   3 - Create `ClusterRole`, `ServiceAccount`, `ClusterRoleBinding` (manifest.yml)
  
   4 - `make docker`