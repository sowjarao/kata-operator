# kata-operator

Deployment:

kubectl create -f deploy/service_account.yaml

kubectl create -f deploy/role.yaml

kubectl create -f deploy/role_binding.yaml

kubectl create -f deploy/crds/mygroup2.mydomain.com_katasets_crd.yaml

kubectl create -f deploy/operator.yaml

kubectl create -f deploy/crds/mygroup2.mydomain.com_v1alpha1_kataset_cr.yaml




