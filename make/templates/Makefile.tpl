SHELL=/bin/bash -o pipefail

.make/maak.mk: maak.yaml
	@maak init makefile

-include custom.mk

PROJECT := {{Name}}
COMPONENTS :={{#Components}} {{Name}}{{/Components}}
CONTAINERS :={{#Containers}} {{Name}}{{/Containers}}

################################################################################
# Generation of the component rules
#
# In addition to the main rules above, each architectural component at least
# have the following standard rules, which are defined below:
#
# - <component>.images: builds the components' docker images, rebuilding source code if needed
# - <component>.down:  shuts down the component
# - <component>.up:    wakes up the component, making sure the last version runs
# - <component>.on:    wakes up the component, taking the last known image
# - <component>.clean: cleans everything, only useful for rebuilding from scratch

{{#Components}}

{{Name}}.build: .make/maak.mk {{Name}}.containers
{{Name}}.clean: .make/maak.mk {{#Containers}}{{Name}}.clean {{/Containers}}
{{Name}}.containers: .make/maak.mk {{#Containers}}{{Name}}.image {{/Containers}}

{{/Components}}


################################################################################
# Generation of the container rules
#

# In addition to the compoennt rules above, each of the component's containers
# have the following rules, which are defined below:
#
# - <container>.image: builds the container image, rebuilding source code if needed
# - <container>.down:  shuts down the component
# - <container>.up:    wakes up the component, making sure the last version runs
# - <container>.on:    wakes up the component, taking the last known image
# - <container>.clean: cleans everything, only useful for rebuilding from scratch

{{#Containers}}
{{Name}}.image: .make/{{Name}}.built
{{Name}}.clean:
	rm -f .make/{{Name}}.*

{{/Containers}}

{{#Containers}}
.make/{{Name}}.built: {{#ComponentDependencies}}.make/{{.}}.built {{/ComponentDependencies}} $(shell maak deps {{Name}})
	@docker build -t {{ImageName}} {{Context}} -f {{Dockerfile}}
	@touch .make/{{Name}}.built

{{/Containers}}

