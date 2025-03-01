// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Existing builds can be imported using the combination of the application name, a colon, and the build ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import heroku:build/build:Build foobar bazbux:4f1db8ef-ed5c-4c42-a3d6-3c28262d5abc
 * ```
 *
 * * `foobar` is the **heroku_build** resource's name
 *
 * * `bazbux` is the Heroku app name (or ID) that the build belongs to
 *
 * * `:` separates the app identifier & the build identifier
 *
 * * `4f1db8ef…` is the build ID
 */
export class Build extends pulumi.CustomResource {
    /**
     * Get an existing Build resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BuildState, opts?: pulumi.CustomResourceOptions): Build {
        return new Build(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'heroku:build/build:Build';

    /**
     * Returns true if the given object is an instance of Build.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Build {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Build.__pulumiType;
    }

    /**
     * Heroku app ID (do not use app name)
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * List of buildpack GitHub URLs
     */
    public readonly buildpacks!: pulumi.Output<string[]>;
    public /*out*/ readonly localChecksum!: pulumi.Output<string>;
    /**
     * URL that [streams the log output from the build](https://devcenter.heroku.com/articles/build-and-release-using-the-api#streaming-build-output)
     */
    public /*out*/ readonly outputStreamUrl!: pulumi.Output<string>;
    /**
     * The Heroku app release created with a build's slug
     */
    public /*out*/ readonly releaseId!: pulumi.Output<string>;
    /**
     * The Heroku slug created by a build
     */
    public /*out*/ readonly slugId!: pulumi.Output<string>;
    /**
     * A block that specifies the source code to build & release:
     */
    public readonly source!: pulumi.Output<outputs.build.BuildSource>;
    /**
     * Name or ID of the [Heroku stack](https://devcenter.heroku.com/articles/stack)
     */
    public /*out*/ readonly stack!: pulumi.Output<string>;
    /**
     * The status of a build. Possible values are `pending`, `successful` and `failed`
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Heroku account that created a build
     */
    public /*out*/ readonly users!: pulumi.Output<outputs.build.BuildUser[]>;
    /**
     * The ID of the build
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;

    /**
     * Create a Build resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BuildArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BuildArgs | BuildState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BuildState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["buildpacks"] = state ? state.buildpacks : undefined;
            resourceInputs["localChecksum"] = state ? state.localChecksum : undefined;
            resourceInputs["outputStreamUrl"] = state ? state.outputStreamUrl : undefined;
            resourceInputs["releaseId"] = state ? state.releaseId : undefined;
            resourceInputs["slugId"] = state ? state.slugId : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["stack"] = state ? state.stack : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as BuildArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["buildpacks"] = args ? args.buildpacks : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["localChecksum"] = undefined /*out*/;
            resourceInputs["outputStreamUrl"] = undefined /*out*/;
            resourceInputs["releaseId"] = undefined /*out*/;
            resourceInputs["slugId"] = undefined /*out*/;
            resourceInputs["stack"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["users"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Build.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Build resources.
 */
export interface BuildState {
    /**
     * Heroku app ID (do not use app name)
     */
    appId?: pulumi.Input<string>;
    /**
     * List of buildpack GitHub URLs
     */
    buildpacks?: pulumi.Input<pulumi.Input<string>[]>;
    localChecksum?: pulumi.Input<string>;
    /**
     * URL that [streams the log output from the build](https://devcenter.heroku.com/articles/build-and-release-using-the-api#streaming-build-output)
     */
    outputStreamUrl?: pulumi.Input<string>;
    /**
     * The Heroku app release created with a build's slug
     */
    releaseId?: pulumi.Input<string>;
    /**
     * The Heroku slug created by a build
     */
    slugId?: pulumi.Input<string>;
    /**
     * A block that specifies the source code to build & release:
     */
    source?: pulumi.Input<inputs.build.BuildSource>;
    /**
     * Name or ID of the [Heroku stack](https://devcenter.heroku.com/articles/stack)
     */
    stack?: pulumi.Input<string>;
    /**
     * The status of a build. Possible values are `pending`, `successful` and `failed`
     */
    status?: pulumi.Input<string>;
    /**
     * Heroku account that created a build
     */
    users?: pulumi.Input<pulumi.Input<inputs.build.BuildUser>[]>;
    /**
     * The ID of the build
     */
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Build resource.
 */
export interface BuildArgs {
    /**
     * Heroku app ID (do not use app name)
     */
    appId: pulumi.Input<string>;
    /**
     * List of buildpack GitHub URLs
     */
    buildpacks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A block that specifies the source code to build & release:
     */
    source: pulumi.Input<inputs.build.BuildSource>;
}
