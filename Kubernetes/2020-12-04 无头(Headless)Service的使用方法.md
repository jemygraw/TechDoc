
一般来讲，Kubernetes里面的Service都有一个同名的Endpoints对象来定义具体的访问端点，这些端点是通过Service中定义的Selector根据Pod的标签选择器来查找指定的一组Pod作为端点的。
但是有些情况下，我们不要能够在Kubernetes的体系之内访问到外部的服务，这种情况下可以定义个静态的Endpoints对象，把相关的实际地址作为端点，然后还是通过Kubernetes的Ingress来访问。

比如后端有一个服务的Nginx配置如下：

```
server {
	listen 80;
	server_name bbs-static.duokexuetang.com;
	root /data0/www/bbs-go/;
}

server {
        listen 80;
        server_name bbs.duokexuetang.com;
        location / {
                proxy_set_header Host $http_host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_pass http://localhost:8081;
        }
}

server {
        listen 80;
        server_name bbs-api.duokexuetang.com;
        location / {
                proxy_set_header Host $http_host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_pass http://localhost:8082;
        }
}
```

上面定义了三个服务，一个静态文件，一个前端服务，一个后端接口服务，我们可以在一个Ingress里面把这三个服务都转发到一个Service中。

```
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: kong
  labels:
    app: cloudbbs
  name: dev-cloudbbs
  namespace: devops
spec:
  rules:
  - host: bbs.duokexuetang.com
    http:
      paths:
      - backend:
          serviceName: cloudbbs-exp
          servicePort: 8080
        path: /
  - host: bbs-api.duokexuetang.com
    http:
      paths:
      - backend:
          serviceName: cloudbbs-exp
          servicePort: 8080
        path: /  
  - host: bbs-static.duokexuetang.com
    http:
      paths:
      - backend:
          serviceName: cloudbbs-exp
          servicePort: 8080
        path: /  
status:
  loadBalancer:
    ingress:
    - ip: 10.23.0.24
```

这里我们再定义一个Service对象：

```
apiVersion: v1
kind: Service
metadata:
  labels:
    app: cloudbbs-exp
  name: cloudbbs-exp
  namespace: devops
spec:
  ports:
  - name: http
    port: 8080
    protocol: TCP
    targetPort: 80
  type: ClusterIP
```

需要注意的是，这里的 targetPort 必须是目标服务的实际监听端口，因为要在Endpoints中使用。最后我们再定一个Endpoints应用就可以了。

```
apiVersion: v1
kind: Endpoints
metadata:
  labels:
    app: cloudbbs
  name: cloudbbs-exp
  namespace: devops
subsets:
- addresses:
  - ip: 10.8.1.127
  ports:
  - name: http
    port: 80
    protocol: TCP
```

以上就是无头（Headless）Service的使用方法。
