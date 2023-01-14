mklink /J cmd\wiki-appengine\templates templates
gcloud app deploy cmd\wiki-appengine\app.yaml
rmdir cmd\wiki-appengine\templates