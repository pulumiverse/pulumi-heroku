// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pipeline

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-heroku/sdk/go/heroku/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "heroku:pipeline/configVar:ConfigVar":
		r = &ConfigVar{}
	case "heroku:pipeline/coupling:Coupling":
		r = &Coupling{}
	case "heroku:pipeline/pipeline:Pipeline":
		r = &Pipeline{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"heroku",
		"pipeline/configVar",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"heroku",
		"pipeline/coupling",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"heroku",
		"pipeline/pipeline",
		&module{version},
	)
}
