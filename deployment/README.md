### Note: Some command

#### 1. Check history of deployment
* kubectl rollout history deploy `<deployment-name>`

#### 2. Rollback previous deployment version
* kubectl rollout undo deployment `<deployment-name>` --to-revision=`<revision-number>`