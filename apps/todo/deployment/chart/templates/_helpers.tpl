{{ define "todo.common.labels" -}}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ .Chart.Name }}
{{- end }}

{{ define "todo.common.selectors" -}}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}


{{ define "todo.server.name" -}}
{{ .Release.Name }}-server
{{- end }}

{{ define "todo.server.labels" -}}
{{ template "todo.common.labels" . }}
app.kubernetes.io/component: server
{{- end }}

{{ define "todo.server.selectors" -}}
{{ template "todo.common.selectors" . }}
app.kubernetes.io/component: server
{{- end }}

{{ define "todo.server.defaultImage" -}}
{{ .Values.docker_registry }}/{{ .Release.Name }}-server:{{ .Values.version }}
{{- end }}

{{ define "todo.middleware.name" -}}
{{ .Release.Name }}-strip-prefix
{{- end }}