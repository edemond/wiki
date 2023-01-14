dev_appserver.py cmd\wiki-appengine\app.yaml ^
  --support_datastore_emulator=True ^
  --env_var WIKI_STATIC_DIR=%~dp0static ^
  --env_var WIKI_TEMPLATE_DIR=%~dp0templates