# Makefile
Hello:
	echo "hello"


PROJECT_NAME = ep

CONFIG := "./config/local.json"
define GetValueFromConfig
$(shell node -p "require('$(CONFIG)').$(1)")
endef


# Deploy Packages - Names of the *.yml files in arch (ordered)
APIBASE := api/apibase
APIROUTE := api/apiroute
APISTAGE := api/apistage

APIHANDLERFUNC := api/handler/handlerfunc


PKGS := $(APIBASE) $(APIHANDLERFUNC) $(APIROUTE) $(APISTAGE)


EXECUTABLES := $(wildcard arch/*/*/*.go)
OUTPUTS := $(patsubst arch/%.go,%,$(EXECUTABLES))

# Files to be built
define build_targets
build $(notdir $(1)): $(1)

$(1):
	@mkdir -p bin/$$@
	@env GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bin/$$@/bootstrap arch/$$@.go 
	@echo "-> Successfully built $$@.go => bin/$$@"

endef

$(foreach output,$(OUTPUTS),$(eval $(call build_targets,$(output))))


S3-BUCKET := $(call GetValueFromConfig, s3Bucket)
STAGE := $(call GetValueFromConfig, Stage)

Hi:
	echo $(S3-BUCKET)
	echo $(STAGE)


# Files to be deployed
define pkg_targets
deploy: $(notdir $(1))-deploy

$(notdir $(1))-deploy:
	@echo "==============================================="
	@echo "Read configuration from $(CONFIG)."
	@echo "Stage - $(STAGE), S3-Bucket - $(S3-BUCKET)."
	@echo "AWS Profile - $(AWS_PROFILE), AWS Default Region - $(AWS_DEFAULT_REGION)."
	@echo "==============================================="
	@[ "$(S3-BUCKET)" ] || ( echo ">> Please enter valid s3 bucket name or create a bucket using command: aws s3 mb <YOUR_BUCKET_NAME>"; exit 1)
	@sam package --template-file arch/$(1).yml --output-template-file arch/$(1)_package.yml --s3-bucket "$(S3-BUCKET)"

	@sam deploy --template-file arch/$(1)_package.yml --stack-name diag-$(notdir $(1)) --capabilities CAPABILITY_NAMED_IAM --no-fail-on-empty-changeset --parameter-overrides "StageName=$(STAGE)" $(2)

remove $(1)-remove: $(notdir $(1))-remove

$(notdir $(1))-remove:
	@echo "==============================================="
	@echo "Removing $(1)."
	@echo "==============================================="
	@aws cloudformation delete-stack --stack-name diag-$(notdir $(1))

cfnlint $(1)-cfnlint: $(notdir $(1))-cfnlint

$(notdir $(1))-cfnlint:
	@echo "==============================================="
	@echo "Linting $(1)."
	@echo "==============================================="
	@cfn-lint arch/$(1).yml
	@sam validate -t arch/$(1).yml

endef

# Files to be published
$(foreach pkg,$(PKGS),$(eval $(call pkg_targets,$(pkg), $(config))))



clean:
	@echo "Cleaning bin folder"
	@rm -rf bin/
	@echo "Cleaning package.yml files"
	@find arch -name \*_package.yml -type f -delete
	@find . -name \*.yml.bak -type f -delete