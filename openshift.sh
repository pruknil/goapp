cd /home/ubuntu/Desktop/OpenShift
./oc cluster up --routing-suffix=127.0.0.1.xip.io --public-hostname=localhost


oc login
oc whoami -t
docker login -u developer -p YvIOVg7U9oZQHBjT--a0-_-jadaYpvVWy0St0Vo78nc 172.30.1.1:5000/myproject

docker push "172.30.1.1:5000/myproject/goapp:latest"
oc get all -o name
oc new-app --search imagestream.image.openshift.io/goapp
oc new-app 172.30.1.1:5000/myproject/goapp:latest --name goapp
oc expose svc/goapp

oc delete deploymentconfig goapp
oc delete imagestream goapp
oc create -f goapp.yaml


ubuntu@ubuntu-H370HD3:~/Desktop/OpenShift$ oc delete all --all