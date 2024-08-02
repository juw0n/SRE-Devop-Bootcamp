#!/bin/bash

# Apply the first set of labels
kubectl label node sre-project type=dependent-services-node
kubectl label node sre-project-m02 type=database-node
kubectl label node sre-project-m03 type=application-node

# Apply the second set of labels
kubectl label node sre-project service=dependent-services-node
kubectl label node sre-project-m02 database=database-node
kubectl label node sre-project-m03 application=application-node

# Make the script executable ==> chmod +x nodes-label.sh
# Run the script: ==> ./nodes-label.sh

# PS: you can also taint the nodes. but this gave me some issues 
# kubectl taint nodes sre-project service=dependent-services-node:NoSchedule
# kubectl taint nodes sre-project-m02 database=database-node:NoSchedule
# kubectl taint nodes sre-project-m03 application=application-node:NoSchedule