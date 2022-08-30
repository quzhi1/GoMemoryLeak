# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_resource', 'helm_resource')

compile_opt = 'GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 '

# Compile example application
local_resource(
  'memory-leak-compile',
  compile_opt + 'go build -o bin/memory-leak main.go',
  deps=['main.go'],
  ignore=['bin', 'helm', 'Dockerfile', 'Tiltfile', 'README.md', 'LICENSE', '.gitignore'],
  labels="memory-leak",
)

# Build example docker image
docker_build_with_restart(
  'memory-leak-image',
  '.',
  entrypoint=['/opt/app/bin/memory-leak'],
  dockerfile='Dockerfile',
  only=[
    './bin',
  ],
  live_update=[
    sync('./bin', '/opt/app/bin'),
  ],
)

# Install example helm chart
helm_resource(
  'memory-leak-service',
  'helm',
  image_deps=['memory-leak-image'],
  image_keys=[('image.repository', 'image.tag')],
  port_forwards=['8080:8080'],
  labels="memory-leak",
)
