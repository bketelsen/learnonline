GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GOBUILDPROD=$(GOBUILD) -ldflags "-linkmode external -extldflags -static" 
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
DOCKER=docker
DOCKERCOMPOSE=docker-compose
SODA=buffalo db
BUFFALO=buffalo
APP=learnonline
DB=learnonline_db

deps: 
	$(GOCMD) get -u -t -v github.com/gobuffalo/buffalo/buffalo
	$(GOGET) -t -v -u github.com/gobuffalo/buffalo  && $(GOINSTALL) github.com/gobuffalo/buffalo
	$(GOGET) -t -v -u github.com/markbates/pop      && $(GOINSTALL) github.com/markbates/pop

build:
	$(BUFFALO) build -o bin/$(APP)

buildprod:
	$(GOBUILDPROD) -v -o $(APP)

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/$(APP)
	rm -rf $(APP)

test:
	$(GOTEST) -v ./grifts -race
	$(GOTEST) -v ./models -race
	$(GOTEST) -v ./actions -race

db:
	$(DOCKER) run --name=$(DB) -d -p 5432:5432 -e POSTGRES_DB=$(APP)_development postgres
	sleep 20
	$(BUFFALO) db migrate up
	$(DOCKER) ps | grep $(DB)

db-down: 
	$(DOCKER) stop $(DB)
	$(DOCKER) rm $(DB)

dev: deps db

teardown-dev: clean
	$(DOCKERCOMPOSE) down

run: dev
	$(BUFFALO) dev

define GIT_ERROR
FATAL: Git (git) is required to download gcon dependencies.
endef

define HG_ERROR
FATAL: Mercurial (hg) is required to download gcon dependencies.
endef

define GLIDE_ERROR
FATAL: Glide (glide) is required to download gcon dependencies.
endef

# check for git
git:
	$(if $(shell git), , $(error $(GIT_ERROR)))

# check for mercurial
hg:
	$(if $(shell hg), , $(error $(HG_ERROR)))
