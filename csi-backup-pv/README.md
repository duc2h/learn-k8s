##### Note
In order to backup persistent volume before deleting pv, we can use [volume snapshot](https://kubernetes.io/docs/concepts/storage/volume-snapshots/)
, [persistent disk csi on gcp](https://github.com/kubernetes-sigs/gcp-compute-persistent-disk-csi-driver/blob/master/docs/kubernetes/user-guides/snapshots.md). 
https://cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/volume-snapshots