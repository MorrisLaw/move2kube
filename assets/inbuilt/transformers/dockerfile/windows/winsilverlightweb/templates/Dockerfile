#   Copyright IBM Corporation 2021
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

# Build stage
FROM hugodahl/silverlightbuild:sl5 As builder
WORKDIR /app
COPY . .
RUN msbuild {{ .AppName }}.sln
RUN mkdir output
RUN cp {{ .AppName }}.Web/*.html output/
RUN cp {{ .AppName }}.Web/*.aspx output/
RUN cp {{ .AppName }}.Web/*.js output/
RUN cp {{ .AppName }}.Web/Web.config output/
RUN cp -r {{ .AppName }}.Web/bin output/
RUN cp -r {{ .AppName }}.Web/ClientBin output/

# Run stage
FROM mcr.microsoft.com/windows/servercore/iis
RUN powershell -NoProfile -Command Remove-Item -Recurse C:\inetpub\wwwroot\*
{{- range .Ports}}
EXPOSE {{.}}
{{- end}}
WORKDIR /inetpub/wwwroot
COPY --from=builder /app/output/ .